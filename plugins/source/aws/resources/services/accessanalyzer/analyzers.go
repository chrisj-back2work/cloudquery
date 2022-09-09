package accessanalyzer

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsmiddleware "github.com/aws/aws-sdk-go-v2/aws/middleware"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
	"github.com/aws/smithy-go/middleware"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Analyzers() *schema.Table {
	return &schema.Table{
		Name:        "aws_access_analyzer_analyzers",
		Description: "Contains information about the analyzer",
		Resolver:    fetchAccessAnalyzerAnalyzers,
		Multiplex:   client.ServiceAccountRegionMultiplexer("access-analyzer"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:            "arn",
				Description:     "The ARN of the analyzer",
				Type:            schema.TypeString,
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "created_at",
				Description: "A timestamp for the time at which the analyzer was created",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:        "name",
				Description: "The name of the analyzer",
				Type:        schema.TypeString,
			},
			{
				Name:        "status",
				Description: "The status of the analyzer",
				Type:        schema.TypeString,
			},
			{
				Name:        "type",
				Description: "The type of analyzer, which corresponds to the zone of trust chosen for the analyzer",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_resource_analyzed",
				Description: "The resource that was most recently analyzed by the analyzer",
				Type:        schema.TypeString,
			},
			{
				Name:        "last_resource_analyzed_at",
				Description: "The time at which the most recently analyzed resource was analyzed",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:     "status_reason",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("StatusReason"),
			},
			{
				Name:        "tags",
				Description: "The tags added to the analyzer",
				Type:        schema.TypeJSON,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_access_analyzer_analyzer_findings",
				Description: "Contains information about a finding",
				Resolver:    fetchAccessAnalyzerAnalyzerFindings,
				Columns: []schema.Column{
					{
						Name:        "analyzer_cq_id",
						Description: "Unique CloudQuery ID of aws_access_analyzer_analyzers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "analyzed_at",
						Description: "The time at which the resource-based policy that generated the finding was analyzed",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "condition",
						Description: "The condition in the analyzed policy statement that resulted in a finding",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "created_at",
						Description: "The time at which the finding was created",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "id",
						Description: "The ID of the finding",
						Type:        schema.TypeString,
					},
					{
						Name:        "resource_owner_account",
						Description: "The Amazon Web Services account ID that owns the resource",
						Type:        schema.TypeString,
					},
					{
						Name:        "resource_type",
						Description: "The type of the resource that the external principal has access to",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "The status of the finding",
						Type:        schema.TypeString,
					},
					{
						Name:        "updated_at",
						Description: "The time at which the finding was most recently updated",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "action",
						Description: "The action in the analyzed policy statement that an external principal has permission to use",
						Type:        schema.TypeStringArray,
					},
					{
						Name:        "error",
						Description: "The error that resulted in an Error finding",
						Type:        schema.TypeString,
					},
					{
						Name:        "is_public",
						Description: "Indicates whether the finding reports a resource that has a policy that allows public access",
						Type:        schema.TypeBool,
					},
					{
						Name:        "principal",
						Description: "The external principal that has access to a resource within the zone of trust",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "resource",
						Description: "The resource that the external principal has access to",
						Type:        schema.TypeString,
					},
					{
						Name:        "sources",
						Description: "The source of the finding",
						Type:        schema.TypeJSON,
					},
				},
			},
			{
				Name:        "aws_access_analyzer_analyzer_archive_rules",
				Description: "Contains information about an archive rule",
				Resolver:    fetchAccessAnalyzerAnalyzerArchiveRules,
				Columns: []schema.Column{
					{
						Name:        "analyzer_cq_id",
						Description: "Unique CloudQuery ID of aws_access_analyzer_analyzers table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "created_at",
						Description: "The time at which the archive rule was created",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "filter",
						Description: "A filter used to define the archive rule",
						Type:        schema.TypeJSON,
					},
					{
						Name:        "rule_name",
						Description: "The name of the archive rule",
						Type:        schema.TypeString,
					},
					{
						Name:        "updated_at",
						Description: "The time at which the archive rule was last updated",
						Type:        schema.TypeTimestamp,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchAccessAnalyzerAnalyzers(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	config := accessanalyzer.ListAnalyzersInput{}
	c := meta.(*client.Client)
	svc := c.Services().Analyzer
	for {
		response, err := svc.ListAnalyzers(ctx, &config, func(options *accessanalyzer.Options) {
			options.APIOptions = append(options.APIOptions, func(stack *middleware.Stack) error {
				if err := stack.Initialize.Add(&awsmiddleware.RegisterServiceMetadata{
					Region:        c.Region,
					ServiceID:     accessanalyzer.ServiceID,
					SigningName:   "access-analyzer",
					OperationName: "ListAnalyzers",
				}, middleware.Before); err != nil {
					return nil
				}
				return nil
			})
		})
		if err != nil {
			return err
		}

		res <- response.Analyzers
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func fetchAccessAnalyzerAnalyzerFindings(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	analyzer := parent.Item.(types.AnalyzerSummary)
	c := meta.(*client.Client)
	svc := c.Services().Analyzer
	config := accessanalyzer.ListFindingsInput{
		AnalyzerArn: analyzer.Arn,
	}
	for {
		response, err := svc.ListFindings(ctx, &config)
		if err != nil {
			return err
		}

		res <- response.Findings
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}
func fetchAccessAnalyzerAnalyzerArchiveRules(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	analyzer := parent.Item.(types.AnalyzerSummary)
	c := meta.(*client.Client)
	svc := c.Services().Analyzer
	config := accessanalyzer.ListArchiveRulesInput{
		AnalyzerName: analyzer.Name,
	}
	for {
		response, err := svc.ListArchiveRules(ctx, &config)
		if err != nil {
			return err
		}

		res <- response.ArchiveRules
		if aws.ToString(response.NextToken) == "" {
			break
		}
		config.NextToken = response.NextToken
	}
	return nil
}

// Code generated by codegen; DO NOT EDIT.

package codepipeline

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Webhooks() *schema.Table {
	return &schema.Table{
		Name:        "aws_codepipeline_webhooks",
		Description: "https://docs.aws.amazon.com/codepipeline/latest/APIReference/API_ListWebhookItem.html",
		Resolver:    fetchCodepipelineWebhooks,
		Multiplex:   client.ServiceAccountRegionMultiplexer("codepipeline"),
		Columns: []schema.Column{
			{
				Name:     "account_id",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSAccount,
			},
			{
				Name:     "region",
				Type:     schema.TypeString,
				Resolver: client.ResolveAWSRegion,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "definition",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Definition"),
			},
			{
				Name:     "url",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Url"),
			},
			{
				Name:     "error_code",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ErrorCode"),
			},
			{
				Name:     "error_message",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ErrorMessage"),
			},
			{
				Name:     "last_triggered",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastTriggered"),
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Tags"),
			},
		},
	}
}

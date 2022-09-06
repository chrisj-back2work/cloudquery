// Code generated by codegen using template resource_list_describe.go.tpl; DO NOT EDIT.

package cloudfront

import (
	"context"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"

	"github.com/aws/aws-sdk-go-v2/service/cloudfront"
)

func CloudfrontDistributions() *schema.Table {
	return &schema.Table{
		Name:      "aws_cloudfront_distributions",
		Resolver:  fetchCloudfrontDistributions,
		Multiplex: client.ServiceAccountRegionMultiplexer("cloudfront"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
				Description: `The AWS Account ID of the resource.`,
			},
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ARN"),
			},
			{
				Name:     "distribution_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DistributionConfig"),
			},
			{
				Name:     "domain_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DomainName"),
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Id"),
			},
			{
				Name:     "in_progress_invalidation_batches",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("InProgressInvalidationBatches"),
			},
			{
				Name:     "last_modified_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastModifiedTime"),
			},
			{
				Name:     "status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status"),
			},
			{
				Name:     "active_trusted_key_groups",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ActiveTrustedKeyGroups"),
			},
			{
				Name:     "active_trusted_signers",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ActiveTrustedSigners"),
			},
			{
				Name:     "alias_icp_recordals",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("AliasICPRecordals"),
			},
		},
	}
}

func fetchCloudfrontDistributions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Cloudfront

	input := cloudfront.ListDistributionsInput{}
	paginator := cloudfront.NewListDistributionsPaginator(svc, &input)
	for paginator.HasMorePages() {
		output, err := paginator.NextPage(ctx)
		if err != nil {

			return errors.WithStack(err)
		}

		for _, item := range output.DistributionList.Items {

			do, err := svc.GetDistribution(ctx, &cloudfront.GetDistributionInput{

				Id: item.Id,
			})
			if err != nil {

				if cl.IsNotFoundError(err) {
					continue
				}
				return errors.WithStack(err)
			}
			res <- do.Distribution
		}
	}
	return nil
}
// Code generated by codegen; DO NOT EDIT.

package xray

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func SamplingRules() *schema.Table {
	return &schema.Table{
		Name:      "aws_xray_sampling_rules",
		Resolver:  fetchXraySamplingRules,
		Multiplex: client.ServiceAccountRegionMultiplexer("xray"),
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
				Resolver: schema.PathResolver("SamplingRule.RuleARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "tags",
				Type:     schema.TypeJSON,
				Resolver: resolveXraySamplingRuleTags,
			},
			{
				Name:     "created_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedAt"),
			},
			{
				Name:     "modified_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ModifiedAt"),
			},
			{
				Name:     "sampling_rule",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SamplingRule"),
			},
		},
	}
}

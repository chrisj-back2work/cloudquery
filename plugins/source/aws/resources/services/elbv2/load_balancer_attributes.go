// Code generated by codegen; DO NOT EDIT.

package elbv2

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func LoadBalancerAttributes() *schema.Table {
	return &schema.Table{
		Name:      "aws_elbv2_load_balancer_attributes",
		Resolver:  fetchElbv2LoadBalancerAttributes,
		Multiplex: client.ServiceAccountRegionMultiplexer("elasticloadbalancing"),
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
				Name:     "load_balancer_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("arn"),
			},
			{
				Name:     "key",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Key"),
			},
			{
				Name:     "value",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Value"),
			},
		},
	}
}
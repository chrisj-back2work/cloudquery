// Code generated by codegen; DO NOT EDIT.

package mq

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func BrokerConfigurationRevisions() *schema.Table {
	return &schema.Table{
		Name:      "aws_mq_broker_configuration_revisions",
		Resolver:  fetchMqBrokerConfigurationRevisions,
		Multiplex: client.ServiceAccountRegionMultiplexer("mq"),
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
				Name:     "broker_configuration_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentColumnResolver("arn"),
			},
			{
				Name:     "data",
				Type:     schema.TypeJSON,
				Resolver: resolveBrokerConfigurationRevisionsData,
			},
			{
				Name:     "configuration_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ConfigurationId"),
			},
			{
				Name:     "created",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Created"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "result_metadata",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ResultMetadata"),
			},
		},
	}
}

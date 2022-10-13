// Code generated by codegen; DO NOT EDIT.

package frauddetector

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Detectors() *schema.Table {
	return &schema.Table{
		Name:      "aws_frauddetector_detectors",
		Resolver:  fetchFrauddetectorDetectors,
		Multiplex: client.ServiceAccountRegionMultiplexer("frauddetector"),
		Columns: []schema.Column{
			{
				Name:     "arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Arn"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "created_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CreatedTime"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "detector_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DetectorId"),
			},
			{
				Name:     "event_type_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EventTypeName"),
			},
			{
				Name:     "last_updated_time",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastUpdatedTime"),
			},
		},
	}
}

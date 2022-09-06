// Code generated by codegen using template resource_get.go.tpl; DO NOT EDIT.

package autoscaling

import (
	"context"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
	"github.com/pkg/errors"

	"github.com/aws/aws-sdk-go-v2/service/autoscaling"
)

func AutoscalingScheduledActions() *schema.Table {
	return &schema.Table{
		Name:      "aws_autoscaling_scheduled_actions",
		Resolver:  fetchAutoscalingScheduledActions,
		Multiplex: client.ServiceAccountRegionMultiplexer("autoscaling"),
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
				Description: `The AWS Account ID of the resource.`,
			},
			{
				Name:        "region",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
				Description: `The AWS Region of the resource.`,
			},
			{
				Name:     "auto_scaling_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AutoScalingGroupName"),
			},
			{
				Name:     "desired_capacity",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DesiredCapacity"),
			},
			{
				Name:     "end_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("EndTime"),
			},
			{
				Name:     "max_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaxSize"),
			},
			{
				Name:     "min_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MinSize"),
			},
			{
				Name:     "recurrence",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Recurrence"),
			},
			{
				Name:     "scheduled_action_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ScheduledActionARN"),
			},
			{
				Name:     "scheduled_action_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ScheduledActionName"),
			},
			{
				Name:     "start_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("StartTime"),
			},
			{
				Name:     "time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Time"),
			},
			{
				Name:     "time_zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TimeZone"),
			},
		},
	}
}

func fetchAutoscalingScheduledActions(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Autoscaling

	input := autoscaling.DescribeScheduledActionsInput{}

	for {
		response, err := svc.DescribeScheduledActions(ctx, &input)
		if err != nil {

			return errors.WithStack(err)
		}

		res <- response.ScheduledUpdateGroupActions

		if aws.ToString(response.NextToken) == "" {
			break
		}
		input.NextToken = response.NextToken
	}
	return nil
}

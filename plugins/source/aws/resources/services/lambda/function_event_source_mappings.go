// Code generated by codegen; DO NOT EDIT.

package lambda

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func FunctionEventSourceMappings() *schema.Table {
	return &schema.Table{
		Name:      "aws_lambda_function_event_source_mappings",
		Resolver:  fetchLambdaFunctionEventSourceMappings,
		Multiplex: client.ServiceAccountRegionMultiplexer("lambda"),
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
				Name:     "function_arn",
				Type:     schema.TypeString,
				Resolver: schema.ParentResourceFieldResolver("arn"),
			},
			{
				Name:     "batch_size",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("BatchSize"),
			},
			{
				Name:     "bisect_batch_on_function_error",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("BisectBatchOnFunctionError"),
			},
			{
				Name:     "destination_config",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("DestinationConfig"),
			},
			{
				Name:     "event_source_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EventSourceArn"),
			},
			{
				Name:     "filter_criteria",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("FilterCriteria"),
			},
			{
				Name:     "function_response_types",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("FunctionResponseTypes"),
			},
			{
				Name:     "last_modified",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastModified"),
			},
			{
				Name:     "last_processing_result",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LastProcessingResult"),
			},
			{
				Name:     "maximum_batching_window_in_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaximumBatchingWindowInSeconds"),
			},
			{
				Name:     "maximum_record_age_in_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaximumRecordAgeInSeconds"),
			},
			{
				Name:     "maximum_retry_attempts",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("MaximumRetryAttempts"),
			},
			{
				Name:     "parallelization_factor",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ParallelizationFactor"),
			},
			{
				Name:     "queues",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Queues"),
			},
			{
				Name:     "self_managed_event_source",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SelfManagedEventSource"),
			},
			{
				Name:     "source_access_configurations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("SourceAccessConfigurations"),
			},
			{
				Name:     "starting_position",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StartingPosition"),
			},
			{
				Name:     "starting_position_timestamp",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("StartingPositionTimestamp"),
			},
			{
				Name:     "state",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("State"),
			},
			{
				Name:     "state_transition_reason",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("StateTransitionReason"),
			},
			{
				Name:     "topics",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Topics"),
			},
			{
				Name:     "tumbling_window_in_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("TumblingWindowInSeconds"),
			},
			{
				Name:     "uuid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UUID"),
			},
		},
	}
}
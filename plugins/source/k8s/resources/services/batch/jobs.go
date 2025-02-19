// Code generated by codegen; DO NOT EDIT.

package batch

import (
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Jobs() *schema.Table {
	return &schema.Table{
		Name:      "k8s_batch_jobs",
		Resolver:  fetchBatchJobs,
		Multiplex: client.ContextMultiplex,
		Columns: []schema.Column{
			{
				Name:     "context",
				Type:     schema.TypeString,
				Resolver: client.ResolveContext,
			},
			{
				Name:     "uid",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UID"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "kind",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Kind"),
			},
			{
				Name:     "api_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("APIVersion"),
			},
			{
				Name:     "name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Name"),
			},
			{
				Name:     "namespace",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Namespace"),
			},
			{
				Name:     "resource_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ResourceVersion"),
			},
			{
				Name:     "generation",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Generation"),
			},
			{
				Name:     "deletion_grace_period_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("DeletionGracePeriodSeconds"),
			},
			{
				Name:     "labels",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Labels"),
			},
			{
				Name:     "annotations",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Annotations"),
			},
			{
				Name:     "owner_references",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("OwnerReferences"),
			},
			{
				Name:     "finalizers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Finalizers"),
			},
			{
				Name:     "spec_parallelism",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Spec.Parallelism"),
			},
			{
				Name:     "spec_completions",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Spec.Completions"),
			},
			{
				Name:     "spec_active_deadline_seconds",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Spec.ActiveDeadlineSeconds"),
			},
			{
				Name:     "spec_pod_failure_policy",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.PodFailurePolicy"),
			},
			{
				Name:     "spec_backoff_limit",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Spec.BackoffLimit"),
			},
			{
				Name:     "spec_selector",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.Selector"),
			},
			{
				Name:     "spec_manual_selector",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Spec.ManualSelector"),
			},
			{
				Name:     "spec_template",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Spec.Template"),
			},
			{
				Name:     "spec_ttl_seconds_after_finished",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Spec.TTLSecondsAfterFinished"),
			},
			{
				Name:     "spec_completion_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Spec.CompletionMode"),
			},
			{
				Name:     "spec_suspend",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("Spec.Suspend"),
			},
			{
				Name:     "status_conditions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status.Conditions"),
			},
			{
				Name:     "status_start_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Status.StartTime"),
			},
			{
				Name:     "status_completion_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("Status.CompletionTime"),
			},
			{
				Name:     "status_active",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Status.Active"),
			},
			{
				Name:     "status_succeeded",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Status.Succeeded"),
			},
			{
				Name:     "status_failed",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Status.Failed"),
			},
			{
				Name:     "status_completed_indexes",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status.CompletedIndexes"),
			},
			{
				Name:     "status_uncounted_terminated_pods",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status.UncountedTerminatedPods"),
			},
			{
				Name:     "status_ready",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Status.Ready"),
			},
		},
	}
}

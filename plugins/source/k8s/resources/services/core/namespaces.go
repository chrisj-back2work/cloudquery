// Code generated by codegen; DO NOT EDIT.

package core

import (
	"github.com/cloudquery/cloudquery/plugins/source/k8s/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Namespaces() *schema.Table {
	return &schema.Table{
		Name:      "k8s_core_namespaces",
		Resolver:  fetchCoreNamespaces,
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
				Name:     "spec_finalizers",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("Spec.Finalizers"),
			},
			{
				Name:     "status_phase",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Status.Phase"),
			},
			{
				Name:     "status_conditions",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Status.Conditions"),
			},
		},
	}
}

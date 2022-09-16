// Code generated by codegen using template resource_manual.go.tpl; DO NOT EDIT.

package waf_packages

import (
	"github.com/cloudquery/plugin-sdk/schema"
)

func wafRules() *schema.Table {
	return &schema.Table{
		Name:     "cloudflare_waf_rules",
		Resolver: fetchWAFRules,
		Columns: []schema.Column{
			{
				Name:     "waf_package_cq_id",
				Type:     schema.TypeUUID,
				Resolver: schema.ParentIDResolver,
			},
			{
				Name:     "id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ID"),
			},
			{
				Name:     "description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Description"),
			},
			{
				Name:     "priority",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Priority"),
			},
			{
				Name:     "package_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PackageID"),
			},
			{
				Name:     "group",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("Group"),
			},
			{
				Name:     "mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Mode"),
			},
			{
				Name:     "default_mode",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DefaultMode"),
			},
			{
				Name:     "allowed_modes",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("AllowedModes"),
			},
		},
	}
}
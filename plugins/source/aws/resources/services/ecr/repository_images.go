// Code generated by codegen; DO NOT EDIT.

package ecr

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func RepositoryImages() *schema.Table {
	return &schema.Table{
		Name:     "aws_ecr_repository_images",
		Resolver: fetchEcrRepositoryImages,
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
				Resolver: resolveImageArn,
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "artifact_media_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ArtifactMediaType"),
			},
			{
				Name:     "image_digest",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ImageDigest"),
			},
			{
				Name:     "image_manifest_media_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ImageManifestMediaType"),
			},
			{
				Name:     "image_pushed_at",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("ImagePushedAt"),
			},
			{
				Name:     "image_scan_findings_summary",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ImageScanFindingsSummary"),
			},
			{
				Name:     "image_scan_status",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("ImageScanStatus"),
			},
			{
				Name:     "image_size_in_bytes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("ImageSizeInBytes"),
			},
			{
				Name:     "image_tags",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ImageTags"),
			},
			{
				Name:     "last_recorded_pull_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("LastRecordedPullTime"),
			},
			{
				Name:     "registry_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RegistryId"),
			},
			{
				Name:     "repository_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RepositoryName"),
			},
		},
	}
}
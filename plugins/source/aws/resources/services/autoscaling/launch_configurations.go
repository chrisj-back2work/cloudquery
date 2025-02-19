// Code generated by codegen; DO NOT EDIT.

package autoscaling

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func LaunchConfigurations() *schema.Table {
	return &schema.Table{
		Name:        "aws_autoscaling_launch_configurations",
		Description: "https://docs.aws.amazon.com/autoscaling/ec2/APIReference/API_LaunchConfiguration.html",
		Resolver:    fetchAutoscalingLaunchConfigurations,
		Multiplex:   client.ServiceAccountRegionMultiplexer("autoscaling"),
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
				Resolver: schema.PathResolver("LaunchConfigurationARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "created_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CreatedTime"),
			},
			{
				Name:     "image_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ImageId"),
			},
			{
				Name:     "instance_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("InstanceType"),
			},
			{
				Name:     "launch_configuration_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("LaunchConfigurationName"),
			},
			{
				Name:     "associate_public_ip_address",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AssociatePublicIpAddress"),
			},
			{
				Name:     "block_device_mappings",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("BlockDeviceMappings"),
			},
			{
				Name:     "classic_link_vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ClassicLinkVPCId"),
			},
			{
				Name:     "classic_link_vpc_security_groups",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("ClassicLinkVPCSecurityGroups"),
			},
			{
				Name:     "ebs_optimized",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("EbsOptimized"),
			},
			{
				Name:     "iam_instance_profile",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("IamInstanceProfile"),
			},
			{
				Name:     "instance_monitoring",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("InstanceMonitoring"),
			},
			{
				Name:     "kernel_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KernelId"),
			},
			{
				Name:     "key_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KeyName"),
			},
			{
				Name:     "metadata_options",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("MetadataOptions"),
			},
			{
				Name:     "placement_tenancy",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PlacementTenancy"),
			},
			{
				Name:     "ramdisk_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("RamdiskId"),
			},
			{
				Name:     "security_groups",
				Type:     schema.TypeStringArray,
				Resolver: schema.PathResolver("SecurityGroups"),
			},
			{
				Name:     "spot_price",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SpotPrice"),
			},
			{
				Name:     "user_data",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("UserData"),
			},
		},
	}
}

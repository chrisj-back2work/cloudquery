// Code generated by codegen; DO NOT EDIT.

package elasticache

import (
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Snapshots() *schema.Table {
	return &schema.Table{
		Name:        "aws_elasticache_snapshots",
		Description: "https://docs.aws.amazon.com/AmazonElastiCache/latest/APIReference/API_Snapshot.html",
		Resolver:    fetchElasticacheSnapshots,
		Multiplex:   client.ServiceAccountRegionMultiplexer("elasticache"),
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
				Resolver: schema.PathResolver("ARN"),
				CreationOptions: schema.ColumnCreationOptions{
					PrimaryKey: true,
				},
			},
			{
				Name:     "auto_minor_version_upgrade",
				Type:     schema.TypeBool,
				Resolver: schema.PathResolver("AutoMinorVersionUpgrade"),
			},
			{
				Name:     "automatic_failover",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("AutomaticFailover"),
			},
			{
				Name:     "cache_cluster_create_time",
				Type:     schema.TypeTimestamp,
				Resolver: schema.PathResolver("CacheClusterCreateTime"),
			},
			{
				Name:     "cache_cluster_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CacheClusterId"),
			},
			{
				Name:     "cache_node_type",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CacheNodeType"),
			},
			{
				Name:     "cache_parameter_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CacheParameterGroupName"),
			},
			{
				Name:     "cache_subnet_group_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("CacheSubnetGroupName"),
			},
			{
				Name:     "data_tiering",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("DataTiering"),
			},
			{
				Name:     "engine",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("Engine"),
			},
			{
				Name:     "engine_version",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("EngineVersion"),
			},
			{
				Name:     "kms_key_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("KmsKeyId"),
			},
			{
				Name:     "node_snapshots",
				Type:     schema.TypeJSON,
				Resolver: schema.PathResolver("NodeSnapshots"),
			},
			{
				Name:     "num_cache_nodes",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NumCacheNodes"),
			},
			{
				Name:     "num_node_groups",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("NumNodeGroups"),
			},
			{
				Name:     "port",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("Port"),
			},
			{
				Name:     "preferred_availability_zone",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PreferredAvailabilityZone"),
			},
			{
				Name:     "preferred_maintenance_window",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PreferredMaintenanceWindow"),
			},
			{
				Name:     "preferred_outpost_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("PreferredOutpostArn"),
			},
			{
				Name:     "replication_group_description",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicationGroupDescription"),
			},
			{
				Name:     "replication_group_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("ReplicationGroupId"),
			},
			{
				Name:     "snapshot_name",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SnapshotName"),
			},
			{
				Name:     "snapshot_retention_limit",
				Type:     schema.TypeInt,
				Resolver: schema.PathResolver("SnapshotRetentionLimit"),
			},
			{
				Name:     "snapshot_source",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SnapshotSource"),
			},
			{
				Name:     "snapshot_status",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SnapshotStatus"),
			},
			{
				Name:     "snapshot_window",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("SnapshotWindow"),
			},
			{
				Name:     "topic_arn",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("TopicArn"),
			},
			{
				Name:     "vpc_id",
				Type:     schema.TypeString,
				Resolver: schema.PathResolver("VpcId"),
			},
		},
	}
}

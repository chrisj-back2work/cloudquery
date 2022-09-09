package backup

import (
	"context"
	"errors"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/aws/arn"
	"github.com/aws/aws-sdk-go-v2/service/backup"
	"github.com/aws/aws-sdk-go-v2/service/backup/types"
	"github.com/aws/smithy-go"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func Vaults() *schema.Table {
	return &schema.Table{
		Name:                 "aws_backup_vaults",
		Description:          "Contains metadata about a backup vault.",
		Resolver:             fetchBackupVaults,
		Multiplex:            client.ServiceAccountRegionMultiplexer("backup"),
		PostResourceResolver: resolveVaultNotifications,
		Columns: []schema.Column{
			{
				Name:        "account_id",
				Description: "The AWS Account ID of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSAccount,
			},
			{
				Name:        "region",
				Description: "The AWS Region of the resource.",
				Type:        schema.TypeString,
				Resolver:    client.ResolveAWSRegion,
			},
			{
				Name:            "arn",
				Description:     "An Amazon Resource Name (ARN) that uniquely identifies a backup vault; for example, arn:aws:backup:us-east-1:123456789012:vault:aBackupVault.",
				Type:            schema.TypeString,
				Resolver:        schema.PathResolver("BackupVaultArn"),
				CreationOptions: schema.ColumnCreationOptions{PrimaryKey: true},
			},
			{
				Name:        "name",
				Description: "The name of a logical container where backups are stored.",
				Type:        schema.TypeString,
				Resolver:    schema.PathResolver("BackupVaultName"),
			},
			{
				Name:        "creation_date",
				Description: "The date and time a resource backup is created, in Unix format and Coordinated Universal Time (UTC).",
				Type:        schema.TypeTimestamp,
			},
			{
				Name:          "creator_request_id",
				Description:   "A unique string that identifies the request and allows failed requests to be retried without the risk of running the operation twice.",
				Type:          schema.TypeString,
				IgnoreInTests: true,
			},
			{
				Name:        "encryption_key_arn",
				Description: "A server-side encryption key you can specify to encrypt your backups from services that support full Backup management.",
				Type:        schema.TypeString,
			},
			{
				Name:          "lock_date",
				Description:   "The date and time when Backup Vault Lock configuration becomes immutable, meaning it cannot be changed or deleted.",
				Type:          schema.TypeTimestamp,
				IgnoreInTests: true,
			},
			{
				Name:        "locked",
				Description: "A Boolean value that indicates whether Backup Vault Lock applies to the selected backup vault.",
				Type:        schema.TypeBool,
			},
			{
				Name:          "max_retention_days",
				Description:   "The Backup Vault Lock setting that specifies the maximum retention period that the vault retains its recovery points.",
				Type:          schema.TypeInt,
				IgnoreInTests: true,
			},
			{
				Name:          "min_retention_days",
				Description:   "The Backup Vault Lock setting that specifies the minimum retention period that the vault retains its recovery points.",
				Type:          schema.TypeInt,
				IgnoreInTests: true,
			},
			{
				Name:        "number_of_recovery_points",
				Description: "The number of recovery points that are stored in a backup vault.",
				Type:        schema.TypeInt,
			},
			{
				Name:          "access_policy",
				Description:   "The backup vault access policy document in JSON format.",
				Type:          schema.TypeJSON,
				Resolver:      resolveVaultAccessPolicy,
				IgnoreInTests: true,
			},
			{
				Name:        "notification_events",
				Description: "An array of events that indicate the status of jobs to back up resources to the backup vault.",
				Type:        schema.TypeStringArray,
			},
			{
				Name:        "notification_sns_topic_arn",
				Description: "An ARN that uniquely identifies an Amazon Simple Notification Service topic.",
				Type:        schema.TypeString,
			},
			{
				Name:        "tags",
				Description: "Resource tags",
				Type:        schema.TypeJSON,
				Resolver:    resolveVaultTags,
			},
		},
		Relations: []*schema.Table{
			{
				Name:        "aws_backup_vault_recovery_points",
				Description: "The recovery points stored in a backup vault.",
				Resolver:    fetchVaultRecoveryPoints,

				IgnoreInTests: true,
				Columns: []schema.Column{
					{
						Name:        "vault_cq_id",
						Description: "Unique CloudQuery ID of aws_backup_vault table (FK)",
						Type:        schema.TypeUUID,
						Resolver:    schema.ParentIdResolver,
					},
					{
						Name:        "backup_size",
						Description: "The size, in bytes, of a backup.",
						Type:        schema.TypeInt,
						Resolver:    schema.PathResolver("BackupSizeInBytes"),
					},
					{
						Name:     "calculated_lifecycle",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("CalculatedLifecycle"),
					},
					{
						Name:        "completion_date",
						Description: "The date and time a job to restore a recovery point is completed, in Unix format and Coordinated Universal Time (UTC).",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "created_by",
						Description: "Contains identifying information about the creation of a recovery point.",
						Type:        schema.TypeJSON,
						Resolver:    schema.PathResolver("CreatedBy"),
					},
					{
						Name:        "creation_date",
						Description: "The date and time a recovery point is created, in Unix format and Coordinated Universal Time (UTC).",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:        "encryption_key_arn",
						Description: "The server-side encryption key that is used to protect your backups.",
						Type:        schema.TypeString,
					},
					{
						Name:        "iam_role_arn",
						Description: "Specifies the IAM role ARN used to create the target recovery point.",
						Type:        schema.TypeString,
					},
					{
						Name:        "is_encrypted",
						Description: "Describes if the recovery point is encrypted.",
						Type:        schema.TypeBool,
					},
					{
						Name:        "last_restore_time",
						Description: "The date and time a recovery point was last restored, in Unix format and Coordinated Universal Time (UTC).",
						Type:        schema.TypeTimestamp,
					},
					{
						Name:     "lifecycle",
						Type:     schema.TypeJSON,
						Resolver: schema.PathResolver("Lifecycle"),
					},
					{
						Name:        "arn",
						Description: "An Amazon Resource Name (ARN) that uniquely identifies a recovery point.",
						Type:        schema.TypeString,
						Resolver:    schema.PathResolver("RecoveryPointArn"),
					},
					{
						Name:        "resource_arn",
						Description: "An ARN that uniquely identifies a resource (saved as a recovery point).",
						Type:        schema.TypeString,
					},
					{
						Name:        "resource_type",
						Description: "The type of Amazon Web Services resource saved as a recovery point.",
						Type:        schema.TypeString,
					},
					{
						Name:        "source_backup_vault_arn",
						Description: "The backup vault where the recovery point was originally copied from.",
						Type:        schema.TypeString,
					},
					{
						Name:        "status",
						Description: "A status code specifying the state of the recovery point.",
						Type:        schema.TypeString,
					},
					{
						Name:        "status_message",
						Description: "A message explaining the reason of the recovery point deletion failure.",
						Type:        schema.TypeString,
					},
					{
						Name:        "tags",
						Description: "Resource tags",
						Type:        schema.TypeJSON,
						Resolver:    resolveRecoveryPointTags,
					},
				},
			},
		},
	}
}

// ====================================================================================================================
//                                               Table Resolver Functions
// ====================================================================================================================

func fetchBackupVaults(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	params := backup.ListBackupVaultsInput{MaxResults: aws.Int32(1000)} // maximum value from https://docs.aws.amazon.com/aws-backup/latest/devguide/API_ListBackupVaults.html
	for {
		result, err := svc.ListBackupVaults(ctx, &params)
		if err != nil {
			return err
		}
		res <- result.BackupVaultList
		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
	}
	return nil
}

func resolveVaultTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	vault := resource.Item.(types.BackupVaultListMember)
	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	params := backup.ListTagsInput{ResourceArn: vault.BackupVaultArn}
	tags := make(map[string]string)
	for {
		result, err := svc.ListTags(ctx, &params, func(o *backup.Options) {
			o.Region = cl.Region
		})
		if result == nil {
			break
		}
		if err != nil {
			return err
		}
		for k, v := range result.Tags {
			tags[k] = v
		}
		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
	}
	return resource.Set(c.Name, tags)
}

func resolveVaultAccessPolicy(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	vault := resource.Item.(types.BackupVaultListMember)
	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	result, err := svc.GetBackupVaultAccessPolicy(
		ctx,
		&backup.GetBackupVaultAccessPolicyInput{BackupVaultName: vault.BackupVaultName},
		func(o *backup.Options) {
			o.Region = cl.Region
		},
	)
	if err != nil {
		if cl.IsNotFoundError(err) {
			return nil
		}
		return err
	}
	return resource.Set(c.Name, result.Policy)
}

func resolveVaultNotifications(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource) error {
	vault := resource.Item.(types.BackupVaultListMember)
	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	result, err := svc.GetBackupVaultNotifications(
		ctx,
		&backup.GetBackupVaultNotificationsInput{BackupVaultName: vault.BackupVaultName},
		func(o *backup.Options) {
			o.Region = cl.Region
		},
	)
	if err != nil {
		var ae smithy.APIError
		if !errors.As(err, &ae) {
			return err
		}
		if ae.ErrorCode() == "ERROR_2106" {
			// trying to ignore "ERROR_2106: Failed reading notifications from database for Backup vault ..."
			return nil
		}
		return err
	}
	if err := resource.Set("notification_events", result.BackupVaultEvents); err != nil {
		return err
	}
	return resource.Set("notification_sns_topic_arn", result.SNSTopicArn)
}

func fetchVaultRecoveryPoints(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	vault := parent.Item.(types.BackupVaultListMember)
	params := backup.ListRecoveryPointsByBackupVaultInput{BackupVaultName: vault.BackupVaultName, MaxResults: aws.Int32(100)}
	for {
		result, err := svc.ListRecoveryPointsByBackupVault(ctx, &params)
		if err != nil {
			return err
		}
		res <- result.RecoveryPoints
		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
	}
	return nil
}

func resolveRecoveryPointTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	rp := resource.Item.(types.RecoveryPointByBackupVault)
	if rp.ResourceArn == nil || rp.RecoveryPointArn == nil {
		return nil
	}
	resourceARN, err := arn.Parse(*rp.ResourceArn)
	if err != nil {
		return err
	}

	// decide if the backed up resource supports tags
	switch client.AWSService(resourceARN.Service) {
	case client.S3Service, client.EFSService, client.DynamoDBService:
		// these services are ok
	case client.EC2Service:
		if !strings.HasPrefix(resourceARN.Resource, "instance/") {
			return nil
		}
	default:
		// full backup management not supported, so no tags
		return nil
	}

	cl := meta.(*client.Client)
	svc := cl.Services().Backup
	params := backup.ListTagsInput{ResourceArn: rp.RecoveryPointArn}
	tags := make(map[string]string)
	for {
		result, err := svc.ListTags(ctx, &params, func(o *backup.Options) {
			o.Region = cl.Region
		})
		if err != nil {
			if client.IsAWSError(err, "ERROR_2603") {
				// ignoring "ERROR_2603: Cannot find recovery point."
				return nil
			}
			if resourceARN.Service == string(client.DynamoDBService) && client.IsAWSError(err, "ERROR_3930") {
				// advanced backup features are not enabled for dynamodb
				return nil
			}
			return err
		}

		if result == nil {
			break
		}

		for k, v := range result.Tags {
			tags[k] = v
		}

		if aws.ToString(result.NextToken) == "" {
			break
		}
		params.NextToken = result.NextToken
	}
	return resource.Set(c.Name, tags)
}

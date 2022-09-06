// Code generated by codegen using template tables.go.tpl; DO NOT EDIT.

package plugin

import (
	"github.com/cloudquery/plugin-sdk/schema"

	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/accessanalyzer"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/acm"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/apigatewayv2"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/applicationautoscaling"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/appsync"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/athena"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/autoscaling"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/backup"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/cloudformation"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/cloudfront"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/cloudtrail"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/cloudwatch"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/cloudwatchlogs"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/codebuild"
	"github.com/cloudquery/cloudquery/plugins/source/aws/codegen/codepipeline"
)

func PluginAutoGeneratedTables() []*schema.Table {
	return []*schema.Table{
		accessanalyzer.AccessAnalyzerAccessanalyzers(),
		accessanalyzer.AccessAnalyzerAccessanalyzersFindings(),
		accessanalyzer.AccessAnalyzerAccessanalyzersArchiveRules(),
		acm.ACMCertificates(),
		apigatewayv2.Apigatewayv2Apis(),
		apigatewayv2.Apigatewayv2ApiAuthorizers(),
		apigatewayv2.Apigatewayv2ApiDeployments(),
		apigatewayv2.Apigatewayv2ApiIntegrations(),
		apigatewayv2.Apigatewayv2ApiIntegrationResponses(),
		apigatewayv2.Apigatewayv2ApiModels(),
		apigatewayv2.Apigatewayv2ApiRoutes(),
		apigatewayv2.Apigatewayv2ApiRouteResponses(),
		apigatewayv2.Apigatewayv2ApiStages(),
		apigatewayv2.Apigatewayv2DomainNames(),
		apigatewayv2.Apigatewayv2DomainNameApiMappings(),
		apigatewayv2.Apigatewayv2VpcLinks(),
		applicationautoscaling.ApplicationAutoscalingPolicies(),
		appsync.AppSyncGraphqlApis(),
		athena.AthenaDataCatalogs(),
		athena.AthenaDataCatalogDatabases(),
		athena.AthenaDataCatalogDatabaseTables(),
		athena.AthenaWorkGroups(),
		athena.AthenaWorkGroupPreparedStatements(),
		athena.AthenaWorkGroupQueryExecutions(),
		athena.AthenaWorkGroupNamedQueries(),
		autoscaling.AutoscalingLaunchConfigurations(),
		autoscaling.AutoscalingScheduledActions(),
		autoscaling.AutoscalingGroups(),
		autoscaling.AutoscalingGroupsScalingPolicies(),
		autoscaling.AutoscalingGroupsLifecycleHooks(),
		backup.BackupGlobalSettings(),
		backup.BackupRegionSettings(),
		backup.BackupVaults(),
		backup.BackupVaultsRecoveryPoints(),
		backup.BackupBackupPlans(),
		backup.BackupBackupPlansBackupSelections(),
		cloudformation.CloudformationStacks(),
		cloudformation.CloudformationStackResources(),
		cloudfront.CloudfrontCachePolicies(),
		cloudfront.CloudfrontDistributions(),
		cloudtrail.CloudtrailTrails(),
		cloudtrail.CloudtrailTrailEventSelectors(),
		cloudwatch.CloudwatchAlarms(),
		cloudwatchlogs.CloudwatchLogsLogGroups(),
		cloudwatchlogs.CloudwatchLogsMetricFilters(),
		codebuild.CodebuildProjects(),
		codepipeline.CodePipelinePipelines(),
		codepipeline.CodePipelineWebhooks(),
	}
}

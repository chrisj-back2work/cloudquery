package iam

import (
	"context"
	"net/url"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/iam"
	"github.com/aws/aws-sdk-go-v2/service/iam/types"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchIamPolicies(ctx context.Context, meta schema.ClientMeta, parent *schema.Resource, res chan<- interface{}) error {
	config := iam.GetAccountAuthorizationDetailsInput{
		Filter: []types.EntityType{
			types.EntityTypeAWSManagedPolicy, types.EntityTypeLocalManagedPolicy,
		},
	}
	svc := meta.(*client.Client).Services().IAM
	for {
		response, err := svc.GetAccountAuthorizationDetails(ctx, &config)
		if err != nil {
			return err
		}
		res <- response.Policies
		if aws.ToString(response.Marker) == "" {
			break
		}
		config.Marker = response.Marker
	}
	return nil
}

func resolveIamPolicyTags(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ManagedPolicyDetail)
	cl := meta.(*client.Client)
	svc := cl.Services().IAM
	response, err := svc.ListPolicyTags(ctx, &iam.ListPolicyTagsInput{PolicyArn: r.Arn})
	if err != nil {
		if cl.IsNotFoundError(err) {
			meta.Logger().Debug().Err(err).Msg("ListPolicyTags: Policy does not exist")
			return nil
		}
		return err
	}
	return resource.Set("tags", client.TagsToMap(response.Tags))
}

func resolveIamPolicyVersionList(ctx context.Context, meta schema.ClientMeta, resource *schema.Resource, c schema.Column) error {
	r := resource.Item.(types.ManagedPolicyDetail)
	for i := range r.PolicyVersionList {
		if v, err := url.PathUnescape(aws.ToString(r.PolicyVersionList[i].Document)); err == nil {
			r.PolicyVersionList[i].Document = &v
		} else {
			meta.Logger().Warn().Err(err).Str("policy_id", aws.ToString(r.PolicyId)).Msg("Failed to unescape policy document, leaving as-is")
		}
	}
	return resource.Set(c.Name, r.PolicyVersionList)
}

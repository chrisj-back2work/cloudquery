package frauddetector

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/frauddetector"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/plugin-sdk/schema"
)

func fetchFrauddetectorEventTypes(ctx context.Context, meta schema.ClientMeta, _ *schema.Resource, res chan<- interface{}) error {
	c := meta.(*client.Client)
	svc := c.Services().FraudDetector

	input := new(frauddetector.GetEventTypesInput)
	for {
		output, err := svc.GetEventTypes(ctx, input)
		if err != nil {
			return err
		}

		res <- output.EventTypes

		if aws.ToString(output.NextToken) == "" {
			break
		}
		input.NextToken = output.NextToken
	}

	return nil
}

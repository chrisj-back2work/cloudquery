// Code generated by codegen using template resource_get_mock_test.go.tpl; DO NOT EDIT.

package cloudwatchlogs

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"

	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs"
	"github.com/aws/aws-sdk-go-v2/service/cloudwatchlogs/types"
)

func buildCloudwatchLogsMetricFilters(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockCloudwatchLogsClient(ctrl)

	item := types.MetricFilter{}

	err := faker.FakeData(&item)
	if err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().DescribeMetricFilters(gomock.Any(), gomock.Any(), gomock.Any()).Return(

		&cloudwatchlogs.DescribeMetricFiltersOutput{
			MetricFilters: []types.MetricFilter{item},
		}, nil)

	return client.Services{
		CloudwatchLogs: mock,
	}
}

func TestCloudwatchLogsMetricFilters(t *testing.T) {
	client.MockTestHelper(t, CloudwatchLogsMetricFilters(), buildCloudwatchLogsMetricFilters, client.TestOptions{})
}
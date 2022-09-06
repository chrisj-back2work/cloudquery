// Code generated by codegen using template resource_get_mock_test.go.tpl; DO NOT EDIT.

package accessanalyzer

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/aws/client"
	"github.com/cloudquery/cloudquery/plugins/source/aws/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/golang/mock/gomock"

	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer"
	"github.com/aws/aws-sdk-go-v2/service/accessanalyzer/types"
)

func buildAccessAnalyzerAccessanalyzers(t *testing.T, ctrl *gomock.Controller) client.Services {
	mock := mocks.NewMockAccessAnalyzerClient(ctrl)

	item := types.AnalyzerSummary{}

	err := faker.FakeData(&item)
	if err != nil {
		t.Fatal(err)
	}
	mock.EXPECT().ListAnalyzers(gomock.Any(), gomock.Any(), gomock.Any()).Return(

		&accessanalyzer.ListAnalyzersOutput{
			Analyzers: []types.AnalyzerSummary{item},
		}, nil)

	return client.Services{
		AccessAnalyzer: mock,
	}
}

func TestAccessAnalyzerAccessanalyzers(t *testing.T) {
	client.MockTestHelper(t, AccessAnalyzerAccessanalyzers(), buildAccessAnalyzerAccessanalyzers, client.TestOptions{})
}
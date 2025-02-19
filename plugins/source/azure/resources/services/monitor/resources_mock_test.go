// Auto generated code - DO NOT EDIT.

package monitor

import (
	"context"
	"testing"

	"github.com/Azure/azure-sdk-for-go/profiles/latest/resources/mgmt/resources"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services"
	"github.com/cloudquery/cloudquery/plugins/source/azure/client/services/mocks"
	"github.com/cloudquery/plugin-sdk/faker"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestMonitorResources(t *testing.T) {
	client.MockTestHelper(t, Resources(), createResourcesMock)
}

func createResourcesMock(t *testing.T, ctrl *gomock.Controller) services.Services {
	mockClient := mocks.NewMockMonitorResourcesClient(ctrl)
	s := services.Services{
		Monitor: services.MonitorClient{
			Resources:          mockClient,
			DiagnosticSettings: createDiagnosticSettingsMock(t, ctrl).Monitor.DiagnosticSettings,
		},
	}

	data := resources.GenericResourceExpanded{}
	require.Nil(t, faker.FakeObject(&data))

	// Ensure name and ID are consistent so we can reference it in other mock
	name := "test"
	data.Name = &name

	// Use correct Azure ID format
	id := "/subscriptions/test/resourceGroups/test/providers/test/test/test"
	data.ID = &id

	result := resources.NewListResultPage(resources.ListResult{Value: &[]resources.GenericResourceExpanded{data}}, func(ctx context.Context, result resources.ListResult) (resources.ListResult, error) {
		return resources.ListResult{}, nil
	})

	mockClient.EXPECT().List(gomock.Any(), "", "", nil).Return(result, nil)
	return s
}

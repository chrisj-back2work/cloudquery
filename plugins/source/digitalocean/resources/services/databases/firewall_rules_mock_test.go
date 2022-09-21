package databases

import (
	"testing"

	"github.com/cloudquery/cloudquery/plugins/source/digitalocean/client/mocks"
	"github.com/cloudquery/faker/v3"
	"github.com/digitalocean/godo"
	"github.com/golang/mock/gomock"
)

func createFirewallRules(t *testing.T, m *mocks.MockDatabasesService) {
	var data []godo.DatabaseFirewallRule
	if err := faker.FakeData(&data); err != nil {
		t.Fatal(err)
	}

	m.EXPECT().GetFirewallRules(gomock.Any(), gomock.Any()).Return(data, &godo.Response{}, nil)
}
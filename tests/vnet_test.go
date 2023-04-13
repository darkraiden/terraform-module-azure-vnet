package tests

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/require"
)

type config struct {
	subscriptionID string
	tenantID       string
}

func getTestConfig(t *testing.T) *config {
	t.Helper()
	config := config{
		subscriptionID: os.Getenv("TEST_AZURE_SUBSCRIPTION_ID"),
		tenantID:       os.Getenv("TEST_AZURE_TENANT_ID"),
	}
	if config.subscriptionID == "" || config.tenantID == "" {
		t.Fatal("AZURE_SUBSCRIPTION_ID and AZURE_TENANT_ID must be set")
	}

	return &config
}

var fixtureFolder = "./fixture"

func TestVirtualNetwork(t *testing.T) {
	config := getTestConfig(t)

	t.Run(`JustVNet`, func(t *testing.T) {
		terraformOptions := &terraform.Options{
			TerraformDir: fixtureFolder,
			Vars: map[string]interface{}{
				"subscription_id": config.subscriptionID,
				"tenant_id":       config.tenantID,
			},
			VarFiles: []string{"just_vnet.tfvars"},
		}
		defer terraform.Destroy(t, terraformOptions)

		terraform.InitAndApply(t, terraformOptions)

		rg_name := terraform.Output(t, terraformOptions, "resource_group_name")
		vnet_name := terraform.Output(t, terraformOptions, "vnet_name")
		vnet_location := terraform.Output(t, terraformOptions, "vnet_location")
		rg_location := terraform.Output(t, terraformOptions, "resource_group_location")

		require.Equal(t, vnet_location, rg_location)
		require.Equal(t, vnet_name, rg_name)
	})

	t.Run(`WithListOfSubnets`, func(t *testing.T) {
		terraformOptions := &terraform.Options{
			TerraformDir: fixtureFolder,
			Vars: map[string]interface{}{
				"subscription_id": config.subscriptionID,
				"tenant_id":       config.tenantID,
			},
			VarFiles: []string{"with_subnets.tfvars"},
		}
		defer terraform.Destroy(t, terraformOptions)

		terraform.InitAndApply(t, terraformOptions)

		subnet_ids := terraform.OutputList(t, terraformOptions, "subnet_ids")
		require.GreaterOrEqual(t, len(subnet_ids), 2)
	})

	t.Run(`WithVNetsLocationDifferentFromRGs`, func(t *testing.T) {
		terraformOptions := &terraform.Options{
			TerraformDir: fixtureFolder,
			Vars: map[string]interface{}{
				"subscription_id": config.subscriptionID,
				"tenant_id":       config.tenantID,
				// we could simply create a new tfvars file
				// but given the small change, I think it's fine passing it as a cli variable
				"vnet_location": "westeurope",
			},
			VarFiles: []string{"just_vnet.tfvars"},
		}
		defer terraform.Destroy(t, terraformOptions)

		terraform.InitAndApply(t, terraformOptions)

		subnet_ids := terraform.OutputList(t, terraformOptions, "subnet_ids")
		rg_location := terraform.Output(t, terraformOptions, "resource_group_location")
		vnet_location := terraform.Output(t, terraformOptions, "vnet_location")

		// we could also just check whether the two locations are not equal
		// but this way we are sure that the vnet is in the right location
		require.Equal(t, vnet_location, "westeurope")
		require.Equal(t, rg_location, "northeurope")
		require.Equal(t, len(subnet_ids), 0)
	})
}

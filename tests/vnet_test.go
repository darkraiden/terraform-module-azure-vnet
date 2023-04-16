package tests

import (
	"os"
	"testing"

	"github.com/gruntwork-io/terratest/modules/terraform"
	test_structure "github.com/gruntwork-io/terratest/modules/test-structure"
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

const fixtureFolder = "tests/fixture"

func TestVirtualNetwork(t *testing.T) {
	config := getTestConfig(t)

	t.Run(`JustVNet`, func(t *testing.T) {
		t.Parallel()

		tmpFolder := test_structure.CopyTerraformFolderToTemp(t, "../", fixtureFolder)
		terraformOptions := &terraform.Options{
			TerraformDir: tmpFolder,
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
		ddos_protection_plan_id := terraform.Output(t, terraformOptions, "ddos_protection_plan_id")

		require.Equal(t, vnet_location, rg_location)
		require.Equal(t, vnet_name, rg_name)
		require.Empty(t, ddos_protection_plan_id)
	})

	t.Run(`WithListOfSubnets`, func(t *testing.T) {
		t.Parallel()

		tmpFolder := test_structure.CopyTerraformFolderToTemp(t, "../", fixtureFolder)
		terraformOptions := &terraform.Options{
			TerraformDir: tmpFolder,
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
		t.Parallel()

		tmpFolder := test_structure.CopyTerraformFolderToTemp(t, "../", fixtureFolder)
		terraformOptions := &terraform.Options{
			TerraformDir: tmpFolder,
			Vars: map[string]interface{}{
				"subscription_id": config.subscriptionID,
				"tenant_id":       config.tenantID,
				"vnet_location":   "westeurope",
			},
			VarFiles: []string{"just_vnet.tfvars"},
		}
		defer terraform.Destroy(t, terraformOptions)

		terraform.InitAndApply(t, terraformOptions)

		subnet_ids := terraform.OutputList(t, terraformOptions, "subnet_ids")
		rg_location := terraform.Output(t, terraformOptions, "resource_group_location")
		vnet_location := terraform.Output(t, terraformOptions, "vnet_location")

		require.Equal(t, vnet_location, "westeurope")
		require.Equal(t, rg_location, "northeurope")
		require.Equal(t, len(subnet_ids), 0)
	})

	t.Run(`WithDDoSProtectionEnabled`, func(t *testing.T) {
		t.Parallel()

		tmpFolder := test_structure.CopyTerraformFolderToTemp(t, "../", fixtureFolder)
		terraformOptions := &terraform.Options{
			TerraformDir: tmpFolder,
			Vars: map[string]interface{}{
				"subscription_id":             config.subscriptionID,
				"tenant_id":                   config.tenantID,
				"enable_ddos_protection_plan": true,
			},
			VarFiles: []string{"just_vnet.tfvars"},
		}
		defer terraform.Destroy(t, terraformOptions)

		terraform.InitAndApply(t, terraformOptions)

		ddosProtectionPlanID := terraform.Output(t, terraformOptions, "ddos_protection_plan_id")

		require.NotEmpty(t, ddosProtectionPlanID)
	})
}

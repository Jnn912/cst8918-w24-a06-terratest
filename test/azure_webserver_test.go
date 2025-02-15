package test

import (
	"fmt"
	"testing"

	"github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

// You normally want to run this under a separate "Testing" subscription
// For lab purposes you will use your assigned subscription under the Cloud Dev/Ops program tenant
var subscriptionID string = "a55a4ed4-e1da-410b-a33d-1840ef02f227"

func TestAzureLinuxVMCreation(t *testing.T) {
	// Define a valid labelPrefix
	labelPrefix := "jin00098" // Replace with your actual college ID or a valid prefix

	// Print the labelPrefix to verify its value
	fmt.Println("labelPrefix:", labelPrefix)

	terraformOptions := &terraform.Options{
		// The path to where our Terraform code is located
		TerraformDir: "../",
		// Override the default terraform variables
		Vars: map[string]interface{}{
			"labelPrefix": labelPrefix, // Use the valid labelPrefix
		},
	}

	// Ensure resources are destroyed after the test
	defer terraform.Destroy(t, terraformOptions)

	// Run `terraform init` and `terraform apply`. Fail the test if there are any errors.
	terraform.InitAndApply(t, terraformOptions)

	// Run `terraform output` to get the value of output variable
	vmName := terraform.Output(t, terraformOptions, "vm_name")
	resourceGroupName := terraform.Output(t, terraformOptions, "resource_group_name")

	// Print the VM name and resource group name for debugging
	fmt.Println("VM Name:", vmName)
	fmt.Println("Resource Group Name:", resourceGroupName)

	// Confirm VM exists
	assert.True(t, azure.VirtualMachineExists(t, vmName, resourceGroupName, subscriptionID))
}

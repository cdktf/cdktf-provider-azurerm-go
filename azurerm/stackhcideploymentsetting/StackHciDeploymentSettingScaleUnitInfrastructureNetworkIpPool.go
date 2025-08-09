// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stackhcideploymentsetting


type StackHciDeploymentSettingScaleUnitInfrastructureNetworkIpPool struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/stack_hci_deployment_setting#ending_address StackHciDeploymentSetting#ending_address}.
	EndingAddress *string `field:"required" json:"endingAddress" yaml:"endingAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/stack_hci_deployment_setting#starting_address StackHciDeploymentSetting#starting_address}.
	StartingAddress *string `field:"required" json:"startingAddress" yaml:"startingAddress"`
}


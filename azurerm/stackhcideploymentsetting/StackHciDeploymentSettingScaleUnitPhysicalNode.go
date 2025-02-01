// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stackhcideploymentsetting


type StackHciDeploymentSettingScaleUnitPhysicalNode struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/stack_hci_deployment_setting#ipv4_address StackHciDeploymentSetting#ipv4_address}.
	Ipv4Address *string `field:"required" json:"ipv4Address" yaml:"ipv4Address"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/stack_hci_deployment_setting#name StackHciDeploymentSetting#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}


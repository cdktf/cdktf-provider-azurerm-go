// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stackhcideploymentsetting


type StackHciDeploymentSettingScaleUnitHostNetworkStorageNetwork struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/stack_hci_deployment_setting#name StackHciDeploymentSetting#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/stack_hci_deployment_setting#network_adapter_name StackHciDeploymentSetting#network_adapter_name}.
	NetworkAdapterName *string `field:"required" json:"networkAdapterName" yaml:"networkAdapterName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/stack_hci_deployment_setting#vlan_id StackHciDeploymentSetting#vlan_id}.
	VlanId *string `field:"required" json:"vlanId" yaml:"vlanId"`
}


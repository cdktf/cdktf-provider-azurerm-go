// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stackhcideploymentsetting


type StackHciDeploymentSettingScaleUnitHostNetwork struct {
	// intent block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/stack_hci_deployment_setting#intent StackHciDeploymentSetting#intent}
	Intent interface{} `field:"required" json:"intent" yaml:"intent"`
	// storage_network block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/stack_hci_deployment_setting#storage_network StackHciDeploymentSetting#storage_network}
	StorageNetwork interface{} `field:"required" json:"storageNetwork" yaml:"storageNetwork"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/stack_hci_deployment_setting#storage_auto_ip_enabled StackHciDeploymentSetting#storage_auto_ip_enabled}.
	StorageAutoIpEnabled interface{} `field:"optional" json:"storageAutoIpEnabled" yaml:"storageAutoIpEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/stack_hci_deployment_setting#storage_connectivity_switchless_enabled StackHciDeploymentSetting#storage_connectivity_switchless_enabled}.
	StorageConnectivitySwitchlessEnabled interface{} `field:"optional" json:"storageConnectivitySwitchlessEnabled" yaml:"storageConnectivitySwitchlessEnabled"`
}


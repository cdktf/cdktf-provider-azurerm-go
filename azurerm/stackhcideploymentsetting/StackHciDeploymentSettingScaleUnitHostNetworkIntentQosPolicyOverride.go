// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stackhcideploymentsetting


type StackHciDeploymentSettingScaleUnitHostNetworkIntentQosPolicyOverride struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/stack_hci_deployment_setting#bandwidth_percentage_smb StackHciDeploymentSetting#bandwidth_percentage_smb}.
	BandwidthPercentageSmb *string `field:"optional" json:"bandwidthPercentageSmb" yaml:"bandwidthPercentageSmb"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/stack_hci_deployment_setting#priority_value8021_action_cluster StackHciDeploymentSetting#priority_value8021_action_cluster}.
	PriorityValue8021ActionCluster *string `field:"optional" json:"priorityValue8021ActionCluster" yaml:"priorityValue8021ActionCluster"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/stack_hci_deployment_setting#priority_value8021_action_smb StackHciDeploymentSetting#priority_value8021_action_smb}.
	PriorityValue8021ActionSmb *string `field:"optional" json:"priorityValue8021ActionSmb" yaml:"priorityValue8021ActionSmb"`
}


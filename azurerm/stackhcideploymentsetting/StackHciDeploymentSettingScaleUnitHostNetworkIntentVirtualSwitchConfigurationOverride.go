// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stackhcideploymentsetting


type StackHciDeploymentSettingScaleUnitHostNetworkIntentVirtualSwitchConfigurationOverride struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/stack_hci_deployment_setting#enable_iov StackHciDeploymentSetting#enable_iov}.
	EnableIov *string `field:"optional" json:"enableIov" yaml:"enableIov"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.17.0/docs/resources/stack_hci_deployment_setting#load_balancing_algorithm StackHciDeploymentSetting#load_balancing_algorithm}.
	LoadBalancingAlgorithm *string `field:"optional" json:"loadBalancingAlgorithm" yaml:"loadBalancingAlgorithm"`
}


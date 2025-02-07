// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stackhcideploymentsetting


type StackHciDeploymentSettingTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/stack_hci_deployment_setting#create StackHciDeploymentSetting#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/stack_hci_deployment_setting#delete StackHciDeploymentSetting#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/stack_hci_deployment_setting#read StackHciDeploymentSetting#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


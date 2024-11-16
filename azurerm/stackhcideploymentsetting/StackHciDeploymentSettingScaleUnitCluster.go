// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package stackhcideploymentsetting


type StackHciDeploymentSettingScaleUnitCluster struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/stack_hci_deployment_setting#azure_service_endpoint StackHciDeploymentSetting#azure_service_endpoint}.
	AzureServiceEndpoint *string `field:"required" json:"azureServiceEndpoint" yaml:"azureServiceEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/stack_hci_deployment_setting#cloud_account_name StackHciDeploymentSetting#cloud_account_name}.
	CloudAccountName *string `field:"required" json:"cloudAccountName" yaml:"cloudAccountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/stack_hci_deployment_setting#name StackHciDeploymentSetting#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/stack_hci_deployment_setting#witness_path StackHciDeploymentSetting#witness_path}.
	WitnessPath *string `field:"required" json:"witnessPath" yaml:"witnessPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/stack_hci_deployment_setting#witness_type StackHciDeploymentSetting#witness_type}.
	WitnessType *string `field:"required" json:"witnessType" yaml:"witnessType"`
}


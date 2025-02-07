// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appservicesourcecontrolslot


type AppServiceSourceControlSlotGithubActionConfiguration struct {
	// code_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/app_service_source_control_slot#code_configuration AppServiceSourceControlSlot#code_configuration}
	CodeConfiguration *AppServiceSourceControlSlotGithubActionConfigurationCodeConfiguration `field:"optional" json:"codeConfiguration" yaml:"codeConfiguration"`
	// container_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/app_service_source_control_slot#container_configuration AppServiceSourceControlSlot#container_configuration}
	ContainerConfiguration *AppServiceSourceControlSlotGithubActionConfigurationContainerConfiguration `field:"optional" json:"containerConfiguration" yaml:"containerConfiguration"`
	// Should the service generate the GitHub Action Workflow file. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/app_service_source_control_slot#generate_workflow_file AppServiceSourceControlSlot#generate_workflow_file}
	GenerateWorkflowFile interface{} `field:"optional" json:"generateWorkflowFile" yaml:"generateWorkflowFile"`
}


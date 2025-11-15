// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package appservicesourcecontrol


type AppServiceSourceControlGithubActionConfiguration struct {
	// code_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/app_service_source_control#code_configuration AppServiceSourceControlA#code_configuration}
	CodeConfiguration *AppServiceSourceControlGithubActionConfigurationCodeConfiguration `field:"optional" json:"codeConfiguration" yaml:"codeConfiguration"`
	// container_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/app_service_source_control#container_configuration AppServiceSourceControlA#container_configuration}
	ContainerConfiguration *AppServiceSourceControlGithubActionConfigurationContainerConfiguration `field:"optional" json:"containerConfiguration" yaml:"containerConfiguration"`
	// Should the service generate the GitHub Action Workflow file. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/app_service_source_control#generate_workflow_file AppServiceSourceControlA#generate_workflow_file}
	GenerateWorkflowFile interface{} `field:"optional" json:"generateWorkflowFile" yaml:"generateWorkflowFile"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerregistrytask


type ContainerRegistryTaskEncodedStep struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/container_registry_task#task_content ContainerRegistryTask#task_content}.
	TaskContent *string `field:"required" json:"taskContent" yaml:"taskContent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/container_registry_task#context_access_token ContainerRegistryTask#context_access_token}.
	ContextAccessToken *string `field:"optional" json:"contextAccessToken" yaml:"contextAccessToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/container_registry_task#context_path ContainerRegistryTask#context_path}.
	ContextPath *string `field:"optional" json:"contextPath" yaml:"contextPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/container_registry_task#secret_values ContainerRegistryTask#secret_values}.
	SecretValues *map[string]*string `field:"optional" json:"secretValues" yaml:"secretValues"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/container_registry_task#value_content ContainerRegistryTask#value_content}.
	ValueContent *string `field:"optional" json:"valueContent" yaml:"valueContent"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/container_registry_task#values ContainerRegistryTask#values}.
	Values *map[string]*string `field:"optional" json:"values" yaml:"values"`
}


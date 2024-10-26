// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerregistrytask


type ContainerRegistryTaskBaseImageTrigger struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/container_registry_task#name ContainerRegistryTask#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/container_registry_task#type ContainerRegistryTask#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/container_registry_task#enabled ContainerRegistryTask#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/container_registry_task#update_trigger_endpoint ContainerRegistryTask#update_trigger_endpoint}.
	UpdateTriggerEndpoint *string `field:"optional" json:"updateTriggerEndpoint" yaml:"updateTriggerEndpoint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/container_registry_task#update_trigger_payload_type ContainerRegistryTask#update_trigger_payload_type}.
	UpdateTriggerPayloadType *string `field:"optional" json:"updateTriggerPayloadType" yaml:"updateTriggerPayloadType"`
}


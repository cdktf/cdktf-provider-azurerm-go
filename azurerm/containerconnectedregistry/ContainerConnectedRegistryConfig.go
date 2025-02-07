// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerconnectedregistry

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerConnectedRegistryConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/container_connected_registry#container_registry_id ContainerConnectedRegistry#container_registry_id}.
	ContainerRegistryId *string `field:"required" json:"containerRegistryId" yaml:"containerRegistryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/container_connected_registry#name ContainerConnectedRegistry#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/container_connected_registry#sync_token_id ContainerConnectedRegistry#sync_token_id}.
	SyncTokenId *string `field:"required" json:"syncTokenId" yaml:"syncTokenId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/container_connected_registry#audit_log_enabled ContainerConnectedRegistry#audit_log_enabled}.
	AuditLogEnabled interface{} `field:"optional" json:"auditLogEnabled" yaml:"auditLogEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/container_connected_registry#client_token_ids ContainerConnectedRegistry#client_token_ids}.
	ClientTokenIds *[]*string `field:"optional" json:"clientTokenIds" yaml:"clientTokenIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/container_connected_registry#id ContainerConnectedRegistry#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/container_connected_registry#log_level ContainerConnectedRegistry#log_level}.
	LogLevel *string `field:"optional" json:"logLevel" yaml:"logLevel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/container_connected_registry#mode ContainerConnectedRegistry#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// notification block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/container_connected_registry#notification ContainerConnectedRegistry#notification}
	Notification interface{} `field:"optional" json:"notification" yaml:"notification"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/container_connected_registry#parent_registry_id ContainerConnectedRegistry#parent_registry_id}.
	ParentRegistryId *string `field:"optional" json:"parentRegistryId" yaml:"parentRegistryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/container_connected_registry#sync_message_ttl ContainerConnectedRegistry#sync_message_ttl}.
	SyncMessageTtl *string `field:"optional" json:"syncMessageTtl" yaml:"syncMessageTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/container_connected_registry#sync_schedule ContainerConnectedRegistry#sync_schedule}.
	SyncSchedule *string `field:"optional" json:"syncSchedule" yaml:"syncSchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/container_connected_registry#sync_window ContainerConnectedRegistry#sync_window}.
	SyncWindow *string `field:"optional" json:"syncWindow" yaml:"syncWindow"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/container_connected_registry#timeouts ContainerConnectedRegistry#timeouts}
	Timeouts *ContainerConnectedRegistryTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package synapseworkspacekey

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SynapseWorkspaceKeyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/synapse_workspace_key#active SynapseWorkspaceKey#active}.
	Active interface{} `field:"required" json:"active" yaml:"active"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/synapse_workspace_key#customer_managed_key_name SynapseWorkspaceKey#customer_managed_key_name}.
	CustomerManagedKeyName *string `field:"required" json:"customerManagedKeyName" yaml:"customerManagedKeyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/synapse_workspace_key#synapse_workspace_id SynapseWorkspaceKey#synapse_workspace_id}.
	SynapseWorkspaceId *string `field:"required" json:"synapseWorkspaceId" yaml:"synapseWorkspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/synapse_workspace_key#customer_managed_key_versionless_id SynapseWorkspaceKey#customer_managed_key_versionless_id}.
	CustomerManagedKeyVersionlessId *string `field:"optional" json:"customerManagedKeyVersionlessId" yaml:"customerManagedKeyVersionlessId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/synapse_workspace_key#id SynapseWorkspaceKey#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.95.0/docs/resources/synapse_workspace_key#timeouts SynapseWorkspaceKey#timeouts}
	Timeouts *SynapseWorkspaceKeyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


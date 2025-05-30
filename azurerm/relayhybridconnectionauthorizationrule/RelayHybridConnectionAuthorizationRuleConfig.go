// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package relayhybridconnectionauthorizationrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RelayHybridConnectionAuthorizationRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/relay_hybrid_connection_authorization_rule#hybrid_connection_name RelayHybridConnectionAuthorizationRule#hybrid_connection_name}.
	HybridConnectionName *string `field:"required" json:"hybridConnectionName" yaml:"hybridConnectionName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/relay_hybrid_connection_authorization_rule#name RelayHybridConnectionAuthorizationRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/relay_hybrid_connection_authorization_rule#namespace_name RelayHybridConnectionAuthorizationRule#namespace_name}.
	NamespaceName *string `field:"required" json:"namespaceName" yaml:"namespaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/relay_hybrid_connection_authorization_rule#resource_group_name RelayHybridConnectionAuthorizationRule#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/relay_hybrid_connection_authorization_rule#id RelayHybridConnectionAuthorizationRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/relay_hybrid_connection_authorization_rule#listen RelayHybridConnectionAuthorizationRule#listen}.
	Listen interface{} `field:"optional" json:"listen" yaml:"listen"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/relay_hybrid_connection_authorization_rule#manage RelayHybridConnectionAuthorizationRule#manage}.
	Manage interface{} `field:"optional" json:"manage" yaml:"manage"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/relay_hybrid_connection_authorization_rule#send RelayHybridConnectionAuthorizationRule#send}.
	Send interface{} `field:"optional" json:"send" yaml:"send"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/relay_hybrid_connection_authorization_rule#timeouts RelayHybridConnectionAuthorizationRule#timeouts}
	Timeouts *RelayHybridConnectionAuthorizationRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


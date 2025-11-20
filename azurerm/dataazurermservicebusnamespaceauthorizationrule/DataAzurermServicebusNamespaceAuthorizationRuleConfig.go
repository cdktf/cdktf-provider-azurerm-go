// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermservicebusnamespaceauthorizationrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermServicebusNamespaceAuthorizationRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/data-sources/servicebus_namespace_authorization_rule#name DataAzurermServicebusNamespaceAuthorizationRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/data-sources/servicebus_namespace_authorization_rule#id DataAzurermServicebusNamespaceAuthorizationRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/data-sources/servicebus_namespace_authorization_rule#namespace_id DataAzurermServicebusNamespaceAuthorizationRule#namespace_id}.
	NamespaceId *string `field:"optional" json:"namespaceId" yaml:"namespaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/data-sources/servicebus_namespace_authorization_rule#namespace_name DataAzurermServicebusNamespaceAuthorizationRule#namespace_name}.
	NamespaceName *string `field:"optional" json:"namespaceName" yaml:"namespaceName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/data-sources/servicebus_namespace_authorization_rule#resource_group_name DataAzurermServicebusNamespaceAuthorizationRule#resource_group_name}.
	ResourceGroupName *string `field:"optional" json:"resourceGroupName" yaml:"resourceGroupName"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/data-sources/servicebus_namespace_authorization_rule#timeouts DataAzurermServicebusNamespaceAuthorizationRule#timeouts}
	Timeouts *DataAzurermServicebusNamespaceAuthorizationRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


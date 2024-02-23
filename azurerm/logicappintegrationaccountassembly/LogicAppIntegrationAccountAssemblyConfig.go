// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logicappintegrationaccountassembly

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogicAppIntegrationAccountAssemblyConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/logic_app_integration_account_assembly#assembly_name LogicAppIntegrationAccountAssembly#assembly_name}.
	AssemblyName *string `field:"required" json:"assemblyName" yaml:"assemblyName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/logic_app_integration_account_assembly#integration_account_name LogicAppIntegrationAccountAssembly#integration_account_name}.
	IntegrationAccountName *string `field:"required" json:"integrationAccountName" yaml:"integrationAccountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/logic_app_integration_account_assembly#name LogicAppIntegrationAccountAssembly#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/logic_app_integration_account_assembly#resource_group_name LogicAppIntegrationAccountAssembly#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/logic_app_integration_account_assembly#assembly_version LogicAppIntegrationAccountAssembly#assembly_version}.
	AssemblyVersion *string `field:"optional" json:"assemblyVersion" yaml:"assemblyVersion"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/logic_app_integration_account_assembly#content LogicAppIntegrationAccountAssembly#content}.
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/logic_app_integration_account_assembly#content_link_uri LogicAppIntegrationAccountAssembly#content_link_uri}.
	ContentLinkUri *string `field:"optional" json:"contentLinkUri" yaml:"contentLinkUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/logic_app_integration_account_assembly#id LogicAppIntegrationAccountAssembly#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/logic_app_integration_account_assembly#metadata LogicAppIntegrationAccountAssembly#metadata}.
	Metadata *map[string]*string `field:"optional" json:"metadata" yaml:"metadata"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/logic_app_integration_account_assembly#timeouts LogicAppIntegrationAccountAssembly#timeouts}
	Timeouts *LogicAppIntegrationAccountAssemblyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


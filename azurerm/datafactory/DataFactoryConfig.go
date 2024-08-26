// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactory

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/data_factory#location DataFactory#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/data_factory#name DataFactory#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/data_factory#resource_group_name DataFactory#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/data_factory#customer_managed_key_id DataFactory#customer_managed_key_id}.
	CustomerManagedKeyId *string `field:"optional" json:"customerManagedKeyId" yaml:"customerManagedKeyId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/data_factory#customer_managed_key_identity_id DataFactory#customer_managed_key_identity_id}.
	CustomerManagedKeyIdentityId *string `field:"optional" json:"customerManagedKeyIdentityId" yaml:"customerManagedKeyIdentityId"`
	// github_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/data_factory#github_configuration DataFactory#github_configuration}
	GithubConfiguration *DataFactoryGithubConfiguration `field:"optional" json:"githubConfiguration" yaml:"githubConfiguration"`
	// global_parameter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/data_factory#global_parameter DataFactory#global_parameter}
	GlobalParameter interface{} `field:"optional" json:"globalParameter" yaml:"globalParameter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/data_factory#id DataFactory#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/data_factory#identity DataFactory#identity}
	Identity *DataFactoryIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/data_factory#managed_virtual_network_enabled DataFactory#managed_virtual_network_enabled}.
	ManagedVirtualNetworkEnabled interface{} `field:"optional" json:"managedVirtualNetworkEnabled" yaml:"managedVirtualNetworkEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/data_factory#public_network_enabled DataFactory#public_network_enabled}.
	PublicNetworkEnabled interface{} `field:"optional" json:"publicNetworkEnabled" yaml:"publicNetworkEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/data_factory#purview_id DataFactory#purview_id}.
	PurviewId *string `field:"optional" json:"purviewId" yaml:"purviewId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/data_factory#tags DataFactory#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/data_factory#timeouts DataFactory#timeouts}
	Timeouts *DataFactoryTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// vsts_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/data_factory#vsts_configuration DataFactory#vsts_configuration}
	VstsConfiguration *DataFactoryVstsConfiguration `field:"optional" json:"vstsConfiguration" yaml:"vstsConfiguration"`
}


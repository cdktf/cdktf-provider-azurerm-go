// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package bastionhost

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BastionHostConfig struct {
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
	// ip_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/bastion_host#ip_configuration BastionHost#ip_configuration}
	IpConfiguration *BastionHostIpConfiguration `field:"required" json:"ipConfiguration" yaml:"ipConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/bastion_host#location BastionHost#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/bastion_host#name BastionHost#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/bastion_host#resource_group_name BastionHost#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/bastion_host#copy_paste_enabled BastionHost#copy_paste_enabled}.
	CopyPasteEnabled interface{} `field:"optional" json:"copyPasteEnabled" yaml:"copyPasteEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/bastion_host#file_copy_enabled BastionHost#file_copy_enabled}.
	FileCopyEnabled interface{} `field:"optional" json:"fileCopyEnabled" yaml:"fileCopyEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/bastion_host#id BastionHost#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/bastion_host#ip_connect_enabled BastionHost#ip_connect_enabled}.
	IpConnectEnabled interface{} `field:"optional" json:"ipConnectEnabled" yaml:"ipConnectEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/bastion_host#scale_units BastionHost#scale_units}.
	ScaleUnits *float64 `field:"optional" json:"scaleUnits" yaml:"scaleUnits"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/bastion_host#shareable_link_enabled BastionHost#shareable_link_enabled}.
	ShareableLinkEnabled interface{} `field:"optional" json:"shareableLinkEnabled" yaml:"shareableLinkEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/bastion_host#sku BastionHost#sku}.
	Sku *string `field:"optional" json:"sku" yaml:"sku"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/bastion_host#tags BastionHost#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/bastion_host#timeouts BastionHost#timeouts}
	Timeouts *BastionHostTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.75.0/docs/resources/bastion_host#tunneling_enabled BastionHost#tunneling_enabled}.
	TunnelingEnabled interface{} `field:"optional" json:"tunnelingEnabled" yaml:"tunnelingEnabled"`
}


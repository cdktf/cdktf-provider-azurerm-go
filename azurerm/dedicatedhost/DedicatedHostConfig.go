// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dedicatedhost

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DedicatedHostConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/dedicated_host#dedicated_host_group_id DedicatedHost#dedicated_host_group_id}.
	DedicatedHostGroupId *string `field:"required" json:"dedicatedHostGroupId" yaml:"dedicatedHostGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/dedicated_host#location DedicatedHost#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/dedicated_host#name DedicatedHost#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/dedicated_host#platform_fault_domain DedicatedHost#platform_fault_domain}.
	PlatformFaultDomain *float64 `field:"required" json:"platformFaultDomain" yaml:"platformFaultDomain"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/dedicated_host#sku_name DedicatedHost#sku_name}.
	SkuName *string `field:"required" json:"skuName" yaml:"skuName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/dedicated_host#auto_replace_on_failure DedicatedHost#auto_replace_on_failure}.
	AutoReplaceOnFailure interface{} `field:"optional" json:"autoReplaceOnFailure" yaml:"autoReplaceOnFailure"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/dedicated_host#id DedicatedHost#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/dedicated_host#license_type DedicatedHost#license_type}.
	LicenseType *string `field:"optional" json:"licenseType" yaml:"licenseType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/dedicated_host#tags DedicatedHost#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/dedicated_host#timeouts DedicatedHost#timeouts}
	Timeouts *DedicatedHostTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


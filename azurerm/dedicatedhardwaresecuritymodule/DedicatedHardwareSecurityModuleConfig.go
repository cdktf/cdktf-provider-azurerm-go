// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dedicatedhardwaresecuritymodule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DedicatedHardwareSecurityModuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/dedicated_hardware_security_module#location DedicatedHardwareSecurityModule#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/dedicated_hardware_security_module#name DedicatedHardwareSecurityModule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// network_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/dedicated_hardware_security_module#network_profile DedicatedHardwareSecurityModule#network_profile}
	NetworkProfile *DedicatedHardwareSecurityModuleNetworkProfile `field:"required" json:"networkProfile" yaml:"networkProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/dedicated_hardware_security_module#resource_group_name DedicatedHardwareSecurityModule#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/dedicated_hardware_security_module#sku_name DedicatedHardwareSecurityModule#sku_name}.
	SkuName *string `field:"required" json:"skuName" yaml:"skuName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/dedicated_hardware_security_module#id DedicatedHardwareSecurityModule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// management_network_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/dedicated_hardware_security_module#management_network_profile DedicatedHardwareSecurityModule#management_network_profile}
	ManagementNetworkProfile *DedicatedHardwareSecurityModuleManagementNetworkProfile `field:"optional" json:"managementNetworkProfile" yaml:"managementNetworkProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/dedicated_hardware_security_module#stamp_id DedicatedHardwareSecurityModule#stamp_id}.
	StampId *string `field:"optional" json:"stampId" yaml:"stampId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/dedicated_hardware_security_module#tags DedicatedHardwareSecurityModule#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/dedicated_hardware_security_module#timeouts DedicatedHardwareSecurityModule#timeouts}
	Timeouts *DedicatedHardwareSecurityModuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/dedicated_hardware_security_module#zones DedicatedHardwareSecurityModule#zones}.
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}


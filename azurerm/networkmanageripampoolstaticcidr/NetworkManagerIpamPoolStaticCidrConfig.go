// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkmanageripampoolstaticcidr

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkManagerIpamPoolStaticCidrConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/network_manager_ipam_pool_static_cidr#ipam_pool_id NetworkManagerIpamPoolStaticCidr#ipam_pool_id}.
	IpamPoolId *string `field:"required" json:"ipamPoolId" yaml:"ipamPoolId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/network_manager_ipam_pool_static_cidr#name NetworkManagerIpamPoolStaticCidr#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/network_manager_ipam_pool_static_cidr#address_prefixes NetworkManagerIpamPoolStaticCidr#address_prefixes}.
	AddressPrefixes *[]*string `field:"optional" json:"addressPrefixes" yaml:"addressPrefixes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/network_manager_ipam_pool_static_cidr#id NetworkManagerIpamPoolStaticCidr#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/network_manager_ipam_pool_static_cidr#number_of_ip_addresses_to_allocate NetworkManagerIpamPoolStaticCidr#number_of_ip_addresses_to_allocate}.
	NumberOfIpAddressesToAllocate *string `field:"optional" json:"numberOfIpAddressesToAllocate" yaml:"numberOfIpAddressesToAllocate"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/network_manager_ipam_pool_static_cidr#timeouts NetworkManagerIpamPoolStaticCidr#timeouts}
	Timeouts *NetworkManagerIpamPoolStaticCidrTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualnetworkpanorama

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PaloAltoNextGenerationFirewallVirtualNetworkPanoramaConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.35.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_panorama#location PaloAltoNextGenerationFirewallVirtualNetworkPanorama#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.35.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_panorama#name PaloAltoNextGenerationFirewallVirtualNetworkPanorama#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// network_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.35.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_panorama#network_profile PaloAltoNextGenerationFirewallVirtualNetworkPanorama#network_profile}
	NetworkProfile *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfile `field:"required" json:"networkProfile" yaml:"networkProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.35.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_panorama#panorama_base64_config PaloAltoNextGenerationFirewallVirtualNetworkPanorama#panorama_base64_config}.
	PanoramaBase64Config *string `field:"required" json:"panoramaBase64Config" yaml:"panoramaBase64Config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.35.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_panorama#resource_group_name PaloAltoNextGenerationFirewallVirtualNetworkPanorama#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// destination_nat block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.35.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_panorama#destination_nat PaloAltoNextGenerationFirewallVirtualNetworkPanorama#destination_nat}
	DestinationNat interface{} `field:"optional" json:"destinationNat" yaml:"destinationNat"`
	// dns_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.35.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_panorama#dns_settings PaloAltoNextGenerationFirewallVirtualNetworkPanorama#dns_settings}
	DnsSettings *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDnsSettings `field:"optional" json:"dnsSettings" yaml:"dnsSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.35.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_panorama#id PaloAltoNextGenerationFirewallVirtualNetworkPanorama#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.35.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_panorama#marketplace_offer_id PaloAltoNextGenerationFirewallVirtualNetworkPanorama#marketplace_offer_id}.
	MarketplaceOfferId *string `field:"optional" json:"marketplaceOfferId" yaml:"marketplaceOfferId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.35.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_panorama#plan_id PaloAltoNextGenerationFirewallVirtualNetworkPanorama#plan_id}.
	PlanId *string `field:"optional" json:"planId" yaml:"planId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.35.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_panorama#tags PaloAltoNextGenerationFirewallVirtualNetworkPanorama#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.35.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_panorama#timeouts PaloAltoNextGenerationFirewallVirtualNetworkPanorama#timeouts}
	Timeouts *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


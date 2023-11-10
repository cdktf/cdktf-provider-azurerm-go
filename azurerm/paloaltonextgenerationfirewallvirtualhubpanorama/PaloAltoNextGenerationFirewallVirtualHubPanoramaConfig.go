// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualhubpanorama

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PaloAltoNextGenerationFirewallVirtualHubPanoramaConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_panorama#location PaloAltoNextGenerationFirewallVirtualHubPanorama#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_panorama#name PaloAltoNextGenerationFirewallVirtualHubPanorama#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// network_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_panorama#network_profile PaloAltoNextGenerationFirewallVirtualHubPanorama#network_profile}
	NetworkProfile *PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfile `field:"required" json:"networkProfile" yaml:"networkProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_panorama#panorama_base64_config PaloAltoNextGenerationFirewallVirtualHubPanorama#panorama_base64_config}.
	PanoramaBase64Config *string `field:"required" json:"panoramaBase64Config" yaml:"panoramaBase64Config"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_panorama#resource_group_name PaloAltoNextGenerationFirewallVirtualHubPanorama#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// destination_nat block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_panorama#destination_nat PaloAltoNextGenerationFirewallVirtualHubPanorama#destination_nat}
	DestinationNat interface{} `field:"optional" json:"destinationNat" yaml:"destinationNat"`
	// dns_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_panorama#dns_settings PaloAltoNextGenerationFirewallVirtualHubPanorama#dns_settings}
	DnsSettings *PaloAltoNextGenerationFirewallVirtualHubPanoramaDnsSettings `field:"optional" json:"dnsSettings" yaml:"dnsSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_panorama#id PaloAltoNextGenerationFirewallVirtualHubPanorama#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_panorama#tags PaloAltoNextGenerationFirewallVirtualHubPanorama#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_panorama#timeouts PaloAltoNextGenerationFirewallVirtualHubPanorama#timeouts}
	Timeouts *PaloAltoNextGenerationFirewallVirtualHubPanoramaTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualnetworklocalrulestack

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestackConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_local_rulestack#name PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// network_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_local_rulestack#network_profile PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack#network_profile}
	NetworkProfile *PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestackNetworkProfile `field:"required" json:"networkProfile" yaml:"networkProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_local_rulestack#resource_group_name PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_local_rulestack#rulestack_id PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack#rulestack_id}.
	RulestackId *string `field:"required" json:"rulestackId" yaml:"rulestackId"`
	// destination_nat block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_local_rulestack#destination_nat PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack#destination_nat}
	DestinationNat interface{} `field:"optional" json:"destinationNat" yaml:"destinationNat"`
	// dns_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_local_rulestack#dns_settings PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack#dns_settings}
	DnsSettings *PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestackDnsSettings `field:"optional" json:"dnsSettings" yaml:"dnsSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_local_rulestack#id PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_local_rulestack#tags PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/palo_alto_next_generation_firewall_virtual_network_local_rulestack#timeouts PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestack#timeouts}
	Timeouts *PaloAltoNextGenerationFirewallVirtualNetworkLocalRulestackTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


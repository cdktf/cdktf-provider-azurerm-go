package paloaltonextgenerationfirewallvirtualhublocalrulestack

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PaloAltoNextGenerationFirewallVirtualHubLocalRulestackConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_local_rulestack#name PaloAltoNextGenerationFirewallVirtualHubLocalRulestack#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// network_profile block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_local_rulestack#network_profile PaloAltoNextGenerationFirewallVirtualHubLocalRulestack#network_profile}
	NetworkProfile *PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfile `field:"required" json:"networkProfile" yaml:"networkProfile"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_local_rulestack#resource_group_name PaloAltoNextGenerationFirewallVirtualHubLocalRulestack#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_local_rulestack#rulestack_id PaloAltoNextGenerationFirewallVirtualHubLocalRulestack#rulestack_id}.
	RulestackId *string `field:"required" json:"rulestackId" yaml:"rulestackId"`
	// destination_nat block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_local_rulestack#destination_nat PaloAltoNextGenerationFirewallVirtualHubLocalRulestack#destination_nat}
	DestinationNat interface{} `field:"optional" json:"destinationNat" yaml:"destinationNat"`
	// dns_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_local_rulestack#dns_settings PaloAltoNextGenerationFirewallVirtualHubLocalRulestack#dns_settings}
	DnsSettings *PaloAltoNextGenerationFirewallVirtualHubLocalRulestackDnsSettings `field:"optional" json:"dnsSettings" yaml:"dnsSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_local_rulestack#id PaloAltoNextGenerationFirewallVirtualHubLocalRulestack#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_local_rulestack#tags PaloAltoNextGenerationFirewallVirtualHubLocalRulestack#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/palo_alto_next_generation_firewall_virtual_hub_local_rulestack#timeouts PaloAltoNextGenerationFirewallVirtualHubLocalRulestack#timeouts}
	Timeouts *PaloAltoNextGenerationFirewallVirtualHubLocalRulestackTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


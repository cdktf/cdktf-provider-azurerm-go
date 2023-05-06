package firewallpolicyrulecollectiongroup


type FirewallPolicyRuleCollectionGroupApplicationRuleCollectionRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/firewall_policy_rule_collection_group#name FirewallPolicyRuleCollectionGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/firewall_policy_rule_collection_group#description FirewallPolicyRuleCollectionGroup#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/firewall_policy_rule_collection_group#destination_addresses FirewallPolicyRuleCollectionGroup#destination_addresses}.
	DestinationAddresses *[]*string `field:"optional" json:"destinationAddresses" yaml:"destinationAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/firewall_policy_rule_collection_group#destination_fqdns FirewallPolicyRuleCollectionGroup#destination_fqdns}.
	DestinationFqdns *[]*string `field:"optional" json:"destinationFqdns" yaml:"destinationFqdns"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/firewall_policy_rule_collection_group#destination_fqdn_tags FirewallPolicyRuleCollectionGroup#destination_fqdn_tags}.
	DestinationFqdnTags *[]*string `field:"optional" json:"destinationFqdnTags" yaml:"destinationFqdnTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/firewall_policy_rule_collection_group#destination_urls FirewallPolicyRuleCollectionGroup#destination_urls}.
	DestinationUrls *[]*string `field:"optional" json:"destinationUrls" yaml:"destinationUrls"`
	// protocols block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/firewall_policy_rule_collection_group#protocols FirewallPolicyRuleCollectionGroup#protocols}
	Protocols interface{} `field:"optional" json:"protocols" yaml:"protocols"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/firewall_policy_rule_collection_group#source_addresses FirewallPolicyRuleCollectionGroup#source_addresses}.
	SourceAddresses *[]*string `field:"optional" json:"sourceAddresses" yaml:"sourceAddresses"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/firewall_policy_rule_collection_group#source_ip_groups FirewallPolicyRuleCollectionGroup#source_ip_groups}.
	SourceIpGroups *[]*string `field:"optional" json:"sourceIpGroups" yaml:"sourceIpGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/firewall_policy_rule_collection_group#terminate_tls FirewallPolicyRuleCollectionGroup#terminate_tls}.
	TerminateTls interface{} `field:"optional" json:"terminateTls" yaml:"terminateTls"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.55.0/docs/resources/firewall_policy_rule_collection_group#web_categories FirewallPolicyRuleCollectionGroup#web_categories}.
	WebCategories *[]*string `field:"optional" json:"webCategories" yaml:"webCategories"`
}


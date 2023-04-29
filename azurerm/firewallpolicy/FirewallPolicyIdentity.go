package firewallpolicy


type FirewallPolicyIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/firewall_policy#identity_ids FirewallPolicy#identity_ids}.
	IdentityIds *[]*string `field:"required" json:"identityIds" yaml:"identityIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/firewall_policy#type FirewallPolicy#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}


package cdnfrontdoorfirewallpolicy


type CdnFrontdoorFirewallPolicyManagedRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_firewall_policy#action CdnFrontdoorFirewallPolicy#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_firewall_policy#type CdnFrontdoorFirewallPolicy#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_firewall_policy#version CdnFrontdoorFirewallPolicy#version}.
	Version *string `field:"required" json:"version" yaml:"version"`
	// exclusion block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_firewall_policy#exclusion CdnFrontdoorFirewallPolicy#exclusion}
	Exclusion interface{} `field:"optional" json:"exclusion" yaml:"exclusion"`
	// override block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/cdn_frontdoor_firewall_policy#override CdnFrontdoorFirewallPolicy#override}
	Override interface{} `field:"optional" json:"override" yaml:"override"`
}

package firewallpolicy


type FirewallPolicyIntrusionDetection struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy#mode FirewallPolicy#mode}.
	Mode *string `field:"optional" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy#private_ranges FirewallPolicy#private_ranges}.
	PrivateRanges *[]*string `field:"optional" json:"privateRanges" yaml:"privateRanges"`
	// signature_overrides block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy#signature_overrides FirewallPolicy#signature_overrides}
	SignatureOverrides interface{} `field:"optional" json:"signatureOverrides" yaml:"signatureOverrides"`
	// traffic_bypass block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/firewall_policy#traffic_bypass FirewallPolicy#traffic_bypass}
	TrafficBypass interface{} `field:"optional" json:"trafficBypass" yaml:"trafficBypass"`
}

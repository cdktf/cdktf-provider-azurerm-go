package frontdoor


type FrontdoorFrontendEndpoint struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#host_name Frontdoor#host_name}.
	HostName *string `field:"required" json:"hostName" yaml:"hostName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#name Frontdoor#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#session_affinity_enabled Frontdoor#session_affinity_enabled}.
	SessionAffinityEnabled interface{} `field:"optional" json:"sessionAffinityEnabled" yaml:"sessionAffinityEnabled"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#session_affinity_ttl_seconds Frontdoor#session_affinity_ttl_seconds}.
	SessionAffinityTtlSeconds *float64 `field:"optional" json:"sessionAffinityTtlSeconds" yaml:"sessionAffinityTtlSeconds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#web_application_firewall_policy_link_id Frontdoor#web_application_firewall_policy_link_id}.
	WebApplicationFirewallPolicyLinkId *string `field:"optional" json:"webApplicationFirewallPolicyLinkId" yaml:"webApplicationFirewallPolicyLinkId"`
}


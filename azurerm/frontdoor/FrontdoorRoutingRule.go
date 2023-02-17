package frontdoor


type FrontdoorRoutingRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#accepted_protocols Frontdoor#accepted_protocols}.
	AcceptedProtocols *[]*string `field:"required" json:"acceptedProtocols" yaml:"acceptedProtocols"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#frontend_endpoints Frontdoor#frontend_endpoints}.
	FrontendEndpoints *[]*string `field:"required" json:"frontendEndpoints" yaml:"frontendEndpoints"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#name Frontdoor#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#patterns_to_match Frontdoor#patterns_to_match}.
	PatternsToMatch *[]*string `field:"required" json:"patternsToMatch" yaml:"patternsToMatch"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#enabled Frontdoor#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// forwarding_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#forwarding_configuration Frontdoor#forwarding_configuration}
	ForwardingConfiguration *FrontdoorRoutingRuleForwardingConfiguration `field:"optional" json:"forwardingConfiguration" yaml:"forwardingConfiguration"`
	// redirect_configuration block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/frontdoor#redirect_configuration Frontdoor#redirect_configuration}
	RedirectConfiguration *FrontdoorRoutingRuleRedirectConfiguration `field:"optional" json:"redirectConfiguration" yaml:"redirectConfiguration"`
}


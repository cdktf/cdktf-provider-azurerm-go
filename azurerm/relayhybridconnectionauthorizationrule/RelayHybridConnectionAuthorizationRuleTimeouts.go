package relayhybridconnectionauthorizationrule


type RelayHybridConnectionAuthorizationRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/relay_hybrid_connection_authorization_rule#create RelayHybridConnectionAuthorizationRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/relay_hybrid_connection_authorization_rule#delete RelayHybridConnectionAuthorizationRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/relay_hybrid_connection_authorization_rule#read RelayHybridConnectionAuthorizationRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/relay_hybrid_connection_authorization_rule#update RelayHybridConnectionAuthorizationRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

package applicationgateway


type ApplicationGatewaySslProfileSslPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#cipher_suites ApplicationGateway#cipher_suites}.
	CipherSuites *[]*string `field:"optional" json:"cipherSuites" yaml:"cipherSuites"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#disabled_protocols ApplicationGateway#disabled_protocols}.
	DisabledProtocols *[]*string `field:"optional" json:"disabledProtocols" yaml:"disabledProtocols"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#min_protocol_version ApplicationGateway#min_protocol_version}.
	MinProtocolVersion *string `field:"optional" json:"minProtocolVersion" yaml:"minProtocolVersion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#policy_name ApplicationGateway#policy_name}.
	PolicyName *string `field:"optional" json:"policyName" yaml:"policyName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/application_gateway#policy_type ApplicationGateway#policy_type}.
	PolicyType *string `field:"optional" json:"policyType" yaml:"policyType"`
}


package iotcentralapplicationnetworkruleset


type IotcentralApplicationNetworkRuleSetIpRule struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iotcentral_application_network_rule_set#ip_mask IotcentralApplicationNetworkRuleSet#ip_mask}.
	IpMask *string `field:"required" json:"ipMask" yaml:"ipMask"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/iotcentral_application_network_rule_set#name IotcentralApplicationNetworkRuleSet#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
}

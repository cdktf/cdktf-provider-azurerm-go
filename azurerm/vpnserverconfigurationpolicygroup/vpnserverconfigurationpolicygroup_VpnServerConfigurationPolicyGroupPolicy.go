package vpnserverconfigurationpolicygroup


type VpnServerConfigurationPolicyGroupPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_server_configuration_policy_group#name VpnServerConfigurationPolicyGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_server_configuration_policy_group#type VpnServerConfigurationPolicyGroup#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/vpn_server_configuration_policy_group#value VpnServerConfigurationPolicyGroup#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

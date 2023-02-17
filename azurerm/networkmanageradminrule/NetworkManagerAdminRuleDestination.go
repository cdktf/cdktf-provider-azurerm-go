package networkmanageradminrule


type NetworkManagerAdminRuleDestination struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_admin_rule#address_prefix NetworkManagerAdminRule#address_prefix}.
	AddressPrefix *string `field:"required" json:"addressPrefix" yaml:"addressPrefix"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_admin_rule#address_prefix_type NetworkManagerAdminRule#address_prefix_type}.
	AddressPrefixType *string `field:"required" json:"addressPrefixType" yaml:"addressPrefixType"`
}


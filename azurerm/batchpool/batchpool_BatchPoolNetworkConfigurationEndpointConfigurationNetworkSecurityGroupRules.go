package batchpool


type BatchPoolNetworkConfigurationEndpointConfigurationNetworkSecurityGroupRules struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#access BatchPool#access}.
	Access *string `field:"required" json:"access" yaml:"access"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#priority BatchPool#priority}.
	Priority *float64 `field:"required" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/batch_pool#source_address_prefix BatchPool#source_address_prefix}.
	SourceAddressPrefix *string `field:"required" json:"sourceAddressPrefix" yaml:"sourceAddressPrefix"`
}


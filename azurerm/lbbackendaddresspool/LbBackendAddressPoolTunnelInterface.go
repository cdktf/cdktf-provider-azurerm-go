package lbbackendaddresspool


type LbBackendAddressPoolTunnelInterface struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lb_backend_address_pool#identifier LbBackendAddressPool#identifier}.
	Identifier *float64 `field:"required" json:"identifier" yaml:"identifier"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lb_backend_address_pool#port LbBackendAddressPool#port}.
	Port *float64 `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lb_backend_address_pool#protocol LbBackendAddressPool#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/lb_backend_address_pool#type LbBackendAddressPool#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}

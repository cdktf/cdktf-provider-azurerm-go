package lbbackendaddresspool


type LbBackendAddressPoolTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/lb_backend_address_pool#create LbBackendAddressPool#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/lb_backend_address_pool#delete LbBackendAddressPool#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/lb_backend_address_pool#read LbBackendAddressPool#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/lb_backend_address_pool#update LbBackendAddressPool#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


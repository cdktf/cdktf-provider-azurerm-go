package networkmanagerscopeconnection


type NetworkManagerScopeConnectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_scope_connection#create NetworkManagerScopeConnection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_scope_connection#delete NetworkManagerScopeConnection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_scope_connection#read NetworkManagerScopeConnection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_scope_connection#update NetworkManagerScopeConnection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


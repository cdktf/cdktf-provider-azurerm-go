package networkmanageradminrule


type NetworkManagerAdminRuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_admin_rule#create NetworkManagerAdminRule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_admin_rule#delete NetworkManagerAdminRule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_admin_rule#read NetworkManagerAdminRule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_admin_rule#update NetworkManagerAdminRule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

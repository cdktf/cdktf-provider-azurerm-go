package networkmanageradminrulecollection


type NetworkManagerAdminRuleCollectionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_admin_rule_collection#create NetworkManagerAdminRuleCollection#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_admin_rule_collection#delete NetworkManagerAdminRuleCollection#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_admin_rule_collection#read NetworkManagerAdminRuleCollection#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_admin_rule_collection#update NetworkManagerAdminRuleCollection#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

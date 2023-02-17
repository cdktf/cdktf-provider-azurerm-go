package networkmanagerstaticmember


type NetworkManagerStaticMemberTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_static_member#create NetworkManagerStaticMember#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_static_member#delete NetworkManagerStaticMember#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_static_member#read NetworkManagerStaticMember#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


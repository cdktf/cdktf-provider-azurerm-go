package networkmanagerstaticmember


type NetworkManagerStaticMemberTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/network_manager_static_member#create NetworkManagerStaticMember#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/network_manager_static_member#delete NetworkManagerStaticMember#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.60.0/docs/resources/network_manager_static_member#read NetworkManagerStaticMember#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


package networkinterfacesecuritygroupassociation


type NetworkInterfaceSecurityGroupAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_interface_security_group_association#create NetworkInterfaceSecurityGroupAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_interface_security_group_association#delete NetworkInterfaceSecurityGroupAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_interface_security_group_association#read NetworkInterfaceSecurityGroupAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_interface_security_group_association#update NetworkInterfaceSecurityGroupAssociation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


package networkinterfaceapplicationsecuritygroupassociation


type NetworkInterfaceApplicationSecurityGroupAssociationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_interface_application_security_group_association#create NetworkInterfaceApplicationSecurityGroupAssociation#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_interface_application_security_group_association#delete NetworkInterfaceApplicationSecurityGroupAssociation#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_interface_application_security_group_association#read NetworkInterfaceApplicationSecurityGroupAssociation#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_interface_application_security_group_association#update NetworkInterfaceApplicationSecurityGroupAssociation#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

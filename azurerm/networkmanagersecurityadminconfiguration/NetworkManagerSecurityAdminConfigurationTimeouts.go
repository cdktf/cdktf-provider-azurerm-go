package networkmanagersecurityadminconfiguration


type NetworkManagerSecurityAdminConfigurationTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_security_admin_configuration#create NetworkManagerSecurityAdminConfiguration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_security_admin_configuration#delete NetworkManagerSecurityAdminConfiguration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_security_admin_configuration#read NetworkManagerSecurityAdminConfiguration#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/network_manager_security_admin_configuration#update NetworkManagerSecurityAdminConfiguration#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

package keyvaultmanagedhardwaresecuritymodule


type KeyVaultManagedHardwareSecurityModuleTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/key_vault_managed_hardware_security_module#create KeyVaultManagedHardwareSecurityModule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/key_vault_managed_hardware_security_module#delete KeyVaultManagedHardwareSecurityModule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/key_vault_managed_hardware_security_module#read KeyVaultManagedHardwareSecurityModule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}

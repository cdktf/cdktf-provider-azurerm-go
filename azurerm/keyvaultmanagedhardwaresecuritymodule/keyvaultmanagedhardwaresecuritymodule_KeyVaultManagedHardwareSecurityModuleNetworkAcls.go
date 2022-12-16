package keyvaultmanagedhardwaresecuritymodule


type KeyVaultManagedHardwareSecurityModuleNetworkAcls struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/key_vault_managed_hardware_security_module#bypass KeyVaultManagedHardwareSecurityModule#bypass}.
	Bypass *string `field:"required" json:"bypass" yaml:"bypass"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/key_vault_managed_hardware_security_module#default_action KeyVaultManagedHardwareSecurityModule#default_action}.
	DefaultAction *string `field:"required" json:"defaultAction" yaml:"defaultAction"`
}


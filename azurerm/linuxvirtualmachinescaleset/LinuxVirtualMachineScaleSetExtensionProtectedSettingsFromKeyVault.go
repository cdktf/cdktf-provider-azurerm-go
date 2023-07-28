package linuxvirtualmachinescaleset


type LinuxVirtualMachineScaleSetExtensionProtectedSettingsFromKeyVault struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_virtual_machine_scale_set#secret_url LinuxVirtualMachineScaleSet#secret_url}.
	SecretUrl *string `field:"required" json:"secretUrl" yaml:"secretUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_virtual_machine_scale_set#source_vault_id LinuxVirtualMachineScaleSet#source_vault_id}.
	SourceVaultId *string `field:"required" json:"sourceVaultId" yaml:"sourceVaultId"`
}


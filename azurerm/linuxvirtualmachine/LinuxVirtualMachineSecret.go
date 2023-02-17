package linuxvirtualmachine


type LinuxVirtualMachineSecret struct {
	// certificate block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_virtual_machine#certificate LinuxVirtualMachine#certificate}
	Certificate interface{} `field:"required" json:"certificate" yaml:"certificate"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_virtual_machine#key_vault_id LinuxVirtualMachine#key_vault_id}.
	KeyVaultId *string `field:"required" json:"keyVaultId" yaml:"keyVaultId"`
}


package linuxvirtualmachine


type LinuxVirtualMachineAdminSshKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_virtual_machine#public_key LinuxVirtualMachine#public_key}.
	PublicKey *string `field:"required" json:"publicKey" yaml:"publicKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_virtual_machine#username LinuxVirtualMachine#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}


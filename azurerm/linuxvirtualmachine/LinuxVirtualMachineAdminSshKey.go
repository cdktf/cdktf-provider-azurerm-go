package linuxvirtualmachine


type LinuxVirtualMachineAdminSshKey struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/linux_virtual_machine#public_key LinuxVirtualMachine#public_key}.
	PublicKey *string `field:"required" json:"publicKey" yaml:"publicKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/linux_virtual_machine#username LinuxVirtualMachine#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}


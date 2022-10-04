package linuxvirtualmachinescaleset


type LinuxVirtualMachineScaleSetAdminSshKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_virtual_machine_scale_set#public_key LinuxVirtualMachineScaleSet#public_key}.
	PublicKey *string `field:"required" json:"publicKey" yaml:"publicKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_virtual_machine_scale_set#username LinuxVirtualMachineScaleSet#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}


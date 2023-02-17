package linuxvirtualmachine


type LinuxVirtualMachinePlan struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_virtual_machine#name LinuxVirtualMachine#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_virtual_machine#product LinuxVirtualMachine#product}.
	Product *string `field:"required" json:"product" yaml:"product"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_virtual_machine#publisher LinuxVirtualMachine#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
}


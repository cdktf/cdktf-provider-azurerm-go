package linuxvirtualmachine


type LinuxVirtualMachineOsDiskDiffDiskSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_virtual_machine#option LinuxVirtualMachine#option}.
	Option *string `field:"required" json:"option" yaml:"option"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_virtual_machine#placement LinuxVirtualMachine#placement}.
	Placement *string `field:"optional" json:"placement" yaml:"placement"`
}


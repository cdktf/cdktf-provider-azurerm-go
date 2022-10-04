package windowsvirtualmachine


type WindowsVirtualMachineOsDiskDiffDiskSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine#option WindowsVirtualMachine#option}.
	Option *string `field:"required" json:"option" yaml:"option"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine#placement WindowsVirtualMachine#placement}.
	Placement *string `field:"optional" json:"placement" yaml:"placement"`
}


package windowsvirtualmachinescaleset


type WindowsVirtualMachineScaleSetOsDiskDiffDiskSettings struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#option WindowsVirtualMachineScaleSet#option}.
	Option *string `field:"required" json:"option" yaml:"option"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#placement WindowsVirtualMachineScaleSet#placement}.
	Placement *string `field:"optional" json:"placement" yaml:"placement"`
}


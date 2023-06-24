package windowsvirtualmachinescaleset


type WindowsVirtualMachineScaleSetOsDiskDiffDiskSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/windows_virtual_machine_scale_set#option WindowsVirtualMachineScaleSet#option}.
	Option *string `field:"required" json:"option" yaml:"option"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/windows_virtual_machine_scale_set#placement WindowsVirtualMachineScaleSet#placement}.
	Placement *string `field:"optional" json:"placement" yaml:"placement"`
}


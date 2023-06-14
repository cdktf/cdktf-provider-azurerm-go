package windowsvirtualmachinescaleset


type WindowsVirtualMachineScaleSetTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/windows_virtual_machine_scale_set#create WindowsVirtualMachineScaleSet#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/windows_virtual_machine_scale_set#delete WindowsVirtualMachineScaleSet#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/windows_virtual_machine_scale_set#read WindowsVirtualMachineScaleSet#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.61.0/docs/resources/windows_virtual_machine_scale_set#update WindowsVirtualMachineScaleSet#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


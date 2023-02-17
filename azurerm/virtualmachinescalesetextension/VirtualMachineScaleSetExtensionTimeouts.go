package virtualmachinescalesetextension


type VirtualMachineScaleSetExtensionTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine_scale_set_extension#create VirtualMachineScaleSetExtensionA#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine_scale_set_extension#delete VirtualMachineScaleSetExtensionA#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine_scale_set_extension#read VirtualMachineScaleSetExtensionA#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine_scale_set_extension#update VirtualMachineScaleSetExtensionA#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


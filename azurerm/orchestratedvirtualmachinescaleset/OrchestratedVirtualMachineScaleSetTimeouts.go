package orchestratedvirtualmachinescaleset


type OrchestratedVirtualMachineScaleSetTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orchestrated_virtual_machine_scale_set#create OrchestratedVirtualMachineScaleSet#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orchestrated_virtual_machine_scale_set#delete OrchestratedVirtualMachineScaleSet#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orchestrated_virtual_machine_scale_set#read OrchestratedVirtualMachineScaleSet#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orchestrated_virtual_machine_scale_set#update OrchestratedVirtualMachineScaleSet#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


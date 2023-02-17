package orchestratedvirtualmachinescaleset


type OrchestratedVirtualMachineScaleSetNetworkInterfaceIpConfigurationPublicIpAddressIpTag struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orchestrated_virtual_machine_scale_set#tag OrchestratedVirtualMachineScaleSet#tag}.
	Tag *string `field:"required" json:"tag" yaml:"tag"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orchestrated_virtual_machine_scale_set#type OrchestratedVirtualMachineScaleSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}


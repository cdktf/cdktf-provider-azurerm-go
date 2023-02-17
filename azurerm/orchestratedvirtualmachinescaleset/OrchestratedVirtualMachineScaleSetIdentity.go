package orchestratedvirtualmachinescaleset


type OrchestratedVirtualMachineScaleSetIdentity struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orchestrated_virtual_machine_scale_set#identity_ids OrchestratedVirtualMachineScaleSet#identity_ids}.
	IdentityIds *[]*string `field:"required" json:"identityIds" yaml:"identityIds"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orchestrated_virtual_machine_scale_set#type OrchestratedVirtualMachineScaleSet#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}


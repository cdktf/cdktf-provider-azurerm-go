package orchestratedvirtualmachinescaleset


type OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationAdminSshKey struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orchestrated_virtual_machine_scale_set#public_key OrchestratedVirtualMachineScaleSet#public_key}.
	PublicKey *string `field:"required" json:"publicKey" yaml:"publicKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orchestrated_virtual_machine_scale_set#username OrchestratedVirtualMachineScaleSet#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}


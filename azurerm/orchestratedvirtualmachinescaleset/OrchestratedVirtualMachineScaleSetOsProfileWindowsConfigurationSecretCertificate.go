package orchestratedvirtualmachinescaleset


type OrchestratedVirtualMachineScaleSetOsProfileWindowsConfigurationSecretCertificate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orchestrated_virtual_machine_scale_set#store OrchestratedVirtualMachineScaleSet#store}.
	Store *string `field:"required" json:"store" yaml:"store"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/orchestrated_virtual_machine_scale_set#url OrchestratedVirtualMachineScaleSet#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
}


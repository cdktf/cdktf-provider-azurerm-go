package windowsvirtualmachinescaleset


type WindowsVirtualMachineScaleSetSecretCertificate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#store WindowsVirtualMachineScaleSet#store}.
	Store *string `field:"required" json:"store" yaml:"store"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine_scale_set#url WindowsVirtualMachineScaleSet#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
}


package windowsvirtualmachinescaleset


type WindowsVirtualMachineScaleSetWinrmListener struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/windows_virtual_machine_scale_set#protocol WindowsVirtualMachineScaleSet#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/windows_virtual_machine_scale_set#certificate_url WindowsVirtualMachineScaleSet#certificate_url}.
	CertificateUrl *string `field:"optional" json:"certificateUrl" yaml:"certificateUrl"`
}


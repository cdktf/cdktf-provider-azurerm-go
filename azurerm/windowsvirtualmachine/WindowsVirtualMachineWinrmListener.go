package windowsvirtualmachine


type WindowsVirtualMachineWinrmListener struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine#protocol WindowsVirtualMachine#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/windows_virtual_machine#certificate_url WindowsVirtualMachine#certificate_url}.
	CertificateUrl *string `field:"optional" json:"certificateUrl" yaml:"certificateUrl"`
}


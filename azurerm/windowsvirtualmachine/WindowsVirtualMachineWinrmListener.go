package windowsvirtualmachine


type WindowsVirtualMachineWinrmListener struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/windows_virtual_machine#protocol WindowsVirtualMachine#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/windows_virtual_machine#certificate_url WindowsVirtualMachine#certificate_url}.
	CertificateUrl *string `field:"optional" json:"certificateUrl" yaml:"certificateUrl"`
}


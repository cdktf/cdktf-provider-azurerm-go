package virtualmachine


type VirtualMachineOsProfileSecretsVaultCertificates struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine#certificate_url VirtualMachine#certificate_url}.
	CertificateUrl *string `field:"required" json:"certificateUrl" yaml:"certificateUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine#certificate_store VirtualMachine#certificate_store}.
	CertificateStore *string `field:"optional" json:"certificateStore" yaml:"certificateStore"`
}


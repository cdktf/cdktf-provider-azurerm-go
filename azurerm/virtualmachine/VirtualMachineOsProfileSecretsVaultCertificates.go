package virtualmachine


type VirtualMachineOsProfileSecretsVaultCertificates struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/virtual_machine#certificate_url VirtualMachine#certificate_url}.
	CertificateUrl *string `field:"required" json:"certificateUrl" yaml:"certificateUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/virtual_machine#certificate_store VirtualMachine#certificate_store}.
	CertificateStore *string `field:"optional" json:"certificateStore" yaml:"certificateStore"`
}


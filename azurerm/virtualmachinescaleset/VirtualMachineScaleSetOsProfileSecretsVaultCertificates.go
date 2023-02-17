package virtualmachinescaleset


type VirtualMachineScaleSetOsProfileSecretsVaultCertificates struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine_scale_set#certificate_url VirtualMachineScaleSet#certificate_url}.
	CertificateUrl *string `field:"required" json:"certificateUrl" yaml:"certificateUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/virtual_machine_scale_set#certificate_store VirtualMachineScaleSet#certificate_store}.
	CertificateStore *string `field:"optional" json:"certificateStore" yaml:"certificateStore"`
}


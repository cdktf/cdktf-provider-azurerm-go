package attestationprovider


type AttestationProviderPolicy struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/attestation_provider#data AttestationProvider#data}.
	Data *string `field:"required" json:"data" yaml:"data"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/attestation_provider#environment_type AttestationProvider#environment_type}.
	EnvironmentType *string `field:"required" json:"environmentType" yaml:"environmentType"`
}


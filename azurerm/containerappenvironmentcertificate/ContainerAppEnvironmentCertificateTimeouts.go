package containerappenvironmentcertificate


type ContainerAppEnvironmentCertificateTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app_environment_certificate#create ContainerAppEnvironmentCertificate#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app_environment_certificate#delete ContainerAppEnvironmentCertificate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app_environment_certificate#read ContainerAppEnvironmentCertificate#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/container_app_environment_certificate#update ContainerAppEnvironmentCertificate#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

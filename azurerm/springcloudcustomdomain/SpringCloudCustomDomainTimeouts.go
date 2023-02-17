package springcloudcustomdomain


type SpringCloudCustomDomainTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_custom_domain#create SpringCloudCustomDomain#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_custom_domain#delete SpringCloudCustomDomain#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_custom_domain#read SpringCloudCustomDomain#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_custom_domain#update SpringCloudCustomDomain#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


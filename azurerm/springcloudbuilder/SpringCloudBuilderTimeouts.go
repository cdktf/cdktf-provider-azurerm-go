package springcloudbuilder


type SpringCloudBuilderTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/spring_cloud_builder#create SpringCloudBuilder#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/spring_cloud_builder#delete SpringCloudBuilder#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/spring_cloud_builder#read SpringCloudBuilder#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/spring_cloud_builder#update SpringCloudBuilder#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


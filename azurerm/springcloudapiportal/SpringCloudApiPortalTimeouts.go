package springcloudapiportal


type SpringCloudApiPortalTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/spring_cloud_api_portal#create SpringCloudApiPortal#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/spring_cloud_api_portal#delete SpringCloudApiPortal#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/spring_cloud_api_portal#read SpringCloudApiPortal#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.53.0/docs/resources/spring_cloud_api_portal#update SpringCloudApiPortal#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


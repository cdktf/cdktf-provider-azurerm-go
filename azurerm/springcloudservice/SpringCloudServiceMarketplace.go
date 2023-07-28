package springcloudservice


type SpringCloudServiceMarketplace struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/spring_cloud_service#plan SpringCloudService#plan}.
	Plan *string `field:"required" json:"plan" yaml:"plan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/spring_cloud_service#product SpringCloudService#product}.
	Product *string `field:"required" json:"product" yaml:"product"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/spring_cloud_service#publisher SpringCloudService#publisher}.
	Publisher *string `field:"required" json:"publisher" yaml:"publisher"`
}


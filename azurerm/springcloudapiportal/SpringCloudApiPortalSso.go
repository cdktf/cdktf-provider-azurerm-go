package springcloudapiportal


type SpringCloudApiPortalSso struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/spring_cloud_api_portal#client_id SpringCloudApiPortal#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/spring_cloud_api_portal#client_secret SpringCloudApiPortal#client_secret}.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/spring_cloud_api_portal#issuer_uri SpringCloudApiPortal#issuer_uri}.
	IssuerUri *string `field:"optional" json:"issuerUri" yaml:"issuerUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/spring_cloud_api_portal#scope SpringCloudApiPortal#scope}.
	Scope *[]*string `field:"optional" json:"scope" yaml:"scope"`
}


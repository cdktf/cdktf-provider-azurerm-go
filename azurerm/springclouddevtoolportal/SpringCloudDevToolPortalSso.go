package springclouddevtoolportal


type SpringCloudDevToolPortalSso struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_dev_tool_portal#client_id SpringCloudDevToolPortal#client_id}.
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_dev_tool_portal#client_secret SpringCloudDevToolPortal#client_secret}.
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_dev_tool_portal#metadata_url SpringCloudDevToolPortal#metadata_url}.
	MetadataUrl *string `field:"optional" json:"metadataUrl" yaml:"metadataUrl"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_dev_tool_portal#scope SpringCloudDevToolPortal#scope}.
	Scope *[]*string `field:"optional" json:"scope" yaml:"scope"`
}

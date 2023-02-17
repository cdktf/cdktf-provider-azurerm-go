package springcloudbuilder


type SpringCloudBuilderBuildPackGroup struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_builder#name SpringCloudBuilder#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_builder#build_pack_ids SpringCloudBuilder#build_pack_ids}.
	BuildPackIds *[]*string `field:"optional" json:"buildPackIds" yaml:"buildPackIds"`
}


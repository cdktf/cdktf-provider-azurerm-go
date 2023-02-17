package springcloudbuildpackbinding


type SpringCloudBuildPackBindingLaunch struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_build_pack_binding#properties SpringCloudBuildPackBinding#properties}.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/spring_cloud_build_pack_binding#secrets SpringCloudBuildPackBinding#secrets}.
	Secrets *map[string]*string `field:"optional" json:"secrets" yaml:"secrets"`
}


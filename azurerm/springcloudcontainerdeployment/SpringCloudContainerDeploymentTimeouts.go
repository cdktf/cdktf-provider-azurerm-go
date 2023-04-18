package springcloudcontainerdeployment


type SpringCloudContainerDeploymentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/spring_cloud_container_deployment#create SpringCloudContainerDeployment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/spring_cloud_container_deployment#delete SpringCloudContainerDeployment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/spring_cloud_container_deployment#read SpringCloudContainerDeployment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.52.0/docs/resources/spring_cloud_container_deployment#update SpringCloudContainerDeployment#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


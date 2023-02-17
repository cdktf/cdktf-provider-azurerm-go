package nginxdeployment


type NginxDeploymentLoggingStorageAccount struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/nginx_deployment#container_name NginxDeployment#container_name}.
	ContainerName *string `field:"optional" json:"containerName" yaml:"containerName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/nginx_deployment#name NginxDeployment#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}


package nginxdeployment


type NginxDeploymentFrontendPublic struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.63.0/docs/resources/nginx_deployment#ip_address NginxDeployment#ip_address}.
	IpAddress *[]*string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
}


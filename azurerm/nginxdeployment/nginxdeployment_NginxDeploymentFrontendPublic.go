package nginxdeployment


type NginxDeploymentFrontendPublic struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/nginx_deployment#ip_address NginxDeployment#ip_address}.
	IpAddress *[]*string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
}


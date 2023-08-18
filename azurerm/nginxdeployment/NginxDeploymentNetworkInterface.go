package nginxdeployment


type NginxDeploymentNetworkInterface struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/nginx_deployment#subnet_id NginxDeployment#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}


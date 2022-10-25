package nginxdeployment


type NginxDeploymentNetworkInterface struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/nginx_deployment#subnet_id NginxDeployment#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}


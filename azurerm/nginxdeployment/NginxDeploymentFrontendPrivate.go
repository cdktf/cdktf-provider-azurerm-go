package nginxdeployment


type NginxDeploymentFrontendPrivate struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/nginx_deployment#allocation_method NginxDeployment#allocation_method}.
	AllocationMethod *string `field:"required" json:"allocationMethod" yaml:"allocationMethod"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/nginx_deployment#ip_address NginxDeployment#ip_address}.
	IpAddress *string `field:"required" json:"ipAddress" yaml:"ipAddress"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/nginx_deployment#subnet_id NginxDeployment#subnet_id}.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
}


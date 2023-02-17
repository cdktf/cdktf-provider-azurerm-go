package appservice


type AppServiceSiteConfigIpRestrictionHeaders struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#x_azure_fdid AppService#x_azure_fdid}.
	XAzureFdid *[]*string `field:"optional" json:"xAzureFdid" yaml:"xAzureFdid"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#x_fd_health_probe AppService#x_fd_health_probe}.
	XFdHealthProbe *[]*string `field:"optional" json:"xFdHealthProbe" yaml:"xFdHealthProbe"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#x_forwarded_for AppService#x_forwarded_for}.
	XForwardedFor *[]*string `field:"optional" json:"xForwardedFor" yaml:"xForwardedFor"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#x_forwarded_host AppService#x_forwarded_host}.
	XForwardedHost *[]*string `field:"optional" json:"xForwardedHost" yaml:"xForwardedHost"`
}


package dataazurermlogicappstandard


type DataAzurermLogicAppStandardSiteConfigScmIpRestrictionHeaders struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/logic_app_standard#x_azure_fdid DataAzurermLogicAppStandard#x_azure_fdid}.
	XAzureFdid *[]*string `field:"optional" json:"xAzureFdid" yaml:"xAzureFdid"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/logic_app_standard#x_fd_health_probe DataAzurermLogicAppStandard#x_fd_health_probe}.
	XFdHealthProbe *[]*string `field:"optional" json:"xFdHealthProbe" yaml:"xFdHealthProbe"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/logic_app_standard#x_forwarded_for DataAzurermLogicAppStandard#x_forwarded_for}.
	XForwardedFor *[]*string `field:"optional" json:"xForwardedFor" yaml:"xForwardedFor"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/d/logic_app_standard#x_forwarded_host DataAzurermLogicAppStandard#x_forwarded_host}.
	XForwardedHost *[]*string `field:"optional" json:"xForwardedHost" yaml:"xForwardedHost"`
}

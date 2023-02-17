package logicappstandard


type LogicAppStandardSiteConfigScmIpRestrictionHeaders struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_standard#x_azure_fdid LogicAppStandard#x_azure_fdid}.
	XAzureFdid *[]*string `field:"optional" json:"xAzureFdid" yaml:"xAzureFdid"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_standard#x_fd_health_probe LogicAppStandard#x_fd_health_probe}.
	XFdHealthProbe *[]*string `field:"optional" json:"xFdHealthProbe" yaml:"xFdHealthProbe"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_standard#x_forwarded_for LogicAppStandard#x_forwarded_for}.
	XForwardedFor *[]*string `field:"optional" json:"xForwardedFor" yaml:"xForwardedFor"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/logic_app_standard#x_forwarded_host LogicAppStandard#x_forwarded_host}.
	XForwardedHost *[]*string `field:"optional" json:"xForwardedHost" yaml:"xForwardedHost"`
}


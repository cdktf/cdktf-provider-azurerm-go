package datadogmonitor


type DatadogMonitorDatadogOrganization struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/datadog_monitor#api_key DatadogMonitor#api_key}.
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/datadog_monitor#application_key DatadogMonitor#application_key}.
	ApplicationKey *string `field:"required" json:"applicationKey" yaml:"applicationKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/datadog_monitor#enterprise_app_id DatadogMonitor#enterprise_app_id}.
	EnterpriseAppId *string `field:"optional" json:"enterpriseAppId" yaml:"enterpriseAppId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/datadog_monitor#linking_auth_code DatadogMonitor#linking_auth_code}.
	LinkingAuthCode *string `field:"optional" json:"linkingAuthCode" yaml:"linkingAuthCode"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/datadog_monitor#linking_client_id DatadogMonitor#linking_client_id}.
	LinkingClientId *string `field:"optional" json:"linkingClientId" yaml:"linkingClientId"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/datadog_monitor#redirect_uri DatadogMonitor#redirect_uri}.
	RedirectUri *string `field:"optional" json:"redirectUri" yaml:"redirectUri"`
}

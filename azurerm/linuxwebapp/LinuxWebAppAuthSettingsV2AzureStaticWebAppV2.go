package linuxwebapp


type LinuxWebAppAuthSettingsV2AzureStaticWebAppV2 struct {
	// The ID of the Client to use to authenticate with Azure Static Web App Authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/linux_web_app#client_id LinuxWebApp#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
}

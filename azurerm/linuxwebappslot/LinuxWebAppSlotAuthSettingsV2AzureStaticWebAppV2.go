package linuxwebappslot


type LinuxWebAppSlotAuthSettingsV2AzureStaticWebAppV2 struct {
	// The ID of the Client to use to authenticate with Azure Static Web App Authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/linux_web_app_slot#client_id LinuxWebAppSlot#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
}


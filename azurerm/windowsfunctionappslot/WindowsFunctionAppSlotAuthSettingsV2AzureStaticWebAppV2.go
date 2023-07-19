package windowsfunctionappslot


type WindowsFunctionAppSlotAuthSettingsV2AzureStaticWebAppV2 struct {
	// The ID of the Client to use to authenticate with Azure Static Web App Authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/windows_function_app_slot#client_id WindowsFunctionAppSlot#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
}


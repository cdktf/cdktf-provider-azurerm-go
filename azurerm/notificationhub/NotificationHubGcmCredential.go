package notificationhub


type NotificationHubGcmCredential struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.54.0/docs/resources/notification_hub#api_key NotificationHub#api_key}.
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
}


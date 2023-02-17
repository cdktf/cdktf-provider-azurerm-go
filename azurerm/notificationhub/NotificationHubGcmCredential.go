package notificationhub


type NotificationHubGcmCredential struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/notification_hub#api_key NotificationHub#api_key}.
	ApiKey *string `field:"required" json:"apiKey" yaml:"apiKey"`
}


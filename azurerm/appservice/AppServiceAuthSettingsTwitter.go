package appservice


type AppServiceAuthSettingsTwitter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/app_service#consumer_key AppService#consumer_key}.
	ConsumerKey *string `field:"required" json:"consumerKey" yaml:"consumerKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/app_service#consumer_secret AppService#consumer_secret}.
	ConsumerSecret *string `field:"required" json:"consumerSecret" yaml:"consumerSecret"`
}


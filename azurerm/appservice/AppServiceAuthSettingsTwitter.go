package appservice


type AppServiceAuthSettingsTwitter struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#consumer_key AppService#consumer_key}.
	ConsumerKey *string `field:"required" json:"consumerKey" yaml:"consumerKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/app_service#consumer_secret AppService#consumer_secret}.
	ConsumerSecret *string `field:"required" json:"consumerSecret" yaml:"consumerSecret"`
}

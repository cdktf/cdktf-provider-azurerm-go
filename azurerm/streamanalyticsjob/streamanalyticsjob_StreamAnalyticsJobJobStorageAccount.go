package streamanalyticsjob


type StreamAnalyticsJobJobStorageAccount struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_job#account_key StreamAnalyticsJob#account_key}.
	AccountKey *string `field:"required" json:"accountKey" yaml:"accountKey"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_job#account_name StreamAnalyticsJob#account_name}.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/stream_analytics_job#authentication_mode StreamAnalyticsJob#authentication_mode}.
	AuthenticationMode *string `field:"optional" json:"authenticationMode" yaml:"authenticationMode"`
}


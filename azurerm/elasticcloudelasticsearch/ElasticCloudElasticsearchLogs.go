package elasticcloudelasticsearch


type ElasticCloudElasticsearchLogs struct {
	// filtering_tag block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/elastic_cloud_elasticsearch#filtering_tag ElasticCloudElasticsearch#filtering_tag}
	FilteringTag interface{} `field:"optional" json:"filteringTag" yaml:"filteringTag"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/elastic_cloud_elasticsearch#send_activity_logs ElasticCloudElasticsearch#send_activity_logs}.
	SendActivityLogs interface{} `field:"optional" json:"sendActivityLogs" yaml:"sendActivityLogs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/elastic_cloud_elasticsearch#send_azuread_logs ElasticCloudElasticsearch#send_azuread_logs}.
	SendAzureadLogs interface{} `field:"optional" json:"sendAzureadLogs" yaml:"sendAzureadLogs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/elastic_cloud_elasticsearch#send_subscription_logs ElasticCloudElasticsearch#send_subscription_logs}.
	SendSubscriptionLogs interface{} `field:"optional" json:"sendSubscriptionLogs" yaml:"sendSubscriptionLogs"`
}


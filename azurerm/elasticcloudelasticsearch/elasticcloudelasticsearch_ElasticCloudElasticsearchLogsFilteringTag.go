package elasticcloudelasticsearch


type ElasticCloudElasticsearchLogsFilteringTag struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/elastic_cloud_elasticsearch#action ElasticCloudElasticsearch#action}.
	Action *string `field:"required" json:"action" yaml:"action"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/elastic_cloud_elasticsearch#name ElasticCloudElasticsearch#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/elastic_cloud_elasticsearch#value ElasticCloudElasticsearch#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}

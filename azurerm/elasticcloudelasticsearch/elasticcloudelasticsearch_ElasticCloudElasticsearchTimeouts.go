package elasticcloudelasticsearch


type ElasticCloudElasticsearchTimeouts struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/elastic_cloud_elasticsearch#create ElasticCloudElasticsearch#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/elastic_cloud_elasticsearch#delete ElasticCloudElasticsearch#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/elastic_cloud_elasticsearch#read ElasticCloudElasticsearch#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm/r/elastic_cloud_elasticsearch#update ElasticCloudElasticsearch#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

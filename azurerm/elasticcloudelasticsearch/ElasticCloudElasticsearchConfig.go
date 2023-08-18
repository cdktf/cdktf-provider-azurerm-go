package elasticcloudelasticsearch

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ElasticCloudElasticsearchConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/elastic_cloud_elasticsearch#elastic_cloud_email_address ElasticCloudElasticsearch#elastic_cloud_email_address}.
	ElasticCloudEmailAddress *string `field:"required" json:"elasticCloudEmailAddress" yaml:"elasticCloudEmailAddress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/elastic_cloud_elasticsearch#location ElasticCloudElasticsearch#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/elastic_cloud_elasticsearch#name ElasticCloudElasticsearch#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/elastic_cloud_elasticsearch#resource_group_name ElasticCloudElasticsearch#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/elastic_cloud_elasticsearch#sku_name ElasticCloudElasticsearch#sku_name}.
	SkuName *string `field:"required" json:"skuName" yaml:"skuName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/elastic_cloud_elasticsearch#id ElasticCloudElasticsearch#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// logs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/elastic_cloud_elasticsearch#logs ElasticCloudElasticsearch#logs}
	Logs *ElasticCloudElasticsearchLogs `field:"optional" json:"logs" yaml:"logs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/elastic_cloud_elasticsearch#monitoring_enabled ElasticCloudElasticsearch#monitoring_enabled}.
	MonitoringEnabled interface{} `field:"optional" json:"monitoringEnabled" yaml:"monitoringEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/elastic_cloud_elasticsearch#tags ElasticCloudElasticsearch#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/elastic_cloud_elasticsearch#timeouts ElasticCloudElasticsearch#timeouts}
	Timeouts *ElasticCloudElasticsearchTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


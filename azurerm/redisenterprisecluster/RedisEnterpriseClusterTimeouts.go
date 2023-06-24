package redisenterprisecluster


type RedisEnterpriseClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/redis_enterprise_cluster#create RedisEnterpriseCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/redis_enterprise_cluster#delete RedisEnterpriseCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/redis_enterprise_cluster#read RedisEnterpriseCluster#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/redis_enterprise_cluster#update RedisEnterpriseCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


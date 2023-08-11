package cosmosdbgremlindatabase

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CosmosdbGremlinDatabaseConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/cosmosdb_gremlin_database#account_name CosmosdbGremlinDatabase#account_name}.
	AccountName *string `field:"required" json:"accountName" yaml:"accountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/cosmosdb_gremlin_database#name CosmosdbGremlinDatabase#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/cosmosdb_gremlin_database#resource_group_name CosmosdbGremlinDatabase#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// autoscale_settings block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/cosmosdb_gremlin_database#autoscale_settings CosmosdbGremlinDatabase#autoscale_settings}
	AutoscaleSettings *CosmosdbGremlinDatabaseAutoscaleSettings `field:"optional" json:"autoscaleSettings" yaml:"autoscaleSettings"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/cosmosdb_gremlin_database#id CosmosdbGremlinDatabase#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/cosmosdb_gremlin_database#throughput CosmosdbGremlinDatabase#throughput}.
	Throughput *float64 `field:"optional" json:"throughput" yaml:"throughput"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.69.0/docs/resources/cosmosdb_gremlin_database#timeouts CosmosdbGremlinDatabase#timeouts}
	Timeouts *CosmosdbGremlinDatabaseTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


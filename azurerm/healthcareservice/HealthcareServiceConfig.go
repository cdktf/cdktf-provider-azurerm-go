package healthcareservice

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HealthcareServiceConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/healthcare_service#location HealthcareService#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/healthcare_service#name HealthcareService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/healthcare_service#resource_group_name HealthcareService#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/healthcare_service#access_policy_object_ids HealthcareService#access_policy_object_ids}.
	AccessPolicyObjectIds *[]*string `field:"optional" json:"accessPolicyObjectIds" yaml:"accessPolicyObjectIds"`
	// authentication_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/healthcare_service#authentication_configuration HealthcareService#authentication_configuration}
	AuthenticationConfiguration *HealthcareServiceAuthenticationConfiguration `field:"optional" json:"authenticationConfiguration" yaml:"authenticationConfiguration"`
	// cors_configuration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/healthcare_service#cors_configuration HealthcareService#cors_configuration}
	CorsConfiguration *HealthcareServiceCorsConfiguration `field:"optional" json:"corsConfiguration" yaml:"corsConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/healthcare_service#cosmosdb_key_vault_key_versionless_id HealthcareService#cosmosdb_key_vault_key_versionless_id}.
	CosmosdbKeyVaultKeyVersionlessId *string `field:"optional" json:"cosmosdbKeyVaultKeyVersionlessId" yaml:"cosmosdbKeyVaultKeyVersionlessId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/healthcare_service#cosmosdb_throughput HealthcareService#cosmosdb_throughput}.
	CosmosdbThroughput *float64 `field:"optional" json:"cosmosdbThroughput" yaml:"cosmosdbThroughput"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/healthcare_service#id HealthcareService#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/healthcare_service#kind HealthcareService#kind}.
	Kind *string `field:"optional" json:"kind" yaml:"kind"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/healthcare_service#public_network_access_enabled HealthcareService#public_network_access_enabled}.
	PublicNetworkAccessEnabled interface{} `field:"optional" json:"publicNetworkAccessEnabled" yaml:"publicNetworkAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/healthcare_service#tags HealthcareService#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/resources/healthcare_service#timeouts HealthcareService#timeouts}
	Timeouts *HealthcareServiceTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


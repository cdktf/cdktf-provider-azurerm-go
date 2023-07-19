package integrationserviceenvironment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IntegrationServiceEnvironmentConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/integration_service_environment#access_endpoint_type IntegrationServiceEnvironment#access_endpoint_type}.
	AccessEndpointType *string `field:"required" json:"accessEndpointType" yaml:"accessEndpointType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/integration_service_environment#location IntegrationServiceEnvironment#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/integration_service_environment#name IntegrationServiceEnvironment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/integration_service_environment#resource_group_name IntegrationServiceEnvironment#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/integration_service_environment#virtual_network_subnet_ids IntegrationServiceEnvironment#virtual_network_subnet_ids}.
	VirtualNetworkSubnetIds *[]*string `field:"required" json:"virtualNetworkSubnetIds" yaml:"virtualNetworkSubnetIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/integration_service_environment#id IntegrationServiceEnvironment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/integration_service_environment#sku_name IntegrationServiceEnvironment#sku_name}.
	SkuName *string `field:"optional" json:"skuName" yaml:"skuName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/integration_service_environment#tags IntegrationServiceEnvironment#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.65.0/docs/resources/integration_service_environment#timeouts IntegrationServiceEnvironment#timeouts}
	Timeouts *IntegrationServiceEnvironmentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


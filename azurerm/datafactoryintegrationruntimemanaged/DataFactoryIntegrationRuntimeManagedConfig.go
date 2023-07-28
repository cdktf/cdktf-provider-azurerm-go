package datafactoryintegrationruntimemanaged

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryIntegrationRuntimeManagedConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/data_factory_integration_runtime_managed#data_factory_id DataFactoryIntegrationRuntimeManaged#data_factory_id}.
	DataFactoryId *string `field:"required" json:"dataFactoryId" yaml:"dataFactoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/data_factory_integration_runtime_managed#location DataFactoryIntegrationRuntimeManaged#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/data_factory_integration_runtime_managed#name DataFactoryIntegrationRuntimeManaged#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/data_factory_integration_runtime_managed#node_size DataFactoryIntegrationRuntimeManaged#node_size}.
	NodeSize *string `field:"required" json:"nodeSize" yaml:"nodeSize"`
	// catalog_info block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/data_factory_integration_runtime_managed#catalog_info DataFactoryIntegrationRuntimeManaged#catalog_info}
	CatalogInfo *DataFactoryIntegrationRuntimeManagedCatalogInfo `field:"optional" json:"catalogInfo" yaml:"catalogInfo"`
	// custom_setup_script block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/data_factory_integration_runtime_managed#custom_setup_script DataFactoryIntegrationRuntimeManaged#custom_setup_script}
	CustomSetupScript *DataFactoryIntegrationRuntimeManagedCustomSetupScript `field:"optional" json:"customSetupScript" yaml:"customSetupScript"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/data_factory_integration_runtime_managed#description DataFactoryIntegrationRuntimeManaged#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/data_factory_integration_runtime_managed#edition DataFactoryIntegrationRuntimeManaged#edition}.
	Edition *string `field:"optional" json:"edition" yaml:"edition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/data_factory_integration_runtime_managed#id DataFactoryIntegrationRuntimeManaged#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/data_factory_integration_runtime_managed#license_type DataFactoryIntegrationRuntimeManaged#license_type}.
	LicenseType *string `field:"optional" json:"licenseType" yaml:"licenseType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/data_factory_integration_runtime_managed#max_parallel_executions_per_node DataFactoryIntegrationRuntimeManaged#max_parallel_executions_per_node}.
	MaxParallelExecutionsPerNode *float64 `field:"optional" json:"maxParallelExecutionsPerNode" yaml:"maxParallelExecutionsPerNode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/data_factory_integration_runtime_managed#number_of_nodes DataFactoryIntegrationRuntimeManaged#number_of_nodes}.
	NumberOfNodes *float64 `field:"optional" json:"numberOfNodes" yaml:"numberOfNodes"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/data_factory_integration_runtime_managed#timeouts DataFactoryIntegrationRuntimeManaged#timeouts}
	Timeouts *DataFactoryIntegrationRuntimeManagedTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// vnet_integration block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/data_factory_integration_runtime_managed#vnet_integration DataFactoryIntegrationRuntimeManaged#vnet_integration}
	VnetIntegration *DataFactoryIntegrationRuntimeManagedVnetIntegration `field:"optional" json:"vnetIntegration" yaml:"vnetIntegration"`
}


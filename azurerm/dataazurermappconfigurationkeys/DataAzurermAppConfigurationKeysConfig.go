package dataazurermappconfigurationkeys

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermAppConfigurationKeysConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/data-sources/app_configuration_keys#configuration_store_id DataAzurermAppConfigurationKeys#configuration_store_id}.
	ConfigurationStoreId *string `field:"required" json:"configurationStoreId" yaml:"configurationStoreId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/data-sources/app_configuration_keys#id DataAzurermAppConfigurationKeys#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/data-sources/app_configuration_keys#key DataAzurermAppConfigurationKeys#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/data-sources/app_configuration_keys#label DataAzurermAppConfigurationKeys#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.62.1/docs/data-sources/app_configuration_keys#timeouts DataAzurermAppConfigurationKeys#timeouts}
	Timeouts *DataAzurermAppConfigurationKeysTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


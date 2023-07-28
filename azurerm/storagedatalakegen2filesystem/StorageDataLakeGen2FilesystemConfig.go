package storagedatalakegen2filesystem

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type StorageDataLakeGen2FilesystemConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/storage_data_lake_gen2_filesystem#name StorageDataLakeGen2Filesystem#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/storage_data_lake_gen2_filesystem#storage_account_id StorageDataLakeGen2Filesystem#storage_account_id}.
	StorageAccountId *string `field:"required" json:"storageAccountId" yaml:"storageAccountId"`
	// ace block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/storage_data_lake_gen2_filesystem#ace StorageDataLakeGen2Filesystem#ace}
	Ace interface{} `field:"optional" json:"ace" yaml:"ace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/storage_data_lake_gen2_filesystem#group StorageDataLakeGen2Filesystem#group}.
	Group *string `field:"optional" json:"group" yaml:"group"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/storage_data_lake_gen2_filesystem#id StorageDataLakeGen2Filesystem#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/storage_data_lake_gen2_filesystem#owner StorageDataLakeGen2Filesystem#owner}.
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/storage_data_lake_gen2_filesystem#properties StorageDataLakeGen2Filesystem#properties}.
	Properties *map[string]*string `field:"optional" json:"properties" yaml:"properties"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.67.0/docs/resources/storage_data_lake_gen2_filesystem#timeouts StorageDataLakeGen2Filesystem#timeouts}
	Timeouts *StorageDataLakeGen2FilesystemTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


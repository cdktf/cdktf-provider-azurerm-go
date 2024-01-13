// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorytriggerblobevent

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataFactoryTriggerBlobEventConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/data_factory_trigger_blob_event#data_factory_id DataFactoryTriggerBlobEvent#data_factory_id}.
	DataFactoryId *string `field:"required" json:"dataFactoryId" yaml:"dataFactoryId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/data_factory_trigger_blob_event#events DataFactoryTriggerBlobEvent#events}.
	Events *[]*string `field:"required" json:"events" yaml:"events"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/data_factory_trigger_blob_event#name DataFactoryTriggerBlobEvent#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// pipeline block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/data_factory_trigger_blob_event#pipeline DataFactoryTriggerBlobEvent#pipeline}
	Pipeline interface{} `field:"required" json:"pipeline" yaml:"pipeline"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/data_factory_trigger_blob_event#storage_account_id DataFactoryTriggerBlobEvent#storage_account_id}.
	StorageAccountId *string `field:"required" json:"storageAccountId" yaml:"storageAccountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/data_factory_trigger_blob_event#activated DataFactoryTriggerBlobEvent#activated}.
	Activated interface{} `field:"optional" json:"activated" yaml:"activated"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/data_factory_trigger_blob_event#additional_properties DataFactoryTriggerBlobEvent#additional_properties}.
	AdditionalProperties *map[string]*string `field:"optional" json:"additionalProperties" yaml:"additionalProperties"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/data_factory_trigger_blob_event#annotations DataFactoryTriggerBlobEvent#annotations}.
	Annotations *[]*string `field:"optional" json:"annotations" yaml:"annotations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/data_factory_trigger_blob_event#blob_path_begins_with DataFactoryTriggerBlobEvent#blob_path_begins_with}.
	BlobPathBeginsWith *string `field:"optional" json:"blobPathBeginsWith" yaml:"blobPathBeginsWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/data_factory_trigger_blob_event#blob_path_ends_with DataFactoryTriggerBlobEvent#blob_path_ends_with}.
	BlobPathEndsWith *string `field:"optional" json:"blobPathEndsWith" yaml:"blobPathEndsWith"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/data_factory_trigger_blob_event#description DataFactoryTriggerBlobEvent#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/data_factory_trigger_blob_event#id DataFactoryTriggerBlobEvent#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/data_factory_trigger_blob_event#ignore_empty_blobs DataFactoryTriggerBlobEvent#ignore_empty_blobs}.
	IgnoreEmptyBlobs interface{} `field:"optional" json:"ignoreEmptyBlobs" yaml:"ignoreEmptyBlobs"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/data_factory_trigger_blob_event#timeouts DataFactoryTriggerBlobEvent#timeouts}
	Timeouts *DataFactoryTriggerBlobEventTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


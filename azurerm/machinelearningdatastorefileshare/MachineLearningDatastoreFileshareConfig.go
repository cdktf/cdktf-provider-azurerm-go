// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package machinelearningdatastorefileshare

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MachineLearningDatastoreFileshareConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/machine_learning_datastore_fileshare#name MachineLearningDatastoreFileshare#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/machine_learning_datastore_fileshare#storage_fileshare_id MachineLearningDatastoreFileshare#storage_fileshare_id}.
	StorageFileshareId *string `field:"required" json:"storageFileshareId" yaml:"storageFileshareId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/machine_learning_datastore_fileshare#workspace_id MachineLearningDatastoreFileshare#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/machine_learning_datastore_fileshare#account_key MachineLearningDatastoreFileshare#account_key}.
	AccountKey *string `field:"optional" json:"accountKey" yaml:"accountKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/machine_learning_datastore_fileshare#description MachineLearningDatastoreFileshare#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/machine_learning_datastore_fileshare#id MachineLearningDatastoreFileshare#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/machine_learning_datastore_fileshare#service_data_identity MachineLearningDatastoreFileshare#service_data_identity}.
	ServiceDataIdentity *string `field:"optional" json:"serviceDataIdentity" yaml:"serviceDataIdentity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/machine_learning_datastore_fileshare#shared_access_signature MachineLearningDatastoreFileshare#shared_access_signature}.
	SharedAccessSignature *string `field:"optional" json:"sharedAccessSignature" yaml:"sharedAccessSignature"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/machine_learning_datastore_fileshare#tags MachineLearningDatastoreFileshare#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/machine_learning_datastore_fileshare#timeouts MachineLearningDatastoreFileshare#timeouts}
	Timeouts *MachineLearningDatastoreFileshareTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


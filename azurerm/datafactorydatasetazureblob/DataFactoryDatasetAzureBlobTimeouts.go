// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorydatasetazureblob


type DataFactoryDatasetAzureBlobTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/data_factory_dataset_azure_blob#create DataFactoryDatasetAzureBlob#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/data_factory_dataset_azure_blob#delete DataFactoryDatasetAzureBlob#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/data_factory_dataset_azure_blob#read DataFactoryDatasetAzureBlob#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/data_factory_dataset_azure_blob#update DataFactoryDatasetAzureBlob#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorydatasetbinary


type DataFactoryDatasetBinaryCompression struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/data_factory_dataset_binary#type DataFactoryDatasetBinary#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.70.0/docs/resources/data_factory_dataset_binary#level DataFactoryDatasetBinary#level}.
	Level *string `field:"optional" json:"level" yaml:"level"`
}


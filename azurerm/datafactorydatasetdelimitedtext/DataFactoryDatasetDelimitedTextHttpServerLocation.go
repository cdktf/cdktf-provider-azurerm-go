// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package datafactorydatasetdelimitedtext


type DataFactoryDatasetDelimitedTextHttpServerLocation struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_dataset_delimited_text#filename DataFactoryDatasetDelimitedText#filename}.
	Filename *string `field:"required" json:"filename" yaml:"filename"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_dataset_delimited_text#path DataFactoryDatasetDelimitedText#path}.
	Path *string `field:"required" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_dataset_delimited_text#relative_url DataFactoryDatasetDelimitedText#relative_url}.
	RelativeUrl *string `field:"required" json:"relativeUrl" yaml:"relativeUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_dataset_delimited_text#dynamic_filename_enabled DataFactoryDatasetDelimitedText#dynamic_filename_enabled}.
	DynamicFilenameEnabled interface{} `field:"optional" json:"dynamicFilenameEnabled" yaml:"dynamicFilenameEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/data_factory_dataset_delimited_text#dynamic_path_enabled DataFactoryDatasetDelimitedText#dynamic_path_enabled}.
	DynamicPathEnabled interface{} `field:"optional" json:"dynamicPathEnabled" yaml:"dynamicPathEnabled"`
}


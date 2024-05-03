// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchpool


type BatchPoolStartTaskContainer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/batch_pool#image_name BatchPool#image_name}.
	ImageName *string `field:"required" json:"imageName" yaml:"imageName"`
	// registry block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/batch_pool#registry BatchPool#registry}
	Registry interface{} `field:"optional" json:"registry" yaml:"registry"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/batch_pool#run_options BatchPool#run_options}.
	RunOptions *string `field:"optional" json:"runOptions" yaml:"runOptions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.102.0/docs/resources/batch_pool#working_directory BatchPool#working_directory}.
	WorkingDirectory *string `field:"optional" json:"workingDirectory" yaml:"workingDirectory"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package synapsesparkpool


type SynapseSparkPoolSparkConfig struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/synapse_spark_pool#content SynapseSparkPool#content}.
	Content *string `field:"required" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/synapse_spark_pool#filename SynapseSparkPool#filename}.
	Filename *string `field:"required" json:"filename" yaml:"filename"`
}


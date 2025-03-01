// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchpool


type BatchPoolAutoScale struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/batch_pool#formula BatchPool#formula}.
	Formula *string `field:"required" json:"formula" yaml:"formula"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/batch_pool#evaluation_interval BatchPool#evaluation_interval}.
	EvaluationInterval *string `field:"optional" json:"evaluationInterval" yaml:"evaluationInterval"`
}


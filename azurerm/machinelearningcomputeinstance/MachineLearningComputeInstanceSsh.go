// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package machinelearningcomputeinstance


type MachineLearningComputeInstanceSsh struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.0/docs/resources/machine_learning_compute_instance#public_key MachineLearningComputeInstance#public_key}.
	PublicKey *string `field:"required" json:"publicKey" yaml:"publicKey"`
}


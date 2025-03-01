// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package machinelearninginferencecluster


type MachineLearningInferenceClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/machine_learning_inference_cluster#create MachineLearningInferenceCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/machine_learning_inference_cluster#delete MachineLearningInferenceCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.1/docs/resources/machine_learning_inference_cluster#read MachineLearningInferenceCluster#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


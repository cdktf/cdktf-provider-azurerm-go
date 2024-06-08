// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package machinelearninginferencecluster


type MachineLearningInferenceClusterSsl struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/machine_learning_inference_cluster#cert MachineLearningInferenceCluster#cert}.
	Cert *string `field:"optional" json:"cert" yaml:"cert"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/machine_learning_inference_cluster#cname MachineLearningInferenceCluster#cname}.
	Cname *string `field:"optional" json:"cname" yaml:"cname"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/machine_learning_inference_cluster#key MachineLearningInferenceCluster#key}.
	Key *string `field:"optional" json:"key" yaml:"key"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/machine_learning_inference_cluster#leaf_domain_label MachineLearningInferenceCluster#leaf_domain_label}.
	LeafDomainLabel *string `field:"optional" json:"leafDomainLabel" yaml:"leafDomainLabel"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.107.0/docs/resources/machine_learning_inference_cluster#overwrite_existing_domain MachineLearningInferenceCluster#overwrite_existing_domain}.
	OverwriteExistingDomain interface{} `field:"optional" json:"overwriteExistingDomain" yaml:"overwriteExistingDomain"`
}


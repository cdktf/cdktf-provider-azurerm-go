// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudcustomizedaccelerator


type SpringCloudCustomizedAcceleratorGitRepositorySshAuth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/spring_cloud_customized_accelerator#private_key SpringCloudCustomizedAccelerator#private_key}.
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/spring_cloud_customized_accelerator#host_key SpringCloudCustomizedAccelerator#host_key}.
	HostKey *string `field:"optional" json:"hostKey" yaml:"hostKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.26.0/docs/resources/spring_cloud_customized_accelerator#host_key_algorithm SpringCloudCustomizedAccelerator#host_key_algorithm}.
	HostKeyAlgorithm *string `field:"optional" json:"hostKeyAlgorithm" yaml:"hostKeyAlgorithm"`
}


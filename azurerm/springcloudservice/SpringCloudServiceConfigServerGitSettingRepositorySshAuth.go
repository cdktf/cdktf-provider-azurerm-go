// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudservice


type SpringCloudServiceConfigServerGitSettingRepositorySshAuth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/spring_cloud_service#private_key SpringCloudService#private_key}.
	PrivateKey *string `field:"required" json:"privateKey" yaml:"privateKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/spring_cloud_service#host_key SpringCloudService#host_key}.
	HostKey *string `field:"optional" json:"hostKey" yaml:"hostKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/spring_cloud_service#host_key_algorithm SpringCloudService#host_key_algorithm}.
	HostKeyAlgorithm *string `field:"optional" json:"hostKeyAlgorithm" yaml:"hostKeyAlgorithm"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/spring_cloud_service#strict_host_key_checking_enabled SpringCloudService#strict_host_key_checking_enabled}.
	StrictHostKeyCheckingEnabled interface{} `field:"optional" json:"strictHostKeyCheckingEnabled" yaml:"strictHostKeyCheckingEnabled"`
}


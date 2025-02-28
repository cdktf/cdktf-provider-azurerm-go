// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudservice


type SpringCloudServiceConfigServerGitSettingRepository struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/spring_cloud_service#name SpringCloudService#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/spring_cloud_service#uri SpringCloudService#uri}.
	Uri *string `field:"required" json:"uri" yaml:"uri"`
	// http_basic_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/spring_cloud_service#http_basic_auth SpringCloudService#http_basic_auth}
	HttpBasicAuth *SpringCloudServiceConfigServerGitSettingRepositoryHttpBasicAuth `field:"optional" json:"httpBasicAuth" yaml:"httpBasicAuth"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/spring_cloud_service#label SpringCloudService#label}.
	Label *string `field:"optional" json:"label" yaml:"label"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/spring_cloud_service#pattern SpringCloudService#pattern}.
	Pattern *[]*string `field:"optional" json:"pattern" yaml:"pattern"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/spring_cloud_service#search_paths SpringCloudService#search_paths}.
	SearchPaths *[]*string `field:"optional" json:"searchPaths" yaml:"searchPaths"`
	// ssh_auth block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/spring_cloud_service#ssh_auth SpringCloudService#ssh_auth}
	SshAuth *SpringCloudServiceConfigServerGitSettingRepositorySshAuth `field:"optional" json:"sshAuth" yaml:"sshAuth"`
}


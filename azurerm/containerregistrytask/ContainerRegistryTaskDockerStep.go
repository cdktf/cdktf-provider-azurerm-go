// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerregistrytask


type ContainerRegistryTaskDockerStep struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/container_registry_task#context_access_token ContainerRegistryTask#context_access_token}.
	ContextAccessToken *string `field:"required" json:"contextAccessToken" yaml:"contextAccessToken"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/container_registry_task#context_path ContainerRegistryTask#context_path}.
	ContextPath *string `field:"required" json:"contextPath" yaml:"contextPath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/container_registry_task#dockerfile_path ContainerRegistryTask#dockerfile_path}.
	DockerfilePath *string `field:"required" json:"dockerfilePath" yaml:"dockerfilePath"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/container_registry_task#arguments ContainerRegistryTask#arguments}.
	Arguments *map[string]*string `field:"optional" json:"arguments" yaml:"arguments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/container_registry_task#cache_enabled ContainerRegistryTask#cache_enabled}.
	CacheEnabled interface{} `field:"optional" json:"cacheEnabled" yaml:"cacheEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/container_registry_task#image_names ContainerRegistryTask#image_names}.
	ImageNames *[]*string `field:"optional" json:"imageNames" yaml:"imageNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/container_registry_task#push_enabled ContainerRegistryTask#push_enabled}.
	PushEnabled interface{} `field:"optional" json:"pushEnabled" yaml:"pushEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/container_registry_task#secret_arguments ContainerRegistryTask#secret_arguments}.
	SecretArguments *map[string]*string `field:"optional" json:"secretArguments" yaml:"secretArguments"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.110.0/docs/resources/container_registry_task#target ContainerRegistryTask#target}.
	Target *string `field:"optional" json:"target" yaml:"target"`
}


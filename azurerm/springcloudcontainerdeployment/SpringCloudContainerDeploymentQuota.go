// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudcontainerdeployment


type SpringCloudContainerDeploymentQuota struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/spring_cloud_container_deployment#cpu SpringCloudContainerDeployment#cpu}.
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.19.0/docs/resources/spring_cloud_container_deployment#memory SpringCloudContainerDeployment#memory}.
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudbuilddeployment


type SpringCloudBuildDeploymentQuota struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/spring_cloud_build_deployment#cpu SpringCloudBuildDeployment#cpu}.
	Cpu *string `field:"optional" json:"cpu" yaml:"cpu"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.92.0/docs/resources/spring_cloud_build_deployment#memory SpringCloudBuildDeployment#memory}.
	Memory *string `field:"optional" json:"memory" yaml:"memory"`
}


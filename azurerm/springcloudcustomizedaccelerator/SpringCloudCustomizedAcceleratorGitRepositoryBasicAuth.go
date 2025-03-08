// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudcustomizedaccelerator


type SpringCloudCustomizedAcceleratorGitRepositoryBasicAuth struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/spring_cloud_customized_accelerator#password SpringCloudCustomizedAccelerator#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/spring_cloud_customized_accelerator#username SpringCloudCustomizedAccelerator#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudcustomizedaccelerator

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpringCloudCustomizedAcceleratorConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// git_repository block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/spring_cloud_customized_accelerator#git_repository SpringCloudCustomizedAccelerator#git_repository}
	GitRepository *SpringCloudCustomizedAcceleratorGitRepository `field:"required" json:"gitRepository" yaml:"gitRepository"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/spring_cloud_customized_accelerator#name SpringCloudCustomizedAccelerator#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/spring_cloud_customized_accelerator#spring_cloud_accelerator_id SpringCloudCustomizedAccelerator#spring_cloud_accelerator_id}.
	SpringCloudAcceleratorId *string `field:"required" json:"springCloudAcceleratorId" yaml:"springCloudAcceleratorId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/spring_cloud_customized_accelerator#accelerator_tags SpringCloudCustomizedAccelerator#accelerator_tags}.
	AcceleratorTags *[]*string `field:"optional" json:"acceleratorTags" yaml:"acceleratorTags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/spring_cloud_customized_accelerator#accelerator_type SpringCloudCustomizedAccelerator#accelerator_type}.
	AcceleratorType *string `field:"optional" json:"acceleratorType" yaml:"acceleratorType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/spring_cloud_customized_accelerator#description SpringCloudCustomizedAccelerator#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/spring_cloud_customized_accelerator#display_name SpringCloudCustomizedAccelerator#display_name}.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/spring_cloud_customized_accelerator#icon_url SpringCloudCustomizedAccelerator#icon_url}.
	IconUrl *string `field:"optional" json:"iconUrl" yaml:"iconUrl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/spring_cloud_customized_accelerator#id SpringCloudCustomizedAccelerator#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.111.0/docs/resources/spring_cloud_customized_accelerator#timeouts SpringCloudCustomizedAccelerator#timeouts}
	Timeouts *SpringCloudCustomizedAcceleratorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


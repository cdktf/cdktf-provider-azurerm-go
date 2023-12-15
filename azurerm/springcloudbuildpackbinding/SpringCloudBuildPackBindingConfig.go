// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudbuildpackbinding

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpringCloudBuildPackBindingConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/spring_cloud_build_pack_binding#name SpringCloudBuildPackBinding#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/spring_cloud_build_pack_binding#spring_cloud_builder_id SpringCloudBuildPackBinding#spring_cloud_builder_id}.
	SpringCloudBuilderId *string `field:"required" json:"springCloudBuilderId" yaml:"springCloudBuilderId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/spring_cloud_build_pack_binding#binding_type SpringCloudBuildPackBinding#binding_type}.
	BindingType *string `field:"optional" json:"bindingType" yaml:"bindingType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/spring_cloud_build_pack_binding#id SpringCloudBuildPackBinding#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// launch block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/spring_cloud_build_pack_binding#launch SpringCloudBuildPackBinding#launch}
	Launch *SpringCloudBuildPackBindingLaunch `field:"optional" json:"launch" yaml:"launch"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.85.0/docs/resources/spring_cloud_build_pack_binding#timeouts SpringCloudBuildPackBinding#timeouts}
	Timeouts *SpringCloudBuildPackBindingTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


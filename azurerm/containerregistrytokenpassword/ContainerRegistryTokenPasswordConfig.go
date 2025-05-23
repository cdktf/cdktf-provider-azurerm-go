// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerregistrytokenpassword

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerRegistryTokenPasswordConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/container_registry_token_password#container_registry_token_id ContainerRegistryTokenPassword#container_registry_token_id}.
	ContainerRegistryTokenId *string `field:"required" json:"containerRegistryTokenId" yaml:"containerRegistryTokenId"`
	// password1 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/container_registry_token_password#password1 ContainerRegistryTokenPassword#password1}
	Password1 *ContainerRegistryTokenPasswordPassword1 `field:"required" json:"password1" yaml:"password1"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/container_registry_token_password#id ContainerRegistryTokenPassword#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// password2 block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/container_registry_token_password#password2 ContainerRegistryTokenPassword#password2}
	Password2 *ContainerRegistryTokenPasswordPassword2 `field:"optional" json:"password2" yaml:"password2"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/container_registry_token_password#timeouts ContainerRegistryTokenPassword#timeouts}
	Timeouts *ContainerRegistryTokenPasswordTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


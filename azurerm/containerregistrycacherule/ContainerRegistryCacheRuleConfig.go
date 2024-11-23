// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerregistrycacherule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerRegistryCacheRuleConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_cache_rule#container_registry_id ContainerRegistryCacheRule#container_registry_id}.
	ContainerRegistryId *string `field:"required" json:"containerRegistryId" yaml:"containerRegistryId"`
	// The name of the cache rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_cache_rule#name ContainerRegistryCacheRule#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The full source repository path such as 'docker.io/library/ubuntu'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_cache_rule#source_repo ContainerRegistryCacheRule#source_repo}
	SourceRepo *string `field:"required" json:"sourceRepo" yaml:"sourceRepo"`
	// The target repository namespace such as 'ubuntu'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_cache_rule#target_repo ContainerRegistryCacheRule#target_repo}
	TargetRepo *string `field:"required" json:"targetRepo" yaml:"targetRepo"`
	// The ARM resource ID of the credential store which is associated with the cache rule.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_cache_rule#credential_set_id ContainerRegistryCacheRule#credential_set_id}
	CredentialSetId *string `field:"optional" json:"credentialSetId" yaml:"credentialSetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_cache_rule#id ContainerRegistryCacheRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.11.0/docs/resources/container_registry_cache_rule#timeouts ContainerRegistryCacheRule#timeouts}
	Timeouts *ContainerRegistryCacheRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package chaosstudiocapability

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ChaosStudioCapabilityConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/chaos_studio_capability#capability_type ChaosStudioCapability#capability_type}.
	CapabilityType *string `field:"required" json:"capabilityType" yaml:"capabilityType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/chaos_studio_capability#chaos_studio_target_id ChaosStudioCapability#chaos_studio_target_id}.
	ChaosStudioTargetId *string `field:"required" json:"chaosStudioTargetId" yaml:"chaosStudioTargetId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/chaos_studio_capability#id ChaosStudioCapability#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/chaos_studio_capability#timeouts ChaosStudioCapability#timeouts}
	Timeouts *ChaosStudioCapabilityTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


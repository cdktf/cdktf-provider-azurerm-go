// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rediscacheaccesspolicyassignment


type RedisCacheAccessPolicyAssignmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/redis_cache_access_policy_assignment#create RedisCacheAccessPolicyAssignment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/redis_cache_access_policy_assignment#delete RedisCacheAccessPolicyAssignment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/redis_cache_access_policy_assignment#read RedisCacheAccessPolicyAssignment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


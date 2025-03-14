// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package rolemanagementpolicy


type RoleManagementPolicyEligibleAssignmentRules struct {
	// Must the assignment have an expiry date.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/role_management_policy#expiration_required RoleManagementPolicy#expiration_required}
	ExpirationRequired interface{} `field:"optional" json:"expirationRequired" yaml:"expirationRequired"`
	// The duration after which assignments expire.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/role_management_policy#expire_after RoleManagementPolicy#expire_after}
	ExpireAfter *string `field:"optional" json:"expireAfter" yaml:"expireAfter"`
}


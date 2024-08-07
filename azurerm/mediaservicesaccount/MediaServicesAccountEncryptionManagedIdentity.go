// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mediaservicesaccount


type MediaServicesAccountEncryptionManagedIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/media_services_account#user_assigned_identity_id MediaServicesAccount#user_assigned_identity_id}.
	UserAssignedIdentityId *string `field:"optional" json:"userAssignedIdentityId" yaml:"userAssignedIdentityId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/media_services_account#use_system_assigned_identity MediaServicesAccount#use_system_assigned_identity}.
	UseSystemAssignedIdentity interface{} `field:"optional" json:"useSystemAssignedIdentity" yaml:"useSystemAssignedIdentity"`
}


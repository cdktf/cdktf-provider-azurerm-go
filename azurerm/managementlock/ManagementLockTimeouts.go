// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managementlock


type ManagementLockTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/management_lock#create ManagementLock#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/management_lock#delete ManagementLock#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/management_lock#read ManagementLock#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


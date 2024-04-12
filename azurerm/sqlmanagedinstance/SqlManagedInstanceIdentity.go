// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqlmanagedinstance


type SqlManagedInstanceIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/sql_managed_instance#type SqlManagedInstance#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}


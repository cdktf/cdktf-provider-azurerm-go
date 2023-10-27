// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sqlserver


type SqlServerIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.78.0/docs/resources/sql_server#type SqlServer#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}


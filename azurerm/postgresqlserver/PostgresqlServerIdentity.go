// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package postgresqlserver


type PostgresqlServerIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/postgresql_server#type PostgresqlServer#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}


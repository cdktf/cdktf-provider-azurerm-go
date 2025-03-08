// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package postgresqlflexibleserver


type PostgresqlFlexibleServerHighAvailability struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/postgresql_flexible_server#mode PostgresqlFlexibleServer#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.22.0/docs/resources/postgresql_flexible_server#standby_availability_zone PostgresqlFlexibleServer#standby_availability_zone}.
	StandbyAvailabilityZone *string `field:"optional" json:"standbyAvailabilityZone" yaml:"standbyAvailabilityZone"`
}


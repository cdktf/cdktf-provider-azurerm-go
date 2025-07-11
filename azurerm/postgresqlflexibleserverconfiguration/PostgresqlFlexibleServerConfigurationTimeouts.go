// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package postgresqlflexibleserverconfiguration


type PostgresqlFlexibleServerConfigurationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/postgresql_flexible_server_configuration#create PostgresqlFlexibleServerConfiguration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/postgresql_flexible_server_configuration#delete PostgresqlFlexibleServerConfiguration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/postgresql_flexible_server_configuration#read PostgresqlFlexibleServerConfiguration#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.36.0/docs/resources/postgresql_flexible_server_configuration#update PostgresqlFlexibleServerConfiguration#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


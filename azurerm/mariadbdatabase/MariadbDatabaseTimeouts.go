// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package mariadbdatabase


type MariadbDatabaseTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/mariadb_database#create MariadbDatabase#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/mariadb_database#delete MariadbDatabase#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/mariadb_database#read MariadbDatabase#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


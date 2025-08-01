// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kustodatabaseprincipalassignment


type KustoDatabasePrincipalAssignmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/kusto_database_principal_assignment#create KustoDatabasePrincipalAssignment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/kusto_database_principal_assignment#delete KustoDatabasePrincipalAssignment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/kusto_database_principal_assignment#read KustoDatabasePrincipalAssignment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


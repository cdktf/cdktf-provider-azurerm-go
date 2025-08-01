// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kustoclusterprincipalassignment


type KustoClusterPrincipalAssignmentTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/kusto_cluster_principal_assignment#create KustoClusterPrincipalAssignment#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/kusto_cluster_principal_assignment#delete KustoClusterPrincipalAssignment#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/kusto_cluster_principal_assignment#read KustoClusterPrincipalAssignment#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


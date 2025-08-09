// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kustoclustercustomermanagedkey


type KustoClusterCustomerManagedKeyTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/kusto_cluster_customer_managed_key#create KustoClusterCustomerManagedKey#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/kusto_cluster_customer_managed_key#delete KustoClusterCustomerManagedKey#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/kusto_cluster_customer_managed_key#read KustoClusterCustomerManagedKey#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.39.0/docs/resources/kusto_cluster_customer_managed_key#update KustoClusterCustomerManagedKey#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


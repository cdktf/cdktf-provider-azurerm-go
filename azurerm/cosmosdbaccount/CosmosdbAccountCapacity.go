// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbaccount


type CosmosdbAccountCapacity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/cosmosdb_account#total_throughput_limit CosmosdbAccount#total_throughput_limit}.
	TotalThroughputLimit *float64 `field:"required" json:"totalThroughputLimit" yaml:"totalThroughputLimit"`
}


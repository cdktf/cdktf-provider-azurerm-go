// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbaccount


type CosmosdbAccountCorsRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/cosmosdb_account#allowed_headers CosmosdbAccount#allowed_headers}.
	AllowedHeaders *[]*string `field:"required" json:"allowedHeaders" yaml:"allowedHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/cosmosdb_account#allowed_methods CosmosdbAccount#allowed_methods}.
	AllowedMethods *[]*string `field:"required" json:"allowedMethods" yaml:"allowedMethods"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/cosmosdb_account#allowed_origins CosmosdbAccount#allowed_origins}.
	AllowedOrigins *[]*string `field:"required" json:"allowedOrigins" yaml:"allowedOrigins"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/cosmosdb_account#exposed_headers CosmosdbAccount#exposed_headers}.
	ExposedHeaders *[]*string `field:"required" json:"exposedHeaders" yaml:"exposedHeaders"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/cosmosdb_account#max_age_in_seconds CosmosdbAccount#max_age_in_seconds}.
	MaxAgeInSeconds *float64 `field:"optional" json:"maxAgeInSeconds" yaml:"maxAgeInSeconds"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package searchservice


type SearchServiceIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/search_service#type SearchService#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}


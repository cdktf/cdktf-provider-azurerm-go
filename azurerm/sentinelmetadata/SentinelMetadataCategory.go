// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelmetadata


type SentinelMetadataCategory struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.90.0/docs/resources/sentinel_metadata#domains SentinelMetadata#domains}.
	Domains *[]*string `field:"optional" json:"domains" yaml:"domains"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.90.0/docs/resources/sentinel_metadata#verticals SentinelMetadata#verticals}.
	Verticals *[]*string `field:"optional" json:"verticals" yaml:"verticals"`
}


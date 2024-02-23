// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelmetadata


type SentinelMetadataSupport struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/sentinel_metadata#tier SentinelMetadata#tier}.
	Tier *string `field:"required" json:"tier" yaml:"tier"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/sentinel_metadata#email SentinelMetadata#email}.
	Email *string `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/sentinel_metadata#link SentinelMetadata#link}.
	Link *string `field:"optional" json:"link" yaml:"link"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/sentinel_metadata#name SentinelMetadata#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}


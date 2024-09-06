// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package sentinelmetadata


type SentinelMetadataAuthor struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/sentinel_metadata#email SentinelMetadata#email}.
	Email *string `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/sentinel_metadata#link SentinelMetadata#link}.
	Link *string `field:"optional" json:"link" yaml:"link"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/sentinel_metadata#name SentinelMetadata#name}.
	Name *string `field:"optional" json:"name" yaml:"name"`
}


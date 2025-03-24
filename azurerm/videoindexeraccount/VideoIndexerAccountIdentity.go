// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package videoindexeraccount


type VideoIndexerAccountIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/video_indexer_account#type VideoIndexerAccount#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.24.0/docs/resources/video_indexer_account#identity_ids VideoIndexerAccount#identity_ids}.
	IdentityIds *[]*string `field:"optional" json:"identityIds" yaml:"identityIds"`
}


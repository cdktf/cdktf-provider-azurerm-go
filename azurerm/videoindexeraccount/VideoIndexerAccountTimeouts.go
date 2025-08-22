// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package videoindexeraccount


type VideoIndexerAccountTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/video_indexer_account#create VideoIndexerAccount#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/video_indexer_account#delete VideoIndexerAccount#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/video_indexer_account#read VideoIndexerAccount#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/video_indexer_account#update VideoIndexerAccount#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


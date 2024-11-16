// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchpool


type BatchPoolWindows struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.10.0/docs/resources/batch_pool#enable_automatic_updates BatchPool#enable_automatic_updates}.
	EnableAutomaticUpdates interface{} `field:"optional" json:"enableAutomaticUpdates" yaml:"enableAutomaticUpdates"`
}


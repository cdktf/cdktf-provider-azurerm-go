// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchpool


type BatchPoolUserAccountsWindowsUserConfiguration struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/batch_pool#login_mode BatchPool#login_mode}.
	LoginMode *string `field:"required" json:"loginMode" yaml:"loginMode"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package batchaccount


type BatchAccountNetworkProfile struct {
	// account_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/batch_account#account_access BatchAccount#account_access}
	AccountAccess *BatchAccountNetworkProfileAccountAccess `field:"optional" json:"accountAccess" yaml:"accountAccess"`
	// node_management_access block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.9.0/docs/resources/batch_account#node_management_access BatchAccount#node_management_access}
	NodeManagementAccess *BatchAccountNetworkProfileNodeManagementAccess `field:"optional" json:"nodeManagementAccess" yaml:"nodeManagementAccess"`
}


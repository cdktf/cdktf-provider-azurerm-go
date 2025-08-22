// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageaccount


type StorageAccountNetworkRulesPrivateLinkAccess struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/storage_account#endpoint_resource_id StorageAccount#endpoint_resource_id}.
	EndpointResourceId *string `field:"required" json:"endpointResourceId" yaml:"endpointResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.41.0/docs/resources/storage_account#endpoint_tenant_id StorageAccount#endpoint_tenant_id}.
	EndpointTenantId *string `field:"optional" json:"endpointTenantId" yaml:"endpointTenantId"`
}


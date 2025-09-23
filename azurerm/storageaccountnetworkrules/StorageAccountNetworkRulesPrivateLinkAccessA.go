// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package storageaccountnetworkrules


type StorageAccountNetworkRulesPrivateLinkAccessA struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/storage_account_network_rules#endpoint_resource_id StorageAccountNetworkRulesA#endpoint_resource_id}.
	EndpointResourceId *string `field:"required" json:"endpointResourceId" yaml:"endpointResourceId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.1/docs/resources/storage_account_network_rules#endpoint_tenant_id StorageAccountNetworkRulesA#endpoint_tenant_id}.
	EndpointTenantId *string `field:"optional" json:"endpointTenantId" yaml:"endpointTenantId"`
}


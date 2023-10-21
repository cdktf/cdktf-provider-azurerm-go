// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hpccache


type HpcCacheDirectoryActiveDirectory struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/hpc_cache#cache_netbios_name HpcCache#cache_netbios_name}.
	CacheNetbiosName *string `field:"required" json:"cacheNetbiosName" yaml:"cacheNetbiosName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/hpc_cache#dns_primary_ip HpcCache#dns_primary_ip}.
	DnsPrimaryIp *string `field:"required" json:"dnsPrimaryIp" yaml:"dnsPrimaryIp"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/hpc_cache#domain_name HpcCache#domain_name}.
	DomainName *string `field:"required" json:"domainName" yaml:"domainName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/hpc_cache#domain_netbios_name HpcCache#domain_netbios_name}.
	DomainNetbiosName *string `field:"required" json:"domainNetbiosName" yaml:"domainNetbiosName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/hpc_cache#password HpcCache#password}.
	Password *string `field:"required" json:"password" yaml:"password"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/hpc_cache#username HpcCache#username}.
	Username *string `field:"required" json:"username" yaml:"username"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.77.0/docs/resources/hpc_cache#dns_secondary_ip HpcCache#dns_secondary_ip}.
	DnsSecondaryIp *string `field:"optional" json:"dnsSecondaryIp" yaml:"dnsSecondaryIp"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package privatednszone


type PrivateDnsZoneSoaRecord struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.73.0/docs/resources/private_dns_zone#email PrivateDnsZone#email}.
	Email *string `field:"required" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.73.0/docs/resources/private_dns_zone#expire_time PrivateDnsZone#expire_time}.
	ExpireTime *float64 `field:"optional" json:"expireTime" yaml:"expireTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.73.0/docs/resources/private_dns_zone#minimum_ttl PrivateDnsZone#minimum_ttl}.
	MinimumTtl *float64 `field:"optional" json:"minimumTtl" yaml:"minimumTtl"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.73.0/docs/resources/private_dns_zone#refresh_time PrivateDnsZone#refresh_time}.
	RefreshTime *float64 `field:"optional" json:"refreshTime" yaml:"refreshTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.73.0/docs/resources/private_dns_zone#retry_time PrivateDnsZone#retry_time}.
	RetryTime *float64 `field:"optional" json:"retryTime" yaml:"retryTime"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.73.0/docs/resources/private_dns_zone#tags PrivateDnsZone#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.73.0/docs/resources/private_dns_zone#ttl PrivateDnsZone#ttl}.
	Ttl *float64 `field:"optional" json:"ttl" yaml:"ttl"`
}


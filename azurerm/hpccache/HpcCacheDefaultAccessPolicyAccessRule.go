// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hpccache


type HpcCacheDefaultAccessPolicyAccessRule struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/hpc_cache#access HpcCache#access}.
	Access *string `field:"required" json:"access" yaml:"access"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/hpc_cache#scope HpcCache#scope}.
	Scope *string `field:"required" json:"scope" yaml:"scope"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/hpc_cache#anonymous_gid HpcCache#anonymous_gid}.
	AnonymousGid *float64 `field:"optional" json:"anonymousGid" yaml:"anonymousGid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/hpc_cache#anonymous_uid HpcCache#anonymous_uid}.
	AnonymousUid *float64 `field:"optional" json:"anonymousUid" yaml:"anonymousUid"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/hpc_cache#filter HpcCache#filter}.
	Filter *string `field:"optional" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/hpc_cache#root_squash_enabled HpcCache#root_squash_enabled}.
	RootSquashEnabled interface{} `field:"optional" json:"rootSquashEnabled" yaml:"rootSquashEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/hpc_cache#submount_access_enabled HpcCache#submount_access_enabled}.
	SubmountAccessEnabled interface{} `field:"optional" json:"submountAccessEnabled" yaml:"submountAccessEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.15.0/docs/resources/hpc_cache#suid_enabled HpcCache#suid_enabled}.
	SuidEnabled interface{} `field:"optional" json:"suidEnabled" yaml:"suidEnabled"`
}


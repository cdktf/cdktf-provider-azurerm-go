// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dnscaarecord


type DnsCaaRecordRecord struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/dns_caa_record#flags DnsCaaRecord#flags}.
	Flags *float64 `field:"required" json:"flags" yaml:"flags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/dns_caa_record#tag DnsCaaRecord#tag}.
	Tag *string `field:"required" json:"tag" yaml:"tag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.116.0/docs/resources/dns_caa_record#value DnsCaaRecord#value}.
	Value *string `field:"required" json:"value" yaml:"value"`
}


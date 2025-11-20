// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package netappvolume


type NetappVolumeCoolAccess struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/netapp_volume#coolness_period_in_days NetappVolume#coolness_period_in_days}.
	CoolnessPeriodInDays *float64 `field:"required" json:"coolnessPeriodInDays" yaml:"coolnessPeriodInDays"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/netapp_volume#retrieval_policy NetappVolume#retrieval_policy}.
	RetrievalPolicy *string `field:"required" json:"retrievalPolicy" yaml:"retrievalPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.54.0/docs/resources/netapp_volume#tiering_policy NetappVolume#tiering_policy}.
	TieringPolicy *string `field:"required" json:"tieringPolicy" yaml:"tieringPolicy"`
}


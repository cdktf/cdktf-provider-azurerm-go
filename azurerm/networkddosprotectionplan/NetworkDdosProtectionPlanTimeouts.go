// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package networkddosprotectionplan


type NetworkDdosProtectionPlanTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/network_ddos_protection_plan#create NetworkDdosProtectionPlan#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/network_ddos_protection_plan#delete NetworkDdosProtectionPlan#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/network_ddos_protection_plan#read NetworkDdosProtectionPlan#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/network_ddos_protection_plan#update NetworkDdosProtectionPlan#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


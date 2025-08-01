// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managementgrouppolicyexemption


type ManagementGroupPolicyExemptionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/management_group_policy_exemption#create ManagementGroupPolicyExemption#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/management_group_policy_exemption#delete ManagementGroupPolicyExemption#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/management_group_policy_exemption#read ManagementGroupPolicyExemption#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.1/docs/resources/management_group_policy_exemption#update ManagementGroupPolicyExemption#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


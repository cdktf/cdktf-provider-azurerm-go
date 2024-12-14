// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package advisorsuppression


type AdvisorSuppressionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.14.0/docs/resources/advisor_suppression#create AdvisorSuppression#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.14.0/docs/resources/advisor_suppression#delete AdvisorSuppression#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.14.0/docs/resources/advisor_suppression#read AdvisorSuppression#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


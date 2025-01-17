// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitycenterassessment


type SecurityCenterAssessmentStatus struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/security_center_assessment#code SecurityCenterAssessment#code}.
	Code *string `field:"required" json:"code" yaml:"code"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/security_center_assessment#cause SecurityCenterAssessment#cause}.
	Cause *string `field:"optional" json:"cause" yaml:"cause"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.16.0/docs/resources/security_center_assessment#description SecurityCenterAssessment#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
}


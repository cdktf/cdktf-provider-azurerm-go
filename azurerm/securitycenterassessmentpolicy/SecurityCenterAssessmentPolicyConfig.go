// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package securitycenterassessmentpolicy

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SecurityCenterAssessmentPolicyConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/security_center_assessment_policy#description SecurityCenterAssessmentPolicy#description}.
	Description *string `field:"required" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/security_center_assessment_policy#display_name SecurityCenterAssessmentPolicy#display_name}.
	DisplayName *string `field:"required" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/security_center_assessment_policy#categories SecurityCenterAssessmentPolicy#categories}.
	Categories *[]*string `field:"optional" json:"categories" yaml:"categories"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/security_center_assessment_policy#id SecurityCenterAssessmentPolicy#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/security_center_assessment_policy#implementation_effort SecurityCenterAssessmentPolicy#implementation_effort}.
	ImplementationEffort *string `field:"optional" json:"implementationEffort" yaml:"implementationEffort"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/security_center_assessment_policy#remediation_description SecurityCenterAssessmentPolicy#remediation_description}.
	RemediationDescription *string `field:"optional" json:"remediationDescription" yaml:"remediationDescription"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/security_center_assessment_policy#severity SecurityCenterAssessmentPolicy#severity}.
	Severity *string `field:"optional" json:"severity" yaml:"severity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/security_center_assessment_policy#threats SecurityCenterAssessmentPolicy#threats}.
	Threats *[]*string `field:"optional" json:"threats" yaml:"threats"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/security_center_assessment_policy#timeouts SecurityCenterAssessmentPolicy#timeouts}
	Timeouts *SecurityCenterAssessmentPolicyTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/security_center_assessment_policy#user_impact SecurityCenterAssessmentPolicy#user_impact}.
	UserImpact *string `field:"optional" json:"userImpact" yaml:"userImpact"`
}


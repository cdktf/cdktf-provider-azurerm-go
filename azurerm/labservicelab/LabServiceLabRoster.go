// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package labservicelab


type LabServiceLabRoster struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/lab_service_lab#active_directory_group_id LabServiceLab#active_directory_group_id}.
	ActiveDirectoryGroupId *string `field:"optional" json:"activeDirectoryGroupId" yaml:"activeDirectoryGroupId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/lab_service_lab#lms_instance LabServiceLab#lms_instance}.
	LmsInstance *string `field:"optional" json:"lmsInstance" yaml:"lmsInstance"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/lab_service_lab#lti_client_id LabServiceLab#lti_client_id}.
	LtiClientId *string `field:"optional" json:"ltiClientId" yaml:"ltiClientId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/lab_service_lab#lti_context_id LabServiceLab#lti_context_id}.
	LtiContextId *string `field:"optional" json:"ltiContextId" yaml:"ltiContextId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.99.0/docs/resources/lab_service_lab#lti_roster_endpoint LabServiceLab#lti_roster_endpoint}.
	LtiRosterEndpoint *string `field:"optional" json:"ltiRosterEndpoint" yaml:"ltiRosterEndpoint"`
}


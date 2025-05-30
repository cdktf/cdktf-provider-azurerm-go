// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automationsoftwareupdateconfiguration


type AutomationSoftwareUpdateConfigurationWindows struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/automation_software_update_configuration#classifications_included AutomationSoftwareUpdateConfiguration#classifications_included}.
	ClassificationsIncluded *[]*string `field:"required" json:"classificationsIncluded" yaml:"classificationsIncluded"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/automation_software_update_configuration#excluded_knowledge_base_numbers AutomationSoftwareUpdateConfiguration#excluded_knowledge_base_numbers}.
	ExcludedKnowledgeBaseNumbers *[]*string `field:"optional" json:"excludedKnowledgeBaseNumbers" yaml:"excludedKnowledgeBaseNumbers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/automation_software_update_configuration#included_knowledge_base_numbers AutomationSoftwareUpdateConfiguration#included_knowledge_base_numbers}.
	IncludedKnowledgeBaseNumbers *[]*string `field:"optional" json:"includedKnowledgeBaseNumbers" yaml:"includedKnowledgeBaseNumbers"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.31.0/docs/resources/automation_software_update_configuration#reboot AutomationSoftwareUpdateConfiguration#reboot}.
	Reboot *string `field:"optional" json:"reboot" yaml:"reboot"`
}


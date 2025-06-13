// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationinsightsworkbooktemplate


type ApplicationInsightsWorkbookTemplateTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/application_insights_workbook_template#create ApplicationInsightsWorkbookTemplate#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/application_insights_workbook_template#delete ApplicationInsightsWorkbookTemplate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/application_insights_workbook_template#read ApplicationInsightsWorkbookTemplate#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/application_insights_workbook_template#update ApplicationInsightsWorkbookTemplate#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


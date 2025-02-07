// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationinsightsworkbooktemplate

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationInsightsWorkbookTemplateConfig struct {
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
	// galleries block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/application_insights_workbook_template#galleries ApplicationInsightsWorkbookTemplate#galleries}
	Galleries interface{} `field:"required" json:"galleries" yaml:"galleries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/application_insights_workbook_template#location ApplicationInsightsWorkbookTemplate#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/application_insights_workbook_template#name ApplicationInsightsWorkbookTemplate#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/application_insights_workbook_template#resource_group_name ApplicationInsightsWorkbookTemplate#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/application_insights_workbook_template#template_data ApplicationInsightsWorkbookTemplate#template_data}.
	TemplateData *string `field:"required" json:"templateData" yaml:"templateData"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/application_insights_workbook_template#author ApplicationInsightsWorkbookTemplate#author}.
	Author *string `field:"optional" json:"author" yaml:"author"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/application_insights_workbook_template#id ApplicationInsightsWorkbookTemplate#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/application_insights_workbook_template#localized ApplicationInsightsWorkbookTemplate#localized}.
	Localized *string `field:"optional" json:"localized" yaml:"localized"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/application_insights_workbook_template#priority ApplicationInsightsWorkbookTemplate#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/application_insights_workbook_template#tags ApplicationInsightsWorkbookTemplate#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.18.0/docs/resources/application_insights_workbook_template#timeouts ApplicationInsightsWorkbookTemplate#timeouts}
	Timeouts *ApplicationInsightsWorkbookTemplateTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


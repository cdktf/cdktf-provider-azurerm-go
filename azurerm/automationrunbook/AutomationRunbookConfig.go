// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package automationrunbook

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutomationRunbookConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/automation_runbook#automation_account_name AutomationRunbook#automation_account_name}.
	AutomationAccountName *string `field:"required" json:"automationAccountName" yaml:"automationAccountName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/automation_runbook#location AutomationRunbook#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/automation_runbook#log_progress AutomationRunbook#log_progress}.
	LogProgress interface{} `field:"required" json:"logProgress" yaml:"logProgress"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/automation_runbook#log_verbose AutomationRunbook#log_verbose}.
	LogVerbose interface{} `field:"required" json:"logVerbose" yaml:"logVerbose"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/automation_runbook#name AutomationRunbook#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/automation_runbook#resource_group_name AutomationRunbook#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/automation_runbook#runbook_type AutomationRunbook#runbook_type}.
	RunbookType *string `field:"required" json:"runbookType" yaml:"runbookType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/automation_runbook#content AutomationRunbook#content}.
	Content *string `field:"optional" json:"content" yaml:"content"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/automation_runbook#description AutomationRunbook#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// draft block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/automation_runbook#draft AutomationRunbook#draft}
	Draft *AutomationRunbookDraft `field:"optional" json:"draft" yaml:"draft"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/automation_runbook#id AutomationRunbook#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/automation_runbook#job_schedule AutomationRunbook#job_schedule}.
	JobSchedule interface{} `field:"optional" json:"jobSchedule" yaml:"jobSchedule"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/automation_runbook#log_activity_trace_level AutomationRunbook#log_activity_trace_level}.
	LogActivityTraceLevel *float64 `field:"optional" json:"logActivityTraceLevel" yaml:"logActivityTraceLevel"`
	// publish_content_link block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/automation_runbook#publish_content_link AutomationRunbook#publish_content_link}
	PublishContentLink *AutomationRunbookPublishContentLink `field:"optional" json:"publishContentLink" yaml:"publishContentLink"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/automation_runbook#tags AutomationRunbook#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.103.1/docs/resources/automation_runbook#timeouts AutomationRunbook#timeouts}
	Timeouts *AutomationRunbookTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


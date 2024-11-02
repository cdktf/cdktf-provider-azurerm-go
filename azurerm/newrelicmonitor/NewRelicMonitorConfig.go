// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package newrelicmonitor

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NewRelicMonitorConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/new_relic_monitor#location NewRelicMonitor#location}.
	Location *string `field:"required" json:"location" yaml:"location"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/new_relic_monitor#name NewRelicMonitor#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// plan block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/new_relic_monitor#plan NewRelicMonitor#plan}
	Plan *NewRelicMonitorPlan `field:"required" json:"plan" yaml:"plan"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/new_relic_monitor#resource_group_name NewRelicMonitor#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// user block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/new_relic_monitor#user NewRelicMonitor#user}
	User *NewRelicMonitorUser `field:"required" json:"user" yaml:"user"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/new_relic_monitor#account_creation_source NewRelicMonitor#account_creation_source}.
	AccountCreationSource *string `field:"optional" json:"accountCreationSource" yaml:"accountCreationSource"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/new_relic_monitor#account_id NewRelicMonitor#account_id}.
	AccountId *string `field:"optional" json:"accountId" yaml:"accountId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/new_relic_monitor#id NewRelicMonitor#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// identity block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/new_relic_monitor#identity NewRelicMonitor#identity}
	Identity *NewRelicMonitorIdentity `field:"optional" json:"identity" yaml:"identity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/new_relic_monitor#ingestion_key NewRelicMonitor#ingestion_key}.
	IngestionKey *string `field:"optional" json:"ingestionKey" yaml:"ingestionKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/new_relic_monitor#organization_id NewRelicMonitor#organization_id}.
	OrganizationId *string `field:"optional" json:"organizationId" yaml:"organizationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/new_relic_monitor#org_creation_source NewRelicMonitor#org_creation_source}.
	OrgCreationSource *string `field:"optional" json:"orgCreationSource" yaml:"orgCreationSource"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/new_relic_monitor#timeouts NewRelicMonitor#timeouts}
	Timeouts *NewRelicMonitorTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.8.0/docs/resources/new_relic_monitor#user_id NewRelicMonitor#user_id}.
	UserId *string `field:"optional" json:"userId" yaml:"userId"`
}


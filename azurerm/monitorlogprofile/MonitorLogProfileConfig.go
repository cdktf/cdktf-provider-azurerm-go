// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitorlogprofile

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorLogProfileConfig struct {
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
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/monitor_log_profile#categories MonitorLogProfile#categories}.
	Categories *[]*string `field:"required" json:"categories" yaml:"categories"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/monitor_log_profile#locations MonitorLogProfile#locations}.
	Locations *[]*string `field:"required" json:"locations" yaml:"locations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/monitor_log_profile#name MonitorLogProfile#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// retention_policy block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/monitor_log_profile#retention_policy MonitorLogProfile#retention_policy}
	RetentionPolicy *MonitorLogProfileRetentionPolicy `field:"required" json:"retentionPolicy" yaml:"retentionPolicy"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/monitor_log_profile#id MonitorLogProfile#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/monitor_log_profile#servicebus_rule_id MonitorLogProfile#servicebus_rule_id}.
	ServicebusRuleId *string `field:"optional" json:"servicebusRuleId" yaml:"servicebusRuleId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/monitor_log_profile#storage_account_id MonitorLogProfile#storage_account_id}.
	StorageAccountId *string `field:"optional" json:"storageAccountId" yaml:"storageAccountId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/monitor_log_profile#timeouts MonitorLogProfile#timeouts}
	Timeouts *MonitorLogProfileTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


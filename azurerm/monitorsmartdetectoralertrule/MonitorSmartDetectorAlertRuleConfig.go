// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitorsmartdetectoralertrule

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorSmartDetectorAlertRuleConfig struct {
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
	// action_group block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/monitor_smart_detector_alert_rule#action_group MonitorSmartDetectorAlertRule#action_group}
	ActionGroup *MonitorSmartDetectorAlertRuleActionGroup `field:"required" json:"actionGroup" yaml:"actionGroup"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/monitor_smart_detector_alert_rule#detector_type MonitorSmartDetectorAlertRule#detector_type}.
	DetectorType *string `field:"required" json:"detectorType" yaml:"detectorType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/monitor_smart_detector_alert_rule#frequency MonitorSmartDetectorAlertRule#frequency}.
	Frequency *string `field:"required" json:"frequency" yaml:"frequency"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/monitor_smart_detector_alert_rule#name MonitorSmartDetectorAlertRule#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/monitor_smart_detector_alert_rule#resource_group_name MonitorSmartDetectorAlertRule#resource_group_name}.
	ResourceGroupName *string `field:"required" json:"resourceGroupName" yaml:"resourceGroupName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/monitor_smart_detector_alert_rule#scope_resource_ids MonitorSmartDetectorAlertRule#scope_resource_ids}.
	ScopeResourceIds *[]*string `field:"required" json:"scopeResourceIds" yaml:"scopeResourceIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/monitor_smart_detector_alert_rule#severity MonitorSmartDetectorAlertRule#severity}.
	Severity *string `field:"required" json:"severity" yaml:"severity"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/monitor_smart_detector_alert_rule#description MonitorSmartDetectorAlertRule#description}.
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/monitor_smart_detector_alert_rule#enabled MonitorSmartDetectorAlertRule#enabled}.
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/monitor_smart_detector_alert_rule#id MonitorSmartDetectorAlertRule#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/monitor_smart_detector_alert_rule#tags MonitorSmartDetectorAlertRule#tags}.
	Tags *map[string]*string `field:"optional" json:"tags" yaml:"tags"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/monitor_smart_detector_alert_rule#throttling_duration MonitorSmartDetectorAlertRule#throttling_duration}.
	ThrottlingDuration *string `field:"optional" json:"throttlingDuration" yaml:"throttlingDuration"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/monitor_smart_detector_alert_rule#timeouts MonitorSmartDetectorAlertRule#timeouts}
	Timeouts *MonitorSmartDetectorAlertRuleTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


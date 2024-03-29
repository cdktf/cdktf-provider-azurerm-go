// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoractionruleactiongroup


type MonitorActionRuleActionGroupScope struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/monitor_action_rule_action_group#resource_ids MonitorActionRuleActionGroup#resource_ids}.
	ResourceIds *[]*string `field:"required" json:"resourceIds" yaml:"resourceIds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/monitor_action_rule_action_group#type MonitorActionRuleActionGroup#type}.
	Type *string `field:"required" json:"type" yaml:"type"`
}


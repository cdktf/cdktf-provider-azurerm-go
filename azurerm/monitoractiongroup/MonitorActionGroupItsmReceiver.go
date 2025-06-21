// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoractiongroup


type MonitorActionGroupItsmReceiver struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/monitor_action_group#connection_id MonitorActionGroup#connection_id}.
	ConnectionId *string `field:"required" json:"connectionId" yaml:"connectionId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/monitor_action_group#name MonitorActionGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/monitor_action_group#region MonitorActionGroup#region}.
	Region *string `field:"required" json:"region" yaml:"region"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/monitor_action_group#ticket_configuration MonitorActionGroup#ticket_configuration}.
	TicketConfiguration *string `field:"required" json:"ticketConfiguration" yaml:"ticketConfiguration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/monitor_action_group#workspace_id MonitorActionGroup#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
}


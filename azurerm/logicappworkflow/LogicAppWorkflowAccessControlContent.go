// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logicappworkflow


type LogicAppWorkflowAccessControlContent struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/logic_app_workflow#allowed_caller_ip_address_range LogicAppWorkflow#allowed_caller_ip_address_range}.
	AllowedCallerIpAddressRange *[]*string `field:"required" json:"allowedCallerIpAddressRange" yaml:"allowedCallerIpAddressRange"`
}


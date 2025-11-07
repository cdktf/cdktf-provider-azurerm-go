// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iotsecuritysolution


type IotSecuritySolutionAdditionalWorkspace struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/iot_security_solution#data_types IotSecuritySolution#data_types}.
	DataTypes *[]*string `field:"required" json:"dataTypes" yaml:"dataTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.52.0/docs/resources/iot_security_solution#workspace_id IotSecuritySolution#workspace_id}.
	WorkspaceId *string `field:"required" json:"workspaceId" yaml:"workspaceId"`
}


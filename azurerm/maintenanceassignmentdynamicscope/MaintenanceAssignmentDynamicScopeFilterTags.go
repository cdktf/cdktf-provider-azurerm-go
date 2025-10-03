// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package maintenanceassignmentdynamicscope


type MaintenanceAssignmentDynamicScopeFilterTags struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/maintenance_assignment_dynamic_scope#tag MaintenanceAssignmentDynamicScope#tag}.
	Tag *string `field:"required" json:"tag" yaml:"tag"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.47.0/docs/resources/maintenance_assignment_dynamic_scope#values MaintenanceAssignmentDynamicScope#values}.
	Values *[]*string `field:"required" json:"values" yaml:"values"`
}


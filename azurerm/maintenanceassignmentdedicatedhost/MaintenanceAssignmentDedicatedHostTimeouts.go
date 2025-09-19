// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package maintenanceassignmentdedicatedhost


type MaintenanceAssignmentDedicatedHostTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/maintenance_assignment_dedicated_host#create MaintenanceAssignmentDedicatedHost#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/maintenance_assignment_dedicated_host#delete MaintenanceAssignmentDedicatedHost#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/maintenance_assignment_dedicated_host#read MaintenanceAssignmentDedicatedHost#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


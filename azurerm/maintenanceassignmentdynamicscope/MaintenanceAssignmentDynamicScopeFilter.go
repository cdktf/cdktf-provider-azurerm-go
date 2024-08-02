// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package maintenanceassignmentdynamicscope


type MaintenanceAssignmentDynamicScopeFilter struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/maintenance_assignment_dynamic_scope#locations MaintenanceAssignmentDynamicScope#locations}.
	Locations *[]*string `field:"optional" json:"locations" yaml:"locations"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/maintenance_assignment_dynamic_scope#os_types MaintenanceAssignmentDynamicScope#os_types}.
	OsTypes *[]*string `field:"optional" json:"osTypes" yaml:"osTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/maintenance_assignment_dynamic_scope#resource_groups MaintenanceAssignmentDynamicScope#resource_groups}.
	ResourceGroups *[]*string `field:"optional" json:"resourceGroups" yaml:"resourceGroups"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/maintenance_assignment_dynamic_scope#resource_types MaintenanceAssignmentDynamicScope#resource_types}.
	ResourceTypes *[]*string `field:"optional" json:"resourceTypes" yaml:"resourceTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/maintenance_assignment_dynamic_scope#tag_filter MaintenanceAssignmentDynamicScope#tag_filter}.
	TagFilter *string `field:"optional" json:"tagFilter" yaml:"tagFilter"`
	// tags block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.114.0/docs/resources/maintenance_assignment_dynamic_scope#tags MaintenanceAssignmentDynamicScope#tags}
	Tags interface{} `field:"optional" json:"tags" yaml:"tags"`
}


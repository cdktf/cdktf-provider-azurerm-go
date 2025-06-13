// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package maintenanceassignmentdynamicscope

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MaintenanceAssignmentDynamicScopeConfig struct {
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
	// filter block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/maintenance_assignment_dynamic_scope#filter MaintenanceAssignmentDynamicScope#filter}
	Filter *MaintenanceAssignmentDynamicScopeFilter `field:"required" json:"filter" yaml:"filter"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/maintenance_assignment_dynamic_scope#maintenance_configuration_id MaintenanceAssignmentDynamicScope#maintenance_configuration_id}.
	MaintenanceConfigurationId *string `field:"required" json:"maintenanceConfigurationId" yaml:"maintenanceConfigurationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/maintenance_assignment_dynamic_scope#name MaintenanceAssignmentDynamicScope#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/maintenance_assignment_dynamic_scope#id MaintenanceAssignmentDynamicScope#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.33.0/docs/resources/maintenance_assignment_dynamic_scope#timeouts MaintenanceAssignmentDynamicScope#timeouts}
	Timeouts *MaintenanceAssignmentDynamicScopeTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


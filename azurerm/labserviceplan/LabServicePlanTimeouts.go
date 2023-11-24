// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package labserviceplan


type LabServicePlanTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/lab_service_plan#create LabServicePlan#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/lab_service_plan#delete LabServicePlan#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/lab_service_plan#read LabServicePlan#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.82.0/docs/resources/lab_service_plan#update LabServicePlan#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


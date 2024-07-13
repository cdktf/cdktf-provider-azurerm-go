// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package labserviceplan


type LabServicePlanSupport struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/lab_service_plan#email LabServicePlan#email}.
	Email *string `field:"optional" json:"email" yaml:"email"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/lab_service_plan#instructions LabServicePlan#instructions}.
	Instructions *string `field:"optional" json:"instructions" yaml:"instructions"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/lab_service_plan#phone LabServicePlan#phone}.
	Phone *string `field:"optional" json:"phone" yaml:"phone"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.112.0/docs/resources/lab_service_plan#url LabServicePlan#url}.
	Url *string `field:"optional" json:"url" yaml:"url"`
}


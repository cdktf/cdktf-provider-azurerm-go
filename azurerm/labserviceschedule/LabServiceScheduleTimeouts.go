// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package labserviceschedule


type LabServiceScheduleTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/lab_service_schedule#create LabServiceSchedule#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/lab_service_schedule#delete LabServiceSchedule#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/lab_service_schedule#read LabServiceSchedule#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.93.0/docs/resources/lab_service_schedule#update LabServiceSchedule#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


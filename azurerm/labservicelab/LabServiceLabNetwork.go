// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package labservicelab


type LabServiceLabNetwork struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.97.1/docs/resources/lab_service_lab#subnet_id LabServiceLab#subnet_id}.
	SubnetId *string `field:"optional" json:"subnetId" yaml:"subnetId"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iotsecuritysolution


type IotSecuritySolutionTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/iot_security_solution#create IotSecuritySolution#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/iot_security_solution#delete IotSecuritySolution#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/iot_security_solution#read IotSecuritySolution#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.80.0/docs/resources/iot_security_solution#update IotSecuritySolution#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


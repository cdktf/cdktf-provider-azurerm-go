// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package devcenterprojectenvironmenttype


type DevCenterProjectEnvironmentTypeUserRoleAssignment struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/dev_center_project_environment_type#roles DevCenterProjectEnvironmentType#roles}.
	Roles *[]*string `field:"required" json:"roles" yaml:"roles"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.0.1/docs/resources/dev_center_project_environment_type#user_id DevCenterProjectEnvironmentType#user_id}.
	UserId *string `field:"required" json:"userId" yaml:"userId"`
}


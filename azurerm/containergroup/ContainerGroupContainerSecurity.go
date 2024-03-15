// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containergroup


type ContainerGroupContainerSecurity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.96.0/docs/resources/container_group#privilege_enabled ContainerGroup#privilege_enabled}.
	PrivilegeEnabled interface{} `field:"required" json:"privilegeEnabled" yaml:"privilegeEnabled"`
}


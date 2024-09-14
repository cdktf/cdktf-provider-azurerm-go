// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerregistrytask


type ContainerRegistryTaskRegistryCredentialSource struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.2.0/docs/resources/container_registry_task#login_mode ContainerRegistryTask#login_mode}.
	LoginMode *string `field:"required" json:"loginMode" yaml:"loginMode"`
}


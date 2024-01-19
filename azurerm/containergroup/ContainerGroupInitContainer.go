// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containergroup


type ContainerGroupInitContainer struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/container_group#image ContainerGroup#image}.
	Image *string `field:"required" json:"image" yaml:"image"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/container_group#name ContainerGroup#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/container_group#commands ContainerGroup#commands}.
	Commands *[]*string `field:"optional" json:"commands" yaml:"commands"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/container_group#environment_variables ContainerGroup#environment_variables}.
	EnvironmentVariables *map[string]*string `field:"optional" json:"environmentVariables" yaml:"environmentVariables"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/container_group#secure_environment_variables ContainerGroup#secure_environment_variables}.
	SecureEnvironmentVariables *map[string]*string `field:"optional" json:"secureEnvironmentVariables" yaml:"secureEnvironmentVariables"`
	// security block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/container_group#security ContainerGroup#security}
	Security interface{} `field:"optional" json:"security" yaml:"security"`
	// volume block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.88.0/docs/resources/container_group#volume ContainerGroup#volume}
	Volume interface{} `field:"optional" json:"volume" yaml:"volume"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerappenvironmentdaprcomponent

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ContainerAppEnvironmentDaprComponentConfig struct {
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
	// The Dapr Component Type. For example `state.azure.blobstorage`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/container_app_environment_dapr_component#component_type ContainerAppEnvironmentDaprComponent#component_type}
	ComponentType *string `field:"required" json:"componentType" yaml:"componentType"`
	// The Container App Managed Environment ID to configure this Dapr component on.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/container_app_environment_dapr_component#container_app_environment_id ContainerAppEnvironmentDaprComponent#container_app_environment_id}
	ContainerAppEnvironmentId *string `field:"required" json:"containerAppEnvironmentId" yaml:"containerAppEnvironmentId"`
	// The name for this Dapr Component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/container_app_environment_dapr_component#name ContainerAppEnvironmentDaprComponent#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The version of the component.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/container_app_environment_dapr_component#version ContainerAppEnvironmentDaprComponent#version}
	Version *string `field:"required" json:"version" yaml:"version"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/container_app_environment_dapr_component#id ContainerAppEnvironmentDaprComponent#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Should the Dapr sidecar to continue initialisation if the component fails to load. Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/container_app_environment_dapr_component#ignore_errors ContainerAppEnvironmentDaprComponent#ignore_errors}
	IgnoreErrors interface{} `field:"optional" json:"ignoreErrors" yaml:"ignoreErrors"`
	// The component initialisation timeout in ISO8601 format. e.g. `5s`, `2h`, `1m`. Defaults to `5s`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/container_app_environment_dapr_component#init_timeout ContainerAppEnvironmentDaprComponent#init_timeout}
	InitTimeout *string `field:"optional" json:"initTimeout" yaml:"initTimeout"`
	// metadata block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/container_app_environment_dapr_component#metadata ContainerAppEnvironmentDaprComponent#metadata}
	Metadata interface{} `field:"optional" json:"metadata" yaml:"metadata"`
	// A list of scopes to which this component applies. e.g. a Container App's `dapr.app_id` value.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/container_app_environment_dapr_component#scopes ContainerAppEnvironmentDaprComponent#scopes}
	Scopes *[]*string `field:"optional" json:"scopes" yaml:"scopes"`
	// secret block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/container_app_environment_dapr_component#secret ContainerAppEnvironmentDaprComponent#secret}
	Secret interface{} `field:"optional" json:"secret" yaml:"secret"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/container_app_environment_dapr_component#timeouts ContainerAppEnvironmentDaprComponent#timeouts}
	Timeouts *ContainerAppEnvironmentDaprComponentTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
}


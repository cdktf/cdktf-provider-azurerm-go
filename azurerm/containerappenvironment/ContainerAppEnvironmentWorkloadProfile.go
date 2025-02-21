// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerappenvironment


type ContainerAppEnvironmentWorkloadProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/container_app_environment#name ContainerAppEnvironment#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/container_app_environment#workload_profile_type ContainerAppEnvironment#workload_profile_type}.
	WorkloadProfileType *string `field:"required" json:"workloadProfileType" yaml:"workloadProfileType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/container_app_environment#maximum_count ContainerAppEnvironment#maximum_count}.
	MaximumCount *float64 `field:"optional" json:"maximumCount" yaml:"maximumCount"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.20.0/docs/resources/container_app_environment#minimum_count ContainerAppEnvironment#minimum_count}.
	MinimumCount *float64 `field:"optional" json:"minimumCount" yaml:"minimumCount"`
}


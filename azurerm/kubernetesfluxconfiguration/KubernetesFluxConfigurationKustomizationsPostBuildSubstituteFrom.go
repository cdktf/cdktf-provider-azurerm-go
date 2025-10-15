// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesfluxconfiguration


type KubernetesFluxConfigurationKustomizationsPostBuildSubstituteFrom struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/kubernetes_flux_configuration#kind KubernetesFluxConfiguration#kind}.
	Kind *string `field:"required" json:"kind" yaml:"kind"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/kubernetes_flux_configuration#name KubernetesFluxConfiguration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.48.0/docs/resources/kubernetes_flux_configuration#optional KubernetesFluxConfiguration#optional}.
	Optional interface{} `field:"optional" json:"optional" yaml:"optional"`
}


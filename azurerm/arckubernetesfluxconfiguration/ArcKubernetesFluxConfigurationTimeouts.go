// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package arckubernetesfluxconfiguration


type ArcKubernetesFluxConfigurationTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/arc_kubernetes_flux_configuration#create ArcKubernetesFluxConfiguration#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/arc_kubernetes_flux_configuration#delete ArcKubernetesFluxConfiguration#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/arc_kubernetes_flux_configuration#read ArcKubernetesFluxConfiguration#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/arc_kubernetes_flux_configuration#update ArcKubernetesFluxConfiguration#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}


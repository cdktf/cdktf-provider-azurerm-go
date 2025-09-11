// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesfluxconfiguration


type KubernetesFluxConfigurationKustomizationsPostBuild struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/kubernetes_flux_configuration#substitute KubernetesFluxConfiguration#substitute}.
	Substitute *map[string]*string `field:"optional" json:"substitute" yaml:"substitute"`
	// substitute_from block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/kubernetes_flux_configuration#substitute_from KubernetesFluxConfiguration#substitute_from}
	SubstituteFrom interface{} `field:"optional" json:"substituteFrom" yaml:"substituteFrom"`
}


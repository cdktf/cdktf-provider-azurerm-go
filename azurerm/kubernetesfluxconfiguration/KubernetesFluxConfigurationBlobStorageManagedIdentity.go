// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesfluxconfiguration


type KubernetesFluxConfigurationBlobStorageManagedIdentity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.90.0/docs/resources/kubernetes_flux_configuration#client_id KubernetesFluxConfiguration#client_id}.
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
}


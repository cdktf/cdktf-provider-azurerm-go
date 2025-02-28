// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package arckubernetesfluxconfiguration


type ArcKubernetesFluxConfigurationKustomizations struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/arc_kubernetes_flux_configuration#name ArcKubernetesFluxConfiguration#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/arc_kubernetes_flux_configuration#depends_on ArcKubernetesFluxConfiguration#depends_on}.
	DependsOn *[]*string `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/arc_kubernetes_flux_configuration#garbage_collection_enabled ArcKubernetesFluxConfiguration#garbage_collection_enabled}.
	GarbageCollectionEnabled interface{} `field:"optional" json:"garbageCollectionEnabled" yaml:"garbageCollectionEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/arc_kubernetes_flux_configuration#path ArcKubernetesFluxConfiguration#path}.
	Path *string `field:"optional" json:"path" yaml:"path"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/arc_kubernetes_flux_configuration#recreating_enabled ArcKubernetesFluxConfiguration#recreating_enabled}.
	RecreatingEnabled interface{} `field:"optional" json:"recreatingEnabled" yaml:"recreatingEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/arc_kubernetes_flux_configuration#retry_interval_in_seconds ArcKubernetesFluxConfiguration#retry_interval_in_seconds}.
	RetryIntervalInSeconds *float64 `field:"optional" json:"retryIntervalInSeconds" yaml:"retryIntervalInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/arc_kubernetes_flux_configuration#sync_interval_in_seconds ArcKubernetesFluxConfiguration#sync_interval_in_seconds}.
	SyncIntervalInSeconds *float64 `field:"optional" json:"syncIntervalInSeconds" yaml:"syncIntervalInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs/resources/arc_kubernetes_flux_configuration#timeout_in_seconds ArcKubernetesFluxConfiguration#timeout_in_seconds}.
	TimeoutInSeconds *float64 `field:"optional" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package arckubernetesfluxconfiguration


type ArcKubernetesFluxConfigurationBlobStorage struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/arc_kubernetes_flux_configuration#container_id ArcKubernetesFluxConfiguration#container_id}.
	ContainerId *string `field:"required" json:"containerId" yaml:"containerId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/arc_kubernetes_flux_configuration#account_key ArcKubernetesFluxConfiguration#account_key}.
	AccountKey *string `field:"optional" json:"accountKey" yaml:"accountKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/arc_kubernetes_flux_configuration#local_auth_reference ArcKubernetesFluxConfiguration#local_auth_reference}.
	LocalAuthReference *string `field:"optional" json:"localAuthReference" yaml:"localAuthReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/arc_kubernetes_flux_configuration#sas_token ArcKubernetesFluxConfiguration#sas_token}.
	SasToken *string `field:"optional" json:"sasToken" yaml:"sasToken"`
	// service_principal block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/arc_kubernetes_flux_configuration#service_principal ArcKubernetesFluxConfiguration#service_principal}
	ServicePrincipal *ArcKubernetesFluxConfigurationBlobStorageServicePrincipal `field:"optional" json:"servicePrincipal" yaml:"servicePrincipal"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/arc_kubernetes_flux_configuration#sync_interval_in_seconds ArcKubernetesFluxConfiguration#sync_interval_in_seconds}.
	SyncIntervalInSeconds *float64 `field:"optional" json:"syncIntervalInSeconds" yaml:"syncIntervalInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.87.0/docs/resources/arc_kubernetes_flux_configuration#timeout_in_seconds ArcKubernetesFluxConfiguration#timeout_in_seconds}.
	TimeoutInSeconds *float64 `field:"optional" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
}


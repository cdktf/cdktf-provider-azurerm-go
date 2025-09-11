// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package arckubernetesfluxconfiguration


type ArcKubernetesFluxConfigurationBucket struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/arc_kubernetes_flux_configuration#bucket_name ArcKubernetesFluxConfiguration#bucket_name}.
	BucketName *string `field:"required" json:"bucketName" yaml:"bucketName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/arc_kubernetes_flux_configuration#url ArcKubernetesFluxConfiguration#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/arc_kubernetes_flux_configuration#access_key ArcKubernetesFluxConfiguration#access_key}.
	AccessKey *string `field:"optional" json:"accessKey" yaml:"accessKey"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/arc_kubernetes_flux_configuration#local_auth_reference ArcKubernetesFluxConfiguration#local_auth_reference}.
	LocalAuthReference *string `field:"optional" json:"localAuthReference" yaml:"localAuthReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/arc_kubernetes_flux_configuration#secret_key_base64 ArcKubernetesFluxConfiguration#secret_key_base64}.
	SecretKeyBase64 *string `field:"optional" json:"secretKeyBase64" yaml:"secretKeyBase64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/arc_kubernetes_flux_configuration#sync_interval_in_seconds ArcKubernetesFluxConfiguration#sync_interval_in_seconds}.
	SyncIntervalInSeconds *float64 `field:"optional" json:"syncIntervalInSeconds" yaml:"syncIntervalInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/arc_kubernetes_flux_configuration#timeout_in_seconds ArcKubernetesFluxConfiguration#timeout_in_seconds}.
	TimeoutInSeconds *float64 `field:"optional" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.44.0/docs/resources/arc_kubernetes_flux_configuration#tls_enabled ArcKubernetesFluxConfiguration#tls_enabled}.
	TlsEnabled interface{} `field:"optional" json:"tlsEnabled" yaml:"tlsEnabled"`
}


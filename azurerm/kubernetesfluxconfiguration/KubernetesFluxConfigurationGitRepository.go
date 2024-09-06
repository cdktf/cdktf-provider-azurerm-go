// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetesfluxconfiguration


type KubernetesFluxConfigurationGitRepository struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/kubernetes_flux_configuration#reference_type KubernetesFluxConfiguration#reference_type}.
	ReferenceType *string `field:"required" json:"referenceType" yaml:"referenceType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/kubernetes_flux_configuration#reference_value KubernetesFluxConfiguration#reference_value}.
	ReferenceValue *string `field:"required" json:"referenceValue" yaml:"referenceValue"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/kubernetes_flux_configuration#url KubernetesFluxConfiguration#url}.
	Url *string `field:"required" json:"url" yaml:"url"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/kubernetes_flux_configuration#https_ca_cert_base64 KubernetesFluxConfiguration#https_ca_cert_base64}.
	HttpsCaCertBase64 *string `field:"optional" json:"httpsCaCertBase64" yaml:"httpsCaCertBase64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/kubernetes_flux_configuration#https_key_base64 KubernetesFluxConfiguration#https_key_base64}.
	HttpsKeyBase64 *string `field:"optional" json:"httpsKeyBase64" yaml:"httpsKeyBase64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/kubernetes_flux_configuration#https_user KubernetesFluxConfiguration#https_user}.
	HttpsUser *string `field:"optional" json:"httpsUser" yaml:"httpsUser"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/kubernetes_flux_configuration#local_auth_reference KubernetesFluxConfiguration#local_auth_reference}.
	LocalAuthReference *string `field:"optional" json:"localAuthReference" yaml:"localAuthReference"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/kubernetes_flux_configuration#ssh_known_hosts_base64 KubernetesFluxConfiguration#ssh_known_hosts_base64}.
	SshKnownHostsBase64 *string `field:"optional" json:"sshKnownHostsBase64" yaml:"sshKnownHostsBase64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/kubernetes_flux_configuration#ssh_private_key_base64 KubernetesFluxConfiguration#ssh_private_key_base64}.
	SshPrivateKeyBase64 *string `field:"optional" json:"sshPrivateKeyBase64" yaml:"sshPrivateKeyBase64"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/kubernetes_flux_configuration#sync_interval_in_seconds KubernetesFluxConfiguration#sync_interval_in_seconds}.
	SyncIntervalInSeconds *float64 `field:"optional" json:"syncIntervalInSeconds" yaml:"syncIntervalInSeconds"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.1.0/docs/resources/kubernetes_flux_configuration#timeout_in_seconds KubernetesFluxConfiguration#timeout_in_seconds}.
	TimeoutInSeconds *float64 `field:"optional" json:"timeoutInSeconds" yaml:"timeoutInSeconds"`
}


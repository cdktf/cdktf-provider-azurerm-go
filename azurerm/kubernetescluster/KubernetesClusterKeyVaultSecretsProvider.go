// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterKeyVaultSecretsProvider struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/kubernetes_cluster#secret_rotation_enabled KubernetesCluster#secret_rotation_enabled}.
	SecretRotationEnabled interface{} `field:"optional" json:"secretRotationEnabled" yaml:"secretRotationEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.79.0/docs/resources/kubernetes_cluster#secret_rotation_interval KubernetesCluster#secret_rotation_interval}.
	SecretRotationInterval *string `field:"optional" json:"secretRotationInterval" yaml:"secretRotationInterval"`
}


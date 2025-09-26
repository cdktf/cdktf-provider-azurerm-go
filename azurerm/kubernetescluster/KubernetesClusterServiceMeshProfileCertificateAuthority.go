// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterServiceMeshProfileCertificateAuthority struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/kubernetes_cluster#cert_chain_object_name KubernetesCluster#cert_chain_object_name}.
	CertChainObjectName *string `field:"required" json:"certChainObjectName" yaml:"certChainObjectName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/kubernetes_cluster#cert_object_name KubernetesCluster#cert_object_name}.
	CertObjectName *string `field:"required" json:"certObjectName" yaml:"certObjectName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/kubernetes_cluster#key_object_name KubernetesCluster#key_object_name}.
	KeyObjectName *string `field:"required" json:"keyObjectName" yaml:"keyObjectName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/kubernetes_cluster#key_vault_id KubernetesCluster#key_vault_id}.
	KeyVaultId *string `field:"required" json:"keyVaultId" yaml:"keyVaultId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.46.0/docs/resources/kubernetes_cluster#root_cert_object_name KubernetesCluster#root_cert_object_name}.
	RootCertObjectName *string `field:"required" json:"rootCertObjectName" yaml:"rootCertObjectName"`
}


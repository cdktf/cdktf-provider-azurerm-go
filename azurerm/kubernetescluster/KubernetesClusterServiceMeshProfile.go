// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetescluster


type KubernetesClusterServiceMeshProfile struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/kubernetes_cluster#mode KubernetesCluster#mode}.
	Mode *string `field:"required" json:"mode" yaml:"mode"`
	// certificate_authority block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/kubernetes_cluster#certificate_authority KubernetesCluster#certificate_authority}
	CertificateAuthority *KubernetesClusterServiceMeshProfileCertificateAuthority `field:"optional" json:"certificateAuthority" yaml:"certificateAuthority"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/kubernetes_cluster#external_ingress_gateway_enabled KubernetesCluster#external_ingress_gateway_enabled}.
	ExternalIngressGatewayEnabled interface{} `field:"optional" json:"externalIngressGatewayEnabled" yaml:"externalIngressGatewayEnabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.115.0/docs/resources/kubernetes_cluster#internal_ingress_gateway_enabled KubernetesCluster#internal_ingress_gateway_enabled}.
	InternalIngressGatewayEnabled interface{} `field:"optional" json:"internalIngressGatewayEnabled" yaml:"internalIngressGatewayEnabled"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicefabriccluster


type ServiceFabricClusterClientCertificateCommonName struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.73.0/docs/resources/service_fabric_cluster#common_name ServiceFabricCluster#common_name}.
	CommonName *string `field:"required" json:"commonName" yaml:"commonName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.73.0/docs/resources/service_fabric_cluster#is_admin ServiceFabricCluster#is_admin}.
	IsAdmin interface{} `field:"required" json:"isAdmin" yaml:"isAdmin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.73.0/docs/resources/service_fabric_cluster#issuer_thumbprint ServiceFabricCluster#issuer_thumbprint}.
	IssuerThumbprint *string `field:"optional" json:"issuerThumbprint" yaml:"issuerThumbprint"`
}


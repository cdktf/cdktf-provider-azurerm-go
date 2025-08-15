// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicefabriccluster


type ServiceFabricClusterClientCertificateThumbprint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/service_fabric_cluster#is_admin ServiceFabricCluster#is_admin}.
	IsAdmin interface{} `field:"required" json:"isAdmin" yaml:"isAdmin"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.40.0/docs/resources/service_fabric_cluster#thumbprint ServiceFabricCluster#thumbprint}.
	Thumbprint *string `field:"required" json:"thumbprint" yaml:"thumbprint"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package servicefabriccluster


type ServiceFabricClusterCertificateCommonNames struct {
	// common_names block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/service_fabric_cluster#common_names ServiceFabricCluster#common_names}
	CommonNames interface{} `field:"required" json:"commonNames" yaml:"commonNames"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.43.0/docs/resources/service_fabric_cluster#x509_store_name ServiceFabricCluster#x509_store_name}.
	X509StoreName *string `field:"required" json:"x509StoreName" yaml:"x509StoreName"`
}


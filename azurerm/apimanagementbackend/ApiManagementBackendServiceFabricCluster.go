// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementbackend


type ApiManagementBackendServiceFabricCluster struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.0/docs/resources/api_management_backend#management_endpoints ApiManagementBackend#management_endpoints}.
	ManagementEndpoints *[]*string `field:"required" json:"managementEndpoints" yaml:"managementEndpoints"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.0/docs/resources/api_management_backend#max_partition_resolution_retries ApiManagementBackend#max_partition_resolution_retries}.
	MaxPartitionResolutionRetries *float64 `field:"required" json:"maxPartitionResolutionRetries" yaml:"maxPartitionResolutionRetries"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.0/docs/resources/api_management_backend#client_certificate_id ApiManagementBackend#client_certificate_id}.
	ClientCertificateId *string `field:"optional" json:"clientCertificateId" yaml:"clientCertificateId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.0/docs/resources/api_management_backend#client_certificate_thumbprint ApiManagementBackend#client_certificate_thumbprint}.
	ClientCertificateThumbprint *string `field:"optional" json:"clientCertificateThumbprint" yaml:"clientCertificateThumbprint"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.0/docs/resources/api_management_backend#server_certificate_thumbprints ApiManagementBackend#server_certificate_thumbprints}.
	ServerCertificateThumbprints *[]*string `field:"optional" json:"serverCertificateThumbprints" yaml:"serverCertificateThumbprints"`
	// server_x509_name block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.38.0/docs/resources/api_management_backend#server_x509_name ApiManagementBackend#server_x509_name}
	ServerX509Name interface{} `field:"optional" json:"serverX509Name" yaml:"serverX509Name"`
}


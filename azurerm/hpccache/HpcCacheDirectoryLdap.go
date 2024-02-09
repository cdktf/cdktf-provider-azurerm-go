// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package hpccache


type HpcCacheDirectoryLdap struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/hpc_cache#base_dn HpcCache#base_dn}.
	BaseDn *string `field:"required" json:"baseDn" yaml:"baseDn"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/hpc_cache#server HpcCache#server}.
	Server *string `field:"required" json:"server" yaml:"server"`
	// bind block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/hpc_cache#bind HpcCache#bind}
	Bind *HpcCacheDirectoryLdapBind `field:"optional" json:"bind" yaml:"bind"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/hpc_cache#certificate_validation_uri HpcCache#certificate_validation_uri}.
	CertificateValidationUri *string `field:"optional" json:"certificateValidationUri" yaml:"certificateValidationUri"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/hpc_cache#download_certificate_automatically HpcCache#download_certificate_automatically}.
	DownloadCertificateAutomatically interface{} `field:"optional" json:"downloadCertificateAutomatically" yaml:"downloadCertificateAutomatically"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.91.0/docs/resources/hpc_cache#encrypted HpcCache#encrypted}.
	Encrypted interface{} `field:"optional" json:"encrypted" yaml:"encrypted"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package containerapp


type ContainerAppIngressCustomDomain struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.113.0/docs/resources/container_app#certificate_id ContainerApp#certificate_id}.
	CertificateId *string `field:"required" json:"certificateId" yaml:"certificateId"`
	// The hostname of the Certificate. Must be the CN or a named SAN in the certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.113.0/docs/resources/container_app#name ContainerApp#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The Binding type. Possible values include `Disabled` and `SniEnabled`. Defaults to `Disabled`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.113.0/docs/resources/container_app#certificate_binding_type ContainerApp#certificate_binding_type}
	CertificateBindingType *string `field:"optional" json:"certificateBindingType" yaml:"certificateBindingType"`
}


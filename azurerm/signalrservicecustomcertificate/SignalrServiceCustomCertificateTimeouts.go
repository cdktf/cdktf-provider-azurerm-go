// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package signalrservicecustomcertificate


type SignalrServiceCustomCertificateTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/signalr_service_custom_certificate#create SignalrServiceCustomCertificate#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/signalr_service_custom_certificate#delete SignalrServiceCustomCertificate#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.42.0/docs/resources/signalr_service_custom_certificate#read SignalrServiceCustomCertificate#read}.
	Read *string `field:"optional" json:"read" yaml:"read"`
}


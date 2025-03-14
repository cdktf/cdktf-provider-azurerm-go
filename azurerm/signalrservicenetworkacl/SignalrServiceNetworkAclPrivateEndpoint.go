// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package signalrservicenetworkacl


type SignalrServiceNetworkAclPrivateEndpoint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/signalr_service_network_acl#id SignalrServiceNetworkAcl#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/signalr_service_network_acl#allowed_request_types SignalrServiceNetworkAcl#allowed_request_types}.
	AllowedRequestTypes *[]*string `field:"optional" json:"allowedRequestTypes" yaml:"allowedRequestTypes"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.23.0/docs/resources/signalr_service_network_acl#denied_request_types SignalrServiceNetworkAcl#denied_request_types}.
	DeniedRequestTypes *[]*string `field:"optional" json:"deniedRequestTypes" yaml:"deniedRequestTypes"`
}


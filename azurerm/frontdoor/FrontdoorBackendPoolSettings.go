// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package frontdoor


type FrontdoorBackendPoolSettings struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/frontdoor#enforce_backend_pools_certificate_name_check Frontdoor#enforce_backend_pools_certificate_name_check}.
	EnforceBackendPoolsCertificateNameCheck interface{} `field:"required" json:"enforceBackendPoolsCertificateNameCheck" yaml:"enforceBackendPoolsCertificateNameCheck"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.104.0/docs/resources/frontdoor#backend_pools_send_receive_timeout_seconds Frontdoor#backend_pools_send_receive_timeout_seconds}.
	BackendPoolsSendReceiveTimeoutSeconds *float64 `field:"optional" json:"backendPoolsSendReceiveTimeoutSeconds" yaml:"backendPoolsSendReceiveTimeoutSeconds"`
}


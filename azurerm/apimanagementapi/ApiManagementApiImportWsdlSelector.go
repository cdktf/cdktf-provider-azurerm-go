// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementapi


type ApiManagementApiImportWsdlSelector struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/api_management_api#endpoint_name ApiManagementApi#endpoint_name}.
	EndpointName *string `field:"required" json:"endpointName" yaml:"endpointName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.30.0/docs/resources/api_management_api#service_name ApiManagementApi#service_name}.
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementbackend


type ApiManagementBackendCircuitBreakerRule struct {
	// failure_condition block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/api_management_backend#failure_condition ApiManagementBackend#failure_condition}
	FailureCondition *ApiManagementBackendCircuitBreakerRuleFailureCondition `field:"required" json:"failureCondition" yaml:"failureCondition"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/api_management_backend#name ApiManagementBackend#name}.
	Name *string `field:"required" json:"name" yaml:"name"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/api_management_backend#trip_duration ApiManagementBackend#trip_duration}.
	TripDuration *string `field:"required" json:"tripDuration" yaml:"tripDuration"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.53.0/docs/resources/api_management_backend#accept_retry_after_enabled ApiManagementBackend#accept_retry_after_enabled}.
	AcceptRetryAfterEnabled interface{} `field:"optional" json:"acceptRetryAfterEnabled" yaml:"acceptRetryAfterEnabled"`
}


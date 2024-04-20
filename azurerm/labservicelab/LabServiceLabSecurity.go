// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package labservicelab


type LabServiceLabSecurity struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.100.0/docs/resources/lab_service_lab#open_access_enabled LabServiceLab#open_access_enabled}.
	OpenAccessEnabled interface{} `field:"required" json:"openAccessEnabled" yaml:"openAccessEnabled"`
}


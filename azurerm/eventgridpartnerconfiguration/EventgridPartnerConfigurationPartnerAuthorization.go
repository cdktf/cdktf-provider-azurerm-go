// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgridpartnerconfiguration


type EventgridPartnerConfigurationPartnerAuthorization struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/eventgrid_partner_configuration#partner_name EventgridPartnerConfiguration#partner_name}.
	PartnerName *string `field:"required" json:"partnerName" yaml:"partnerName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/eventgrid_partner_configuration#partner_registration_id EventgridPartnerConfiguration#partner_registration_id}.
	PartnerRegistrationId *string `field:"required" json:"partnerRegistrationId" yaml:"partnerRegistrationId"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.34.0/docs/resources/eventgrid_partner_configuration#authorization_expiration_time_in_utc EventgridPartnerConfiguration#authorization_expiration_time_in_utc}.
	AuthorizationExpirationTimeInUtc *string `field:"optional" json:"authorizationExpirationTimeInUtc" yaml:"authorizationExpirationTimeInUtc"`
}


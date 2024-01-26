// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package orbitalcontactprofile


type OrbitalContactProfileLinksChannelsEndPoint struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.89.0/docs/resources/orbital_contact_profile#end_point_name OrbitalContactProfile#end_point_name}.
	EndPointName *string `field:"required" json:"endPointName" yaml:"endPointName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.89.0/docs/resources/orbital_contact_profile#port OrbitalContactProfile#port}.
	Port *string `field:"required" json:"port" yaml:"port"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.89.0/docs/resources/orbital_contact_profile#protocol OrbitalContactProfile#protocol}.
	Protocol *string `field:"required" json:"protocol" yaml:"protocol"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.89.0/docs/resources/orbital_contact_profile#ip_address OrbitalContactProfile#ip_address}.
	IpAddress *string `field:"optional" json:"ipAddress" yaml:"ipAddress"`
}


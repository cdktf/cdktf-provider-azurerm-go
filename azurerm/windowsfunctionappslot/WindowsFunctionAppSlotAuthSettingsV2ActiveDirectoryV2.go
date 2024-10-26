// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsfunctionappslot


type WindowsFunctionAppSlotAuthSettingsV2ActiveDirectoryV2 struct {
	// The ID of the Client to use to authenticate with Azure Active Directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/windows_function_app_slot#client_id WindowsFunctionAppSlot#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The Azure Tenant Endpoint for the Authenticating Tenant. e.g. `https://login.microsoftonline.com/v2.0/{tenant-guid}/`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/windows_function_app_slot#tenant_auth_endpoint WindowsFunctionAppSlot#tenant_auth_endpoint}
	TenantAuthEndpoint *string `field:"required" json:"tenantAuthEndpoint" yaml:"tenantAuthEndpoint"`
	// The list of allowed Applications for the Default Authorisation Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/windows_function_app_slot#allowed_applications WindowsFunctionAppSlot#allowed_applications}
	AllowedApplications *[]*string `field:"optional" json:"allowedApplications" yaml:"allowedApplications"`
	// Specifies a list of Allowed audience values to consider when validating JWTs issued by Azure Active Directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/windows_function_app_slot#allowed_audiences WindowsFunctionAppSlot#allowed_audiences}
	AllowedAudiences *[]*string `field:"optional" json:"allowedAudiences" yaml:"allowedAudiences"`
	// The list of allowed Group Names for the Default Authorisation Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/windows_function_app_slot#allowed_groups WindowsFunctionAppSlot#allowed_groups}
	AllowedGroups *[]*string `field:"optional" json:"allowedGroups" yaml:"allowedGroups"`
	// The list of allowed Identities for the Default Authorisation Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/windows_function_app_slot#allowed_identities WindowsFunctionAppSlot#allowed_identities}
	AllowedIdentities *[]*string `field:"optional" json:"allowedIdentities" yaml:"allowedIdentities"`
	// The thumbprint of the certificate used for signing purposes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/windows_function_app_slot#client_secret_certificate_thumbprint WindowsFunctionAppSlot#client_secret_certificate_thumbprint}
	ClientSecretCertificateThumbprint *string `field:"optional" json:"clientSecretCertificateThumbprint" yaml:"clientSecretCertificateThumbprint"`
	// The App Setting name that contains the client secret of the Client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/windows_function_app_slot#client_secret_setting_name WindowsFunctionAppSlot#client_secret_setting_name}
	ClientSecretSettingName *string `field:"optional" json:"clientSecretSettingName" yaml:"clientSecretSettingName"`
	// A list of Allowed Client Applications in the JWT Claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/windows_function_app_slot#jwt_allowed_client_applications WindowsFunctionAppSlot#jwt_allowed_client_applications}
	JwtAllowedClientApplications *[]*string `field:"optional" json:"jwtAllowedClientApplications" yaml:"jwtAllowedClientApplications"`
	// A list of Allowed Groups in the JWT Claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/windows_function_app_slot#jwt_allowed_groups WindowsFunctionAppSlot#jwt_allowed_groups}
	JwtAllowedGroups *[]*string `field:"optional" json:"jwtAllowedGroups" yaml:"jwtAllowedGroups"`
	// A map of key-value pairs to send to the Authorisation Endpoint when a user logs in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/windows_function_app_slot#login_parameters WindowsFunctionAppSlot#login_parameters}
	LoginParameters *map[string]*string `field:"optional" json:"loginParameters" yaml:"loginParameters"`
	// Should the www-authenticate provider should be omitted from the request? Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.7.0/docs/resources/windows_function_app_slot#www_authentication_disabled WindowsFunctionAppSlot#www_authentication_disabled}
	WwwAuthenticationDisabled interface{} `field:"optional" json:"wwwAuthenticationDisabled" yaml:"wwwAuthenticationDisabled"`
}


package windowswebapp


type WindowsWebAppAuthSettingsV2ActiveDirectoryV2 struct {
	// The ID of the Client to use to authenticate with Azure Active Directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/windows_web_app#client_id WindowsWebApp#client_id}
	ClientId *string `field:"required" json:"clientId" yaml:"clientId"`
	// The Azure Tenant Endpoint for the Authenticating Tenant. e.g. `https://login.microsoftonline.com/v2.0/{tenant-guid}/`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/windows_web_app#tenant_auth_endpoint WindowsWebApp#tenant_auth_endpoint}
	TenantAuthEndpoint *string `field:"required" json:"tenantAuthEndpoint" yaml:"tenantAuthEndpoint"`
	// The list of allowed Applications for the Default Authorisation Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/windows_web_app#allowed_applications WindowsWebApp#allowed_applications}
	AllowedApplications *[]*string `field:"optional" json:"allowedApplications" yaml:"allowedApplications"`
	// Specifies a list of Allowed audience values to consider when validating JWTs issued by Azure Active Directory.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/windows_web_app#allowed_audiences WindowsWebApp#allowed_audiences}
	AllowedAudiences *[]*string `field:"optional" json:"allowedAudiences" yaml:"allowedAudiences"`
	// The list of allowed Group Names for the Default Authorisation Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/windows_web_app#allowed_groups WindowsWebApp#allowed_groups}
	AllowedGroups *[]*string `field:"optional" json:"allowedGroups" yaml:"allowedGroups"`
	// The list of allowed Identities for the Default Authorisation Policy.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/windows_web_app#allowed_identities WindowsWebApp#allowed_identities}
	AllowedIdentities *[]*string `field:"optional" json:"allowedIdentities" yaml:"allowedIdentities"`
	// The thumbprint of the certificate used for signing purposes.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/windows_web_app#client_secret_certificate_thumbprint WindowsWebApp#client_secret_certificate_thumbprint}
	ClientSecretCertificateThumbprint *string `field:"optional" json:"clientSecretCertificateThumbprint" yaml:"clientSecretCertificateThumbprint"`
	// The App Setting name that contains the client secret of the Client.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/windows_web_app#client_secret_setting_name WindowsWebApp#client_secret_setting_name}
	ClientSecretSettingName *string `field:"optional" json:"clientSecretSettingName" yaml:"clientSecretSettingName"`
	// A list of Allowed Client Applications in the JWT Claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/windows_web_app#jwt_allowed_client_applications WindowsWebApp#jwt_allowed_client_applications}
	JwtAllowedClientApplications *[]*string `field:"optional" json:"jwtAllowedClientApplications" yaml:"jwtAllowedClientApplications"`
	// A list of Allowed Groups in the JWT Claim.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/windows_web_app#jwt_allowed_groups WindowsWebApp#jwt_allowed_groups}
	JwtAllowedGroups *[]*string `field:"optional" json:"jwtAllowedGroups" yaml:"jwtAllowedGroups"`
	// A map of key-value pairs to send to the Authorisation Endpoint when a user logs in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/windows_web_app#login_parameters WindowsWebApp#login_parameters}
	LoginParameters *map[string]*string `field:"optional" json:"loginParameters" yaml:"loginParameters"`
	// Should the www-authenticate provider should be omitted from the request? Defaults to `false`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.58.0/docs/resources/windows_web_app#www_authentication_disabled WindowsWebApp#www_authentication_disabled}
	WwwAuthenticationDisabled interface{} `field:"optional" json:"wwwAuthenticationDisabled" yaml:"wwwAuthenticationDisabled"`
}


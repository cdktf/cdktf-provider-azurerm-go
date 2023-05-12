package provider


type AzurermProviderConfig struct {
	// features block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs#features AzurermProvider#features}
	Features *AzurermProviderFeatures `field:"required" json:"features" yaml:"features"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs#alias AzurermProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs#auxiliary_tenant_ids AzurermProvider#auxiliary_tenant_ids}.
	AuxiliaryTenantIds *[]*string `field:"optional" json:"auxiliaryTenantIds" yaml:"auxiliaryTenantIds"`
	// Base64 encoded PKCS#12 certificate bundle to use when authenticating as a Service Principal using a Client Certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs#client_certificate AzurermProvider#client_certificate}
	ClientCertificate *string `field:"optional" json:"clientCertificate" yaml:"clientCertificate"`
	// The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client Certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs#client_certificate_password AzurermProvider#client_certificate_password}
	ClientCertificatePassword *string `field:"optional" json:"clientCertificatePassword" yaml:"clientCertificatePassword"`
	// The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service Principal using a Client Certificate.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs#client_certificate_path AzurermProvider#client_certificate_path}
	ClientCertificatePath *string `field:"optional" json:"clientCertificatePath" yaml:"clientCertificatePath"`
	// The Client ID which should be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs#client_id AzurermProvider#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs#client_secret AzurermProvider#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// This will disable the x-ms-correlation-request-id header.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs#disable_correlation_request_id AzurermProvider#disable_correlation_request_id}
	DisableCorrelationRequestId interface{} `field:"optional" json:"disableCorrelationRequestId" yaml:"disableCorrelationRequestId"`
	// This will disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs#disable_terraform_partner_id AzurermProvider#disable_terraform_partner_id}
	DisableTerraformPartnerId interface{} `field:"optional" json:"disableTerraformPartnerId" yaml:"disableTerraformPartnerId"`
	// The Cloud Environment which should be used. Possible values are public, usgovernment, and china. Defaults to public.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs#environment AzurermProvider#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// The Hostname which should be used for the Azure Metadata Service.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs#metadata_host AzurermProvider#metadata_host}
	MetadataHost *string `field:"optional" json:"metadataHost" yaml:"metadataHost"`
	// The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected automatically.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs#msi_endpoint AzurermProvider#msi_endpoint}
	MsiEndpoint *string `field:"optional" json:"msiEndpoint" yaml:"msiEndpoint"`
	// The bearer token for the request to the OIDC provider.
	//
	// For use when authenticating as a Service Principal using OpenID Connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs#oidc_request_token AzurermProvider#oidc_request_token}
	OidcRequestToken *string `field:"optional" json:"oidcRequestToken" yaml:"oidcRequestToken"`
	// The URL for the OIDC provider from which to request an ID token.
	//
	// For use when authenticating as a Service Principal using OpenID Connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs#oidc_request_url AzurermProvider#oidc_request_url}
	OidcRequestUrl *string `field:"optional" json:"oidcRequestUrl" yaml:"oidcRequestUrl"`
	// The OIDC ID token for use when authenticating as a Service Principal using OpenID Connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs#oidc_token AzurermProvider#oidc_token}
	OidcToken *string `field:"optional" json:"oidcToken" yaml:"oidcToken"`
	// The path to a file containing an OIDC ID token for use when authenticating as a Service Principal using OpenID Connect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs#oidc_token_file_path AzurermProvider#oidc_token_file_path}
	OidcTokenFilePath *string `field:"optional" json:"oidcTokenFilePath" yaml:"oidcTokenFilePath"`
	// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs#partner_id AzurermProvider#partner_id}
	PartnerId *string `field:"optional" json:"partnerId" yaml:"partnerId"`
	// Should the AzureRM Provider skip registering all of the Resource Providers that it supports, if they're not already registered?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs#skip_provider_registration AzurermProvider#skip_provider_registration}
	SkipProviderRegistration interface{} `field:"optional" json:"skipProviderRegistration" yaml:"skipProviderRegistration"`
	// Should the AzureRM Provider use AzureAD to access the Storage Data Plane API's?
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs#storage_use_azuread AzurermProvider#storage_use_azuread}
	StorageUseAzuread interface{} `field:"optional" json:"storageUseAzuread" yaml:"storageUseAzuread"`
	// The Subscription ID which should be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs#subscription_id AzurermProvider#subscription_id}
	SubscriptionId *string `field:"optional" json:"subscriptionId" yaml:"subscriptionId"`
	// The Tenant ID which should be used.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs#tenant_id AzurermProvider#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// Allow Azure CLI to be used for Authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs#use_cli AzurermProvider#use_cli}
	UseCli interface{} `field:"optional" json:"useCli" yaml:"useCli"`
	// Allow Managed Service Identity to be used for Authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs#use_msi AzurermProvider#use_msi}
	UseMsi interface{} `field:"optional" json:"useMsi" yaml:"useMsi"`
	// Allow OpenID Connect to be used for authentication.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.56.0/docs#use_oidc AzurermProvider#use_oidc}
	UseOidc interface{} `field:"optional" json:"useOidc" yaml:"useOidc"`
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs azurerm}.
type AzurermProvider interface {
	cdktf.TerraformProvider
	AdoPipelineServiceConnectionId() *string
	SetAdoPipelineServiceConnectionId(val *string)
	AdoPipelineServiceConnectionIdInput() *string
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	AuxiliaryTenantIds() *[]*string
	SetAuxiliaryTenantIds(val *[]*string)
	AuxiliaryTenantIdsInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientCertificate() *string
	SetClientCertificate(val *string)
	ClientCertificateInput() *string
	ClientCertificatePassword() *string
	SetClientCertificatePassword(val *string)
	ClientCertificatePasswordInput() *string
	ClientCertificatePath() *string
	SetClientCertificatePath(val *string)
	ClientCertificatePathInput() *string
	ClientId() *string
	SetClientId(val *string)
	ClientIdFilePath() *string
	SetClientIdFilePath(val *string)
	ClientIdFilePathInput() *string
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
	ClientSecretFilePath() *string
	SetClientSecretFilePath(val *string)
	ClientSecretFilePathInput() *string
	ClientSecretInput() *string
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	DisableCorrelationRequestId() interface{}
	SetDisableCorrelationRequestId(val interface{})
	DisableCorrelationRequestIdInput() interface{}
	DisableTerraformPartnerId() interface{}
	SetDisableTerraformPartnerId(val interface{})
	DisableTerraformPartnerIdInput() interface{}
	Environment() *string
	SetEnvironment(val *string)
	EnvironmentInput() *string
	Features() interface{}
	SetFeatures(val interface{})
	FeaturesInput() interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	// Experimental.
	MetaAttributes() *map[string]interface{}
	MetadataHost() *string
	SetMetadataHost(val *string)
	MetadataHostInput() *string
	MsiEndpoint() *string
	SetMsiEndpoint(val *string)
	MsiEndpointInput() *string
	// The tree node.
	Node() constructs.Node
	OidcRequestToken() *string
	SetOidcRequestToken(val *string)
	OidcRequestTokenInput() *string
	OidcRequestUrl() *string
	SetOidcRequestUrl(val *string)
	OidcRequestUrlInput() *string
	OidcToken() *string
	SetOidcToken(val *string)
	OidcTokenFilePath() *string
	SetOidcTokenFilePath(val *string)
	OidcTokenFilePathInput() *string
	OidcTokenInput() *string
	PartnerId() *string
	SetPartnerId(val *string)
	PartnerIdInput() *string
	// Experimental.
	RawOverrides() interface{}
	ResourceProviderRegistrations() *string
	SetResourceProviderRegistrations(val *string)
	ResourceProviderRegistrationsInput() *string
	ResourceProvidersToRegister() *[]*string
	SetResourceProvidersToRegister(val *[]*string)
	ResourceProvidersToRegisterInput() *[]*string
	SkipProviderRegistration() interface{}
	SetSkipProviderRegistration(val interface{})
	SkipProviderRegistrationInput() interface{}
	StorageUseAzuread() interface{}
	SetStorageUseAzuread(val interface{})
	StorageUseAzureadInput() interface{}
	SubscriptionId() *string
	SetSubscriptionId(val *string)
	SubscriptionIdInput() *string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	UseAksWorkloadIdentity() interface{}
	SetUseAksWorkloadIdentity(val interface{})
	UseAksWorkloadIdentityInput() interface{}
	UseCli() interface{}
	SetUseCli(val interface{})
	UseCliInput() interface{}
	UseMsi() interface{}
	SetUseMsi(val interface{})
	UseMsiInput() interface{}
	UseOidc() interface{}
	SetUseOidc(val interface{})
	UseOidcInput() interface{}
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAdoPipelineServiceConnectionId()
	ResetAlias()
	ResetAuxiliaryTenantIds()
	ResetClientCertificate()
	ResetClientCertificatePassword()
	ResetClientCertificatePath()
	ResetClientId()
	ResetClientIdFilePath()
	ResetClientSecret()
	ResetClientSecretFilePath()
	ResetDisableCorrelationRequestId()
	ResetDisableTerraformPartnerId()
	ResetEnvironment()
	ResetFeatures()
	ResetMetadataHost()
	ResetMsiEndpoint()
	ResetOidcRequestToken()
	ResetOidcRequestUrl()
	ResetOidcToken()
	ResetOidcTokenFilePath()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPartnerId()
	ResetResourceProviderRegistrations()
	ResetResourceProvidersToRegister()
	ResetSkipProviderRegistration()
	ResetStorageUseAzuread()
	ResetSubscriptionId()
	ResetTenantId()
	ResetUseAksWorkloadIdentity()
	ResetUseCli()
	ResetUseMsi()
	ResetUseOidc()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AzurermProvider
type jsiiProxy_AzurermProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_AzurermProvider) AdoPipelineServiceConnectionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adoPipelineServiceConnectionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) AdoPipelineServiceConnectionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adoPipelineServiceConnectionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) AuxiliaryTenantIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auxiliaryTenantIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) AuxiliaryTenantIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auxiliaryTenantIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) ClientCertificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) ClientCertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) ClientCertificatePassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificatePassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) ClientCertificatePasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificatePasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) ClientCertificatePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificatePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) ClientCertificatePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientCertificatePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) ClientIdFilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) ClientIdFilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) ClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) ClientSecretFilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) ClientSecretFilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) ClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) DisableCorrelationRequestId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableCorrelationRequestId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) DisableCorrelationRequestIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableCorrelationRequestIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) DisableTerraformPartnerId() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableTerraformPartnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) DisableTerraformPartnerIdInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableTerraformPartnerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) EnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) Features() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"features",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) FeaturesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"featuresInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) MetadataHost() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) MetadataHostInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataHostInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) MsiEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"msiEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) MsiEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"msiEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) OidcRequestToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcRequestToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) OidcRequestTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcRequestTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) OidcRequestUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcRequestUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) OidcRequestUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcRequestUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) OidcToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) OidcTokenFilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcTokenFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) OidcTokenFilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcTokenFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) OidcTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) PartnerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) PartnerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) ResourceProviderRegistrations() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceProviderRegistrations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) ResourceProviderRegistrationsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceProviderRegistrationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) ResourceProvidersToRegister() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceProvidersToRegister",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) ResourceProvidersToRegisterInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceProvidersToRegisterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) SkipProviderRegistration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipProviderRegistration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) SkipProviderRegistrationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipProviderRegistrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) StorageUseAzuread() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageUseAzuread",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) StorageUseAzureadInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageUseAzureadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) SubscriptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) SubscriptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) UseAksWorkloadIdentity() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAksWorkloadIdentity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) UseAksWorkloadIdentityInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useAksWorkloadIdentityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) UseCli() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCli",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) UseCliInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useCliInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) UseMsi() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useMsi",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) UseMsiInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useMsiInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) UseOidc() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useOidc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) UseOidcInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useOidcInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs azurerm} Resource.
func NewAzurermProvider(scope constructs.Construct, id *string, config *AzurermProviderConfig) AzurermProvider {
	_init_.Initialize()

	if err := validateNewAzurermProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AzurermProvider{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.provider.AzurermProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.21.0/docs azurerm} Resource.
func NewAzurermProvider_Override(a AzurermProvider, scope constructs.Construct, id *string, config *AzurermProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.provider.AzurermProvider",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AzurermProvider)SetAdoPipelineServiceConnectionId(val *string) {
	_jsii_.Set(
		j,
		"adoPipelineServiceConnectionId",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetAuxiliaryTenantIds(val *[]*string) {
	_jsii_.Set(
		j,
		"auxiliaryTenantIds",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetClientCertificate(val *string) {
	_jsii_.Set(
		j,
		"clientCertificate",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetClientCertificatePassword(val *string) {
	_jsii_.Set(
		j,
		"clientCertificatePassword",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetClientCertificatePath(val *string) {
	_jsii_.Set(
		j,
		"clientCertificatePath",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetClientIdFilePath(val *string) {
	_jsii_.Set(
		j,
		"clientIdFilePath",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetClientSecret(val *string) {
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetClientSecretFilePath(val *string) {
	_jsii_.Set(
		j,
		"clientSecretFilePath",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetDisableCorrelationRequestId(val interface{}) {
	if err := j.validateSetDisableCorrelationRequestIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableCorrelationRequestId",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetDisableTerraformPartnerId(val interface{}) {
	if err := j.validateSetDisableTerraformPartnerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableTerraformPartnerId",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetEnvironment(val *string) {
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetFeatures(val interface{}) {
	if err := j.validateSetFeaturesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"features",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetMetadataHost(val *string) {
	_jsii_.Set(
		j,
		"metadataHost",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetMsiEndpoint(val *string) {
	_jsii_.Set(
		j,
		"msiEndpoint",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetOidcRequestToken(val *string) {
	_jsii_.Set(
		j,
		"oidcRequestToken",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetOidcRequestUrl(val *string) {
	_jsii_.Set(
		j,
		"oidcRequestUrl",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetOidcToken(val *string) {
	_jsii_.Set(
		j,
		"oidcToken",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetOidcTokenFilePath(val *string) {
	_jsii_.Set(
		j,
		"oidcTokenFilePath",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetPartnerId(val *string) {
	_jsii_.Set(
		j,
		"partnerId",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetResourceProviderRegistrations(val *string) {
	_jsii_.Set(
		j,
		"resourceProviderRegistrations",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetResourceProvidersToRegister(val *[]*string) {
	_jsii_.Set(
		j,
		"resourceProvidersToRegister",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetSkipProviderRegistration(val interface{}) {
	if err := j.validateSetSkipProviderRegistrationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipProviderRegistration",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetStorageUseAzuread(val interface{}) {
	if err := j.validateSetStorageUseAzureadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageUseAzuread",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetSubscriptionId(val *string) {
	_jsii_.Set(
		j,
		"subscriptionId",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetTenantId(val *string) {
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetUseAksWorkloadIdentity(val interface{}) {
	if err := j.validateSetUseAksWorkloadIdentityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useAksWorkloadIdentity",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetUseCli(val interface{}) {
	if err := j.validateSetUseCliParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useCli",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetUseMsi(val interface{}) {
	if err := j.validateSetUseMsiParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useMsi",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider)SetUseOidc(val interface{}) {
	if err := j.validateSetUseOidcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useOidc",
		val,
	)
}

// Generates CDKTF code for importing a AzurermProvider resource upon running "cdktf plan <stack-name>".
func AzurermProvider_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateAzurermProvider_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.provider.AzurermProvider",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func AzurermProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAzurermProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.provider.AzurermProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AzurermProvider_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAzurermProvider_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.provider.AzurermProvider",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AzurermProvider_IsTerraformProvider(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAzurermProvider_IsTerraformProviderParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.provider.AzurermProvider",
		"isTerraformProvider",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AzurermProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.provider.AzurermProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AzurermProvider) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AzurermProvider) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AzurermProvider) ResetAdoPipelineServiceConnectionId() {
	_jsii_.InvokeVoid(
		a,
		"resetAdoPipelineServiceConnectionId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		a,
		"resetAlias",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetAuxiliaryTenantIds() {
	_jsii_.InvokeVoid(
		a,
		"resetAuxiliaryTenantIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetClientCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetClientCertificatePassword() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificatePassword",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetClientCertificatePath() {
	_jsii_.InvokeVoid(
		a,
		"resetClientCertificatePath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetClientId() {
	_jsii_.InvokeVoid(
		a,
		"resetClientId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetClientIdFilePath() {
	_jsii_.InvokeVoid(
		a,
		"resetClientIdFilePath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetClientSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecret",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetClientSecretFilePath() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecretFilePath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetDisableCorrelationRequestId() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableCorrelationRequestId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetDisableTerraformPartnerId() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableTerraformPartnerId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetEnvironment() {
	_jsii_.InvokeVoid(
		a,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetFeatures() {
	_jsii_.InvokeVoid(
		a,
		"resetFeatures",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetMetadataHost() {
	_jsii_.InvokeVoid(
		a,
		"resetMetadataHost",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetMsiEndpoint() {
	_jsii_.InvokeVoid(
		a,
		"resetMsiEndpoint",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetOidcRequestToken() {
	_jsii_.InvokeVoid(
		a,
		"resetOidcRequestToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetOidcRequestUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetOidcRequestUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetOidcToken() {
	_jsii_.InvokeVoid(
		a,
		"resetOidcToken",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetOidcTokenFilePath() {
	_jsii_.InvokeVoid(
		a,
		"resetOidcTokenFilePath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetPartnerId() {
	_jsii_.InvokeVoid(
		a,
		"resetPartnerId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetResourceProviderRegistrations() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceProviderRegistrations",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetResourceProvidersToRegister() {
	_jsii_.InvokeVoid(
		a,
		"resetResourceProvidersToRegister",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetSkipProviderRegistration() {
	_jsii_.InvokeVoid(
		a,
		"resetSkipProviderRegistration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetStorageUseAzuread() {
	_jsii_.InvokeVoid(
		a,
		"resetStorageUseAzuread",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetSubscriptionId() {
	_jsii_.InvokeVoid(
		a,
		"resetSubscriptionId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetTenantId() {
	_jsii_.InvokeVoid(
		a,
		"resetTenantId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetUseAksWorkloadIdentity() {
	_jsii_.InvokeVoid(
		a,
		"resetUseAksWorkloadIdentity",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetUseCli() {
	_jsii_.InvokeVoid(
		a,
		"resetUseCli",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetUseMsi() {
	_jsii_.InvokeVoid(
		a,
		"resetUseMsi",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) ResetUseOidc() {
	_jsii_.InvokeVoid(
		a,
		"resetUseOidc",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzurermProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurermProvider) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurermProvider) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurermProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurermProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzurermProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


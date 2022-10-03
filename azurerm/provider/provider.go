package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/azurerm azurerm}.
type AzurermProvider interface {
	cdktf.TerraformProvider
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	AuxiliaryTenantIds() *[]*string
	SetAuxiliaryTenantIds(val *[]*string)
	AuxiliaryTenantIdsInput() *[]*string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientCertificatePassword() *string
	SetClientCertificatePassword(val *string)
	ClientCertificatePasswordInput() *string
	ClientCertificatePath() *string
	SetClientCertificatePath(val *string)
	ClientCertificatePathInput() *string
	ClientId() *string
	SetClientId(val *string)
	ClientIdInput() *string
	ClientSecret() *string
	SetClientSecret(val *string)
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
	Features() *AzurermProviderFeatures
	SetFeatures(val *AzurermProviderFeatures)
	FeaturesInput() *AzurermProviderFeatures
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
	ResetAlias()
	ResetAuxiliaryTenantIds()
	ResetClientCertificatePassword()
	ResetClientCertificatePath()
	ResetClientId()
	ResetClientSecret()
	ResetDisableCorrelationRequestId()
	ResetDisableTerraformPartnerId()
	ResetEnvironment()
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
	ResetSkipProviderRegistration()
	ResetStorageUseAzuread()
	ResetSubscriptionId()
	ResetTenantId()
	ResetUseMsi()
	ResetUseOidc()
	SynthesizeAttributes() *map[string]interface{}
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

func (j *jsiiProxy_AzurermProvider) Features() *AzurermProviderFeatures {
	var returns *AzurermProviderFeatures
	_jsii_.Get(
		j,
		"features",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzurermProvider) FeaturesInput() *AzurermProviderFeatures {
	var returns *AzurermProviderFeatures
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


// Create a new {@link https://www.terraform.io/docs/providers/azurerm azurerm} Resource.
func NewAzurermProvider(scope constructs.Construct, id *string, config *AzurermProviderConfig) AzurermProvider {
	_init_.Initialize()

	j := jsiiProxy_AzurermProvider{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.provider.AzurermProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/azurerm azurerm} Resource.
func NewAzurermProvider_Override(a AzurermProvider, scope constructs.Construct, id *string, config *AzurermProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.provider.AzurermProvider",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AzurermProvider) SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider) SetAuxiliaryTenantIds(val *[]*string) {
	_jsii_.Set(
		j,
		"auxiliaryTenantIds",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider) SetClientCertificatePassword(val *string) {
	_jsii_.Set(
		j,
		"clientCertificatePassword",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider) SetClientCertificatePath(val *string) {
	_jsii_.Set(
		j,
		"clientCertificatePath",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider) SetClientId(val *string) {
	_jsii_.Set(
		j,
		"clientId",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider) SetClientSecret(val *string) {
	_jsii_.Set(
		j,
		"clientSecret",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider) SetDisableCorrelationRequestId(val interface{}) {
	_jsii_.Set(
		j,
		"disableCorrelationRequestId",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider) SetDisableTerraformPartnerId(val interface{}) {
	_jsii_.Set(
		j,
		"disableTerraformPartnerId",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider) SetEnvironment(val *string) {
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider) SetFeatures(val *AzurermProviderFeatures) {
	_jsii_.Set(
		j,
		"features",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider) SetMetadataHost(val *string) {
	_jsii_.Set(
		j,
		"metadataHost",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider) SetMsiEndpoint(val *string) {
	_jsii_.Set(
		j,
		"msiEndpoint",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider) SetOidcRequestToken(val *string) {
	_jsii_.Set(
		j,
		"oidcRequestToken",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider) SetOidcRequestUrl(val *string) {
	_jsii_.Set(
		j,
		"oidcRequestUrl",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider) SetOidcToken(val *string) {
	_jsii_.Set(
		j,
		"oidcToken",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider) SetOidcTokenFilePath(val *string) {
	_jsii_.Set(
		j,
		"oidcTokenFilePath",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider) SetPartnerId(val *string) {
	_jsii_.Set(
		j,
		"partnerId",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider) SetSkipProviderRegistration(val interface{}) {
	_jsii_.Set(
		j,
		"skipProviderRegistration",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider) SetStorageUseAzuread(val interface{}) {
	_jsii_.Set(
		j,
		"storageUseAzuread",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider) SetSubscriptionId(val *string) {
	_jsii_.Set(
		j,
		"subscriptionId",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider) SetTenantId(val *string) {
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider) SetUseMsi(val interface{}) {
	_jsii_.Set(
		j,
		"useMsi",
		val,
	)
}

func (j *jsiiProxy_AzurermProvider) SetUseOidc(val interface{}) {
	_jsii_.Set(
		j,
		"useOidc",
		val,
	)
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

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.provider.AzurermProvider",
		"isConstruct",
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
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AzurermProvider) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
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

func (a *jsiiProxy_AzurermProvider) ResetClientSecret() {
	_jsii_.InvokeVoid(
		a,
		"resetClientSecret",
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

type AzurermProviderConfig struct {
	// features block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#features AzurermProvider#features}
	Features *AzurermProviderFeatures `field:"required" json:"features" yaml:"features"`
	// Alias name.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#alias AzurermProvider#alias}
	Alias *string `field:"optional" json:"alias" yaml:"alias"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#auxiliary_tenant_ids AzurermProvider#auxiliary_tenant_ids}.
	AuxiliaryTenantIds *[]*string `field:"optional" json:"auxiliaryTenantIds" yaml:"auxiliaryTenantIds"`
	// The password associated with the Client Certificate. For use when authenticating as a Service Principal using a Client Certificate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#client_certificate_password AzurermProvider#client_certificate_password}
	ClientCertificatePassword *string `field:"optional" json:"clientCertificatePassword" yaml:"clientCertificatePassword"`
	// The path to the Client Certificate associated with the Service Principal for use when authenticating as a Service Principal using a Client Certificate.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#client_certificate_path AzurermProvider#client_certificate_path}
	ClientCertificatePath *string `field:"optional" json:"clientCertificatePath" yaml:"clientCertificatePath"`
	// The Client ID which should be used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#client_id AzurermProvider#client_id}
	ClientId *string `field:"optional" json:"clientId" yaml:"clientId"`
	// The Client Secret which should be used. For use When authenticating as a Service Principal using a Client Secret.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#client_secret AzurermProvider#client_secret}
	ClientSecret *string `field:"optional" json:"clientSecret" yaml:"clientSecret"`
	// This will disable the x-ms-correlation-request-id header.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#disable_correlation_request_id AzurermProvider#disable_correlation_request_id}
	DisableCorrelationRequestId interface{} `field:"optional" json:"disableCorrelationRequestId" yaml:"disableCorrelationRequestId"`
	// This will disable the Terraform Partner ID which is used if a custom `partner_id` isn't specified.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#disable_terraform_partner_id AzurermProvider#disable_terraform_partner_id}
	DisableTerraformPartnerId interface{} `field:"optional" json:"disableTerraformPartnerId" yaml:"disableTerraformPartnerId"`
	// The Cloud Environment which should be used. Possible values are public, usgovernment, and china. Defaults to public.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#environment AzurermProvider#environment}
	Environment *string `field:"optional" json:"environment" yaml:"environment"`
	// The Hostname which should be used for the Azure Metadata Service.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#metadata_host AzurermProvider#metadata_host}
	MetadataHost *string `field:"optional" json:"metadataHost" yaml:"metadataHost"`
	// The path to a custom endpoint for Managed Service Identity - in most circumstances this should be detected automatically.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#msi_endpoint AzurermProvider#msi_endpoint}
	MsiEndpoint *string `field:"optional" json:"msiEndpoint" yaml:"msiEndpoint"`
	// The bearer token for the request to the OIDC provider.
	//
	// For use when authenticating as a Service Principal using OpenID Connect.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#oidc_request_token AzurermProvider#oidc_request_token}
	OidcRequestToken *string `field:"optional" json:"oidcRequestToken" yaml:"oidcRequestToken"`
	// The URL for the OIDC provider from which to request an ID token.
	//
	// For use when authenticating as a Service Principal using OpenID Connect.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#oidc_request_url AzurermProvider#oidc_request_url}
	OidcRequestUrl *string `field:"optional" json:"oidcRequestUrl" yaml:"oidcRequestUrl"`
	// The OIDC ID token for use when authenticating as a Service Principal using OpenID Connect.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#oidc_token AzurermProvider#oidc_token}
	OidcToken *string `field:"optional" json:"oidcToken" yaml:"oidcToken"`
	// The path to a file containing an OIDC ID token for use when authenticating as a Service Principal using OpenID Connect.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#oidc_token_file_path AzurermProvider#oidc_token_file_path}
	OidcTokenFilePath *string `field:"optional" json:"oidcTokenFilePath" yaml:"oidcTokenFilePath"`
	// A GUID/UUID that is registered with Microsoft to facilitate partner resource usage attribution.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#partner_id AzurermProvider#partner_id}
	PartnerId *string `field:"optional" json:"partnerId" yaml:"partnerId"`
	// Should the AzureRM Provider skip registering all of the Resource Providers that it supports, if they're not already registered?
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#skip_provider_registration AzurermProvider#skip_provider_registration}
	SkipProviderRegistration interface{} `field:"optional" json:"skipProviderRegistration" yaml:"skipProviderRegistration"`
	// Should the AzureRM Provider use AzureAD to access the Storage Data Plane API's?
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#storage_use_azuread AzurermProvider#storage_use_azuread}
	StorageUseAzuread interface{} `field:"optional" json:"storageUseAzuread" yaml:"storageUseAzuread"`
	// The Subscription ID which should be used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#subscription_id AzurermProvider#subscription_id}
	SubscriptionId *string `field:"optional" json:"subscriptionId" yaml:"subscriptionId"`
	// The Tenant ID which should be used.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#tenant_id AzurermProvider#tenant_id}
	TenantId *string `field:"optional" json:"tenantId" yaml:"tenantId"`
	// Allowed Managed Service Identity be used for Authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#use_msi AzurermProvider#use_msi}
	UseMsi interface{} `field:"optional" json:"useMsi" yaml:"useMsi"`
	// Allow OpenID Connect to be used for authentication.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#use_oidc AzurermProvider#use_oidc}
	UseOidc interface{} `field:"optional" json:"useOidc" yaml:"useOidc"`
}

type AzurermProviderFeatures struct {
	// api_management block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#api_management AzurermProvider#api_management}
	ApiManagement *AzurermProviderFeaturesApiManagement `field:"optional" json:"apiManagement" yaml:"apiManagement"`
	// application_insights block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#application_insights AzurermProvider#application_insights}
	ApplicationInsights *AzurermProviderFeaturesApplicationInsights `field:"optional" json:"applicationInsights" yaml:"applicationInsights"`
	// cognitive_account block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#cognitive_account AzurermProvider#cognitive_account}
	CognitiveAccount *AzurermProviderFeaturesCognitiveAccount `field:"optional" json:"cognitiveAccount" yaml:"cognitiveAccount"`
	// key_vault block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#key_vault AzurermProvider#key_vault}
	KeyVault *AzurermProviderFeaturesKeyVault `field:"optional" json:"keyVault" yaml:"keyVault"`
	// log_analytics_workspace block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#log_analytics_workspace AzurermProvider#log_analytics_workspace}
	LogAnalyticsWorkspace *AzurermProviderFeaturesLogAnalyticsWorkspace `field:"optional" json:"logAnalyticsWorkspace" yaml:"logAnalyticsWorkspace"`
	// network block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#network AzurermProvider#network}
	Network *AzurermProviderFeaturesNetwork `field:"optional" json:"network" yaml:"network"`
	// resource_group block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#resource_group AzurermProvider#resource_group}
	ResourceGroup *AzurermProviderFeaturesResourceGroup `field:"optional" json:"resourceGroup" yaml:"resourceGroup"`
	// template_deployment block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#template_deployment AzurermProvider#template_deployment}
	TemplateDeployment *AzurermProviderFeaturesTemplateDeployment `field:"optional" json:"templateDeployment" yaml:"templateDeployment"`
	// virtual_machine block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#virtual_machine AzurermProvider#virtual_machine}
	VirtualMachine *AzurermProviderFeaturesVirtualMachine `field:"optional" json:"virtualMachine" yaml:"virtualMachine"`
	// virtual_machine_scale_set block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#virtual_machine_scale_set AzurermProvider#virtual_machine_scale_set}
	VirtualMachineScaleSet *AzurermProviderFeaturesVirtualMachineScaleSet `field:"optional" json:"virtualMachineScaleSet" yaml:"virtualMachineScaleSet"`
}

type AzurermProviderFeaturesApiManagement struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#purge_soft_delete_on_destroy AzurermProvider#purge_soft_delete_on_destroy}.
	PurgeSoftDeleteOnDestroy interface{} `field:"optional" json:"purgeSoftDeleteOnDestroy" yaml:"purgeSoftDeleteOnDestroy"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#recover_soft_deleted AzurermProvider#recover_soft_deleted}.
	RecoverSoftDeleted interface{} `field:"optional" json:"recoverSoftDeleted" yaml:"recoverSoftDeleted"`
}

type AzurermProviderFeaturesApplicationInsights struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#disable_generated_rule AzurermProvider#disable_generated_rule}.
	DisableGeneratedRule interface{} `field:"optional" json:"disableGeneratedRule" yaml:"disableGeneratedRule"`
}

type AzurermProviderFeaturesCognitiveAccount struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#purge_soft_delete_on_destroy AzurermProvider#purge_soft_delete_on_destroy}.
	PurgeSoftDeleteOnDestroy interface{} `field:"optional" json:"purgeSoftDeleteOnDestroy" yaml:"purgeSoftDeleteOnDestroy"`
}

type AzurermProviderFeaturesKeyVault struct {
	// When enabled soft-deleted `azurerm_key_vault_certificate` resources will be permanently deleted (e.g purged), when destroyed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#purge_soft_deleted_certificates_on_destroy AzurermProvider#purge_soft_deleted_certificates_on_destroy}
	PurgeSoftDeletedCertificatesOnDestroy interface{} `field:"optional" json:"purgeSoftDeletedCertificatesOnDestroy" yaml:"purgeSoftDeletedCertificatesOnDestroy"`
	// When enabled soft-deleted `azurerm_key_vault_managed_hardware_security_module` resources will be permanently deleted (e.g purged), when destroyed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#purge_soft_deleted_hardware_security_modules_on_destroy AzurermProvider#purge_soft_deleted_hardware_security_modules_on_destroy}
	PurgeSoftDeletedHardwareSecurityModulesOnDestroy interface{} `field:"optional" json:"purgeSoftDeletedHardwareSecurityModulesOnDestroy" yaml:"purgeSoftDeletedHardwareSecurityModulesOnDestroy"`
	// When enabled soft-deleted `azurerm_key_vault_key` resources will be permanently deleted (e.g purged), when destroyed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#purge_soft_deleted_keys_on_destroy AzurermProvider#purge_soft_deleted_keys_on_destroy}
	PurgeSoftDeletedKeysOnDestroy interface{} `field:"optional" json:"purgeSoftDeletedKeysOnDestroy" yaml:"purgeSoftDeletedKeysOnDestroy"`
	// When enabled soft-deleted `azurerm_key_vault_secret` resources will be permanently deleted (e.g purged), when destroyed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#purge_soft_deleted_secrets_on_destroy AzurermProvider#purge_soft_deleted_secrets_on_destroy}
	PurgeSoftDeletedSecretsOnDestroy interface{} `field:"optional" json:"purgeSoftDeletedSecretsOnDestroy" yaml:"purgeSoftDeletedSecretsOnDestroy"`
	// When enabled soft-deleted `azurerm_key_vault` resources will be permanently deleted (e.g purged), when destroyed.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#purge_soft_delete_on_destroy AzurermProvider#purge_soft_delete_on_destroy}
	PurgeSoftDeleteOnDestroy interface{} `field:"optional" json:"purgeSoftDeleteOnDestroy" yaml:"purgeSoftDeleteOnDestroy"`
	// When enabled soft-deleted `azurerm_key_vault_certificate` resources will be restored, instead of creating new ones.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#recover_soft_deleted_certificates AzurermProvider#recover_soft_deleted_certificates}
	RecoverSoftDeletedCertificates interface{} `field:"optional" json:"recoverSoftDeletedCertificates" yaml:"recoverSoftDeletedCertificates"`
	// When enabled soft-deleted `azurerm_key_vault_key` resources will be restored, instead of creating new ones.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#recover_soft_deleted_keys AzurermProvider#recover_soft_deleted_keys}
	RecoverSoftDeletedKeys interface{} `field:"optional" json:"recoverSoftDeletedKeys" yaml:"recoverSoftDeletedKeys"`
	// When enabled soft-deleted `azurerm_key_vault` resources will be restored, instead of creating new ones.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#recover_soft_deleted_key_vaults AzurermProvider#recover_soft_deleted_key_vaults}
	RecoverSoftDeletedKeyVaults interface{} `field:"optional" json:"recoverSoftDeletedKeyVaults" yaml:"recoverSoftDeletedKeyVaults"`
	// When enabled soft-deleted `azurerm_key_vault_secret` resources will be restored, instead of creating new ones.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#recover_soft_deleted_secrets AzurermProvider#recover_soft_deleted_secrets}
	RecoverSoftDeletedSecrets interface{} `field:"optional" json:"recoverSoftDeletedSecrets" yaml:"recoverSoftDeletedSecrets"`
}

type AzurermProviderFeaturesLogAnalyticsWorkspace struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#permanently_delete_on_destroy AzurermProvider#permanently_delete_on_destroy}.
	PermanentlyDeleteOnDestroy interface{} `field:"optional" json:"permanentlyDeleteOnDestroy" yaml:"permanentlyDeleteOnDestroy"`
}

type AzurermProviderFeaturesNetwork struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#relaxed_locking AzurermProvider#relaxed_locking}.
	RelaxedLocking interface{} `field:"required" json:"relaxedLocking" yaml:"relaxedLocking"`
}

type AzurermProviderFeaturesResourceGroup struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#prevent_deletion_if_contains_resources AzurermProvider#prevent_deletion_if_contains_resources}.
	PreventDeletionIfContainsResources interface{} `field:"optional" json:"preventDeletionIfContainsResources" yaml:"preventDeletionIfContainsResources"`
}

type AzurermProviderFeaturesTemplateDeployment struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#delete_nested_items_during_deletion AzurermProvider#delete_nested_items_during_deletion}.
	DeleteNestedItemsDuringDeletion interface{} `field:"required" json:"deleteNestedItemsDuringDeletion" yaml:"deleteNestedItemsDuringDeletion"`
}

type AzurermProviderFeaturesVirtualMachine struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#delete_os_disk_on_deletion AzurermProvider#delete_os_disk_on_deletion}.
	DeleteOsDiskOnDeletion interface{} `field:"optional" json:"deleteOsDiskOnDeletion" yaml:"deleteOsDiskOnDeletion"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#graceful_shutdown AzurermProvider#graceful_shutdown}.
	GracefulShutdown interface{} `field:"optional" json:"gracefulShutdown" yaml:"gracefulShutdown"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#skip_shutdown_and_force_delete AzurermProvider#skip_shutdown_and_force_delete}.
	SkipShutdownAndForceDelete interface{} `field:"optional" json:"skipShutdownAndForceDelete" yaml:"skipShutdownAndForceDelete"`
}

type AzurermProviderFeaturesVirtualMachineScaleSet struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#roll_instances_when_required AzurermProvider#roll_instances_when_required}.
	RollInstancesWhenRequired interface{} `field:"required" json:"rollInstancesWhenRequired" yaml:"rollInstancesWhenRequired"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#force_delete AzurermProvider#force_delete}.
	ForceDelete interface{} `field:"optional" json:"forceDelete" yaml:"forceDelete"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/azurerm#scale_to_zero_before_deletion AzurermProvider#scale_to_zero_before_deletion}.
	ScaleToZeroBeforeDeletion interface{} `field:"optional" json:"scaleToZeroBeforeDeletion" yaml:"scaleToZeroBeforeDeletion"`
}


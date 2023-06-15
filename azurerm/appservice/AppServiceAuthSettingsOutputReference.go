package appservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/appservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppServiceAuthSettingsOutputReference interface {
	cdktf.ComplexObject
	ActiveDirectory() AppServiceAuthSettingsActiveDirectoryOutputReference
	ActiveDirectoryInput() *AppServiceAuthSettingsActiveDirectory
	AdditionalLoginParams() *map[string]*string
	SetAdditionalLoginParams(val *map[string]*string)
	AdditionalLoginParamsInput() *map[string]*string
	AllowedExternalRedirectUrls() *[]*string
	SetAllowedExternalRedirectUrls(val *[]*string)
	AllowedExternalRedirectUrlsInput() *[]*string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultProvider() *string
	SetDefaultProvider(val *string)
	DefaultProviderInput() *string
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	Facebook() AppServiceAuthSettingsFacebookOutputReference
	FacebookInput() *AppServiceAuthSettingsFacebook
	// Experimental.
	Fqn() *string
	Google() AppServiceAuthSettingsGoogleOutputReference
	GoogleInput() *AppServiceAuthSettingsGoogle
	InternalValue() *AppServiceAuthSettings
	SetInternalValue(val *AppServiceAuthSettings)
	Issuer() *string
	SetIssuer(val *string)
	IssuerInput() *string
	Microsoft() AppServiceAuthSettingsMicrosoftOutputReference
	MicrosoftInput() *AppServiceAuthSettingsMicrosoft
	RuntimeVersion() *string
	SetRuntimeVersion(val *string)
	RuntimeVersionInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TokenRefreshExtensionHours() *float64
	SetTokenRefreshExtensionHours(val *float64)
	TokenRefreshExtensionHoursInput() *float64
	TokenStoreEnabled() interface{}
	SetTokenStoreEnabled(val interface{})
	TokenStoreEnabledInput() interface{}
	Twitter() AppServiceAuthSettingsTwitterOutputReference
	TwitterInput() *AppServiceAuthSettingsTwitter
	UnauthenticatedClientAction() *string
	SetUnauthenticatedClientAction(val *string)
	UnauthenticatedClientActionInput() *string
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	PutActiveDirectory(value *AppServiceAuthSettingsActiveDirectory)
	PutFacebook(value *AppServiceAuthSettingsFacebook)
	PutGoogle(value *AppServiceAuthSettingsGoogle)
	PutMicrosoft(value *AppServiceAuthSettingsMicrosoft)
	PutTwitter(value *AppServiceAuthSettingsTwitter)
	ResetActiveDirectory()
	ResetAdditionalLoginParams()
	ResetAllowedExternalRedirectUrls()
	ResetDefaultProvider()
	ResetFacebook()
	ResetGoogle()
	ResetIssuer()
	ResetMicrosoft()
	ResetRuntimeVersion()
	ResetTokenRefreshExtensionHours()
	ResetTokenStoreEnabled()
	ResetTwitter()
	ResetUnauthenticatedClientAction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AppServiceAuthSettingsOutputReference
type jsiiProxy_AppServiceAuthSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) ActiveDirectory() AppServiceAuthSettingsActiveDirectoryOutputReference {
	var returns AppServiceAuthSettingsActiveDirectoryOutputReference
	_jsii_.Get(
		j,
		"activeDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) ActiveDirectoryInput() *AppServiceAuthSettingsActiveDirectory {
	var returns *AppServiceAuthSettingsActiveDirectory
	_jsii_.Get(
		j,
		"activeDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) AdditionalLoginParams() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalLoginParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) AdditionalLoginParamsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalLoginParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) AllowedExternalRedirectUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedExternalRedirectUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) AllowedExternalRedirectUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedExternalRedirectUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) DefaultProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) DefaultProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) Facebook() AppServiceAuthSettingsFacebookOutputReference {
	var returns AppServiceAuthSettingsFacebookOutputReference
	_jsii_.Get(
		j,
		"facebook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) FacebookInput() *AppServiceAuthSettingsFacebook {
	var returns *AppServiceAuthSettingsFacebook
	_jsii_.Get(
		j,
		"facebookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) Google() AppServiceAuthSettingsGoogleOutputReference {
	var returns AppServiceAuthSettingsGoogleOutputReference
	_jsii_.Get(
		j,
		"google",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) GoogleInput() *AppServiceAuthSettingsGoogle {
	var returns *AppServiceAuthSettingsGoogle
	_jsii_.Get(
		j,
		"googleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) InternalValue() *AppServiceAuthSettings {
	var returns *AppServiceAuthSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) Microsoft() AppServiceAuthSettingsMicrosoftOutputReference {
	var returns AppServiceAuthSettingsMicrosoftOutputReference
	_jsii_.Get(
		j,
		"microsoft",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) MicrosoftInput() *AppServiceAuthSettingsMicrosoft {
	var returns *AppServiceAuthSettingsMicrosoft
	_jsii_.Get(
		j,
		"microsoftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) RuntimeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) TokenRefreshExtensionHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenRefreshExtensionHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) TokenRefreshExtensionHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenRefreshExtensionHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) TokenStoreEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenStoreEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) TokenStoreEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenStoreEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) Twitter() AppServiceAuthSettingsTwitterOutputReference {
	var returns AppServiceAuthSettingsTwitterOutputReference
	_jsii_.Get(
		j,
		"twitter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) TwitterInput() *AppServiceAuthSettingsTwitter {
	var returns *AppServiceAuthSettingsTwitter
	_jsii_.Get(
		j,
		"twitterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) UnauthenticatedClientAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unauthenticatedClientAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference) UnauthenticatedClientActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unauthenticatedClientActionInput",
		&returns,
	)
	return returns
}


func NewAppServiceAuthSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceAuthSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewAppServiceAuthSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppServiceAuthSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appService.AppServiceAuthSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceAuthSettingsOutputReference_Override(a AppServiceAuthSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appService.AppServiceAuthSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference)SetAdditionalLoginParams(val *map[string]*string) {
	if err := j.validateSetAdditionalLoginParamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalLoginParams",
		val,
	)
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference)SetAllowedExternalRedirectUrls(val *[]*string) {
	if err := j.validateSetAllowedExternalRedirectUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedExternalRedirectUrls",
		val,
	)
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference)SetDefaultProvider(val *string) {
	if err := j.validateSetDefaultProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultProvider",
		val,
	)
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference)SetInternalValue(val *AppServiceAuthSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference)SetIssuer(val *string) {
	if err := j.validateSetIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference)SetRuntimeVersion(val *string) {
	if err := j.validateSetRuntimeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference)SetTokenRefreshExtensionHours(val *float64) {
	if err := j.validateSetTokenRefreshExtensionHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenRefreshExtensionHours",
		val,
	)
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference)SetTokenStoreEnabled(val interface{}) {
	if err := j.validateSetTokenStoreEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenStoreEnabled",
		val,
	)
}

func (j *jsiiProxy_AppServiceAuthSettingsOutputReference)SetUnauthenticatedClientAction(val *string) {
	if err := j.validateSetUnauthenticatedClientActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unauthenticatedClientAction",
		val,
	)
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) PutActiveDirectory(value *AppServiceAuthSettingsActiveDirectory) {
	if err := a.validatePutActiveDirectoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putActiveDirectory",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) PutFacebook(value *AppServiceAuthSettingsFacebook) {
	if err := a.validatePutFacebookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFacebook",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) PutGoogle(value *AppServiceAuthSettingsGoogle) {
	if err := a.validatePutGoogleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGoogle",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) PutMicrosoft(value *AppServiceAuthSettingsMicrosoft) {
	if err := a.validatePutMicrosoftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMicrosoft",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) PutTwitter(value *AppServiceAuthSettingsTwitter) {
	if err := a.validatePutTwitterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTwitter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) ResetActiveDirectory() {
	_jsii_.InvokeVoid(
		a,
		"resetActiveDirectory",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) ResetAdditionalLoginParams() {
	_jsii_.InvokeVoid(
		a,
		"resetAdditionalLoginParams",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) ResetAllowedExternalRedirectUrls() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedExternalRedirectUrls",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) ResetDefaultProvider() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultProvider",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) ResetFacebook() {
	_jsii_.InvokeVoid(
		a,
		"resetFacebook",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) ResetGoogle() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) ResetIssuer() {
	_jsii_.InvokeVoid(
		a,
		"resetIssuer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) ResetMicrosoft() {
	_jsii_.InvokeVoid(
		a,
		"resetMicrosoft",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) ResetRuntimeVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetRuntimeVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) ResetTokenRefreshExtensionHours() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenRefreshExtensionHours",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) ResetTokenStoreEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenStoreEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) ResetTwitter() {
	_jsii_.InvokeVoid(
		a,
		"resetTwitter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) ResetUnauthenticatedClientAction() {
	_jsii_.InvokeVoid(
		a,
		"resetUnauthenticatedClientAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceAuthSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package appserviceslot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/appserviceslot/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AppServiceSlotAuthSettingsOutputReference interface {
	cdktf.ComplexObject
	ActiveDirectory() AppServiceSlotAuthSettingsActiveDirectoryOutputReference
	ActiveDirectoryInput() *AppServiceSlotAuthSettingsActiveDirectory
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
	Facebook() AppServiceSlotAuthSettingsFacebookOutputReference
	FacebookInput() *AppServiceSlotAuthSettingsFacebook
	// Experimental.
	Fqn() *string
	Google() AppServiceSlotAuthSettingsGoogleOutputReference
	GoogleInput() *AppServiceSlotAuthSettingsGoogle
	InternalValue() *AppServiceSlotAuthSettings
	SetInternalValue(val *AppServiceSlotAuthSettings)
	Issuer() *string
	SetIssuer(val *string)
	IssuerInput() *string
	Microsoft() AppServiceSlotAuthSettingsMicrosoftOutputReference
	MicrosoftInput() *AppServiceSlotAuthSettingsMicrosoft
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
	Twitter() AppServiceSlotAuthSettingsTwitterOutputReference
	TwitterInput() *AppServiceSlotAuthSettingsTwitter
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
	PutActiveDirectory(value *AppServiceSlotAuthSettingsActiveDirectory)
	PutFacebook(value *AppServiceSlotAuthSettingsFacebook)
	PutGoogle(value *AppServiceSlotAuthSettingsGoogle)
	PutMicrosoft(value *AppServiceSlotAuthSettingsMicrosoft)
	PutTwitter(value *AppServiceSlotAuthSettingsTwitter)
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

// The jsii proxy struct for AppServiceSlotAuthSettingsOutputReference
type jsiiProxy_AppServiceSlotAuthSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ActiveDirectory() AppServiceSlotAuthSettingsActiveDirectoryOutputReference {
	var returns AppServiceSlotAuthSettingsActiveDirectoryOutputReference
	_jsii_.Get(
		j,
		"activeDirectory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ActiveDirectoryInput() *AppServiceSlotAuthSettingsActiveDirectory {
	var returns *AppServiceSlotAuthSettingsActiveDirectory
	_jsii_.Get(
		j,
		"activeDirectoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) AdditionalLoginParams() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalLoginParams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) AdditionalLoginParamsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"additionalLoginParamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) AllowedExternalRedirectUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedExternalRedirectUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) AllowedExternalRedirectUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedExternalRedirectUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) DefaultProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) DefaultProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) Facebook() AppServiceSlotAuthSettingsFacebookOutputReference {
	var returns AppServiceSlotAuthSettingsFacebookOutputReference
	_jsii_.Get(
		j,
		"facebook",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) FacebookInput() *AppServiceSlotAuthSettingsFacebook {
	var returns *AppServiceSlotAuthSettingsFacebook
	_jsii_.Get(
		j,
		"facebookInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) Google() AppServiceSlotAuthSettingsGoogleOutputReference {
	var returns AppServiceSlotAuthSettingsGoogleOutputReference
	_jsii_.Get(
		j,
		"google",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) GoogleInput() *AppServiceSlotAuthSettingsGoogle {
	var returns *AppServiceSlotAuthSettingsGoogle
	_jsii_.Get(
		j,
		"googleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) InternalValue() *AppServiceSlotAuthSettings {
	var returns *AppServiceSlotAuthSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) Microsoft() AppServiceSlotAuthSettingsMicrosoftOutputReference {
	var returns AppServiceSlotAuthSettingsMicrosoftOutputReference
	_jsii_.Get(
		j,
		"microsoft",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) MicrosoftInput() *AppServiceSlotAuthSettingsMicrosoft {
	var returns *AppServiceSlotAuthSettingsMicrosoft
	_jsii_.Get(
		j,
		"microsoftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) RuntimeVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) RuntimeVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"runtimeVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) TokenRefreshExtensionHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenRefreshExtensionHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) TokenRefreshExtensionHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenRefreshExtensionHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) TokenStoreEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenStoreEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) TokenStoreEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenStoreEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) Twitter() AppServiceSlotAuthSettingsTwitterOutputReference {
	var returns AppServiceSlotAuthSettingsTwitterOutputReference
	_jsii_.Get(
		j,
		"twitter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) TwitterInput() *AppServiceSlotAuthSettingsTwitter {
	var returns *AppServiceSlotAuthSettingsTwitter
	_jsii_.Get(
		j,
		"twitterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) UnauthenticatedClientAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unauthenticatedClientAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) UnauthenticatedClientActionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"unauthenticatedClientActionInput",
		&returns,
	)
	return returns
}


func NewAppServiceSlotAuthSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AppServiceSlotAuthSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewAppServiceSlotAuthSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AppServiceSlotAuthSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotAuthSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAppServiceSlotAuthSettingsOutputReference_Override(a AppServiceSlotAuthSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.appServiceSlot.AppServiceSlotAuthSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference)SetAdditionalLoginParams(val *map[string]*string) {
	if err := j.validateSetAdditionalLoginParamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalLoginParams",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference)SetAllowedExternalRedirectUrls(val *[]*string) {
	if err := j.validateSetAllowedExternalRedirectUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedExternalRedirectUrls",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference)SetDefaultProvider(val *string) {
	if err := j.validateSetDefaultProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultProvider",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference)SetInternalValue(val *AppServiceSlotAuthSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference)SetIssuer(val *string) {
	if err := j.validateSetIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference)SetRuntimeVersion(val *string) {
	if err := j.validateSetRuntimeVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeVersion",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference)SetTokenRefreshExtensionHours(val *float64) {
	if err := j.validateSetTokenRefreshExtensionHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenRefreshExtensionHours",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference)SetTokenStoreEnabled(val interface{}) {
	if err := j.validateSetTokenStoreEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenStoreEnabled",
		val,
	)
}

func (j *jsiiProxy_AppServiceSlotAuthSettingsOutputReference)SetUnauthenticatedClientAction(val *string) {
	if err := j.validateSetUnauthenticatedClientActionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unauthenticatedClientAction",
		val,
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) PutActiveDirectory(value *AppServiceSlotAuthSettingsActiveDirectory) {
	if err := a.validatePutActiveDirectoryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putActiveDirectory",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) PutFacebook(value *AppServiceSlotAuthSettingsFacebook) {
	if err := a.validatePutFacebookParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFacebook",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) PutGoogle(value *AppServiceSlotAuthSettingsGoogle) {
	if err := a.validatePutGoogleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putGoogle",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) PutMicrosoft(value *AppServiceSlotAuthSettingsMicrosoft) {
	if err := a.validatePutMicrosoftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putMicrosoft",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) PutTwitter(value *AppServiceSlotAuthSettingsTwitter) {
	if err := a.validatePutTwitterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putTwitter",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ResetActiveDirectory() {
	_jsii_.InvokeVoid(
		a,
		"resetActiveDirectory",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ResetAdditionalLoginParams() {
	_jsii_.InvokeVoid(
		a,
		"resetAdditionalLoginParams",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ResetAllowedExternalRedirectUrls() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowedExternalRedirectUrls",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ResetDefaultProvider() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultProvider",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ResetFacebook() {
	_jsii_.InvokeVoid(
		a,
		"resetFacebook",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ResetGoogle() {
	_jsii_.InvokeVoid(
		a,
		"resetGoogle",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ResetIssuer() {
	_jsii_.InvokeVoid(
		a,
		"resetIssuer",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ResetMicrosoft() {
	_jsii_.InvokeVoid(
		a,
		"resetMicrosoft",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ResetRuntimeVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetRuntimeVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ResetTokenRefreshExtensionHours() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenRefreshExtensionHours",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ResetTokenStoreEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenStoreEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ResetTwitter() {
	_jsii_.InvokeVoid(
		a,
		"resetTwitter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ResetUnauthenticatedClientAction() {
	_jsii_.InvokeVoid(
		a,
		"resetUnauthenticatedClientAction",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_AppServiceSlotAuthSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


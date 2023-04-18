package windowswebappslot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/windowswebappslot/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WindowsWebAppSlotAuthSettingsV2LoginOutputReference interface {
	cdktf.ComplexObject
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
	CookieExpirationConvention() *string
	SetCookieExpirationConvention(val *string)
	CookieExpirationConventionInput() *string
	CookieExpirationTime() *string
	SetCookieExpirationTime(val *string)
	CookieExpirationTimeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *WindowsWebAppSlotAuthSettingsV2Login
	SetInternalValue(val *WindowsWebAppSlotAuthSettingsV2Login)
	LogoutEndpoint() *string
	SetLogoutEndpoint(val *string)
	LogoutEndpointInput() *string
	NonceExpirationTime() *string
	SetNonceExpirationTime(val *string)
	NonceExpirationTimeInput() *string
	PreserveUrlFragmentsForLogins() interface{}
	SetPreserveUrlFragmentsForLogins(val interface{})
	PreserveUrlFragmentsForLoginsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TokenRefreshExtensionTime() *float64
	SetTokenRefreshExtensionTime(val *float64)
	TokenRefreshExtensionTimeInput() *float64
	TokenStoreEnabled() interface{}
	SetTokenStoreEnabled(val interface{})
	TokenStoreEnabledInput() interface{}
	TokenStorePath() *string
	SetTokenStorePath(val *string)
	TokenStorePathInput() *string
	TokenStoreSasSettingName() *string
	SetTokenStoreSasSettingName(val *string)
	TokenStoreSasSettingNameInput() *string
	ValidateNonce() interface{}
	SetValidateNonce(val interface{})
	ValidateNonceInput() interface{}
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
	ResetAllowedExternalRedirectUrls()
	ResetCookieExpirationConvention()
	ResetCookieExpirationTime()
	ResetLogoutEndpoint()
	ResetNonceExpirationTime()
	ResetPreserveUrlFragmentsForLogins()
	ResetTokenRefreshExtensionTime()
	ResetTokenStoreEnabled()
	ResetTokenStorePath()
	ResetTokenStoreSasSettingName()
	ResetValidateNonce()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WindowsWebAppSlotAuthSettingsV2LoginOutputReference
type jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) AllowedExternalRedirectUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedExternalRedirectUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) AllowedExternalRedirectUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedExternalRedirectUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) CookieExpirationConvention() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieExpirationConvention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) CookieExpirationConventionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieExpirationConventionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) CookieExpirationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieExpirationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) CookieExpirationTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieExpirationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) InternalValue() *WindowsWebAppSlotAuthSettingsV2Login {
	var returns *WindowsWebAppSlotAuthSettingsV2Login
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) LogoutEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoutEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) LogoutEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoutEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) NonceExpirationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nonceExpirationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) NonceExpirationTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nonceExpirationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) PreserveUrlFragmentsForLogins() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveUrlFragmentsForLogins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) PreserveUrlFragmentsForLoginsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveUrlFragmentsForLoginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) TokenRefreshExtensionTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenRefreshExtensionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) TokenRefreshExtensionTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenRefreshExtensionTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) TokenStoreEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenStoreEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) TokenStoreEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenStoreEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) TokenStorePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenStorePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) TokenStorePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenStorePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) TokenStoreSasSettingName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenStoreSasSettingName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) TokenStoreSasSettingNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenStoreSasSettingNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) ValidateNonce() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateNonce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) ValidateNonceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateNonceInput",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSlotAuthSettingsV2LoginOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppSlotAuthSettingsV2LoginOutputReference {
	_init_.Initialize()

	if err := validateNewWindowsWebAppSlotAuthSettingsV2LoginOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebAppSlot.WindowsWebAppSlotAuthSettingsV2LoginOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppSlotAuthSettingsV2LoginOutputReference_Override(w WindowsWebAppSlotAuthSettingsV2LoginOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebAppSlot.WindowsWebAppSlotAuthSettingsV2LoginOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference)SetAllowedExternalRedirectUrls(val *[]*string) {
	if err := j.validateSetAllowedExternalRedirectUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedExternalRedirectUrls",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference)SetCookieExpirationConvention(val *string) {
	if err := j.validateSetCookieExpirationConventionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cookieExpirationConvention",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference)SetCookieExpirationTime(val *string) {
	if err := j.validateSetCookieExpirationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cookieExpirationTime",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference)SetInternalValue(val *WindowsWebAppSlotAuthSettingsV2Login) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference)SetLogoutEndpoint(val *string) {
	if err := j.validateSetLogoutEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logoutEndpoint",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference)SetNonceExpirationTime(val *string) {
	if err := j.validateSetNonceExpirationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nonceExpirationTime",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference)SetPreserveUrlFragmentsForLogins(val interface{}) {
	if err := j.validateSetPreserveUrlFragmentsForLoginsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveUrlFragmentsForLogins",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference)SetTokenRefreshExtensionTime(val *float64) {
	if err := j.validateSetTokenRefreshExtensionTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenRefreshExtensionTime",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference)SetTokenStoreEnabled(val interface{}) {
	if err := j.validateSetTokenStoreEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenStoreEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference)SetTokenStorePath(val *string) {
	if err := j.validateSetTokenStorePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenStorePath",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference)SetTokenStoreSasSettingName(val *string) {
	if err := j.validateSetTokenStoreSasSettingNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenStoreSasSettingName",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference)SetValidateNonce(val interface{}) {
	if err := j.validateSetValidateNonceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validateNonce",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) ResetAllowedExternalRedirectUrls() {
	_jsii_.InvokeVoid(
		w,
		"resetAllowedExternalRedirectUrls",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) ResetCookieExpirationConvention() {
	_jsii_.InvokeVoid(
		w,
		"resetCookieExpirationConvention",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) ResetCookieExpirationTime() {
	_jsii_.InvokeVoid(
		w,
		"resetCookieExpirationTime",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) ResetLogoutEndpoint() {
	_jsii_.InvokeVoid(
		w,
		"resetLogoutEndpoint",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) ResetNonceExpirationTime() {
	_jsii_.InvokeVoid(
		w,
		"resetNonceExpirationTime",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) ResetPreserveUrlFragmentsForLogins() {
	_jsii_.InvokeVoid(
		w,
		"resetPreserveUrlFragmentsForLogins",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) ResetTokenRefreshExtensionTime() {
	_jsii_.InvokeVoid(
		w,
		"resetTokenRefreshExtensionTime",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) ResetTokenStoreEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetTokenStoreEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) ResetTokenStorePath() {
	_jsii_.InvokeVoid(
		w,
		"resetTokenStorePath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) ResetTokenStoreSasSettingName() {
	_jsii_.InvokeVoid(
		w,
		"resetTokenStoreSasSettingName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) ResetValidateNonce() {
	_jsii_.InvokeVoid(
		w,
		"resetValidateNonce",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotAuthSettingsV2LoginOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


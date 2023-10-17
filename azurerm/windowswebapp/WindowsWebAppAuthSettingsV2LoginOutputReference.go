// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/windowswebapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WindowsWebAppAuthSettingsV2LoginOutputReference interface {
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
	InternalValue() *WindowsWebAppAuthSettingsV2Login
	SetInternalValue(val *WindowsWebAppAuthSettingsV2Login)
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

// The jsii proxy struct for WindowsWebAppAuthSettingsV2LoginOutputReference
type jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) AllowedExternalRedirectUrls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedExternalRedirectUrls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) AllowedExternalRedirectUrlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedExternalRedirectUrlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) CookieExpirationConvention() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieExpirationConvention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) CookieExpirationConventionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieExpirationConventionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) CookieExpirationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieExpirationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) CookieExpirationTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cookieExpirationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) InternalValue() *WindowsWebAppAuthSettingsV2Login {
	var returns *WindowsWebAppAuthSettingsV2Login
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) LogoutEndpoint() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoutEndpoint",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) LogoutEndpointInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logoutEndpointInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) NonceExpirationTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nonceExpirationTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) NonceExpirationTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nonceExpirationTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) PreserveUrlFragmentsForLogins() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveUrlFragmentsForLogins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) PreserveUrlFragmentsForLoginsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"preserveUrlFragmentsForLoginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) TokenRefreshExtensionTime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenRefreshExtensionTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) TokenRefreshExtensionTimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenRefreshExtensionTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) TokenStoreEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenStoreEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) TokenStoreEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenStoreEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) TokenStorePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenStorePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) TokenStorePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenStorePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) TokenStoreSasSettingName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenStoreSasSettingName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) TokenStoreSasSettingNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenStoreSasSettingNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) ValidateNonce() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateNonce",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) ValidateNonceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateNonceInput",
		&returns,
	)
	return returns
}


func NewWindowsWebAppAuthSettingsV2LoginOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WindowsWebAppAuthSettingsV2LoginOutputReference {
	_init_.Initialize()

	if err := validateNewWindowsWebAppAuthSettingsV2LoginOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppAuthSettingsV2LoginOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWindowsWebAppAuthSettingsV2LoginOutputReference_Override(w WindowsWebAppAuthSettingsV2LoginOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebApp.WindowsWebAppAuthSettingsV2LoginOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference)SetAllowedExternalRedirectUrls(val *[]*string) {
	if err := j.validateSetAllowedExternalRedirectUrlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedExternalRedirectUrls",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference)SetCookieExpirationConvention(val *string) {
	if err := j.validateSetCookieExpirationConventionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cookieExpirationConvention",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference)SetCookieExpirationTime(val *string) {
	if err := j.validateSetCookieExpirationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cookieExpirationTime",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference)SetInternalValue(val *WindowsWebAppAuthSettingsV2Login) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference)SetLogoutEndpoint(val *string) {
	if err := j.validateSetLogoutEndpointParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logoutEndpoint",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference)SetNonceExpirationTime(val *string) {
	if err := j.validateSetNonceExpirationTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nonceExpirationTime",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference)SetPreserveUrlFragmentsForLogins(val interface{}) {
	if err := j.validateSetPreserveUrlFragmentsForLoginsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preserveUrlFragmentsForLogins",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference)SetTokenRefreshExtensionTime(val *float64) {
	if err := j.validateSetTokenRefreshExtensionTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenRefreshExtensionTime",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference)SetTokenStoreEnabled(val interface{}) {
	if err := j.validateSetTokenStoreEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenStoreEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference)SetTokenStorePath(val *string) {
	if err := j.validateSetTokenStorePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenStorePath",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference)SetTokenStoreSasSettingName(val *string) {
	if err := j.validateSetTokenStoreSasSettingNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenStoreSasSettingName",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference)SetValidateNonce(val interface{}) {
	if err := j.validateSetValidateNonceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"validateNonce",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) ResetAllowedExternalRedirectUrls() {
	_jsii_.InvokeVoid(
		w,
		"resetAllowedExternalRedirectUrls",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) ResetCookieExpirationConvention() {
	_jsii_.InvokeVoid(
		w,
		"resetCookieExpirationConvention",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) ResetCookieExpirationTime() {
	_jsii_.InvokeVoid(
		w,
		"resetCookieExpirationTime",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) ResetLogoutEndpoint() {
	_jsii_.InvokeVoid(
		w,
		"resetLogoutEndpoint",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) ResetNonceExpirationTime() {
	_jsii_.InvokeVoid(
		w,
		"resetNonceExpirationTime",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) ResetPreserveUrlFragmentsForLogins() {
	_jsii_.InvokeVoid(
		w,
		"resetPreserveUrlFragmentsForLogins",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) ResetTokenRefreshExtensionTime() {
	_jsii_.InvokeVoid(
		w,
		"resetTokenRefreshExtensionTime",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) ResetTokenStoreEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetTokenStoreEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) ResetTokenStorePath() {
	_jsii_.InvokeVoid(
		w,
		"resetTokenStorePath",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) ResetTokenStoreSasSettingName() {
	_jsii_.InvokeVoid(
		w,
		"resetTokenStoreSasSettingName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) ResetValidateNonce() {
	_jsii_.InvokeVoid(
		w,
		"resetValidateNonce",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_WindowsWebAppAuthSettingsV2LoginOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


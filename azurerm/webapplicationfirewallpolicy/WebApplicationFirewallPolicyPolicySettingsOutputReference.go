// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package webapplicationfirewallpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/webapplicationfirewallpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WebApplicationFirewallPolicyPolicySettingsOutputReference interface {
	cdktf.ComplexObject
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	FileUploadEnforcement() interface{}
	SetFileUploadEnforcement(val interface{})
	FileUploadEnforcementInput() interface{}
	FileUploadLimitInMb() *float64
	SetFileUploadLimitInMb(val *float64)
	FileUploadLimitInMbInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *WebApplicationFirewallPolicyPolicySettings
	SetInternalValue(val *WebApplicationFirewallPolicyPolicySettings)
	JsChallengeCookieExpirationInMinutes() *float64
	SetJsChallengeCookieExpirationInMinutes(val *float64)
	JsChallengeCookieExpirationInMinutesInput() *float64
	LogScrubbing() WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference
	LogScrubbingInput() *WebApplicationFirewallPolicyPolicySettingsLogScrubbing
	MaxRequestBodySizeInKb() *float64
	SetMaxRequestBodySizeInKb(val *float64)
	MaxRequestBodySizeInKbInput() *float64
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	RequestBodyCheck() interface{}
	SetRequestBodyCheck(val interface{})
	RequestBodyCheckInput() interface{}
	RequestBodyEnforcement() interface{}
	SetRequestBodyEnforcement(val interface{})
	RequestBodyEnforcementInput() interface{}
	RequestBodyInspectLimitInKb() *float64
	SetRequestBodyInspectLimitInKb(val *float64)
	RequestBodyInspectLimitInKbInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
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
	PutLogScrubbing(value *WebApplicationFirewallPolicyPolicySettingsLogScrubbing)
	ResetEnabled()
	ResetFileUploadEnforcement()
	ResetFileUploadLimitInMb()
	ResetJsChallengeCookieExpirationInMinutes()
	ResetLogScrubbing()
	ResetMaxRequestBodySizeInKb()
	ResetMode()
	ResetRequestBodyCheck()
	ResetRequestBodyEnforcement()
	ResetRequestBodyInspectLimitInKb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for WebApplicationFirewallPolicyPolicySettingsOutputReference
type jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) FileUploadEnforcement() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fileUploadEnforcement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) FileUploadEnforcementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"fileUploadEnforcementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) FileUploadLimitInMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fileUploadLimitInMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) FileUploadLimitInMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fileUploadLimitInMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) InternalValue() *WebApplicationFirewallPolicyPolicySettings {
	var returns *WebApplicationFirewallPolicyPolicySettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) JsChallengeCookieExpirationInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jsChallengeCookieExpirationInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) JsChallengeCookieExpirationInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"jsChallengeCookieExpirationInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) LogScrubbing() WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference {
	var returns WebApplicationFirewallPolicyPolicySettingsLogScrubbingOutputReference
	_jsii_.Get(
		j,
		"logScrubbing",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) LogScrubbingInput() *WebApplicationFirewallPolicyPolicySettingsLogScrubbing {
	var returns *WebApplicationFirewallPolicyPolicySettingsLogScrubbing
	_jsii_.Get(
		j,
		"logScrubbingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) MaxRequestBodySizeInKb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRequestBodySizeInKb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) MaxRequestBodySizeInKbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRequestBodySizeInKbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) RequestBodyCheck() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestBodyCheck",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) RequestBodyCheckInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestBodyCheckInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) RequestBodyEnforcement() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestBodyEnforcement",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) RequestBodyEnforcementInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestBodyEnforcementInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) RequestBodyInspectLimitInKb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestBodyInspectLimitInKb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) RequestBodyInspectLimitInKbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestBodyInspectLimitInKbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewWebApplicationFirewallPolicyPolicySettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) WebApplicationFirewallPolicyPolicySettingsOutputReference {
	_init_.Initialize()

	if err := validateNewWebApplicationFirewallPolicyPolicySettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.webApplicationFirewallPolicy.WebApplicationFirewallPolicyPolicySettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewWebApplicationFirewallPolicyPolicySettingsOutputReference_Override(w WebApplicationFirewallPolicyPolicySettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.webApplicationFirewallPolicy.WebApplicationFirewallPolicyPolicySettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		w,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference)SetFileUploadEnforcement(val interface{}) {
	if err := j.validateSetFileUploadEnforcementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileUploadEnforcement",
		val,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference)SetFileUploadLimitInMb(val *float64) {
	if err := j.validateSetFileUploadLimitInMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fileUploadLimitInMb",
		val,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference)SetInternalValue(val *WebApplicationFirewallPolicyPolicySettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference)SetJsChallengeCookieExpirationInMinutes(val *float64) {
	if err := j.validateSetJsChallengeCookieExpirationInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jsChallengeCookieExpirationInMinutes",
		val,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference)SetMaxRequestBodySizeInKb(val *float64) {
	if err := j.validateSetMaxRequestBodySizeInKbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRequestBodySizeInKb",
		val,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference)SetRequestBodyCheck(val interface{}) {
	if err := j.validateSetRequestBodyCheckParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestBodyCheck",
		val,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference)SetRequestBodyEnforcement(val interface{}) {
	if err := j.validateSetRequestBodyEnforcementParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestBodyEnforcement",
		val,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference)SetRequestBodyInspectLimitInKb(val *float64) {
	if err := j.validateSetRequestBodyInspectLimitInKbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestBodyInspectLimitInKb",
		val,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) PutLogScrubbing(value *WebApplicationFirewallPolicyPolicySettingsLogScrubbing) {
	if err := w.validatePutLogScrubbingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putLogScrubbing",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) ResetEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) ResetFileUploadEnforcement() {
	_jsii_.InvokeVoid(
		w,
		"resetFileUploadEnforcement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) ResetFileUploadLimitInMb() {
	_jsii_.InvokeVoid(
		w,
		"resetFileUploadLimitInMb",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) ResetJsChallengeCookieExpirationInMinutes() {
	_jsii_.InvokeVoid(
		w,
		"resetJsChallengeCookieExpirationInMinutes",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) ResetLogScrubbing() {
	_jsii_.InvokeVoid(
		w,
		"resetLogScrubbing",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) ResetMaxRequestBodySizeInKb() {
	_jsii_.InvokeVoid(
		w,
		"resetMaxRequestBodySizeInKb",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) ResetMode() {
	_jsii_.InvokeVoid(
		w,
		"resetMode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) ResetRequestBodyCheck() {
	_jsii_.InvokeVoid(
		w,
		"resetRequestBodyCheck",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) ResetRequestBodyEnforcement() {
	_jsii_.InvokeVoid(
		w,
		"resetRequestBodyEnforcement",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) ResetRequestBodyInspectLimitInKb() {
	_jsii_.InvokeVoid(
		w,
		"resetRequestBodyInspectLimitInKb",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (w *jsiiProxy_WebApplicationFirewallPolicyPolicySettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


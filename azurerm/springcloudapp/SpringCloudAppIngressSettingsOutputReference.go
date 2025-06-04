// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/springcloudapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpringCloudAppIngressSettingsOutputReference interface {
	cdktf.ComplexObject
	BackendProtocol() *string
	SetBackendProtocol(val *string)
	BackendProtocolInput() *string
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
	// Experimental.
	Fqn() *string
	InternalValue() *SpringCloudAppIngressSettings
	SetInternalValue(val *SpringCloudAppIngressSettings)
	ReadTimeoutInSeconds() *float64
	SetReadTimeoutInSeconds(val *float64)
	ReadTimeoutInSecondsInput() *float64
	SendTimeoutInSeconds() *float64
	SetSendTimeoutInSeconds(val *float64)
	SendTimeoutInSecondsInput() *float64
	SessionAffinity() *string
	SetSessionAffinity(val *string)
	SessionAffinityInput() *string
	SessionCookieMaxAge() *float64
	SetSessionCookieMaxAge(val *float64)
	SessionCookieMaxAgeInput() *float64
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
	ResetBackendProtocol()
	ResetReadTimeoutInSeconds()
	ResetSendTimeoutInSeconds()
	ResetSessionAffinity()
	ResetSessionCookieMaxAge()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SpringCloudAppIngressSettingsOutputReference
type jsiiProxy_SpringCloudAppIngressSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) BackendProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) BackendProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) InternalValue() *SpringCloudAppIngressSettings {
	var returns *SpringCloudAppIngressSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) ReadTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) ReadTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"readTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) SendTimeoutInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sendTimeoutInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) SendTimeoutInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sendTimeoutInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) SessionAffinity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionAffinity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) SessionAffinityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sessionAffinityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) SessionCookieMaxAge() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionCookieMaxAge",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) SessionCookieMaxAgeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"sessionCookieMaxAgeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSpringCloudAppIngressSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SpringCloudAppIngressSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewSpringCloudAppIngressSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpringCloudAppIngressSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudApp.SpringCloudAppIngressSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSpringCloudAppIngressSettingsOutputReference_Override(s SpringCloudAppIngressSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudApp.SpringCloudAppIngressSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference)SetBackendProtocol(val *string) {
	if err := j.validateSetBackendProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backendProtocol",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference)SetInternalValue(val *SpringCloudAppIngressSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference)SetReadTimeoutInSeconds(val *float64) {
	if err := j.validateSetReadTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference)SetSendTimeoutInSeconds(val *float64) {
	if err := j.validateSetSendTimeoutInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sendTimeoutInSeconds",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference)SetSessionAffinity(val *string) {
	if err := j.validateSetSessionAffinityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionAffinity",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference)SetSessionCookieMaxAge(val *float64) {
	if err := j.validateSetSessionCookieMaxAgeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sessionCookieMaxAge",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpringCloudAppIngressSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) ResetBackendProtocol() {
	_jsii_.InvokeVoid(
		s,
		"resetBackendProtocol",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) ResetReadTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetReadTimeoutInSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) ResetSendTimeoutInSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetSendTimeoutInSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) ResetSessionAffinity() {
	_jsii_.InvokeVoid(
		s,
		"resetSessionAffinity",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) ResetSessionCookieMaxAge() {
	_jsii_.InvokeVoid(
		s,
		"resetSessionCookieMaxAge",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudAppIngressSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


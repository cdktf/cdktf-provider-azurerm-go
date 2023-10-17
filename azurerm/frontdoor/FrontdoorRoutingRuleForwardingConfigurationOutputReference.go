// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package frontdoor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/frontdoor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FrontdoorRoutingRuleForwardingConfigurationOutputReference interface {
	cdktf.ComplexObject
	BackendPoolName() *string
	SetBackendPoolName(val *string)
	BackendPoolNameInput() *string
	CacheDuration() *string
	SetCacheDuration(val *string)
	CacheDurationInput() *string
	CacheEnabled() interface{}
	SetCacheEnabled(val interface{})
	CacheEnabledInput() interface{}
	CacheQueryParameters() *[]*string
	SetCacheQueryParameters(val *[]*string)
	CacheQueryParametersInput() *[]*string
	CacheQueryParameterStripDirective() *string
	SetCacheQueryParameterStripDirective(val *string)
	CacheQueryParameterStripDirectiveInput() *string
	CacheUseDynamicCompression() interface{}
	SetCacheUseDynamicCompression(val interface{})
	CacheUseDynamicCompressionInput() interface{}
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
	CustomForwardingPath() *string
	SetCustomForwardingPath(val *string)
	CustomForwardingPathInput() *string
	ForwardingProtocol() *string
	SetForwardingProtocol(val *string)
	ForwardingProtocolInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *FrontdoorRoutingRuleForwardingConfiguration
	SetInternalValue(val *FrontdoorRoutingRuleForwardingConfiguration)
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
	ResetCacheDuration()
	ResetCacheEnabled()
	ResetCacheQueryParameters()
	ResetCacheQueryParameterStripDirective()
	ResetCacheUseDynamicCompression()
	ResetCustomForwardingPath()
	ResetForwardingProtocol()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FrontdoorRoutingRuleForwardingConfigurationOutputReference
type jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) BackendPoolName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendPoolName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) BackendPoolNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendPoolNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) CacheDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) CacheDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) CacheEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) CacheEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) CacheQueryParameters() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheQueryParameters",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) CacheQueryParametersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cacheQueryParametersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) CacheQueryParameterStripDirective() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheQueryParameterStripDirective",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) CacheQueryParameterStripDirectiveInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cacheQueryParameterStripDirectiveInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) CacheUseDynamicCompression() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheUseDynamicCompression",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) CacheUseDynamicCompressionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cacheUseDynamicCompressionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) CustomForwardingPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customForwardingPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) CustomForwardingPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customForwardingPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) ForwardingProtocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardingProtocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) ForwardingProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"forwardingProtocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) InternalValue() *FrontdoorRoutingRuleForwardingConfiguration {
	var returns *FrontdoorRoutingRuleForwardingConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewFrontdoorRoutingRuleForwardingConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FrontdoorRoutingRuleForwardingConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewFrontdoorRoutingRuleForwardingConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorRoutingRuleForwardingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFrontdoorRoutingRuleForwardingConfigurationOutputReference_Override(f FrontdoorRoutingRuleForwardingConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.frontdoor.FrontdoorRoutingRuleForwardingConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference)SetBackendPoolName(val *string) {
	if err := j.validateSetBackendPoolNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backendPoolName",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference)SetCacheDuration(val *string) {
	if err := j.validateSetCacheDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheDuration",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference)SetCacheEnabled(val interface{}) {
	if err := j.validateSetCacheEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheEnabled",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference)SetCacheQueryParameters(val *[]*string) {
	if err := j.validateSetCacheQueryParametersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheQueryParameters",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference)SetCacheQueryParameterStripDirective(val *string) {
	if err := j.validateSetCacheQueryParameterStripDirectiveParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheQueryParameterStripDirective",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference)SetCacheUseDynamicCompression(val interface{}) {
	if err := j.validateSetCacheUseDynamicCompressionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cacheUseDynamicCompression",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference)SetCustomForwardingPath(val *string) {
	if err := j.validateSetCustomForwardingPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customForwardingPath",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference)SetForwardingProtocol(val *string) {
	if err := j.validateSetForwardingProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"forwardingProtocol",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference)SetInternalValue(val *FrontdoorRoutingRuleForwardingConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) ResetCacheDuration() {
	_jsii_.InvokeVoid(
		f,
		"resetCacheDuration",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) ResetCacheEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetCacheEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) ResetCacheQueryParameters() {
	_jsii_.InvokeVoid(
		f,
		"resetCacheQueryParameters",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) ResetCacheQueryParameterStripDirective() {
	_jsii_.InvokeVoid(
		f,
		"resetCacheQueryParameterStripDirective",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) ResetCacheUseDynamicCompression() {
	_jsii_.InvokeVoid(
		f,
		"resetCacheUseDynamicCompression",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) ResetCustomForwardingPath() {
	_jsii_.InvokeVoid(
		f,
		"resetCustomForwardingPath",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) ResetForwardingProtocol() {
	_jsii_.InvokeVoid(
		f,
		"resetForwardingProtocol",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FrontdoorRoutingRuleForwardingConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


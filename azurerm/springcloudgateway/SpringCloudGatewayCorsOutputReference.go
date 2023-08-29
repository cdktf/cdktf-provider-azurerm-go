// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package springcloudgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/springcloudgateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SpringCloudGatewayCorsOutputReference interface {
	cdktf.ComplexObject
	AllowedHeaders() *[]*string
	SetAllowedHeaders(val *[]*string)
	AllowedHeadersInput() *[]*string
	AllowedMethods() *[]*string
	SetAllowedMethods(val *[]*string)
	AllowedMethodsInput() *[]*string
	AllowedOriginPatterns() *[]*string
	SetAllowedOriginPatterns(val *[]*string)
	AllowedOriginPatternsInput() *[]*string
	AllowedOrigins() *[]*string
	SetAllowedOrigins(val *[]*string)
	AllowedOriginsInput() *[]*string
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
	CredentialsAllowed() interface{}
	SetCredentialsAllowed(val interface{})
	CredentialsAllowedInput() interface{}
	ExposedHeaders() *[]*string
	SetExposedHeaders(val *[]*string)
	ExposedHeadersInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *SpringCloudGatewayCors
	SetInternalValue(val *SpringCloudGatewayCors)
	MaxAgeSeconds() *float64
	SetMaxAgeSeconds(val *float64)
	MaxAgeSecondsInput() *float64
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
	ResetAllowedHeaders()
	ResetAllowedMethods()
	ResetAllowedOriginPatterns()
	ResetAllowedOrigins()
	ResetCredentialsAllowed()
	ResetExposedHeaders()
	ResetMaxAgeSeconds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SpringCloudGatewayCorsOutputReference
type jsiiProxy_SpringCloudGatewayCorsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference) AllowedHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference) AllowedHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference) AllowedMethods() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMethods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference) AllowedMethodsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMethodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference) AllowedOriginPatterns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOriginPatterns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference) AllowedOriginPatternsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOriginPatternsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference) AllowedOrigins() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOrigins",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference) AllowedOriginsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedOriginsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference) CredentialsAllowed() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"credentialsAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference) CredentialsAllowedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"credentialsAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference) ExposedHeaders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exposedHeaders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference) ExposedHeadersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"exposedHeadersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference) InternalValue() *SpringCloudGatewayCors {
	var returns *SpringCloudGatewayCors
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference) MaxAgeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAgeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference) MaxAgeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxAgeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewSpringCloudGatewayCorsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) SpringCloudGatewayCorsOutputReference {
	_init_.Initialize()

	if err := validateNewSpringCloudGatewayCorsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_SpringCloudGatewayCorsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudGateway.SpringCloudGatewayCorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewSpringCloudGatewayCorsOutputReference_Override(s SpringCloudGatewayCorsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.springCloudGateway.SpringCloudGatewayCorsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		s,
	)
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference)SetAllowedHeaders(val *[]*string) {
	if err := j.validateSetAllowedHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedHeaders",
		val,
	)
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference)SetAllowedMethods(val *[]*string) {
	if err := j.validateSetAllowedMethodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedMethods",
		val,
	)
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference)SetAllowedOriginPatterns(val *[]*string) {
	if err := j.validateSetAllowedOriginPatternsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedOriginPatterns",
		val,
	)
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference)SetAllowedOrigins(val *[]*string) {
	if err := j.validateSetAllowedOriginsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedOrigins",
		val,
	)
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference)SetCredentialsAllowed(val interface{}) {
	if err := j.validateSetCredentialsAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentialsAllowed",
		val,
	)
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference)SetExposedHeaders(val *[]*string) {
	if err := j.validateSetExposedHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"exposedHeaders",
		val,
	)
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference)SetInternalValue(val *SpringCloudGatewayCors) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference)SetMaxAgeSeconds(val *float64) {
	if err := j.validateSetMaxAgeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxAgeSeconds",
		val,
	)
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SpringCloudGatewayCorsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (s *jsiiProxy_SpringCloudGatewayCorsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudGatewayCorsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (s *jsiiProxy_SpringCloudGatewayCorsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpringCloudGatewayCorsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (s *jsiiProxy_SpringCloudGatewayCorsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (s *jsiiProxy_SpringCloudGatewayCorsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (s *jsiiProxy_SpringCloudGatewayCorsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (s *jsiiProxy_SpringCloudGatewayCorsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (s *jsiiProxy_SpringCloudGatewayCorsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (s *jsiiProxy_SpringCloudGatewayCorsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (s *jsiiProxy_SpringCloudGatewayCorsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SpringCloudGatewayCorsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (s *jsiiProxy_SpringCloudGatewayCorsOutputReference) ResetAllowedHeaders() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedHeaders",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudGatewayCorsOutputReference) ResetAllowedMethods() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedMethods",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudGatewayCorsOutputReference) ResetAllowedOriginPatterns() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedOriginPatterns",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudGatewayCorsOutputReference) ResetAllowedOrigins() {
	_jsii_.InvokeVoid(
		s,
		"resetAllowedOrigins",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudGatewayCorsOutputReference) ResetCredentialsAllowed() {
	_jsii_.InvokeVoid(
		s,
		"resetCredentialsAllowed",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudGatewayCorsOutputReference) ResetExposedHeaders() {
	_jsii_.InvokeVoid(
		s,
		"resetExposedHeaders",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudGatewayCorsOutputReference) ResetMaxAgeSeconds() {
	_jsii_.InvokeVoid(
		s,
		"resetMaxAgeSeconds",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SpringCloudGatewayCorsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (s *jsiiProxy_SpringCloudGatewayCorsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


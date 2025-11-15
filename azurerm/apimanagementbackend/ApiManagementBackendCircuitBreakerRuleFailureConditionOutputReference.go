// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementbackend

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/apimanagementbackend/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference interface {
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
	Count() *float64
	SetCount(val *float64)
	CountInput() *float64
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ErrorReasons() *[]*string
	SetErrorReasons(val *[]*string)
	ErrorReasonsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *ApiManagementBackendCircuitBreakerRuleFailureCondition
	SetInternalValue(val *ApiManagementBackendCircuitBreakerRuleFailureCondition)
	IntervalDuration() *string
	SetIntervalDuration(val *string)
	IntervalDurationInput() *string
	Percentage() *float64
	SetPercentage(val *float64)
	PercentageInput() *float64
	StatusCodeRange() ApiManagementBackendCircuitBreakerRuleFailureConditionStatusCodeRangeList
	StatusCodeRangeInput() interface{}
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutStatusCodeRange(value interface{})
	ResetCount()
	ResetErrorReasons()
	ResetPercentage()
	ResetStatusCodeRange()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference
type jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) CountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) ErrorReasons() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"errorReasons",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) ErrorReasonsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"errorReasonsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) InternalValue() *ApiManagementBackendCircuitBreakerRuleFailureCondition {
	var returns *ApiManagementBackendCircuitBreakerRuleFailureCondition
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) IntervalDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) IntervalDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"intervalDurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) Percentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) PercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"percentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) StatusCodeRange() ApiManagementBackendCircuitBreakerRuleFailureConditionStatusCodeRangeList {
	var returns ApiManagementBackendCircuitBreakerRuleFailureConditionStatusCodeRangeList
	_jsii_.Get(
		j,
		"statusCodeRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) StatusCodeRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statusCodeRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference {
	_init_.Initialize()

	if err := validateNewApiManagementBackendCircuitBreakerRuleFailureConditionOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagementBackend.ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference_Override(a ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagementBackend.ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference)SetCount(val *float64) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference)SetErrorReasons(val *[]*string) {
	if err := j.validateSetErrorReasonsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"errorReasons",
		val,
	)
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference)SetInternalValue(val *ApiManagementBackendCircuitBreakerRuleFailureCondition) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference)SetIntervalDuration(val *string) {
	if err := j.validateSetIntervalDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"intervalDuration",
		val,
	)
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference)SetPercentage(val *float64) {
	if err := j.validateSetPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"percentage",
		val,
	)
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) PutStatusCodeRange(value interface{}) {
	if err := a.validatePutStatusCodeRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putStatusCodeRange",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) ResetCount() {
	_jsii_.InvokeVoid(
		a,
		"resetCount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) ResetErrorReasons() {
	_jsii_.InvokeVoid(
		a,
		"resetErrorReasons",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) ResetPercentage() {
	_jsii_.InvokeVoid(
		a,
		"resetPercentage",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) ResetStatusCodeRange() {
	_jsii_.InvokeVoid(
		a,
		"resetStatusCodeRange",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


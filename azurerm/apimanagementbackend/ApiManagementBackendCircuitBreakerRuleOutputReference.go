// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package apimanagementbackend

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/apimanagementbackend/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApiManagementBackendCircuitBreakerRuleOutputReference interface {
	cdktf.ComplexObject
	AcceptRetryAfterEnabled() interface{}
	SetAcceptRetryAfterEnabled(val interface{})
	AcceptRetryAfterEnabledInput() interface{}
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
	FailureCondition() ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference
	FailureConditionInput() *ApiManagementBackendCircuitBreakerRuleFailureCondition
	// Experimental.
	Fqn() *string
	InternalValue() *ApiManagementBackendCircuitBreakerRule
	SetInternalValue(val *ApiManagementBackendCircuitBreakerRule)
	Name() *string
	SetName(val *string)
	NameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TripDuration() *string
	SetTripDuration(val *string)
	TripDurationInput() *string
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
	PutFailureCondition(value *ApiManagementBackendCircuitBreakerRuleFailureCondition)
	ResetAcceptRetryAfterEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApiManagementBackendCircuitBreakerRuleOutputReference
type jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) AcceptRetryAfterEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceptRetryAfterEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) AcceptRetryAfterEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceptRetryAfterEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) FailureCondition() ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference {
	var returns ApiManagementBackendCircuitBreakerRuleFailureConditionOutputReference
	_jsii_.Get(
		j,
		"failureCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) FailureConditionInput() *ApiManagementBackendCircuitBreakerRuleFailureCondition {
	var returns *ApiManagementBackendCircuitBreakerRuleFailureCondition
	_jsii_.Get(
		j,
		"failureConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) InternalValue() *ApiManagementBackendCircuitBreakerRule {
	var returns *ApiManagementBackendCircuitBreakerRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) TripDuration() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tripDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) TripDurationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tripDurationInput",
		&returns,
	)
	return returns
}


func NewApiManagementBackendCircuitBreakerRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApiManagementBackendCircuitBreakerRuleOutputReference {
	_init_.Initialize()

	if err := validateNewApiManagementBackendCircuitBreakerRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagementBackend.ApiManagementBackendCircuitBreakerRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApiManagementBackendCircuitBreakerRuleOutputReference_Override(a ApiManagementBackendCircuitBreakerRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.apiManagementBackend.ApiManagementBackendCircuitBreakerRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference)SetAcceptRetryAfterEnabled(val interface{}) {
	if err := j.validateSetAcceptRetryAfterEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceptRetryAfterEnabled",
		val,
	)
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference)SetInternalValue(val *ApiManagementBackendCircuitBreakerRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference)SetTripDuration(val *string) {
	if err := j.validateSetTripDurationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tripDuration",
		val,
	)
}

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) PutFailureCondition(value *ApiManagementBackendCircuitBreakerRuleFailureCondition) {
	if err := a.validatePutFailureConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putFailureCondition",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) ResetAcceptRetryAfterEnabled() {
	_jsii_.InvokeVoid(
		a,
		"resetAcceptRetryAfterEnabled",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApiManagementBackendCircuitBreakerRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


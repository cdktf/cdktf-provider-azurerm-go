// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgridpartnerconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/eventgridpartnerconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventgridPartnerConfigurationPartnerAuthorizationOutputReference interface {
	cdktf.ComplexObject
	AuthorizationExpirationTimeInUtc() *string
	SetAuthorizationExpirationTimeInUtc(val *string)
	AuthorizationExpirationTimeInUtcInput() *string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PartnerName() *string
	SetPartnerName(val *string)
	PartnerNameInput() *string
	PartnerRegistrationId() *string
	SetPartnerRegistrationId(val *string)
	PartnerRegistrationIdInput() *string
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
	ResetAuthorizationExpirationTimeInUtc()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridPartnerConfigurationPartnerAuthorizationOutputReference
type jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) AuthorizationExpirationTimeInUtc() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationExpirationTimeInUtc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) AuthorizationExpirationTimeInUtcInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authorizationExpirationTimeInUtcInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) PartnerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) PartnerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) PartnerRegistrationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerRegistrationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) PartnerRegistrationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"partnerRegistrationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventgridPartnerConfigurationPartnerAuthorizationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) EventgridPartnerConfigurationPartnerAuthorizationOutputReference {
	_init_.Initialize()

	if err := validateNewEventgridPartnerConfigurationPartnerAuthorizationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridPartnerConfiguration.EventgridPartnerConfigurationPartnerAuthorizationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewEventgridPartnerConfigurationPartnerAuthorizationOutputReference_Override(e EventgridPartnerConfigurationPartnerAuthorizationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridPartnerConfiguration.EventgridPartnerConfigurationPartnerAuthorizationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		e,
	)
}

func (j *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference)SetAuthorizationExpirationTimeInUtc(val *string) {
	if err := j.validateSetAuthorizationExpirationTimeInUtcParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authorizationExpirationTimeInUtc",
		val,
	)
}

func (j *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference)SetPartnerName(val *string) {
	if err := j.validateSetPartnerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partnerName",
		val,
	)
}

func (j *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference)SetPartnerRegistrationId(val *string) {
	if err := j.validateSetPartnerRegistrationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"partnerRegistrationId",
		val,
	)
}

func (j *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) ResetAuthorizationExpirationTimeInUtc() {
	_jsii_.InvokeVoid(
		e,
		"resetAuthorizationExpirationTimeInUtc",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridPartnerConfigurationPartnerAuthorizationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


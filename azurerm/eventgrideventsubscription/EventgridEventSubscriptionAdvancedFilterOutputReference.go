// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package eventgrideventsubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/eventgrideventsubscription/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventgridEventSubscriptionAdvancedFilterOutputReference interface {
	cdktf.ComplexObject
	BoolEquals() EventgridEventSubscriptionAdvancedFilterBoolEqualsList
	BoolEqualsInput() interface{}
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
	InternalValue() *EventgridEventSubscriptionAdvancedFilter
	SetInternalValue(val *EventgridEventSubscriptionAdvancedFilter)
	IsNotNull() EventgridEventSubscriptionAdvancedFilterIsNotNullList
	IsNotNullInput() interface{}
	IsNullOrUndefined() EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList
	IsNullOrUndefinedInput() interface{}
	NumberGreaterThan() EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList
	NumberGreaterThanInput() interface{}
	NumberGreaterThanOrEquals() EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList
	NumberGreaterThanOrEqualsInput() interface{}
	NumberIn() EventgridEventSubscriptionAdvancedFilterNumberInList
	NumberInInput() interface{}
	NumberInRange() EventgridEventSubscriptionAdvancedFilterNumberInRangeList
	NumberInRangeInput() interface{}
	NumberLessThan() EventgridEventSubscriptionAdvancedFilterNumberLessThanList
	NumberLessThanInput() interface{}
	NumberLessThanOrEquals() EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList
	NumberLessThanOrEqualsInput() interface{}
	NumberNotIn() EventgridEventSubscriptionAdvancedFilterNumberNotInList
	NumberNotInInput() interface{}
	NumberNotInRange() EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList
	NumberNotInRangeInput() interface{}
	StringBeginsWith() EventgridEventSubscriptionAdvancedFilterStringBeginsWithList
	StringBeginsWithInput() interface{}
	StringContains() EventgridEventSubscriptionAdvancedFilterStringContainsList
	StringContainsInput() interface{}
	StringEndsWith() EventgridEventSubscriptionAdvancedFilterStringEndsWithList
	StringEndsWithInput() interface{}
	StringIn() EventgridEventSubscriptionAdvancedFilterStringInList
	StringInInput() interface{}
	StringNotBeginsWith() EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList
	StringNotBeginsWithInput() interface{}
	StringNotContains() EventgridEventSubscriptionAdvancedFilterStringNotContainsList
	StringNotContainsInput() interface{}
	StringNotEndsWith() EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList
	StringNotEndsWithInput() interface{}
	StringNotIn() EventgridEventSubscriptionAdvancedFilterStringNotInList
	StringNotInInput() interface{}
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
	PutBoolEquals(value interface{})
	PutIsNotNull(value interface{})
	PutIsNullOrUndefined(value interface{})
	PutNumberGreaterThan(value interface{})
	PutNumberGreaterThanOrEquals(value interface{})
	PutNumberIn(value interface{})
	PutNumberInRange(value interface{})
	PutNumberLessThan(value interface{})
	PutNumberLessThanOrEquals(value interface{})
	PutNumberNotIn(value interface{})
	PutNumberNotInRange(value interface{})
	PutStringBeginsWith(value interface{})
	PutStringContains(value interface{})
	PutStringEndsWith(value interface{})
	PutStringIn(value interface{})
	PutStringNotBeginsWith(value interface{})
	PutStringNotContains(value interface{})
	PutStringNotEndsWith(value interface{})
	PutStringNotIn(value interface{})
	ResetBoolEquals()
	ResetIsNotNull()
	ResetIsNullOrUndefined()
	ResetNumberGreaterThan()
	ResetNumberGreaterThanOrEquals()
	ResetNumberIn()
	ResetNumberInRange()
	ResetNumberLessThan()
	ResetNumberLessThanOrEquals()
	ResetNumberNotIn()
	ResetNumberNotInRange()
	ResetStringBeginsWith()
	ResetStringContains()
	ResetStringEndsWith()
	ResetStringIn()
	ResetStringNotBeginsWith()
	ResetStringNotContains()
	ResetStringNotEndsWith()
	ResetStringNotIn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for EventgridEventSubscriptionAdvancedFilterOutputReference
type jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) BoolEquals() EventgridEventSubscriptionAdvancedFilterBoolEqualsList {
	var returns EventgridEventSubscriptionAdvancedFilterBoolEqualsList
	_jsii_.Get(
		j,
		"boolEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) BoolEqualsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"boolEqualsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) InternalValue() *EventgridEventSubscriptionAdvancedFilter {
	var returns *EventgridEventSubscriptionAdvancedFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) IsNotNull() EventgridEventSubscriptionAdvancedFilterIsNotNullList {
	var returns EventgridEventSubscriptionAdvancedFilterIsNotNullList
	_jsii_.Get(
		j,
		"isNotNull",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) IsNotNullInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isNotNullInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) IsNullOrUndefined() EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList {
	var returns EventgridEventSubscriptionAdvancedFilterIsNullOrUndefinedList
	_jsii_.Get(
		j,
		"isNullOrUndefined",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) IsNullOrUndefinedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isNullOrUndefinedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberGreaterThan() EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList {
	var returns EventgridEventSubscriptionAdvancedFilterNumberGreaterThanList
	_jsii_.Get(
		j,
		"numberGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberGreaterThanInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberGreaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberGreaterThanOrEquals() EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList {
	var returns EventgridEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList
	_jsii_.Get(
		j,
		"numberGreaterThanOrEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberGreaterThanOrEqualsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberGreaterThanOrEqualsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberIn() EventgridEventSubscriptionAdvancedFilterNumberInList {
	var returns EventgridEventSubscriptionAdvancedFilterNumberInList
	_jsii_.Get(
		j,
		"numberIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberInRange() EventgridEventSubscriptionAdvancedFilterNumberInRangeList {
	var returns EventgridEventSubscriptionAdvancedFilterNumberInRangeList
	_jsii_.Get(
		j,
		"numberInRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberInRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberInRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberLessThan() EventgridEventSubscriptionAdvancedFilterNumberLessThanList {
	var returns EventgridEventSubscriptionAdvancedFilterNumberLessThanList
	_jsii_.Get(
		j,
		"numberLessThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberLessThanInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberLessThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberLessThanOrEquals() EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList {
	var returns EventgridEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList
	_jsii_.Get(
		j,
		"numberLessThanOrEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberLessThanOrEqualsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberLessThanOrEqualsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberNotIn() EventgridEventSubscriptionAdvancedFilterNumberNotInList {
	var returns EventgridEventSubscriptionAdvancedFilterNumberNotInList
	_jsii_.Get(
		j,
		"numberNotIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberNotInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberNotInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberNotInRange() EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList {
	var returns EventgridEventSubscriptionAdvancedFilterNumberNotInRangeList
	_jsii_.Get(
		j,
		"numberNotInRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) NumberNotInRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberNotInRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringBeginsWith() EventgridEventSubscriptionAdvancedFilterStringBeginsWithList {
	var returns EventgridEventSubscriptionAdvancedFilterStringBeginsWithList
	_jsii_.Get(
		j,
		"stringBeginsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringBeginsWithInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringBeginsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringContains() EventgridEventSubscriptionAdvancedFilterStringContainsList {
	var returns EventgridEventSubscriptionAdvancedFilterStringContainsList
	_jsii_.Get(
		j,
		"stringContains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringContainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringContainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringEndsWith() EventgridEventSubscriptionAdvancedFilterStringEndsWithList {
	var returns EventgridEventSubscriptionAdvancedFilterStringEndsWithList
	_jsii_.Get(
		j,
		"stringEndsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringEndsWithInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringEndsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringIn() EventgridEventSubscriptionAdvancedFilterStringInList {
	var returns EventgridEventSubscriptionAdvancedFilterStringInList
	_jsii_.Get(
		j,
		"stringIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringNotBeginsWith() EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList {
	var returns EventgridEventSubscriptionAdvancedFilterStringNotBeginsWithList
	_jsii_.Get(
		j,
		"stringNotBeginsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringNotBeginsWithInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringNotBeginsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringNotContains() EventgridEventSubscriptionAdvancedFilterStringNotContainsList {
	var returns EventgridEventSubscriptionAdvancedFilterStringNotContainsList
	_jsii_.Get(
		j,
		"stringNotContains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringNotContainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringNotContainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringNotEndsWith() EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList {
	var returns EventgridEventSubscriptionAdvancedFilterStringNotEndsWithList
	_jsii_.Get(
		j,
		"stringNotEndsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringNotEndsWithInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringNotEndsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringNotIn() EventgridEventSubscriptionAdvancedFilterStringNotInList {
	var returns EventgridEventSubscriptionAdvancedFilterStringNotInList
	_jsii_.Get(
		j,
		"stringNotIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) StringNotInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringNotInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventgridEventSubscriptionAdvancedFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventgridEventSubscriptionAdvancedFilterOutputReference {
	_init_.Initialize()

	if err := validateNewEventgridEventSubscriptionAdvancedFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventgridEventSubscriptionAdvancedFilterOutputReference_Override(e EventgridEventSubscriptionAdvancedFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridEventSubscription.EventgridEventSubscriptionAdvancedFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference)SetInternalValue(val *EventgridEventSubscriptionAdvancedFilter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutBoolEquals(value interface{}) {
	if err := e.validatePutBoolEqualsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putBoolEquals",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutIsNotNull(value interface{}) {
	if err := e.validatePutIsNotNullParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIsNotNull",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutIsNullOrUndefined(value interface{}) {
	if err := e.validatePutIsNullOrUndefinedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIsNullOrUndefined",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutNumberGreaterThan(value interface{}) {
	if err := e.validatePutNumberGreaterThanParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNumberGreaterThan",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutNumberGreaterThanOrEquals(value interface{}) {
	if err := e.validatePutNumberGreaterThanOrEqualsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNumberGreaterThanOrEquals",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutNumberIn(value interface{}) {
	if err := e.validatePutNumberInParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNumberIn",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutNumberInRange(value interface{}) {
	if err := e.validatePutNumberInRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNumberInRange",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutNumberLessThan(value interface{}) {
	if err := e.validatePutNumberLessThanParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNumberLessThan",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutNumberLessThanOrEquals(value interface{}) {
	if err := e.validatePutNumberLessThanOrEqualsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNumberLessThanOrEquals",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutNumberNotIn(value interface{}) {
	if err := e.validatePutNumberNotInParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNumberNotIn",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutNumberNotInRange(value interface{}) {
	if err := e.validatePutNumberNotInRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNumberNotInRange",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutStringBeginsWith(value interface{}) {
	if err := e.validatePutStringBeginsWithParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStringBeginsWith",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutStringContains(value interface{}) {
	if err := e.validatePutStringContainsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStringContains",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutStringEndsWith(value interface{}) {
	if err := e.validatePutStringEndsWithParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStringEndsWith",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutStringIn(value interface{}) {
	if err := e.validatePutStringInParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStringIn",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutStringNotBeginsWith(value interface{}) {
	if err := e.validatePutStringNotBeginsWithParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStringNotBeginsWith",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutStringNotContains(value interface{}) {
	if err := e.validatePutStringNotContainsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStringNotContains",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutStringNotEndsWith(value interface{}) {
	if err := e.validatePutStringNotEndsWithParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStringNotEndsWith",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) PutStringNotIn(value interface{}) {
	if err := e.validatePutStringNotInParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStringNotIn",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetBoolEquals() {
	_jsii_.InvokeVoid(
		e,
		"resetBoolEquals",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetIsNotNull() {
	_jsii_.InvokeVoid(
		e,
		"resetIsNotNull",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetIsNullOrUndefined() {
	_jsii_.InvokeVoid(
		e,
		"resetIsNullOrUndefined",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetNumberGreaterThan() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberGreaterThan",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetNumberGreaterThanOrEquals() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberGreaterThanOrEquals",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetNumberIn() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberIn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetNumberInRange() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberInRange",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetNumberLessThan() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberLessThan",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetNumberLessThanOrEquals() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberLessThanOrEquals",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetNumberNotIn() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberNotIn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetNumberNotInRange() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberNotInRange",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetStringBeginsWith() {
	_jsii_.InvokeVoid(
		e,
		"resetStringBeginsWith",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetStringContains() {
	_jsii_.InvokeVoid(
		e,
		"resetStringContains",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetStringEndsWith() {
	_jsii_.InvokeVoid(
		e,
		"resetStringEndsWith",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetStringIn() {
	_jsii_.InvokeVoid(
		e,
		"resetStringIn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetStringNotBeginsWith() {
	_jsii_.InvokeVoid(
		e,
		"resetStringNotBeginsWith",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetStringNotContains() {
	_jsii_.InvokeVoid(
		e,
		"resetStringNotContains",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetStringNotEndsWith() {
	_jsii_.InvokeVoid(
		e,
		"resetStringNotEndsWith",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ResetStringNotIn() {
	_jsii_.InvokeVoid(
		e,
		"resetStringNotIn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EventgridEventSubscriptionAdvancedFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


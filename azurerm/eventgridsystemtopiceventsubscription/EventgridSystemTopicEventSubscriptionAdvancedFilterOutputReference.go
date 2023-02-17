package eventgridsystemtopiceventsubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/eventgridsystemtopiceventsubscription/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference interface {
	cdktf.ComplexObject
	BoolEquals() EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList
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
	InternalValue() *EventgridSystemTopicEventSubscriptionAdvancedFilter
	SetInternalValue(val *EventgridSystemTopicEventSubscriptionAdvancedFilter)
	IsNotNull() EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList
	IsNotNullInput() interface{}
	IsNullOrUndefined() EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList
	IsNullOrUndefinedInput() interface{}
	NumberGreaterThan() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList
	NumberGreaterThanInput() interface{}
	NumberGreaterThanOrEquals() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList
	NumberGreaterThanOrEqualsInput() interface{}
	NumberIn() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList
	NumberInInput() interface{}
	NumberInRange() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList
	NumberInRangeInput() interface{}
	NumberLessThan() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList
	NumberLessThanInput() interface{}
	NumberLessThanOrEquals() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList
	NumberLessThanOrEqualsInput() interface{}
	NumberNotIn() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList
	NumberNotInInput() interface{}
	NumberNotInRange() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList
	NumberNotInRangeInput() interface{}
	StringBeginsWith() EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList
	StringBeginsWithInput() interface{}
	StringContains() EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList
	StringContainsInput() interface{}
	StringEndsWith() EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList
	StringEndsWithInput() interface{}
	StringIn() EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList
	StringInInput() interface{}
	StringNotBeginsWith() EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList
	StringNotBeginsWithInput() interface{}
	StringNotContains() EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList
	StringNotContainsInput() interface{}
	StringNotEndsWith() EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList
	StringNotEndsWithInput() interface{}
	StringNotIn() EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList
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

// The jsii proxy struct for EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference
type jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) BoolEquals() EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterBoolEqualsList
	_jsii_.Get(
		j,
		"boolEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) BoolEqualsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"boolEqualsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) InternalValue() *EventgridSystemTopicEventSubscriptionAdvancedFilter {
	var returns *EventgridSystemTopicEventSubscriptionAdvancedFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) IsNotNull() EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterIsNotNullList
	_jsii_.Get(
		j,
		"isNotNull",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) IsNotNullInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isNotNullInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) IsNullOrUndefined() EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterIsNullOrUndefinedList
	_jsii_.Get(
		j,
		"isNullOrUndefined",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) IsNullOrUndefinedInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isNullOrUndefinedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberGreaterThan() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanList
	_jsii_.Get(
		j,
		"numberGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberGreaterThanInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberGreaterThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberGreaterThanOrEquals() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterNumberGreaterThanOrEqualsList
	_jsii_.Get(
		j,
		"numberGreaterThanOrEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberGreaterThanOrEqualsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberGreaterThanOrEqualsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberIn() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInList
	_jsii_.Get(
		j,
		"numberIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberInRange() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterNumberInRangeList
	_jsii_.Get(
		j,
		"numberInRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberInRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberInRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberLessThan() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanList
	_jsii_.Get(
		j,
		"numberLessThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberLessThanInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberLessThanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberLessThanOrEquals() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterNumberLessThanOrEqualsList
	_jsii_.Get(
		j,
		"numberLessThanOrEquals",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberLessThanOrEqualsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberLessThanOrEqualsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberNotIn() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInList
	_jsii_.Get(
		j,
		"numberNotIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberNotInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberNotInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberNotInRange() EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterNumberNotInRangeList
	_jsii_.Get(
		j,
		"numberNotInRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) NumberNotInRangeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"numberNotInRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringBeginsWith() EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterStringBeginsWithList
	_jsii_.Get(
		j,
		"stringBeginsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringBeginsWithInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringBeginsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringContains() EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterStringContainsList
	_jsii_.Get(
		j,
		"stringContains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringContainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringContainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringEndsWith() EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterStringEndsWithList
	_jsii_.Get(
		j,
		"stringEndsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringEndsWithInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringEndsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringIn() EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterStringInList
	_jsii_.Get(
		j,
		"stringIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringNotBeginsWith() EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotBeginsWithList
	_jsii_.Get(
		j,
		"stringNotBeginsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringNotBeginsWithInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringNotBeginsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringNotContains() EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotContainsList
	_jsii_.Get(
		j,
		"stringNotContains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringNotContainsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringNotContainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringNotEndsWith() EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotEndsWithList
	_jsii_.Get(
		j,
		"stringNotEndsWith",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringNotEndsWithInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringNotEndsWithInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringNotIn() EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList {
	var returns EventgridSystemTopicEventSubscriptionAdvancedFilterStringNotInList
	_jsii_.Get(
		j,
		"stringNotIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) StringNotInInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stringNotInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewEventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference {
	_init_.Initialize()

	if err := validateNewEventgridSystemTopicEventSubscriptionAdvancedFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewEventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference_Override(e EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.eventgridSystemTopicEventSubscription.EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference)SetInternalValue(val *EventgridSystemTopicEventSubscriptionAdvancedFilter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutBoolEquals(value interface{}) {
	if err := e.validatePutBoolEqualsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putBoolEquals",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutIsNotNull(value interface{}) {
	if err := e.validatePutIsNotNullParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIsNotNull",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutIsNullOrUndefined(value interface{}) {
	if err := e.validatePutIsNullOrUndefinedParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putIsNullOrUndefined",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutNumberGreaterThan(value interface{}) {
	if err := e.validatePutNumberGreaterThanParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNumberGreaterThan",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutNumberGreaterThanOrEquals(value interface{}) {
	if err := e.validatePutNumberGreaterThanOrEqualsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNumberGreaterThanOrEquals",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutNumberIn(value interface{}) {
	if err := e.validatePutNumberInParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNumberIn",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutNumberInRange(value interface{}) {
	if err := e.validatePutNumberInRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNumberInRange",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutNumberLessThan(value interface{}) {
	if err := e.validatePutNumberLessThanParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNumberLessThan",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutNumberLessThanOrEquals(value interface{}) {
	if err := e.validatePutNumberLessThanOrEqualsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNumberLessThanOrEquals",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutNumberNotIn(value interface{}) {
	if err := e.validatePutNumberNotInParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNumberNotIn",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutNumberNotInRange(value interface{}) {
	if err := e.validatePutNumberNotInRangeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putNumberNotInRange",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutStringBeginsWith(value interface{}) {
	if err := e.validatePutStringBeginsWithParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStringBeginsWith",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutStringContains(value interface{}) {
	if err := e.validatePutStringContainsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStringContains",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutStringEndsWith(value interface{}) {
	if err := e.validatePutStringEndsWithParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStringEndsWith",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutStringIn(value interface{}) {
	if err := e.validatePutStringInParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStringIn",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutStringNotBeginsWith(value interface{}) {
	if err := e.validatePutStringNotBeginsWithParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStringNotBeginsWith",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutStringNotContains(value interface{}) {
	if err := e.validatePutStringNotContainsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStringNotContains",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutStringNotEndsWith(value interface{}) {
	if err := e.validatePutStringNotEndsWithParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStringNotEndsWith",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) PutStringNotIn(value interface{}) {
	if err := e.validatePutStringNotInParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"putStringNotIn",
		[]interface{}{value},
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetBoolEquals() {
	_jsii_.InvokeVoid(
		e,
		"resetBoolEquals",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetIsNotNull() {
	_jsii_.InvokeVoid(
		e,
		"resetIsNotNull",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetIsNullOrUndefined() {
	_jsii_.InvokeVoid(
		e,
		"resetIsNullOrUndefined",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetNumberGreaterThan() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberGreaterThan",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetNumberGreaterThanOrEquals() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberGreaterThanOrEquals",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetNumberIn() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberIn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetNumberInRange() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberInRange",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetNumberLessThan() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberLessThan",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetNumberLessThanOrEquals() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberLessThanOrEquals",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetNumberNotIn() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberNotIn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetNumberNotInRange() {
	_jsii_.InvokeVoid(
		e,
		"resetNumberNotInRange",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetStringBeginsWith() {
	_jsii_.InvokeVoid(
		e,
		"resetStringBeginsWith",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetStringContains() {
	_jsii_.InvokeVoid(
		e,
		"resetStringContains",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetStringEndsWith() {
	_jsii_.InvokeVoid(
		e,
		"resetStringEndsWith",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetStringIn() {
	_jsii_.InvokeVoid(
		e,
		"resetStringIn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetStringNotBeginsWith() {
	_jsii_.InvokeVoid(
		e,
		"resetStringNotBeginsWith",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetStringNotContains() {
	_jsii_.InvokeVoid(
		e,
		"resetStringNotContains",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetStringNotEndsWith() {
	_jsii_.InvokeVoid(
		e,
		"resetStringNotEndsWith",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ResetStringNotIn() {
	_jsii_.InvokeVoid(
		e,
		"resetStringNotIn",
		nil, // no parameters
	)
}

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (e *jsiiProxy_EventgridSystemTopicEventSubscriptionAdvancedFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


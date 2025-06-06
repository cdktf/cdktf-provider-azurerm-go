// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pimeligibleroleassignment

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/pimeligibleroleassignment/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PimEligibleRoleAssignmentScheduleExpirationOutputReference interface {
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
	DurationDays() *float64
	SetDurationDays(val *float64)
	DurationDaysInput() *float64
	DurationHours() *float64
	SetDurationHours(val *float64)
	DurationHoursInput() *float64
	EndDateTime() *string
	SetEndDateTime(val *string)
	EndDateTimeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *PimEligibleRoleAssignmentScheduleExpiration
	SetInternalValue(val *PimEligibleRoleAssignmentScheduleExpiration)
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
	ResetDurationDays()
	ResetDurationHours()
	ResetEndDateTime()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PimEligibleRoleAssignmentScheduleExpirationOutputReference
type jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) DurationDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) DurationDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) DurationHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) DurationHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"durationHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) EndDateTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endDateTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) EndDateTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endDateTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) InternalValue() *PimEligibleRoleAssignmentScheduleExpiration {
	var returns *PimEligibleRoleAssignmentScheduleExpiration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPimEligibleRoleAssignmentScheduleExpirationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PimEligibleRoleAssignmentScheduleExpirationOutputReference {
	_init_.Initialize()

	if err := validateNewPimEligibleRoleAssignmentScheduleExpirationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.pimEligibleRoleAssignment.PimEligibleRoleAssignmentScheduleExpirationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPimEligibleRoleAssignmentScheduleExpirationOutputReference_Override(p PimEligibleRoleAssignmentScheduleExpirationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.pimEligibleRoleAssignment.PimEligibleRoleAssignmentScheduleExpirationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference)SetDurationDays(val *float64) {
	if err := j.validateSetDurationDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"durationDays",
		val,
	)
}

func (j *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference)SetDurationHours(val *float64) {
	if err := j.validateSetDurationHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"durationHours",
		val,
	)
}

func (j *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference)SetEndDateTime(val *string) {
	if err := j.validateSetEndDateTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endDateTime",
		val,
	)
}

func (j *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference)SetInternalValue(val *PimEligibleRoleAssignmentScheduleExpiration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) ResetDurationDays() {
	_jsii_.InvokeVoid(
		p,
		"resetDurationDays",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) ResetDurationHours() {
	_jsii_.InvokeVoid(
		p,
		"resetDurationHours",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) ResetEndDateTime() {
	_jsii_.InvokeVoid(
		p,
		"resetEndDateTime",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		p,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PimEligibleRoleAssignmentScheduleExpirationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


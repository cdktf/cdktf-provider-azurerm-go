// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionappslot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/functionappslot/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference interface {
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
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	XAzureFdid() *[]*string
	SetXAzureFdid(val *[]*string)
	XAzureFdidInput() *[]*string
	XFdHealthProbe() *[]*string
	SetXFdHealthProbe(val *[]*string)
	XFdHealthProbeInput() *[]*string
	XForwardedFor() *[]*string
	SetXForwardedFor(val *[]*string)
	XForwardedForInput() *[]*string
	XForwardedHost() *[]*string
	SetXForwardedHost(val *[]*string)
	XForwardedHostInput() *[]*string
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
	ResetXAzureFdid()
	ResetXFdHealthProbe()
	ResetXForwardedFor()
	ResetXForwardedHost()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference
type jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) XAzureFdid() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xAzureFdid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) XAzureFdidInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xAzureFdidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) XFdHealthProbe() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xFdHealthProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) XFdHealthProbeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xFdHealthProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) XForwardedFor() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xForwardedFor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) XForwardedForInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xForwardedForInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) XForwardedHost() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xForwardedHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) XForwardedHostInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xForwardedHostInput",
		&returns,
	)
	return returns
}


func NewFunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference {
	_init_.Initialize()

	if err := validateNewFunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.functionAppSlot.FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewFunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference_Override(f FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.functionAppSlot.FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		f,
	)
}

func (j *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference)SetXAzureFdid(val *[]*string) {
	if err := j.validateSetXAzureFdidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xAzureFdid",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference)SetXFdHealthProbe(val *[]*string) {
	if err := j.validateSetXFdHealthProbeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xFdHealthProbe",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference)SetXForwardedFor(val *[]*string) {
	if err := j.validateSetXForwardedForParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xForwardedFor",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference)SetXForwardedHost(val *[]*string) {
	if err := j.validateSetXForwardedHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xForwardedHost",
		val,
	)
}

func (f *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (f *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (f *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (f *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (f *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (f *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (f *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (f *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (f *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (f *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) ResetXAzureFdid() {
	_jsii_.InvokeVoid(
		f,
		"resetXAzureFdid",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) ResetXFdHealthProbe() {
	_jsii_.InvokeVoid(
		f,
		"resetXFdHealthProbe",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) ResetXForwardedFor() {
	_jsii_.InvokeVoid(
		f,
		"resetXForwardedFor",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) ResetXForwardedHost() {
	_jsii_.InvokeVoid(
		f,
		"resetXForwardedHost",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (f *jsiiProxy_FunctionAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


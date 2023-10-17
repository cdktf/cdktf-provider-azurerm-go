// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowswebappslot

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/windowswebappslot/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference interface {
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

// The jsii proxy struct for WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference
type jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) XAzureFdid() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xAzureFdid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) XAzureFdidInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xAzureFdidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) XFdHealthProbe() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xFdHealthProbe",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) XFdHealthProbeInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xFdHealthProbeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) XForwardedFor() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xForwardedFor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) XForwardedForInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xForwardedForInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) XForwardedHost() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xForwardedHost",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) XForwardedHostInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"xForwardedHostInput",
		&returns,
	)
	return returns
}


func NewWindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference {
	_init_.Initialize()

	if err := validateNewWindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebAppSlot.WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference_Override(w WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsWebAppSlot.WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference)SetXAzureFdid(val *[]*string) {
	if err := j.validateSetXAzureFdidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xAzureFdid",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference)SetXFdHealthProbe(val *[]*string) {
	if err := j.validateSetXFdHealthProbeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xFdHealthProbe",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference)SetXForwardedFor(val *[]*string) {
	if err := j.validateSetXForwardedForParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xForwardedFor",
		val,
	)
}

func (j *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference)SetXForwardedHost(val *[]*string) {
	if err := j.validateSetXForwardedHostParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"xForwardedHost",
		val,
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) ResetXAzureFdid() {
	_jsii_.InvokeVoid(
		w,
		"resetXAzureFdid",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) ResetXFdHealthProbe() {
	_jsii_.InvokeVoid(
		w,
		"resetXFdHealthProbe",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) ResetXForwardedFor() {
	_jsii_.InvokeVoid(
		w,
		"resetXForwardedFor",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) ResetXForwardedHost() {
	_jsii_.InvokeVoid(
		w,
		"resetXForwardedHost",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsWebAppSlotSiteConfigScmIpRestrictionHeadersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


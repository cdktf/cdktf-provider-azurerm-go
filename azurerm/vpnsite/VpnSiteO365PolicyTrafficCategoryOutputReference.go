// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnsite

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/vpnsite/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpnSiteO365PolicyTrafficCategoryOutputReference interface {
	cdktf.ComplexObject
	AllowEndpointEnabled() interface{}
	SetAllowEndpointEnabled(val interface{})
	AllowEndpointEnabledInput() interface{}
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
	DefaultEndpointEnabled() interface{}
	SetDefaultEndpointEnabled(val interface{})
	DefaultEndpointEnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *VpnSiteO365PolicyTrafficCategory
	SetInternalValue(val *VpnSiteO365PolicyTrafficCategory)
	OptimizeEndpointEnabled() interface{}
	SetOptimizeEndpointEnabled(val interface{})
	OptimizeEndpointEnabledInput() interface{}
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
	ResetAllowEndpointEnabled()
	ResetDefaultEndpointEnabled()
	ResetOptimizeEndpointEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpnSiteO365PolicyTrafficCategoryOutputReference
type jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) AllowEndpointEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowEndpointEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) AllowEndpointEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowEndpointEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) DefaultEndpointEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultEndpointEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) DefaultEndpointEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"defaultEndpointEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) InternalValue() *VpnSiteO365PolicyTrafficCategory {
	var returns *VpnSiteO365PolicyTrafficCategory
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) OptimizeEndpointEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optimizeEndpointEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) OptimizeEndpointEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"optimizeEndpointEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVpnSiteO365PolicyTrafficCategoryOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VpnSiteO365PolicyTrafficCategoryOutputReference {
	_init_.Initialize()

	if err := validateNewVpnSiteO365PolicyTrafficCategoryOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnSite.VpnSiteO365PolicyTrafficCategoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVpnSiteO365PolicyTrafficCategoryOutputReference_Override(v VpnSiteO365PolicyTrafficCategoryOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnSite.VpnSiteO365PolicyTrafficCategoryOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference)SetAllowEndpointEnabled(val interface{}) {
	if err := j.validateSetAllowEndpointEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowEndpointEnabled",
		val,
	)
}

func (j *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference)SetDefaultEndpointEnabled(val interface{}) {
	if err := j.validateSetDefaultEndpointEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultEndpointEnabled",
		val,
	)
}

func (j *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference)SetInternalValue(val *VpnSiteO365PolicyTrafficCategory) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference)SetOptimizeEndpointEnabled(val interface{}) {
	if err := j.validateSetOptimizeEndpointEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"optimizeEndpointEnabled",
		val,
	)
}

func (j *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) ResetAllowEndpointEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetAllowEndpointEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) ResetDefaultEndpointEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetDefaultEndpointEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) ResetOptimizeEndpointEnabled() {
	_jsii_.InvokeVoid(
		v,
		"resetOptimizeEndpointEnabled",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnSiteO365PolicyTrafficCategoryOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


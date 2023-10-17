// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package routemap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/routemap/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type RouteMapRuleMatchCriterionOutputReference interface {
	cdktf.ComplexObject
	AsPath() *[]*string
	SetAsPath(val *[]*string)
	AsPathInput() *[]*string
	Community() *[]*string
	SetCommunity(val *[]*string)
	CommunityInput() *[]*string
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
	MatchCondition() *string
	SetMatchCondition(val *string)
	MatchConditionInput() *string
	RoutePrefix() *[]*string
	SetRoutePrefix(val *[]*string)
	RoutePrefixInput() *[]*string
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
	ResetAsPath()
	ResetCommunity()
	ResetRoutePrefix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for RouteMapRuleMatchCriterionOutputReference
type jsiiProxy_RouteMapRuleMatchCriterionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) AsPath() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"asPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) AsPathInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"asPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) Community() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"community",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) CommunityInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"communityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) MatchCondition() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) MatchConditionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"matchConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) RoutePrefix() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"routePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) RoutePrefixInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"routePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewRouteMapRuleMatchCriterionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) RouteMapRuleMatchCriterionOutputReference {
	_init_.Initialize()

	if err := validateNewRouteMapRuleMatchCriterionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_RouteMapRuleMatchCriterionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.routeMap.RouteMapRuleMatchCriterionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewRouteMapRuleMatchCriterionOutputReference_Override(r RouteMapRuleMatchCriterionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.routeMap.RouteMapRuleMatchCriterionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		r,
	)
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionOutputReference)SetAsPath(val *[]*string) {
	if err := j.validateSetAsPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"asPath",
		val,
	)
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionOutputReference)SetCommunity(val *[]*string) {
	if err := j.validateSetCommunityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"community",
		val,
	)
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionOutputReference)SetMatchCondition(val *string) {
	if err := j.validateSetMatchConditionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"matchCondition",
		val,
	)
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionOutputReference)SetRoutePrefix(val *[]*string) {
	if err := j.validateSetRoutePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"routePrefix",
		val,
	)
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_RouteMapRuleMatchCriterionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) ResetAsPath() {
	_jsii_.InvokeVoid(
		r,
		"resetAsPath",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) ResetCommunity() {
	_jsii_.InvokeVoid(
		r,
		"resetCommunity",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) ResetRoutePrefix() {
	_jsii_.InvokeVoid(
		r,
		"resetRoutePrefix",
		nil, // no parameters
	)
}

func (r *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_RouteMapRuleMatchCriterionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package hpccache

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/hpccache/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HpcCacheDefaultAccessPolicyAccessRuleOutputReference interface {
	cdktf.ComplexObject
	Access() *string
	SetAccess(val *string)
	AccessInput() *string
	AnonymousGid() *float64
	SetAnonymousGid(val *float64)
	AnonymousGidInput() *float64
	AnonymousUid() *float64
	SetAnonymousUid(val *float64)
	AnonymousUidInput() *float64
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
	Filter() *string
	SetFilter(val *string)
	FilterInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	RootSquashEnabled() interface{}
	SetRootSquashEnabled(val interface{})
	RootSquashEnabledInput() interface{}
	Scope() *string
	SetScope(val *string)
	ScopeInput() *string
	SubmountAccessEnabled() interface{}
	SetSubmountAccessEnabled(val interface{})
	SubmountAccessEnabledInput() interface{}
	SuidEnabled() interface{}
	SetSuidEnabled(val interface{})
	SuidEnabledInput() interface{}
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
	ResetAnonymousGid()
	ResetAnonymousUid()
	ResetFilter()
	ResetRootSquashEnabled()
	ResetSubmountAccessEnabled()
	ResetSuidEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HpcCacheDefaultAccessPolicyAccessRuleOutputReference
type jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) Access() *string {
	var returns *string
	_jsii_.Get(
		j,
		"access",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) AccessInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) AnonymousGid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"anonymousGid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) AnonymousGidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"anonymousGidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) AnonymousUid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"anonymousUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) AnonymousUidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"anonymousUidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) Filter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) FilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"filterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) RootSquashEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootSquashEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) RootSquashEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rootSquashEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) SubmountAccessEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"submountAccessEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) SubmountAccessEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"submountAccessEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) SuidEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suidEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) SuidEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"suidEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewHpcCacheDefaultAccessPolicyAccessRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) HpcCacheDefaultAccessPolicyAccessRuleOutputReference {
	_init_.Initialize()

	if err := validateNewHpcCacheDefaultAccessPolicyAccessRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCacheDefaultAccessPolicyAccessRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewHpcCacheDefaultAccessPolicyAccessRuleOutputReference_Override(h HpcCacheDefaultAccessPolicyAccessRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hpcCache.HpcCacheDefaultAccessPolicyAccessRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		h,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference)SetAccess(val *string) {
	if err := j.validateSetAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"access",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference)SetAnonymousGid(val *float64) {
	if err := j.validateSetAnonymousGidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anonymousGid",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference)SetAnonymousUid(val *float64) {
	if err := j.validateSetAnonymousUidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anonymousUid",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference)SetFilter(val *string) {
	if err := j.validateSetFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"filter",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference)SetRootSquashEnabled(val interface{}) {
	if err := j.validateSetRootSquashEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootSquashEnabled",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference)SetSubmountAccessEnabled(val interface{}) {
	if err := j.validateSetSubmountAccessEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"submountAccessEnabled",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference)SetSuidEnabled(val interface{}) {
	if err := j.validateSetSuidEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"suidEnabled",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) ResetAnonymousGid() {
	_jsii_.InvokeVoid(
		h,
		"resetAnonymousGid",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) ResetAnonymousUid() {
	_jsii_.InvokeVoid(
		h,
		"resetAnonymousUid",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) ResetFilter() {
	_jsii_.InvokeVoid(
		h,
		"resetFilter",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) ResetRootSquashEnabled() {
	_jsii_.InvokeVoid(
		h,
		"resetRootSquashEnabled",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) ResetSubmountAccessEnabled() {
	_jsii_.InvokeVoid(
		h,
		"resetSubmountAccessEnabled",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) ResetSuidEnabled() {
	_jsii_.InvokeVoid(
		h,
		"resetSuidEnabled",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := h.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HpcCacheDefaultAccessPolicyAccessRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


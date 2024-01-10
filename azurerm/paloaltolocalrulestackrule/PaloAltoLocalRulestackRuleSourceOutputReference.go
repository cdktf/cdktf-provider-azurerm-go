// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltolocalrulestackrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/paloaltolocalrulestackrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PaloAltoLocalRulestackRuleSourceOutputReference interface {
	cdktf.ComplexObject
	Cidrs() *[]*string
	SetCidrs(val *[]*string)
	CidrsInput() *[]*string
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
	Countries() *[]*string
	SetCountries(val *[]*string)
	CountriesInput() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Feeds() *[]*string
	SetFeeds(val *[]*string)
	FeedsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *PaloAltoLocalRulestackRuleSource
	SetInternalValue(val *PaloAltoLocalRulestackRuleSource)
	LocalRulestackPrefixListIds() *[]*string
	SetLocalRulestackPrefixListIds(val *[]*string)
	LocalRulestackPrefixListIdsInput() *[]*string
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
	ResetCidrs()
	ResetCountries()
	ResetFeeds()
	ResetLocalRulestackPrefixListIds()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PaloAltoLocalRulestackRuleSourceOutputReference
type jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) Cidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) CidrsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) Countries() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"countries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) CountriesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"countriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) Feeds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"feeds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) FeedsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"feedsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) InternalValue() *PaloAltoLocalRulestackRuleSource {
	var returns *PaloAltoLocalRulestackRuleSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) LocalRulestackPrefixListIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localRulestackPrefixListIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) LocalRulestackPrefixListIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localRulestackPrefixListIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPaloAltoLocalRulestackRuleSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PaloAltoLocalRulestackRuleSourceOutputReference {
	_init_.Initialize()

	if err := validateNewPaloAltoLocalRulestackRuleSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.paloAltoLocalRulestackRule.PaloAltoLocalRulestackRuleSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPaloAltoLocalRulestackRuleSourceOutputReference_Override(p PaloAltoLocalRulestackRuleSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.paloAltoLocalRulestackRule.PaloAltoLocalRulestackRuleSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference)SetCidrs(val *[]*string) {
	if err := j.validateSetCidrsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cidrs",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference)SetCountries(val *[]*string) {
	if err := j.validateSetCountriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"countries",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference)SetFeeds(val *[]*string) {
	if err := j.validateSetFeedsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"feeds",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference)SetInternalValue(val *PaloAltoLocalRulestackRuleSource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference)SetLocalRulestackPrefixListIds(val *[]*string) {
	if err := j.validateSetLocalRulestackPrefixListIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localRulestackPrefixListIds",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) ResetCidrs() {
	_jsii_.InvokeVoid(
		p,
		"resetCidrs",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) ResetCountries() {
	_jsii_.InvokeVoid(
		p,
		"resetCountries",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) ResetFeeds() {
	_jsii_.InvokeVoid(
		p,
		"resetFeeds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) ResetLocalRulestackPrefixListIds() {
	_jsii_.InvokeVoid(
		p,
		"resetLocalRulestackPrefixListIds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PaloAltoLocalRulestackRuleSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


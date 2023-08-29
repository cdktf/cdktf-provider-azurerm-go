// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltolocalrulestackrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/paloaltolocalrulestackrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PaloAltoLocalRulestackRuleDestinationOutputReference interface {
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
	InternalValue() *PaloAltoLocalRulestackRuleDestination
	SetInternalValue(val *PaloAltoLocalRulestackRuleDestination)
	LocalRulestackFqdnListIds() *[]*string
	SetLocalRulestackFqdnListIds(val *[]*string)
	LocalRulestackFqdnListIdsInput() *[]*string
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
	ResetLocalRulestackFqdnListIds()
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

// The jsii proxy struct for PaloAltoLocalRulestackRuleDestinationOutputReference
type jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) Cidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) CidrsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) Countries() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"countries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) CountriesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"countriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) Feeds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"feeds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) FeedsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"feedsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) InternalValue() *PaloAltoLocalRulestackRuleDestination {
	var returns *PaloAltoLocalRulestackRuleDestination
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) LocalRulestackFqdnListIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localRulestackFqdnListIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) LocalRulestackFqdnListIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localRulestackFqdnListIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) LocalRulestackPrefixListIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localRulestackPrefixListIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) LocalRulestackPrefixListIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localRulestackPrefixListIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPaloAltoLocalRulestackRuleDestinationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PaloAltoLocalRulestackRuleDestinationOutputReference {
	_init_.Initialize()

	if err := validateNewPaloAltoLocalRulestackRuleDestinationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.paloAltoLocalRulestackRule.PaloAltoLocalRulestackRuleDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPaloAltoLocalRulestackRuleDestinationOutputReference_Override(p PaloAltoLocalRulestackRuleDestinationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.paloAltoLocalRulestackRule.PaloAltoLocalRulestackRuleDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference)SetCidrs(val *[]*string) {
	if err := j.validateSetCidrsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cidrs",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference)SetCountries(val *[]*string) {
	if err := j.validateSetCountriesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"countries",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference)SetFeeds(val *[]*string) {
	if err := j.validateSetFeedsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"feeds",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference)SetInternalValue(val *PaloAltoLocalRulestackRuleDestination) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference)SetLocalRulestackFqdnListIds(val *[]*string) {
	if err := j.validateSetLocalRulestackFqdnListIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localRulestackFqdnListIds",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference)SetLocalRulestackPrefixListIds(val *[]*string) {
	if err := j.validateSetLocalRulestackPrefixListIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localRulestackPrefixListIds",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) ResetCidrs() {
	_jsii_.InvokeVoid(
		p,
		"resetCidrs",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) ResetCountries() {
	_jsii_.InvokeVoid(
		p,
		"resetCountries",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) ResetFeeds() {
	_jsii_.InvokeVoid(
		p,
		"resetFeeds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) ResetLocalRulestackFqdnListIds() {
	_jsii_.InvokeVoid(
		p,
		"resetLocalRulestackFqdnListIds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) ResetLocalRulestackPrefixListIds() {
	_jsii_.InvokeVoid(
		p,
		"resetLocalRulestackPrefixListIds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PaloAltoLocalRulestackRuleDestinationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


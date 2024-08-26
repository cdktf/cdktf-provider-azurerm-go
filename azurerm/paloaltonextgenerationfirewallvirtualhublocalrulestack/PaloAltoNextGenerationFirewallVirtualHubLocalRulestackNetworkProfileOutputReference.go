// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualhublocalrulestack

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/paloaltonextgenerationfirewallvirtualhublocalrulestack/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference interface {
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
	EgressNatIpAddresses() *[]*string
	EgressNatIpAddressIds() *[]*string
	SetEgressNatIpAddressIds(val *[]*string)
	EgressNatIpAddressIdsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfile
	SetInternalValue(val *PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfile)
	IpOfTrustForUserDefinedRoutes() *string
	NetworkVirtualApplianceId() *string
	SetNetworkVirtualApplianceId(val *string)
	NetworkVirtualApplianceIdInput() *string
	PublicIpAddresses() *[]*string
	PublicIpAddressIds() *[]*string
	SetPublicIpAddressIds(val *[]*string)
	PublicIpAddressIdsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrustedAddressRanges() *[]*string
	SetTrustedAddressRanges(val *[]*string)
	TrustedAddressRangesInput() *[]*string
	TrustedSubnetId() *string
	UntrustedSubnetId() *string
	VirtualHubId() *string
	SetVirtualHubId(val *string)
	VirtualHubIdInput() *string
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
	ResetEgressNatIpAddressIds()
	ResetTrustedAddressRanges()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference
type jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) EgressNatIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"egressNatIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) EgressNatIpAddressIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"egressNatIpAddressIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) EgressNatIpAddressIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"egressNatIpAddressIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) InternalValue() *PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfile {
	var returns *PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) IpOfTrustForUserDefinedRoutes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipOfTrustForUserDefinedRoutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) NetworkVirtualApplianceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkVirtualApplianceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) NetworkVirtualApplianceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkVirtualApplianceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) PublicIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) PublicIpAddressIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicIpAddressIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) PublicIpAddressIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicIpAddressIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) TrustedAddressRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedAddressRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) TrustedAddressRangesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedAddressRangesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) TrustedSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustedSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) UntrustedSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"untrustedSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) VirtualHubId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualHubId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) VirtualHubIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualHubIdInput",
		&returns,
	)
	return returns
}


func NewPaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference {
	_init_.Initialize()

	if err := validateNewPaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.paloAltoNextGenerationFirewallVirtualHubLocalRulestack.PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference_Override(p PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.paloAltoNextGenerationFirewallVirtualHubLocalRulestack.PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference)SetEgressNatIpAddressIds(val *[]*string) {
	if err := j.validateSetEgressNatIpAddressIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"egressNatIpAddressIds",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference)SetInternalValue(val *PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference)SetNetworkVirtualApplianceId(val *string) {
	if err := j.validateSetNetworkVirtualApplianceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkVirtualApplianceId",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference)SetPublicIpAddressIds(val *[]*string) {
	if err := j.validateSetPublicIpAddressIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicIpAddressIds",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference)SetTrustedAddressRanges(val *[]*string) {
	if err := j.validateSetTrustedAddressRangesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustedAddressRanges",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference)SetVirtualHubId(val *string) {
	if err := j.validateSetVirtualHubIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualHubId",
		val,
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) ResetEgressNatIpAddressIds() {
	_jsii_.InvokeVoid(
		p,
		"resetEgressNatIpAddressIds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) ResetTrustedAddressRanges() {
	_jsii_.InvokeVoid(
		p,
		"resetTrustedAddressRanges",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubLocalRulestackNetworkProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


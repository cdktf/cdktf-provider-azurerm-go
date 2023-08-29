// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualhubpanorama

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/paloaltonextgenerationfirewallvirtualhubpanorama/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference interface {
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
	InternalValue() *PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfile
	SetInternalValue(val *PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfile)
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference
type jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) EgressNatIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"egressNatIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) EgressNatIpAddressIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"egressNatIpAddressIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) EgressNatIpAddressIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"egressNatIpAddressIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) InternalValue() *PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfile {
	var returns *PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) IpOfTrustForUserDefinedRoutes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipOfTrustForUserDefinedRoutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) NetworkVirtualApplianceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkVirtualApplianceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) NetworkVirtualApplianceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkVirtualApplianceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) PublicIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) PublicIpAddressIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicIpAddressIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) PublicIpAddressIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicIpAddressIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) TrustedSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustedSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) UntrustedSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"untrustedSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) VirtualHubId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualHubId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) VirtualHubIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualHubIdInput",
		&returns,
	)
	return returns
}


func NewPaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference {
	_init_.Initialize()

	if err := validateNewPaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.paloAltoNextGenerationFirewallVirtualHubPanorama.PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference_Override(p PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.paloAltoNextGenerationFirewallVirtualHubPanorama.PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference)SetEgressNatIpAddressIds(val *[]*string) {
	if err := j.validateSetEgressNatIpAddressIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"egressNatIpAddressIds",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference)SetInternalValue(val *PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference)SetNetworkVirtualApplianceId(val *string) {
	if err := j.validateSetNetworkVirtualApplianceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkVirtualApplianceId",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference)SetPublicIpAddressIds(val *[]*string) {
	if err := j.validateSetPublicIpAddressIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicIpAddressIds",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference)SetVirtualHubId(val *string) {
	if err := j.validateSetVirtualHubIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualHubId",
		val,
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) ResetEgressNatIpAddressIds() {
	_jsii_.InvokeVoid(
		p,
		"resetEgressNatIpAddressIds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualHubPanoramaNetworkProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package paloaltonextgenerationfirewallvirtualnetworkpanorama

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/paloaltonextgenerationfirewallvirtualnetworkpanorama/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference interface {
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
	InternalValue() *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfiguration
	SetInternalValue(val *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfiguration)
	IpOfTrustForUserDefinedRoutes() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrustedSubnetId() *string
	SetTrustedSubnetId(val *string)
	TrustedSubnetIdInput() *string
	UntrustedSubnetId() *string
	SetUntrustedSubnetId(val *string)
	UntrustedSubnetIdInput() *string
	VirtualNetworkId() *string
	SetVirtualNetworkId(val *string)
	VirtualNetworkIdInput() *string
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
	ResetTrustedSubnetId()
	ResetUntrustedSubnetId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference
type jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) InternalValue() *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfiguration {
	var returns *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) IpOfTrustForUserDefinedRoutes() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipOfTrustForUserDefinedRoutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) TrustedSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustedSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) TrustedSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"trustedSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) UntrustedSubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"untrustedSubnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) UntrustedSubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"untrustedSubnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) VirtualNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) VirtualNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkIdInput",
		&returns,
	)
	return returns
}


func NewPaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewPaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkPanorama.PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference_Override(p PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkPanorama.PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference)SetInternalValue(val *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference)SetTrustedSubnetId(val *string) {
	if err := j.validateSetTrustedSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustedSubnetId",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference)SetUntrustedSubnetId(val *string) {
	if err := j.validateSetUntrustedSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"untrustedSubnetId",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference)SetVirtualNetworkId(val *string) {
	if err := j.validateSetVirtualNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualNetworkId",
		val,
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) ResetTrustedSubnetId() {
	_jsii_.InvokeVoid(
		p,
		"resetTrustedSubnetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) ResetUntrustedSubnetId() {
	_jsii_.InvokeVoid(
		p,
		"resetUntrustedSubnetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


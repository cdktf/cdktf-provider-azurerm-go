package paloaltonextgenerationfirewallvirtualnetworkpanorama

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/paloaltonextgenerationfirewallvirtualnetworkpanorama/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference interface {
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
	InternalValue() *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfile
	SetInternalValue(val *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfile)
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
	VnetConfiguration() PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference
	VnetConfigurationInput() *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfiguration
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
	PutVnetConfiguration(value *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfiguration)
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

// The jsii proxy struct for PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference
type jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) EgressNatIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"egressNatIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) EgressNatIpAddressIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"egressNatIpAddressIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) EgressNatIpAddressIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"egressNatIpAddressIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) InternalValue() *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfile {
	var returns *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) PublicIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) PublicIpAddressIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicIpAddressIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) PublicIpAddressIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicIpAddressIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) VnetConfiguration() PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference {
	var returns PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfigurationOutputReference
	_jsii_.Get(
		j,
		"vnetConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) VnetConfigurationInput() *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfiguration {
	var returns *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfiguration
	_jsii_.Get(
		j,
		"vnetConfigurationInput",
		&returns,
	)
	return returns
}


func NewPaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference {
	_init_.Initialize()

	if err := validateNewPaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkPanorama.PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference_Override(p PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkPanorama.PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference)SetEgressNatIpAddressIds(val *[]*string) {
	if err := j.validateSetEgressNatIpAddressIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"egressNatIpAddressIds",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference)SetInternalValue(val *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference)SetPublicIpAddressIds(val *[]*string) {
	if err := j.validateSetPublicIpAddressIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicIpAddressIds",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) PutVnetConfiguration(value *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileVnetConfiguration) {
	if err := p.validatePutVnetConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putVnetConfiguration",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) ResetEgressNatIpAddressIds() {
	_jsii_.InvokeVoid(
		p,
		"resetEgressNatIpAddressIds",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaNetworkProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


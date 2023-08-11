package paloaltonextgenerationfirewallvirtualnetworkpanorama

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/paloaltonextgenerationfirewallvirtualnetworkpanorama/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference interface {
	cdktf.ComplexObject
	BackendConfig() PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatBackendConfigOutputReference
	BackendConfigInput() *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatBackendConfig
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
	FrontendConfig() PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatFrontendConfigOutputReference
	FrontendConfigInput() *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatFrontendConfig
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Protocol() *string
	SetProtocol(val *string)
	ProtocolInput() *string
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
	PutBackendConfig(value *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatBackendConfig)
	PutFrontendConfig(value *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatFrontendConfig)
	ResetBackendConfig()
	ResetFrontendConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference
type jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) BackendConfig() PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatBackendConfigOutputReference {
	var returns PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatBackendConfigOutputReference
	_jsii_.Get(
		j,
		"backendConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) BackendConfigInput() *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatBackendConfig {
	var returns *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatBackendConfig
	_jsii_.Get(
		j,
		"backendConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) FrontendConfig() PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatFrontendConfigOutputReference {
	var returns PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatFrontendConfigOutputReference
	_jsii_.Get(
		j,
		"frontendConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) FrontendConfigInput() *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatFrontendConfig {
	var returns *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatFrontendConfig
	_jsii_.Get(
		j,
		"frontendConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) Protocol() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocol",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) ProtocolInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"protocolInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewPaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference {
	_init_.Initialize()

	if err := validateNewPaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkPanorama.PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewPaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference_Override(p PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.paloAltoNextGenerationFirewallVirtualNetworkPanorama.PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		p,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference)SetProtocol(val *string) {
	if err := j.validateSetProtocolParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"protocol",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) PutBackendConfig(value *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatBackendConfig) {
	if err := p.validatePutBackendConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putBackendConfig",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) PutFrontendConfig(value *PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatFrontendConfig) {
	if err := p.validatePutFrontendConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putFrontendConfig",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) ResetBackendConfig() {
	_jsii_.InvokeVoid(
		p,
		"resetBackendConfig",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) ResetFrontendConfig() {
	_jsii_.InvokeVoid(
		p,
		"resetFrontendConfig",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (p *jsiiProxy_PaloAltoNextGenerationFirewallVirtualNetworkPanoramaDestinationNatOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


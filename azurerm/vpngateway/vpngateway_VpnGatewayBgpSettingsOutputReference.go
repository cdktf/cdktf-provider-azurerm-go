package vpngateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v5/vpngateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpnGatewayBgpSettingsOutputReference interface {
	cdktf.ComplexObject
	Asn() *float64
	SetAsn(val *float64)
	AsnInput() *float64
	BgpPeeringAddress() *string
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
	Instance0BgpPeeringAddress() VpnGatewayBgpSettingsInstance0BgpPeeringAddressOutputReference
	Instance0BgpPeeringAddressInput() *VpnGatewayBgpSettingsInstance0BgpPeeringAddress
	Instance1BgpPeeringAddress() VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference
	Instance1BgpPeeringAddressInput() *VpnGatewayBgpSettingsInstance1BgpPeeringAddress
	InternalValue() *VpnGatewayBgpSettings
	SetInternalValue(val *VpnGatewayBgpSettings)
	PeerWeight() *float64
	SetPeerWeight(val *float64)
	PeerWeightInput() *float64
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
	PutInstance0BgpPeeringAddress(value *VpnGatewayBgpSettingsInstance0BgpPeeringAddress)
	PutInstance1BgpPeeringAddress(value *VpnGatewayBgpSettingsInstance1BgpPeeringAddress)
	ResetInstance0BgpPeeringAddress()
	ResetInstance1BgpPeeringAddress()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpnGatewayBgpSettingsOutputReference
type jsiiProxy_VpnGatewayBgpSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpnGatewayBgpSettingsOutputReference) Asn() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"asn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayBgpSettingsOutputReference) AsnInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"asnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayBgpSettingsOutputReference) BgpPeeringAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bgpPeeringAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayBgpSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayBgpSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayBgpSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayBgpSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayBgpSettingsOutputReference) Instance0BgpPeeringAddress() VpnGatewayBgpSettingsInstance0BgpPeeringAddressOutputReference {
	var returns VpnGatewayBgpSettingsInstance0BgpPeeringAddressOutputReference
	_jsii_.Get(
		j,
		"instance0BgpPeeringAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayBgpSettingsOutputReference) Instance0BgpPeeringAddressInput() *VpnGatewayBgpSettingsInstance0BgpPeeringAddress {
	var returns *VpnGatewayBgpSettingsInstance0BgpPeeringAddress
	_jsii_.Get(
		j,
		"instance0BgpPeeringAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayBgpSettingsOutputReference) Instance1BgpPeeringAddress() VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference {
	var returns VpnGatewayBgpSettingsInstance1BgpPeeringAddressOutputReference
	_jsii_.Get(
		j,
		"instance1BgpPeeringAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayBgpSettingsOutputReference) Instance1BgpPeeringAddressInput() *VpnGatewayBgpSettingsInstance1BgpPeeringAddress {
	var returns *VpnGatewayBgpSettingsInstance1BgpPeeringAddress
	_jsii_.Get(
		j,
		"instance1BgpPeeringAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayBgpSettingsOutputReference) InternalValue() *VpnGatewayBgpSettings {
	var returns *VpnGatewayBgpSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayBgpSettingsOutputReference) PeerWeight() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"peerWeight",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayBgpSettingsOutputReference) PeerWeightInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"peerWeightInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayBgpSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayBgpSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVpnGatewayBgpSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VpnGatewayBgpSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewVpnGatewayBgpSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpnGatewayBgpSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGateway.VpnGatewayBgpSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVpnGatewayBgpSettingsOutputReference_Override(v VpnGatewayBgpSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGateway.VpnGatewayBgpSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VpnGatewayBgpSettingsOutputReference)SetAsn(val *float64) {
	if err := j.validateSetAsnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"asn",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayBgpSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayBgpSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayBgpSettingsOutputReference)SetInternalValue(val *VpnGatewayBgpSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayBgpSettingsOutputReference)SetPeerWeight(val *float64) {
	if err := j.validateSetPeerWeightParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peerWeight",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayBgpSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayBgpSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VpnGatewayBgpSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayBgpSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VpnGatewayBgpSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpnGatewayBgpSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VpnGatewayBgpSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VpnGatewayBgpSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VpnGatewayBgpSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VpnGatewayBgpSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VpnGatewayBgpSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VpnGatewayBgpSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VpnGatewayBgpSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayBgpSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpnGatewayBgpSettingsOutputReference) PutInstance0BgpPeeringAddress(value *VpnGatewayBgpSettingsInstance0BgpPeeringAddress) {
	if err := v.validatePutInstance0BgpPeeringAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putInstance0BgpPeeringAddress",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnGatewayBgpSettingsOutputReference) PutInstance1BgpPeeringAddress(value *VpnGatewayBgpSettingsInstance1BgpPeeringAddress) {
	if err := v.validatePutInstance1BgpPeeringAddressParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putInstance1BgpPeeringAddress",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VpnGatewayBgpSettingsOutputReference) ResetInstance0BgpPeeringAddress() {
	_jsii_.InvokeVoid(
		v,
		"resetInstance0BgpPeeringAddress",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayBgpSettingsOutputReference) ResetInstance1BgpPeeringAddress() {
	_jsii_.InvokeVoid(
		v,
		"resetInstance1BgpPeeringAddress",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VpnGatewayBgpSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VpnGatewayBgpSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


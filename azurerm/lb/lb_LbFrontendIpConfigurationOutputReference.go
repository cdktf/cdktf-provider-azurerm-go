package lb

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/lb/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LbFrontendIpConfigurationOutputReference interface {
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
	GatewayLoadBalancerFrontendIpConfigurationId() *string
	SetGatewayLoadBalancerFrontendIpConfigurationId(val *string)
	GatewayLoadBalancerFrontendIpConfigurationIdInput() *string
	Id() *string
	InboundNatRules() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LoadBalancerRules() *[]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	OutboundRules() *[]*string
	PrivateIpAddress() *string
	SetPrivateIpAddress(val *string)
	PrivateIpAddressAllocation() *string
	SetPrivateIpAddressAllocation(val *string)
	PrivateIpAddressAllocationInput() *string
	PrivateIpAddressInput() *string
	PrivateIpAddressVersion() *string
	SetPrivateIpAddressVersion(val *string)
	PrivateIpAddressVersionInput() *string
	PublicIpAddressId() *string
	SetPublicIpAddressId(val *string)
	PublicIpAddressIdInput() *string
	PublicIpPrefixId() *string
	SetPublicIpPrefixId(val *string)
	PublicIpPrefixIdInput() *string
	SubnetId() *string
	SetSubnetId(val *string)
	SubnetIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Zones() *[]*string
	SetZones(val *[]*string)
	ZonesInput() *[]*string
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
	ResetGatewayLoadBalancerFrontendIpConfigurationId()
	ResetPrivateIpAddress()
	ResetPrivateIpAddressAllocation()
	ResetPrivateIpAddressVersion()
	ResetPublicIpAddressId()
	ResetPublicIpPrefixId()
	ResetSubnetId()
	ResetZones()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LbFrontendIpConfigurationOutputReference
type jsiiProxy_LbFrontendIpConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) GatewayLoadBalancerFrontendIpConfigurationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayLoadBalancerFrontendIpConfigurationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) GatewayLoadBalancerFrontendIpConfigurationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gatewayLoadBalancerFrontendIpConfigurationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) InboundNatRules() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inboundNatRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) LoadBalancerRules() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) OutboundRules() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"outboundRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) PrivateIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) PrivateIpAddressAllocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddressAllocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) PrivateIpAddressAllocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddressAllocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) PrivateIpAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) PrivateIpAddressVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddressVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) PrivateIpAddressVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddressVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) PublicIpAddressId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpAddressId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) PublicIpAddressIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpAddressIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) PublicIpPrefixId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpPrefixId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) PublicIpPrefixIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpPrefixIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) Zones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference) ZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zonesInput",
		&returns,
	)
	return returns
}


func NewLbFrontendIpConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LbFrontendIpConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewLbFrontendIpConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LbFrontendIpConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.lb.LbFrontendIpConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLbFrontendIpConfigurationOutputReference_Override(l LbFrontendIpConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.lb.LbFrontendIpConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference)SetGatewayLoadBalancerFrontendIpConfigurationId(val *string) {
	if err := j.validateSetGatewayLoadBalancerFrontendIpConfigurationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gatewayLoadBalancerFrontendIpConfigurationId",
		val,
	)
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference)SetPrivateIpAddress(val *string) {
	if err := j.validateSetPrivateIpAddressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpAddress",
		val,
	)
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference)SetPrivateIpAddressAllocation(val *string) {
	if err := j.validateSetPrivateIpAddressAllocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpAddressAllocation",
		val,
	)
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference)SetPrivateIpAddressVersion(val *string) {
	if err := j.validateSetPrivateIpAddressVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateIpAddressVersion",
		val,
	)
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference)SetPublicIpAddressId(val *string) {
	if err := j.validateSetPublicIpAddressIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicIpAddressId",
		val,
	)
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference)SetPublicIpPrefixId(val *string) {
	if err := j.validateSetPublicIpPrefixIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicIpPrefixId",
		val,
	)
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LbFrontendIpConfigurationOutputReference)SetZones(val *[]*string) {
	if err := j.validateSetZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zones",
		val,
	)
}

func (l *jsiiProxy_LbFrontendIpConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbFrontendIpConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbFrontendIpConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbFrontendIpConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbFrontendIpConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbFrontendIpConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbFrontendIpConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbFrontendIpConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbFrontendIpConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbFrontendIpConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbFrontendIpConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbFrontendIpConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbFrontendIpConfigurationOutputReference) ResetGatewayLoadBalancerFrontendIpConfigurationId() {
	_jsii_.InvokeVoid(
		l,
		"resetGatewayLoadBalancerFrontendIpConfigurationId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbFrontendIpConfigurationOutputReference) ResetPrivateIpAddress() {
	_jsii_.InvokeVoid(
		l,
		"resetPrivateIpAddress",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbFrontendIpConfigurationOutputReference) ResetPrivateIpAddressAllocation() {
	_jsii_.InvokeVoid(
		l,
		"resetPrivateIpAddressAllocation",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbFrontendIpConfigurationOutputReference) ResetPrivateIpAddressVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetPrivateIpAddressVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbFrontendIpConfigurationOutputReference) ResetPublicIpAddressId() {
	_jsii_.InvokeVoid(
		l,
		"resetPublicIpAddressId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbFrontendIpConfigurationOutputReference) ResetPublicIpPrefixId() {
	_jsii_.InvokeVoid(
		l,
		"resetPublicIpPrefixId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbFrontendIpConfigurationOutputReference) ResetSubnetId() {
	_jsii_.InvokeVoid(
		l,
		"resetSubnetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbFrontendIpConfigurationOutputReference) ResetZones() {
	_jsii_.InvokeVoid(
		l,
		"resetZones",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LbFrontendIpConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LbFrontendIpConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


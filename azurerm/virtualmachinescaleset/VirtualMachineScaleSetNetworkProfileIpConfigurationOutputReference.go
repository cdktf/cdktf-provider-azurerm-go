package virtualmachinescaleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/virtualmachinescaleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference interface {
	cdktf.ComplexObject
	ApplicationGatewayBackendAddressPoolIds() *[]*string
	SetApplicationGatewayBackendAddressPoolIds(val *[]*string)
	ApplicationGatewayBackendAddressPoolIdsInput() *[]*string
	ApplicationSecurityGroupIds() *[]*string
	SetApplicationSecurityGroupIds(val *[]*string)
	ApplicationSecurityGroupIdsInput() *[]*string
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	LoadBalancerBackendAddressPoolIds() *[]*string
	SetLoadBalancerBackendAddressPoolIds(val *[]*string)
	LoadBalancerBackendAddressPoolIdsInput() *[]*string
	LoadBalancerInboundNatRulesIds() *[]*string
	SetLoadBalancerInboundNatRulesIds(val *[]*string)
	LoadBalancerInboundNatRulesIdsInput() *[]*string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Primary() interface{}
	SetPrimary(val interface{})
	PrimaryInput() interface{}
	PublicIpAddressConfiguration() VirtualMachineScaleSetNetworkProfileIpConfigurationPublicIpAddressConfigurationOutputReference
	PublicIpAddressConfigurationInput() *VirtualMachineScaleSetNetworkProfileIpConfigurationPublicIpAddressConfiguration
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
	PutPublicIpAddressConfiguration(value *VirtualMachineScaleSetNetworkProfileIpConfigurationPublicIpAddressConfiguration)
	ResetApplicationGatewayBackendAddressPoolIds()
	ResetApplicationSecurityGroupIds()
	ResetLoadBalancerBackendAddressPoolIds()
	ResetLoadBalancerInboundNatRulesIds()
	ResetPublicIpAddressConfiguration()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference
type jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) ApplicationGatewayBackendAddressPoolIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"applicationGatewayBackendAddressPoolIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) ApplicationGatewayBackendAddressPoolIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"applicationGatewayBackendAddressPoolIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) ApplicationSecurityGroupIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"applicationSecurityGroupIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) ApplicationSecurityGroupIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"applicationSecurityGroupIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) LoadBalancerBackendAddressPoolIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerBackendAddressPoolIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) LoadBalancerBackendAddressPoolIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerBackendAddressPoolIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) LoadBalancerInboundNatRulesIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerInboundNatRulesIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) LoadBalancerInboundNatRulesIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"loadBalancerInboundNatRulesIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) Primary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) PrimaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) PublicIpAddressConfiguration() VirtualMachineScaleSetNetworkProfileIpConfigurationPublicIpAddressConfigurationOutputReference {
	var returns VirtualMachineScaleSetNetworkProfileIpConfigurationPublicIpAddressConfigurationOutputReference
	_jsii_.Get(
		j,
		"publicIpAddressConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) PublicIpAddressConfigurationInput() *VirtualMachineScaleSetNetworkProfileIpConfigurationPublicIpAddressConfiguration {
	var returns *VirtualMachineScaleSetNetworkProfileIpConfigurationPublicIpAddressConfiguration
	_jsii_.Get(
		j,
		"publicIpAddressConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) SubnetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewVirtualMachineScaleSetNetworkProfileIpConfigurationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualMachineScaleSet.VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference_Override(v VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualMachineScaleSet.VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference)SetApplicationGatewayBackendAddressPoolIds(val *[]*string) {
	if err := j.validateSetApplicationGatewayBackendAddressPoolIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationGatewayBackendAddressPoolIds",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference)SetApplicationSecurityGroupIds(val *[]*string) {
	if err := j.validateSetApplicationSecurityGroupIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationSecurityGroupIds",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference)SetLoadBalancerBackendAddressPoolIds(val *[]*string) {
	if err := j.validateSetLoadBalancerBackendAddressPoolIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerBackendAddressPoolIds",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference)SetLoadBalancerInboundNatRulesIds(val *[]*string) {
	if err := j.validateSetLoadBalancerInboundNatRulesIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"loadBalancerInboundNatRulesIds",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference)SetPrimary(val interface{}) {
	if err := j.validateSetPrimaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primary",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) PutPublicIpAddressConfiguration(value *VirtualMachineScaleSetNetworkProfileIpConfigurationPublicIpAddressConfiguration) {
	if err := v.validatePutPublicIpAddressConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putPublicIpAddressConfiguration",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) ResetApplicationGatewayBackendAddressPoolIds() {
	_jsii_.InvokeVoid(
		v,
		"resetApplicationGatewayBackendAddressPoolIds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) ResetApplicationSecurityGroupIds() {
	_jsii_.InvokeVoid(
		v,
		"resetApplicationSecurityGroupIds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) ResetLoadBalancerBackendAddressPoolIds() {
	_jsii_.InvokeVoid(
		v,
		"resetLoadBalancerBackendAddressPoolIds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) ResetLoadBalancerInboundNatRulesIds() {
	_jsii_.InvokeVoid(
		v,
		"resetLoadBalancerInboundNatRulesIds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) ResetPublicIpAddressConfiguration() {
	_jsii_.InvokeVoid(
		v,
		"resetPublicIpAddressConfiguration",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileIpConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package virtualmachinescaleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/virtualmachinescaleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualMachineScaleSetNetworkProfileOutputReference interface {
	cdktf.ComplexObject
	AcceleratedNetworking() interface{}
	SetAcceleratedNetworking(val interface{})
	AcceleratedNetworkingInput() interface{}
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
	DnsSettings() VirtualMachineScaleSetNetworkProfileDnsSettingsOutputReference
	DnsSettingsInput() *VirtualMachineScaleSetNetworkProfileDnsSettings
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	IpConfiguration() VirtualMachineScaleSetNetworkProfileIpConfigurationList
	IpConfigurationInput() interface{}
	IpForwarding() interface{}
	SetIpForwarding(val interface{})
	IpForwardingInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkSecurityGroupId() *string
	SetNetworkSecurityGroupId(val *string)
	NetworkSecurityGroupIdInput() *string
	Primary() interface{}
	SetPrimary(val interface{})
	PrimaryInput() interface{}
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
	PutDnsSettings(value *VirtualMachineScaleSetNetworkProfileDnsSettings)
	PutIpConfiguration(value interface{})
	ResetAcceleratedNetworking()
	ResetDnsSettings()
	ResetIpForwarding()
	ResetNetworkSecurityGroupId()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualMachineScaleSetNetworkProfileOutputReference
type jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) AcceleratedNetworking() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceleratedNetworking",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) AcceleratedNetworkingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"acceleratedNetworkingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) DnsSettings() VirtualMachineScaleSetNetworkProfileDnsSettingsOutputReference {
	var returns VirtualMachineScaleSetNetworkProfileDnsSettingsOutputReference
	_jsii_.Get(
		j,
		"dnsSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) DnsSettingsInput() *VirtualMachineScaleSetNetworkProfileDnsSettings {
	var returns *VirtualMachineScaleSetNetworkProfileDnsSettings
	_jsii_.Get(
		j,
		"dnsSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) IpConfiguration() VirtualMachineScaleSetNetworkProfileIpConfigurationList {
	var returns VirtualMachineScaleSetNetworkProfileIpConfigurationList
	_jsii_.Get(
		j,
		"ipConfiguration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) IpConfigurationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipConfigurationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) IpForwarding() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipForwarding",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) IpForwardingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipForwardingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) NetworkSecurityGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkSecurityGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) NetworkSecurityGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkSecurityGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) Primary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) PrimaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVirtualMachineScaleSetNetworkProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VirtualMachineScaleSetNetworkProfileOutputReference {
	_init_.Initialize()

	if err := validateNewVirtualMachineScaleSetNetworkProfileOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualMachineScaleSet.VirtualMachineScaleSetNetworkProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVirtualMachineScaleSetNetworkProfileOutputReference_Override(v VirtualMachineScaleSetNetworkProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualMachineScaleSet.VirtualMachineScaleSetNetworkProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference)SetAcceleratedNetworking(val interface{}) {
	if err := j.validateSetAcceleratedNetworkingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"acceleratedNetworking",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference)SetIpForwarding(val interface{}) {
	if err := j.validateSetIpForwardingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipForwarding",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference)SetNetworkSecurityGroupId(val *string) {
	if err := j.validateSetNetworkSecurityGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkSecurityGroupId",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference)SetPrimary(val interface{}) {
	if err := j.validateSetPrimaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primary",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) PutDnsSettings(value *VirtualMachineScaleSetNetworkProfileDnsSettings) {
	if err := v.validatePutDnsSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putDnsSettings",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) PutIpConfiguration(value interface{}) {
	if err := v.validatePutIpConfigurationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"putIpConfiguration",
		[]interface{}{value},
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) ResetAcceleratedNetworking() {
	_jsii_.InvokeVoid(
		v,
		"resetAcceleratedNetworking",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) ResetDnsSettings() {
	_jsii_.InvokeVoid(
		v,
		"resetDnsSettings",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) ResetIpForwarding() {
	_jsii_.InvokeVoid(
		v,
		"resetIpForwarding",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) ResetNetworkSecurityGroupId() {
	_jsii_.InvokeVoid(
		v,
		"resetNetworkSecurityGroupId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VirtualMachineScaleSetNetworkProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


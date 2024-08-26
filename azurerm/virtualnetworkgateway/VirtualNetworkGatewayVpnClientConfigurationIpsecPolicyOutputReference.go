// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package virtualnetworkgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/virtualnetworkgateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference interface {
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
	DhGroup() *string
	SetDhGroup(val *string)
	DhGroupInput() *string
	// Experimental.
	Fqn() *string
	IkeEncryption() *string
	SetIkeEncryption(val *string)
	IkeEncryptionInput() *string
	IkeIntegrity() *string
	SetIkeIntegrity(val *string)
	IkeIntegrityInput() *string
	InternalValue() *VirtualNetworkGatewayVpnClientConfigurationIpsecPolicy
	SetInternalValue(val *VirtualNetworkGatewayVpnClientConfigurationIpsecPolicy)
	IpsecEncryption() *string
	SetIpsecEncryption(val *string)
	IpsecEncryptionInput() *string
	IpsecIntegrity() *string
	SetIpsecIntegrity(val *string)
	IpsecIntegrityInput() *string
	PfsGroup() *string
	SetPfsGroup(val *string)
	PfsGroupInput() *string
	SaDataSizeInKilobytes() *float64
	SetSaDataSizeInKilobytes(val *float64)
	SaDataSizeInKilobytesInput() *float64
	SaLifetimeInSeconds() *float64
	SetSaLifetimeInSeconds(val *float64)
	SaLifetimeInSecondsInput() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference
type jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) DhGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dhGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) DhGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dhGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) IkeEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) IkeEncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) IkeIntegrity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeIntegrity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) IkeIntegrityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeIntegrityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) InternalValue() *VirtualNetworkGatewayVpnClientConfigurationIpsecPolicy {
	var returns *VirtualNetworkGatewayVpnClientConfigurationIpsecPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) IpsecEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipsecEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) IpsecEncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipsecEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) IpsecIntegrity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipsecIntegrity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) IpsecIntegrityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipsecIntegrityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) PfsGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) PfsGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) SaDataSizeInKilobytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"saDataSizeInKilobytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) SaDataSizeInKilobytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"saDataSizeInKilobytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) SaLifetimeInSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"saLifetimeInSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) SaLifetimeInSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"saLifetimeInSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewVirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference_Override(v VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualNetworkGateway.VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference)SetDhGroup(val *string) {
	if err := j.validateSetDhGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dhGroup",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference)SetIkeEncryption(val *string) {
	if err := j.validateSetIkeEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ikeEncryption",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference)SetIkeIntegrity(val *string) {
	if err := j.validateSetIkeIntegrityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ikeIntegrity",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference)SetInternalValue(val *VirtualNetworkGatewayVpnClientConfigurationIpsecPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference)SetIpsecEncryption(val *string) {
	if err := j.validateSetIpsecEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipsecEncryption",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference)SetIpsecIntegrity(val *string) {
	if err := j.validateSetIpsecIntegrityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipsecIntegrity",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference)SetPfsGroup(val *string) {
	if err := j.validateSetPfsGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pfsGroup",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference)SetSaDataSizeInKilobytes(val *float64) {
	if err := j.validateSetSaDataSizeInKilobytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saDataSizeInKilobytes",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference)SetSaLifetimeInSeconds(val *float64) {
	if err := j.validateSetSaLifetimeInSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saLifetimeInSeconds",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VirtualNetworkGatewayVpnClientConfigurationIpsecPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpnserverconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/vpnserverconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpnServerConfigurationIpsecPolicyOutputReference interface {
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
	InternalValue() *VpnServerConfigurationIpsecPolicy
	SetInternalValue(val *VpnServerConfigurationIpsecPolicy)
	IpsecEncryption() *string
	SetIpsecEncryption(val *string)
	IpsecEncryptionInput() *string
	IpsecIntegrity() *string
	SetIpsecIntegrity(val *string)
	IpsecIntegrityInput() *string
	PfsGroup() *string
	SetPfsGroup(val *string)
	PfsGroupInput() *string
	SaDataSizeKilobytes() *float64
	SetSaDataSizeKilobytes(val *float64)
	SaDataSizeKilobytesInput() *float64
	SaLifetimeSeconds() *float64
	SetSaLifetimeSeconds(val *float64)
	SaLifetimeSecondsInput() *float64
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

// The jsii proxy struct for VpnServerConfigurationIpsecPolicyOutputReference
type jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) DhGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dhGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) DhGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dhGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) IkeEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) IkeEncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) IkeIntegrity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeIntegrity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) IkeIntegrityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeIntegrityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) InternalValue() *VpnServerConfigurationIpsecPolicy {
	var returns *VpnServerConfigurationIpsecPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) IpsecEncryption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipsecEncryption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) IpsecEncryptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipsecEncryptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) IpsecIntegrity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipsecIntegrity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) IpsecIntegrityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipsecIntegrityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) PfsGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) PfsGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) SaDataSizeKilobytes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"saDataSizeKilobytes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) SaDataSizeKilobytesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"saDataSizeKilobytesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) SaLifetimeSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"saLifetimeSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) SaLifetimeSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"saLifetimeSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVpnServerConfigurationIpsecPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) VpnServerConfigurationIpsecPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewVpnServerConfigurationIpsecPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnServerConfiguration.VpnServerConfigurationIpsecPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewVpnServerConfigurationIpsecPolicyOutputReference_Override(v VpnServerConfigurationIpsecPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnServerConfiguration.VpnServerConfigurationIpsecPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		v,
	)
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference)SetDhGroup(val *string) {
	if err := j.validateSetDhGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dhGroup",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference)SetIkeEncryption(val *string) {
	if err := j.validateSetIkeEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ikeEncryption",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference)SetIkeIntegrity(val *string) {
	if err := j.validateSetIkeIntegrityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ikeIntegrity",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference)SetInternalValue(val *VpnServerConfigurationIpsecPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference)SetIpsecEncryption(val *string) {
	if err := j.validateSetIpsecEncryptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipsecEncryption",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference)SetIpsecIntegrity(val *string) {
	if err := j.validateSetIpsecIntegrityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ipsecIntegrity",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference)SetPfsGroup(val *string) {
	if err := j.validateSetPfsGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pfsGroup",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference)SetSaDataSizeKilobytes(val *float64) {
	if err := j.validateSetSaDataSizeKilobytesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saDataSizeKilobytes",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference)SetSaLifetimeSeconds(val *float64) {
	if err := j.validateSetSaLifetimeSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saLifetimeSeconds",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VpnServerConfigurationIpsecPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


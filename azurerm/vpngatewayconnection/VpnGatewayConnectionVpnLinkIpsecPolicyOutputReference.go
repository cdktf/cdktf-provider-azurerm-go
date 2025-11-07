// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package vpngatewayconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/vpngatewayconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference interface {
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
	EncryptionAlgorithm() *string
	SetEncryptionAlgorithm(val *string)
	EncryptionAlgorithmInput() *string
	// Experimental.
	Fqn() *string
	IkeEncryptionAlgorithm() *string
	SetIkeEncryptionAlgorithm(val *string)
	IkeEncryptionAlgorithmInput() *string
	IkeIntegrityAlgorithm() *string
	SetIkeIntegrityAlgorithm(val *string)
	IkeIntegrityAlgorithmInput() *string
	IntegrityAlgorithm() *string
	SetIntegrityAlgorithm(val *string)
	IntegrityAlgorithmInput() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PfsGroup() *string
	SetPfsGroup(val *string)
	PfsGroupInput() *string
	SaDataSizeKb() *float64
	SetSaDataSizeKb(val *float64)
	SaDataSizeKbInput() *float64
	SaLifetimeSec() *float64
	SetSaLifetimeSec(val *float64)
	SaLifetimeSecInput() *float64
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference
type jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) DhGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dhGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) DhGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dhGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) EncryptionAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) EncryptionAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"encryptionAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) IkeEncryptionAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeEncryptionAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) IkeEncryptionAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeEncryptionAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) IkeIntegrityAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeIntegrityAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) IkeIntegrityAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ikeIntegrityAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) IntegrityAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrityAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) IntegrityAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrityAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) PfsGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfsGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) PfsGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pfsGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) SaDataSizeKb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"saDataSizeKb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) SaDataSizeKbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"saDataSizeKbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) SaLifetimeSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"saLifetimeSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) SaLifetimeSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"saLifetimeSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVpnGatewayConnectionVpnLinkIpsecPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewVpnGatewayConnectionVpnLinkIpsecPolicyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVpnGatewayConnectionVpnLinkIpsecPolicyOutputReference_Override(v VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.vpnGatewayConnection.VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference)SetDhGroup(val *string) {
	if err := j.validateSetDhGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dhGroup",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference)SetEncryptionAlgorithm(val *string) {
	if err := j.validateSetEncryptionAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionAlgorithm",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference)SetIkeEncryptionAlgorithm(val *string) {
	if err := j.validateSetIkeEncryptionAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ikeEncryptionAlgorithm",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference)SetIkeIntegrityAlgorithm(val *string) {
	if err := j.validateSetIkeIntegrityAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ikeIntegrityAlgorithm",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference)SetIntegrityAlgorithm(val *string) {
	if err := j.validateSetIntegrityAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrityAlgorithm",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference)SetPfsGroup(val *string) {
	if err := j.validateSetPfsGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pfsGroup",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference)SetSaDataSizeKb(val *float64) {
	if err := j.validateSetSaDataSizeKbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saDataSizeKb",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference)SetSaLifetimeSec(val *float64) {
	if err := j.validateSetSaLifetimeSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"saLifetimeSec",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VpnGatewayConnectionVpnLinkIpsecPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


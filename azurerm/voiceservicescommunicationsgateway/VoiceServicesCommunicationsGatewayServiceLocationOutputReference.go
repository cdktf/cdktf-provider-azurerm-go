// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package voiceservicescommunicationsgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/voiceservicescommunicationsgateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VoiceServicesCommunicationsGatewayServiceLocationOutputReference interface {
	cdktf.ComplexObject
	AllowedMediaSourceAddressPrefixes() *[]*string
	SetAllowedMediaSourceAddressPrefixes(val *[]*string)
	AllowedMediaSourceAddressPrefixesInput() *[]*string
	AllowedSignalingSourceAddressPrefixes() *[]*string
	SetAllowedSignalingSourceAddressPrefixes(val *[]*string)
	AllowedSignalingSourceAddressPrefixesInput() *[]*string
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
	EsrpAddresses() *[]*string
	SetEsrpAddresses(val *[]*string)
	EsrpAddressesInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	OperatorAddresses() *[]*string
	SetOperatorAddresses(val *[]*string)
	OperatorAddressesInput() *[]*string
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
	ResetAllowedMediaSourceAddressPrefixes()
	ResetAllowedSignalingSourceAddressPrefixes()
	ResetEsrpAddresses()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VoiceServicesCommunicationsGatewayServiceLocationOutputReference
type jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) AllowedMediaSourceAddressPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMediaSourceAddressPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) AllowedMediaSourceAddressPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedMediaSourceAddressPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) AllowedSignalingSourceAddressPrefixes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedSignalingSourceAddressPrefixes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) AllowedSignalingSourceAddressPrefixesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedSignalingSourceAddressPrefixesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) EsrpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"esrpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) EsrpAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"esrpAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) OperatorAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"operatorAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) OperatorAddressesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"operatorAddressesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVoiceServicesCommunicationsGatewayServiceLocationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VoiceServicesCommunicationsGatewayServiceLocationOutputReference {
	_init_.Initialize()

	if err := validateNewVoiceServicesCommunicationsGatewayServiceLocationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.voiceServicesCommunicationsGateway.VoiceServicesCommunicationsGatewayServiceLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVoiceServicesCommunicationsGatewayServiceLocationOutputReference_Override(v VoiceServicesCommunicationsGatewayServiceLocationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.voiceServicesCommunicationsGateway.VoiceServicesCommunicationsGatewayServiceLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference)SetAllowedMediaSourceAddressPrefixes(val *[]*string) {
	if err := j.validateSetAllowedMediaSourceAddressPrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedMediaSourceAddressPrefixes",
		val,
	)
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference)SetAllowedSignalingSourceAddressPrefixes(val *[]*string) {
	if err := j.validateSetAllowedSignalingSourceAddressPrefixesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedSignalingSourceAddressPrefixes",
		val,
	)
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference)SetEsrpAddresses(val *[]*string) {
	if err := j.validateSetEsrpAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"esrpAddresses",
		val,
	)
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference)SetOperatorAddresses(val *[]*string) {
	if err := j.validateSetOperatorAddressesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operatorAddresses",
		val,
	)
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (v *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (v *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (v *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (v *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (v *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (v *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (v *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (v *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (v *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) ResetAllowedMediaSourceAddressPrefixes() {
	_jsii_.InvokeVoid(
		v,
		"resetAllowedMediaSourceAddressPrefixes",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) ResetAllowedSignalingSourceAddressPrefixes() {
	_jsii_.InvokeVoid(
		v,
		"resetAllowedSignalingSourceAddressPrefixes",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) ResetEsrpAddresses() {
	_jsii_.InvokeVoid(
		v,
		"resetEsrpAddresses",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (v *jsiiProxy_VoiceServicesCommunicationsGatewayServiceLocationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


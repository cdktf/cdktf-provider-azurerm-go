// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/applicationgateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationGatewaySslProfileSslPolicyOutputReference interface {
	cdktf.ComplexObject
	CipherSuites() *[]*string
	SetCipherSuites(val *[]*string)
	CipherSuitesInput() *[]*string
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
	DisabledProtocols() *[]*string
	SetDisabledProtocols(val *[]*string)
	DisabledProtocolsInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *ApplicationGatewaySslProfileSslPolicy
	SetInternalValue(val *ApplicationGatewaySslProfileSslPolicy)
	MinProtocolVersion() *string
	SetMinProtocolVersion(val *string)
	MinProtocolVersionInput() *string
	PolicyName() *string
	SetPolicyName(val *string)
	PolicyNameInput() *string
	PolicyType() *string
	SetPolicyType(val *string)
	PolicyTypeInput() *string
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
	ResetCipherSuites()
	ResetDisabledProtocols()
	ResetMinProtocolVersion()
	ResetPolicyName()
	ResetPolicyType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationGatewaySslProfileSslPolicyOutputReference
type jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) CipherSuites() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cipherSuites",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) CipherSuitesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cipherSuitesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) DisabledProtocols() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disabledProtocols",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) DisabledProtocolsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disabledProtocolsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) InternalValue() *ApplicationGatewaySslProfileSslPolicy {
	var returns *ApplicationGatewaySslProfileSslPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) MinProtocolVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minProtocolVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) MinProtocolVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minProtocolVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) PolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) PolicyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) PolicyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) PolicyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewApplicationGatewaySslProfileSslPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ApplicationGatewaySslProfileSslPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationGatewaySslProfileSslPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGatewaySslProfileSslPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewApplicationGatewaySslProfileSslPolicyOutputReference_Override(a ApplicationGatewaySslProfileSslPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGatewaySslProfileSslPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference)SetCipherSuites(val *[]*string) {
	if err := j.validateSetCipherSuitesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cipherSuites",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference)SetDisabledProtocols(val *[]*string) {
	if err := j.validateSetDisabledProtocolsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disabledProtocols",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference)SetInternalValue(val *ApplicationGatewaySslProfileSslPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference)SetMinProtocolVersion(val *string) {
	if err := j.validateSetMinProtocolVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minProtocolVersion",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference)SetPolicyName(val *string) {
	if err := j.validateSetPolicyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyName",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference)SetPolicyType(val *string) {
	if err := j.validateSetPolicyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyType",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (a *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) ResetCipherSuites() {
	_jsii_.InvokeVoid(
		a,
		"resetCipherSuites",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) ResetDisabledProtocols() {
	_jsii_.InvokeVoid(
		a,
		"resetDisabledProtocols",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) ResetMinProtocolVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetMinProtocolVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) ResetPolicyName() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicyName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) ResetPolicyType() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicyType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewaySslProfileSslPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


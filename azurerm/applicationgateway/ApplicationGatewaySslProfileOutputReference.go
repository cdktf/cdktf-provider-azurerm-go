// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package applicationgateway

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/applicationgateway/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ApplicationGatewaySslProfileOutputReference interface {
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
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	SslPolicy() ApplicationGatewaySslProfileSslPolicyOutputReference
	SslPolicyInput() *ApplicationGatewaySslProfileSslPolicy
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TrustedClientCertificateNames() *[]*string
	SetTrustedClientCertificateNames(val *[]*string)
	TrustedClientCertificateNamesInput() *[]*string
	VerifyClientCertIssuerDn() interface{}
	SetVerifyClientCertIssuerDn(val interface{})
	VerifyClientCertIssuerDnInput() interface{}
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
	PutSslPolicy(value *ApplicationGatewaySslProfileSslPolicy)
	ResetSslPolicy()
	ResetTrustedClientCertificateNames()
	ResetVerifyClientCertIssuerDn()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ApplicationGatewaySslProfileOutputReference
type jsiiProxy_ApplicationGatewaySslProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ApplicationGatewaySslProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileOutputReference) SslPolicy() ApplicationGatewaySslProfileSslPolicyOutputReference {
	var returns ApplicationGatewaySslProfileSslPolicyOutputReference
	_jsii_.Get(
		j,
		"sslPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileOutputReference) SslPolicyInput() *ApplicationGatewaySslProfileSslPolicy {
	var returns *ApplicationGatewaySslProfileSslPolicy
	_jsii_.Get(
		j,
		"sslPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileOutputReference) TrustedClientCertificateNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedClientCertificateNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileOutputReference) TrustedClientCertificateNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"trustedClientCertificateNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileOutputReference) VerifyClientCertIssuerDn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifyClientCertIssuerDn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApplicationGatewaySslProfileOutputReference) VerifyClientCertIssuerDnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifyClientCertIssuerDnInput",
		&returns,
	)
	return returns
}


func NewApplicationGatewaySslProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ApplicationGatewaySslProfileOutputReference {
	_init_.Initialize()

	if err := validateNewApplicationGatewaySslProfileOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ApplicationGatewaySslProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGatewaySslProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewApplicationGatewaySslProfileOutputReference_Override(a ApplicationGatewaySslProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.applicationGateway.ApplicationGatewaySslProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		a,
	)
}

func (j *jsiiProxy_ApplicationGatewaySslProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewaySslProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewaySslProfileOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewaySslProfileOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewaySslProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewaySslProfileOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewaySslProfileOutputReference)SetTrustedClientCertificateNames(val *[]*string) {
	if err := j.validateSetTrustedClientCertificateNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"trustedClientCertificateNames",
		val,
	)
}

func (j *jsiiProxy_ApplicationGatewaySslProfileOutputReference)SetVerifyClientCertIssuerDn(val interface{}) {
	if err := j.validateSetVerifyClientCertIssuerDnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifyClientCertIssuerDn",
		val,
	)
}

func (a *jsiiProxy_ApplicationGatewaySslProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewaySslProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_ApplicationGatewaySslProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationGatewaySslProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_ApplicationGatewaySslProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_ApplicationGatewaySslProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_ApplicationGatewaySslProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_ApplicationGatewaySslProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_ApplicationGatewaySslProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_ApplicationGatewaySslProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_ApplicationGatewaySslProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApplicationGatewaySslProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (a *jsiiProxy_ApplicationGatewaySslProfileOutputReference) PutSslPolicy(value *ApplicationGatewaySslProfileSslPolicy) {
	if err := a.validatePutSslPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSslPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_ApplicationGatewaySslProfileOutputReference) ResetSslPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetSslPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewaySslProfileOutputReference) ResetTrustedClientCertificateNames() {
	_jsii_.InvokeVoid(
		a,
		"resetTrustedClientCertificateNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewaySslProfileOutputReference) ResetVerifyClientCertIssuerDn() {
	_jsii_.InvokeVoid(
		a,
		"resetVerifyClientCertIssuerDn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApplicationGatewaySslProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (a *jsiiProxy_ApplicationGatewaySslProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


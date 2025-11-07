// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package expressrouteport

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/expressrouteport/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ExpressRoutePortLink2OutputReference interface {
	cdktf.ComplexObject
	AdminEnabled() interface{}
	SetAdminEnabled(val interface{})
	AdminEnabledInput() interface{}
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
	ConnectorType() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	Id() *string
	InterfaceName() *string
	InternalValue() *ExpressRoutePortLink2
	SetInternalValue(val *ExpressRoutePortLink2)
	MacsecCakKeyvaultSecretId() *string
	SetMacsecCakKeyvaultSecretId(val *string)
	MacsecCakKeyvaultSecretIdInput() *string
	MacsecCipher() *string
	SetMacsecCipher(val *string)
	MacsecCipherInput() *string
	MacsecCknKeyvaultSecretId() *string
	SetMacsecCknKeyvaultSecretId(val *string)
	MacsecCknKeyvaultSecretIdInput() *string
	MacsecSciEnabled() interface{}
	SetMacsecSciEnabled(val interface{})
	MacsecSciEnabledInput() interface{}
	PatchPanelId() *string
	RackId() *string
	RouterName() *string
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
	ResetAdminEnabled()
	ResetMacsecCakKeyvaultSecretId()
	ResetMacsecCipher()
	ResetMacsecCknKeyvaultSecretId()
	ResetMacsecSciEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ExpressRoutePortLink2OutputReference
type jsiiProxy_ExpressRoutePortLink2OutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference) AdminEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference) AdminEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference) ConnectorType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectorType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference) InterfaceName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"interfaceName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference) InternalValue() *ExpressRoutePortLink2 {
	var returns *ExpressRoutePortLink2
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference) MacsecCakKeyvaultSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macsecCakKeyvaultSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference) MacsecCakKeyvaultSecretIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macsecCakKeyvaultSecretIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference) MacsecCipher() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macsecCipher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference) MacsecCipherInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macsecCipherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference) MacsecCknKeyvaultSecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macsecCknKeyvaultSecretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference) MacsecCknKeyvaultSecretIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"macsecCknKeyvaultSecretIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference) MacsecSciEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"macsecSciEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference) MacsecSciEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"macsecSciEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference) PatchPanelId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchPanelId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference) RackId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rackId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference) RouterName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"routerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewExpressRoutePortLink2OutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ExpressRoutePortLink2OutputReference {
	_init_.Initialize()

	if err := validateNewExpressRoutePortLink2OutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ExpressRoutePortLink2OutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.expressRoutePort.ExpressRoutePortLink2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewExpressRoutePortLink2OutputReference_Override(e ExpressRoutePortLink2OutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.expressRoutePort.ExpressRoutePortLink2OutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		e,
	)
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference)SetAdminEnabled(val interface{}) {
	if err := j.validateSetAdminEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminEnabled",
		val,
	)
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference)SetInternalValue(val *ExpressRoutePortLink2) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference)SetMacsecCakKeyvaultSecretId(val *string) {
	if err := j.validateSetMacsecCakKeyvaultSecretIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"macsecCakKeyvaultSecretId",
		val,
	)
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference)SetMacsecCipher(val *string) {
	if err := j.validateSetMacsecCipherParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"macsecCipher",
		val,
	)
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference)SetMacsecCknKeyvaultSecretId(val *string) {
	if err := j.validateSetMacsecCknKeyvaultSecretIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"macsecCknKeyvaultSecretId",
		val,
	)
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference)SetMacsecSciEnabled(val interface{}) {
	if err := j.validateSetMacsecSciEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"macsecSciEnabled",
		val,
	)
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ExpressRoutePortLink2OutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (e *jsiiProxy_ExpressRoutePortLink2OutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRoutePortLink2OutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := e.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRoutePortLink2OutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRoutePortLink2OutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := e.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		e,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRoutePortLink2OutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := e.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		e,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRoutePortLink2OutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := e.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		e,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRoutePortLink2OutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := e.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		e,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRoutePortLink2OutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := e.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		e,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRoutePortLink2OutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := e.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRoutePortLink2OutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := e.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		e,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRoutePortLink2OutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRoutePortLink2OutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := e.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRoutePortLink2OutputReference) ResetAdminEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetAdminEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRoutePortLink2OutputReference) ResetMacsecCakKeyvaultSecretId() {
	_jsii_.InvokeVoid(
		e,
		"resetMacsecCakKeyvaultSecretId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRoutePortLink2OutputReference) ResetMacsecCipher() {
	_jsii_.InvokeVoid(
		e,
		"resetMacsecCipher",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRoutePortLink2OutputReference) ResetMacsecCknKeyvaultSecretId() {
	_jsii_.InvokeVoid(
		e,
		"resetMacsecCknKeyvaultSecretId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRoutePortLink2OutputReference) ResetMacsecSciEnabled() {
	_jsii_.InvokeVoid(
		e,
		"resetMacsecSciEnabled",
		nil, // no parameters
	)
}

func (e *jsiiProxy_ExpressRoutePortLink2OutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := e.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		e,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_ExpressRoutePortLink2OutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


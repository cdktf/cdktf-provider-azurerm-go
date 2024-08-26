// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package iotsecuritydevicegroup

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/iotsecuritydevicegroup/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IotSecurityDeviceGroupAllowRuleOutputReference interface {
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
	ConnectionFromIpsNotAllowed() *[]*string
	SetConnectionFromIpsNotAllowed(val *[]*string)
	ConnectionFromIpsNotAllowedInput() *[]*string
	ConnectionToIpsNotAllowed() *[]*string
	SetConnectionToIpsNotAllowed(val *[]*string)
	ConnectionToIpsNotAllowedInput() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *IotSecurityDeviceGroupAllowRule
	SetInternalValue(val *IotSecurityDeviceGroupAllowRule)
	LocalUsersNotAllowed() *[]*string
	SetLocalUsersNotAllowed(val *[]*string)
	LocalUsersNotAllowedInput() *[]*string
	ProcessesNotAllowed() *[]*string
	SetProcessesNotAllowed(val *[]*string)
	ProcessesNotAllowedInput() *[]*string
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
	ResetConnectionFromIpsNotAllowed()
	ResetConnectionToIpsNotAllowed()
	ResetLocalUsersNotAllowed()
	ResetProcessesNotAllowed()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for IotSecurityDeviceGroupAllowRuleOutputReference
type jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) ConnectionFromIpsNotAllowed() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"connectionFromIpsNotAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) ConnectionFromIpsNotAllowedInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"connectionFromIpsNotAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) ConnectionToIpsNotAllowed() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"connectionToIpsNotAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) ConnectionToIpsNotAllowedInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"connectionToIpsNotAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) InternalValue() *IotSecurityDeviceGroupAllowRule {
	var returns *IotSecurityDeviceGroupAllowRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) LocalUsersNotAllowed() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localUsersNotAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) LocalUsersNotAllowedInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"localUsersNotAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) ProcessesNotAllowed() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"processesNotAllowed",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) ProcessesNotAllowedInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"processesNotAllowedInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewIotSecurityDeviceGroupAllowRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) IotSecurityDeviceGroupAllowRuleOutputReference {
	_init_.Initialize()

	if err := validateNewIotSecurityDeviceGroupAllowRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.iotSecurityDeviceGroup.IotSecurityDeviceGroupAllowRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewIotSecurityDeviceGroupAllowRuleOutputReference_Override(i IotSecurityDeviceGroupAllowRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.iotSecurityDeviceGroup.IotSecurityDeviceGroupAllowRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		i,
	)
}

func (j *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference)SetConnectionFromIpsNotAllowed(val *[]*string) {
	if err := j.validateSetConnectionFromIpsNotAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionFromIpsNotAllowed",
		val,
	)
}

func (j *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference)SetConnectionToIpsNotAllowed(val *[]*string) {
	if err := j.validateSetConnectionToIpsNotAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionToIpsNotAllowed",
		val,
	)
}

func (j *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference)SetInternalValue(val *IotSecurityDeviceGroupAllowRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference)SetLocalUsersNotAllowed(val *[]*string) {
	if err := j.validateSetLocalUsersNotAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"localUsersNotAllowed",
		val,
	)
}

func (j *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference)SetProcessesNotAllowed(val *[]*string) {
	if err := j.validateSetProcessesNotAllowedParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"processesNotAllowed",
		val,
	)
}

func (j *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (i *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) ResetConnectionFromIpsNotAllowed() {
	_jsii_.InvokeVoid(
		i,
		"resetConnectionFromIpsNotAllowed",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) ResetConnectionToIpsNotAllowed() {
	_jsii_.InvokeVoid(
		i,
		"resetConnectionToIpsNotAllowed",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) ResetLocalUsersNotAllowed() {
	_jsii_.InvokeVoid(
		i,
		"resetLocalUsersNotAllowed",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) ResetProcessesNotAllowed() {
	_jsii_.InvokeVoid(
		i,
		"resetProcessesNotAllowed",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := i.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		i,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IotSecurityDeviceGroupAllowRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


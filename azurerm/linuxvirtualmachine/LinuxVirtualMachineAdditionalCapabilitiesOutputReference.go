// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxvirtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/linuxvirtualmachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LinuxVirtualMachineAdditionalCapabilitiesOutputReference interface {
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
	HibernationEnabled() interface{}
	SetHibernationEnabled(val interface{})
	HibernationEnabledInput() interface{}
	InternalValue() *LinuxVirtualMachineAdditionalCapabilities
	SetInternalValue(val *LinuxVirtualMachineAdditionalCapabilities)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UltraSsdEnabled() interface{}
	SetUltraSsdEnabled(val interface{})
	UltraSsdEnabledInput() interface{}
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
	ResetHibernationEnabled()
	ResetUltraSsdEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LinuxVirtualMachineAdditionalCapabilitiesOutputReference
type jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) HibernationEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hibernationEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) HibernationEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hibernationEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) InternalValue() *LinuxVirtualMachineAdditionalCapabilities {
	var returns *LinuxVirtualMachineAdditionalCapabilities
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) UltraSsdEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ultraSsdEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) UltraSsdEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ultraSsdEnabledInput",
		&returns,
	)
	return returns
}


func NewLinuxVirtualMachineAdditionalCapabilitiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LinuxVirtualMachineAdditionalCapabilitiesOutputReference {
	_init_.Initialize()

	if err := validateNewLinuxVirtualMachineAdditionalCapabilitiesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxVirtualMachine.LinuxVirtualMachineAdditionalCapabilitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLinuxVirtualMachineAdditionalCapabilitiesOutputReference_Override(l LinuxVirtualMachineAdditionalCapabilitiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxVirtualMachine.LinuxVirtualMachineAdditionalCapabilitiesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference)SetHibernationEnabled(val interface{}) {
	if err := j.validateSetHibernationEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hibernationEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference)SetInternalValue(val *LinuxVirtualMachineAdditionalCapabilities) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference)SetUltraSsdEnabled(val interface{}) {
	if err := j.validateSetUltraSsdEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ultraSsdEnabled",
		val,
	)
}

func (l *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) ResetHibernationEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetHibernationEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) ResetUltraSsdEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetUltraSsdEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineAdditionalCapabilitiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


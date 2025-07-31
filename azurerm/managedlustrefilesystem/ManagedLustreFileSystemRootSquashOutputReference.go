// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package managedlustrefilesystem

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/managedlustrefilesystem/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ManagedLustreFileSystemRootSquashOutputReference interface {
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
	InternalValue() *ManagedLustreFileSystemRootSquash
	SetInternalValue(val *ManagedLustreFileSystemRootSquash)
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	NoSquashNids() *string
	SetNoSquashNids(val *string)
	NoSquashNidsInput() *string
	SquashGid() *float64
	SetSquashGid(val *float64)
	SquashGidInput() *float64
	SquashUid() *float64
	SetSquashUid(val *float64)
	SquashUidInput() *float64
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
	ResetSquashGid()
	ResetSquashUid()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ManagedLustreFileSystemRootSquashOutputReference
type jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) InternalValue() *ManagedLustreFileSystemRootSquash {
	var returns *ManagedLustreFileSystemRootSquash
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) NoSquashNids() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noSquashNids",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) NoSquashNidsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"noSquashNidsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) SquashGid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"squashGid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) SquashGidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"squashGidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) SquashUid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"squashUid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) SquashUidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"squashUidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewManagedLustreFileSystemRootSquashOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ManagedLustreFileSystemRootSquashOutputReference {
	_init_.Initialize()

	if err := validateNewManagedLustreFileSystemRootSquashOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.managedLustreFileSystem.ManagedLustreFileSystemRootSquashOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewManagedLustreFileSystemRootSquashOutputReference_Override(m ManagedLustreFileSystemRootSquashOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.managedLustreFileSystem.ManagedLustreFileSystemRootSquashOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference)SetInternalValue(val *ManagedLustreFileSystemRootSquash) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference)SetNoSquashNids(val *string) {
	if err := j.validateSetNoSquashNidsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noSquashNids",
		val,
	)
}

func (j *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference)SetSquashGid(val *float64) {
	if err := j.validateSetSquashGidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"squashGid",
		val,
	)
}

func (j *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference)SetSquashUid(val *float64) {
	if err := j.validateSetSquashUidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"squashUid",
		val,
	)
}

func (j *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) ResetSquashGid() {
	_jsii_.InvokeVoid(
		m,
		"resetSquashGid",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) ResetSquashUid() {
	_jsii_.InvokeVoid(
		m,
		"resetSquashUid",
		nil, // no parameters
	)
}

func (m *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_ManagedLustreFileSystemRootSquashOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


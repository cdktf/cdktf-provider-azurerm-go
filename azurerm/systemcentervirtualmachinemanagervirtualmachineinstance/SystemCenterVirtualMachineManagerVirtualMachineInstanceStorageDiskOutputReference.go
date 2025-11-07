// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package systemcentervirtualmachinemanagervirtualmachineinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/systemcentervirtualmachinemanagervirtualmachineinstance/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference interface {
	cdktf.ComplexObject
	Bus() *float64
	SetBus(val *float64)
	BusInput() *float64
	BusType() *string
	SetBusType(val *string)
	BusTypeInput() *string
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
	DiskSizeGb() *float64
	SetDiskSizeGb(val *float64)
	DiskSizeGbInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Lun() *float64
	SetLun(val *float64)
	LunInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	StorageQosPolicyName() *string
	SetStorageQosPolicyName(val *string)
	StorageQosPolicyNameInput() *string
	TemplateDiskId() *string
	SetTemplateDiskId(val *string)
	TemplateDiskIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VhdType() *string
	SetVhdType(val *string)
	VhdTypeInput() *string
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
	ResetBus()
	ResetBusType()
	ResetDiskSizeGb()
	ResetLun()
	ResetName()
	ResetStorageQosPolicyName()
	ResetTemplateDiskId()
	ResetVhdType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference
type jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) Bus() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) BusInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"busInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) BusType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"busType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) BusTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"busTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) DiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) Lun() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lun",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) LunInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lunInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) StorageQosPolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageQosPolicyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) StorageQosPolicyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageQosPolicyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) TemplateDiskId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateDiskId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) TemplateDiskIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"templateDiskIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) VhdType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vhdType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) VhdTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vhdTypeInput",
		&returns,
	)
	return returns
}


func NewSystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference {
	_init_.Initialize()

	if err := validateNewSystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.systemCenterVirtualMachineManagerVirtualMachineInstance.SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewSystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference_Override(s SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.systemCenterVirtualMachineManagerVirtualMachineInstance.SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference)SetBus(val *float64) {
	if err := j.validateSetBusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bus",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference)SetBusType(val *string) {
	if err := j.validateSetBusTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"busType",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference)SetDiskSizeGb(val *float64) {
	if err := j.validateSetDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSizeGb",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference)SetLun(val *float64) {
	if err := j.validateSetLunParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lun",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference)SetStorageQosPolicyName(val *string) {
	if err := j.validateSetStorageQosPolicyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageQosPolicyName",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference)SetTemplateDiskId(val *string) {
	if err := j.validateSetTemplateDiskIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"templateDiskId",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference)SetVhdType(val *string) {
	if err := j.validateSetVhdTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vhdType",
		val,
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) ResetBus() {
	_jsii_.InvokeVoid(
		s,
		"resetBus",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) ResetBusType() {
	_jsii_.InvokeVoid(
		s,
		"resetBusType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) ResetDiskSizeGb() {
	_jsii_.InvokeVoid(
		s,
		"resetDiskSizeGb",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) ResetLun() {
	_jsii_.InvokeVoid(
		s,
		"resetLun",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) ResetName() {
	_jsii_.InvokeVoid(
		s,
		"resetName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) ResetStorageQosPolicyName() {
	_jsii_.InvokeVoid(
		s,
		"resetStorageQosPolicyName",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) ResetTemplateDiskId() {
	_jsii_.InvokeVoid(
		s,
		"resetTemplateDiskId",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) ResetVhdType() {
	_jsii_.InvokeVoid(
		s,
		"resetVhdType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_SystemCenterVirtualMachineManagerVirtualMachineInstanceStorageDiskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


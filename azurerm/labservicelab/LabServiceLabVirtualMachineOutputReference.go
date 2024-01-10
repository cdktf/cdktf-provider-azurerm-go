// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package labservicelab

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/labservicelab/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LabServiceLabVirtualMachineOutputReference interface {
	cdktf.ComplexObject
	AdditionalCapabilityGpuDriversInstalled() interface{}
	SetAdditionalCapabilityGpuDriversInstalled(val interface{})
	AdditionalCapabilityGpuDriversInstalledInput() interface{}
	AdminUser() LabServiceLabVirtualMachineAdminUserOutputReference
	AdminUserInput() *LabServiceLabVirtualMachineAdminUser
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
	CreateOption() *string
	SetCreateOption(val *string)
	CreateOptionInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	ImageReference() LabServiceLabVirtualMachineImageReferenceOutputReference
	ImageReferenceInput() *LabServiceLabVirtualMachineImageReference
	InternalValue() *LabServiceLabVirtualMachine
	SetInternalValue(val *LabServiceLabVirtualMachine)
	NonAdminUser() LabServiceLabVirtualMachineNonAdminUserOutputReference
	NonAdminUserInput() *LabServiceLabVirtualMachineNonAdminUser
	SharedPasswordEnabled() interface{}
	SetSharedPasswordEnabled(val interface{})
	SharedPasswordEnabledInput() interface{}
	Sku() LabServiceLabVirtualMachineSkuOutputReference
	SkuInput() *LabServiceLabVirtualMachineSku
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UsageQuota() *string
	SetUsageQuota(val *string)
	UsageQuotaInput() *string
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
	PutAdminUser(value *LabServiceLabVirtualMachineAdminUser)
	PutImageReference(value *LabServiceLabVirtualMachineImageReference)
	PutNonAdminUser(value *LabServiceLabVirtualMachineNonAdminUser)
	PutSku(value *LabServiceLabVirtualMachineSku)
	ResetAdditionalCapabilityGpuDriversInstalled()
	ResetCreateOption()
	ResetNonAdminUser()
	ResetSharedPasswordEnabled()
	ResetUsageQuota()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LabServiceLabVirtualMachineOutputReference
type jsiiProxy_LabServiceLabVirtualMachineOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference) AdditionalCapabilityGpuDriversInstalled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalCapabilityGpuDriversInstalled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference) AdditionalCapabilityGpuDriversInstalledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalCapabilityGpuDriversInstalledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference) AdminUser() LabServiceLabVirtualMachineAdminUserOutputReference {
	var returns LabServiceLabVirtualMachineAdminUserOutputReference
	_jsii_.Get(
		j,
		"adminUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference) AdminUserInput() *LabServiceLabVirtualMachineAdminUser {
	var returns *LabServiceLabVirtualMachineAdminUser
	_jsii_.Get(
		j,
		"adminUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference) CreateOption() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createOption",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference) CreateOptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"createOptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference) ImageReference() LabServiceLabVirtualMachineImageReferenceOutputReference {
	var returns LabServiceLabVirtualMachineImageReferenceOutputReference
	_jsii_.Get(
		j,
		"imageReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference) ImageReferenceInput() *LabServiceLabVirtualMachineImageReference {
	var returns *LabServiceLabVirtualMachineImageReference
	_jsii_.Get(
		j,
		"imageReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference) InternalValue() *LabServiceLabVirtualMachine {
	var returns *LabServiceLabVirtualMachine
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference) NonAdminUser() LabServiceLabVirtualMachineNonAdminUserOutputReference {
	var returns LabServiceLabVirtualMachineNonAdminUserOutputReference
	_jsii_.Get(
		j,
		"nonAdminUser",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference) NonAdminUserInput() *LabServiceLabVirtualMachineNonAdminUser {
	var returns *LabServiceLabVirtualMachineNonAdminUser
	_jsii_.Get(
		j,
		"nonAdminUserInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference) SharedPasswordEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharedPasswordEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference) SharedPasswordEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sharedPasswordEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference) Sku() LabServiceLabVirtualMachineSkuOutputReference {
	var returns LabServiceLabVirtualMachineSkuOutputReference
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference) SkuInput() *LabServiceLabVirtualMachineSku {
	var returns *LabServiceLabVirtualMachineSku
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference) UsageQuota() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageQuota",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference) UsageQuotaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageQuotaInput",
		&returns,
	)
	return returns
}


func NewLabServiceLabVirtualMachineOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LabServiceLabVirtualMachineOutputReference {
	_init_.Initialize()

	if err := validateNewLabServiceLabVirtualMachineOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LabServiceLabVirtualMachineOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.labServiceLab.LabServiceLabVirtualMachineOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLabServiceLabVirtualMachineOutputReference_Override(l LabServiceLabVirtualMachineOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.labServiceLab.LabServiceLabVirtualMachineOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference)SetAdditionalCapabilityGpuDriversInstalled(val interface{}) {
	if err := j.validateSetAdditionalCapabilityGpuDriversInstalledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"additionalCapabilityGpuDriversInstalled",
		val,
	)
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference)SetCreateOption(val *string) {
	if err := j.validateSetCreateOptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"createOption",
		val,
	)
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference)SetInternalValue(val *LabServiceLabVirtualMachine) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference)SetSharedPasswordEnabled(val interface{}) {
	if err := j.validateSetSharedPasswordEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sharedPasswordEnabled",
		val,
	)
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LabServiceLabVirtualMachineOutputReference)SetUsageQuota(val *string) {
	if err := j.validateSetUsageQuotaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usageQuota",
		val,
	)
}

func (l *jsiiProxy_LabServiceLabVirtualMachineOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLabVirtualMachineOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LabServiceLabVirtualMachineOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LabServiceLabVirtualMachineOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LabServiceLabVirtualMachineOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LabServiceLabVirtualMachineOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LabServiceLabVirtualMachineOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LabServiceLabVirtualMachineOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LabServiceLabVirtualMachineOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LabServiceLabVirtualMachineOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LabServiceLabVirtualMachineOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LabServiceLabVirtualMachineOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LabServiceLabVirtualMachineOutputReference) PutAdminUser(value *LabServiceLabVirtualMachineAdminUser) {
	if err := l.validatePutAdminUserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAdminUser",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LabServiceLabVirtualMachineOutputReference) PutImageReference(value *LabServiceLabVirtualMachineImageReference) {
	if err := l.validatePutImageReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putImageReference",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LabServiceLabVirtualMachineOutputReference) PutNonAdminUser(value *LabServiceLabVirtualMachineNonAdminUser) {
	if err := l.validatePutNonAdminUserParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putNonAdminUser",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LabServiceLabVirtualMachineOutputReference) PutSku(value *LabServiceLabVirtualMachineSku) {
	if err := l.validatePutSkuParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSku",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LabServiceLabVirtualMachineOutputReference) ResetAdditionalCapabilityGpuDriversInstalled() {
	_jsii_.InvokeVoid(
		l,
		"resetAdditionalCapabilityGpuDriversInstalled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LabServiceLabVirtualMachineOutputReference) ResetCreateOption() {
	_jsii_.InvokeVoid(
		l,
		"resetCreateOption",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LabServiceLabVirtualMachineOutputReference) ResetNonAdminUser() {
	_jsii_.InvokeVoid(
		l,
		"resetNonAdminUser",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LabServiceLabVirtualMachineOutputReference) ResetSharedPasswordEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetSharedPasswordEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LabServiceLabVirtualMachineOutputReference) ResetUsageQuota() {
	_jsii_.InvokeVoid(
		l,
		"resetUsageQuota",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LabServiceLabVirtualMachineOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LabServiceLabVirtualMachineOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


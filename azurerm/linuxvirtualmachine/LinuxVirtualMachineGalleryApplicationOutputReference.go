// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxvirtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/linuxvirtualmachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LinuxVirtualMachineGalleryApplicationOutputReference interface {
	cdktf.ComplexObject
	AutomaticUpgradeEnabled() interface{}
	SetAutomaticUpgradeEnabled(val interface{})
	AutomaticUpgradeEnabledInput() interface{}
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
	ConfigurationBlobUri() *string
	SetConfigurationBlobUri(val *string)
	ConfigurationBlobUriInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Order() *float64
	SetOrder(val *float64)
	OrderInput() *float64
	Tag() *string
	SetTag(val *string)
	TagInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TreatFailureAsDeploymentFailureEnabled() interface{}
	SetTreatFailureAsDeploymentFailureEnabled(val interface{})
	TreatFailureAsDeploymentFailureEnabledInput() interface{}
	VersionId() *string
	SetVersionId(val *string)
	VersionIdInput() *string
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
	ResetAutomaticUpgradeEnabled()
	ResetConfigurationBlobUri()
	ResetOrder()
	ResetTag()
	ResetTreatFailureAsDeploymentFailureEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LinuxVirtualMachineGalleryApplicationOutputReference
type jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) AutomaticUpgradeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticUpgradeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) AutomaticUpgradeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticUpgradeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) ConfigurationBlobUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationBlobUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) ConfigurationBlobUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationBlobUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) Order() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) OrderInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"orderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) Tag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) TagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) TreatFailureAsDeploymentFailureEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"treatFailureAsDeploymentFailureEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) TreatFailureAsDeploymentFailureEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"treatFailureAsDeploymentFailureEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) VersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) VersionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionIdInput",
		&returns,
	)
	return returns
}


func NewLinuxVirtualMachineGalleryApplicationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) LinuxVirtualMachineGalleryApplicationOutputReference {
	_init_.Initialize()

	if err := validateNewLinuxVirtualMachineGalleryApplicationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxVirtualMachine.LinuxVirtualMachineGalleryApplicationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewLinuxVirtualMachineGalleryApplicationOutputReference_Override(l LinuxVirtualMachineGalleryApplicationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxVirtualMachine.LinuxVirtualMachineGalleryApplicationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		l,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference)SetAutomaticUpgradeEnabled(val interface{}) {
	if err := j.validateSetAutomaticUpgradeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticUpgradeEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference)SetConfigurationBlobUri(val *string) {
	if err := j.validateSetConfigurationBlobUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationBlobUri",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference)SetOrder(val *float64) {
	if err := j.validateSetOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"order",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference)SetTag(val *string) {
	if err := j.validateSetTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tag",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference)SetTreatFailureAsDeploymentFailureEnabled(val interface{}) {
	if err := j.validateSetTreatFailureAsDeploymentFailureEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"treatFailureAsDeploymentFailureEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference)SetVersionId(val *string) {
	if err := j.validateSetVersionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionId",
		val,
	)
}

func (l *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) ResetAutomaticUpgradeEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetAutomaticUpgradeEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) ResetConfigurationBlobUri() {
	_jsii_.InvokeVoid(
		l,
		"resetConfigurationBlobUri",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) ResetOrder() {
	_jsii_.InvokeVoid(
		l,
		"resetOrder",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) ResetTag() {
	_jsii_.InvokeVoid(
		l,
		"resetTag",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) ResetTreatFailureAsDeploymentFailureEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetTreatFailureAsDeploymentFailureEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LinuxVirtualMachineGalleryApplicationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


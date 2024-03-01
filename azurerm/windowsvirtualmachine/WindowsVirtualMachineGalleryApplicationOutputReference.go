// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsvirtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/windowsvirtualmachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type WindowsVirtualMachineGalleryApplicationOutputReference interface {
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

// The jsii proxy struct for WindowsVirtualMachineGalleryApplicationOutputReference
type jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) AutomaticUpgradeEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticUpgradeEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) AutomaticUpgradeEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticUpgradeEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) ConfigurationBlobUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationBlobUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) ConfigurationBlobUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"configurationBlobUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) Order() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) OrderInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"orderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) Tag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) TagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) TreatFailureAsDeploymentFailureEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"treatFailureAsDeploymentFailureEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) TreatFailureAsDeploymentFailureEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"treatFailureAsDeploymentFailureEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) VersionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) VersionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"versionIdInput",
		&returns,
	)
	return returns
}


func NewWindowsVirtualMachineGalleryApplicationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) WindowsVirtualMachineGalleryApplicationOutputReference {
	_init_.Initialize()

	if err := validateNewWindowsVirtualMachineGalleryApplicationOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachine.WindowsVirtualMachineGalleryApplicationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewWindowsVirtualMachineGalleryApplicationOutputReference_Override(w WindowsVirtualMachineGalleryApplicationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachine.WindowsVirtualMachineGalleryApplicationOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference)SetAutomaticUpgradeEnabled(val interface{}) {
	if err := j.validateSetAutomaticUpgradeEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automaticUpgradeEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference)SetConfigurationBlobUri(val *string) {
	if err := j.validateSetConfigurationBlobUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"configurationBlobUri",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference)SetOrder(val *float64) {
	if err := j.validateSetOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"order",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference)SetTag(val *string) {
	if err := j.validateSetTagParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tag",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference)SetTreatFailureAsDeploymentFailureEnabled(val interface{}) {
	if err := j.validateSetTreatFailureAsDeploymentFailureEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"treatFailureAsDeploymentFailureEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference)SetVersionId(val *string) {
	if err := j.validateSetVersionIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"versionId",
		val,
	)
}

func (w *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) ResetAutomaticUpgradeEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetAutomaticUpgradeEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) ResetConfigurationBlobUri() {
	_jsii_.InvokeVoid(
		w,
		"resetConfigurationBlobUri",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) ResetOrder() {
	_jsii_.InvokeVoid(
		w,
		"resetOrder",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) ResetTag() {
	_jsii_.InvokeVoid(
		w,
		"resetTag",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) ResetTreatFailureAsDeploymentFailureEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetTreatFailureAsDeploymentFailureEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := w.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		w,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineGalleryApplicationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


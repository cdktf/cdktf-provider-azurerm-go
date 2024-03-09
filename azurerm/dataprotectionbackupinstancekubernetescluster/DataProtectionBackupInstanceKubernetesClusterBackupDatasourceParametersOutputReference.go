// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataprotectionbackupinstancekubernetescluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/dataprotectionbackupinstancekubernetescluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference interface {
	cdktf.ComplexObject
	ClusterScopedResourcesEnabled() interface{}
	SetClusterScopedResourcesEnabled(val interface{})
	ClusterScopedResourcesEnabledInput() interface{}
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
	ExcludedNamespaces() *[]*string
	SetExcludedNamespaces(val *[]*string)
	ExcludedNamespacesInput() *[]*string
	ExcludedResourceTypes() *[]*string
	SetExcludedResourceTypes(val *[]*string)
	ExcludedResourceTypesInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludedNamespaces() *[]*string
	SetIncludedNamespaces(val *[]*string)
	IncludedNamespacesInput() *[]*string
	IncludedResourceTypes() *[]*string
	SetIncludedResourceTypes(val *[]*string)
	IncludedResourceTypesInput() *[]*string
	InternalValue() *DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParameters
	SetInternalValue(val *DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParameters)
	LabelSelectors() *[]*string
	SetLabelSelectors(val *[]*string)
	LabelSelectorsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VolumeSnapshotEnabled() interface{}
	SetVolumeSnapshotEnabled(val interface{})
	VolumeSnapshotEnabledInput() interface{}
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
	ResetClusterScopedResourcesEnabled()
	ResetExcludedNamespaces()
	ResetExcludedResourceTypes()
	ResetIncludedNamespaces()
	ResetIncludedResourceTypes()
	ResetLabelSelectors()
	ResetVolumeSnapshotEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference
type jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) ClusterScopedResourcesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusterScopedResourcesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) ClusterScopedResourcesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clusterScopedResourcesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) ExcludedNamespaces() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) ExcludedNamespacesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) ExcludedResourceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedResourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) ExcludedResourceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedResourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) IncludedNamespaces() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) IncludedNamespacesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) IncludedResourceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedResourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) IncludedResourceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedResourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) InternalValue() *DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParameters {
	var returns *DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) LabelSelectors() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labelSelectors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) LabelSelectorsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"labelSelectorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) VolumeSnapshotEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeSnapshotEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) VolumeSnapshotEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"volumeSnapshotEnabledInput",
		&returns,
	)
	return returns
}


func NewDataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference {
	_init_.Initialize()

	if err := validateNewDataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataProtectionBackupInstanceKubernetesCluster.DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference_Override(d DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataProtectionBackupInstanceKubernetesCluster.DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference)SetClusterScopedResourcesEnabled(val interface{}) {
	if err := j.validateSetClusterScopedResourcesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clusterScopedResourcesEnabled",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference)SetExcludedNamespaces(val *[]*string) {
	if err := j.validateSetExcludedNamespacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedNamespaces",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference)SetExcludedResourceTypes(val *[]*string) {
	if err := j.validateSetExcludedResourceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedResourceTypes",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference)SetIncludedNamespaces(val *[]*string) {
	if err := j.validateSetIncludedNamespacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedNamespaces",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference)SetIncludedResourceTypes(val *[]*string) {
	if err := j.validateSetIncludedResourceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedResourceTypes",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference)SetInternalValue(val *DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference)SetLabelSelectors(val *[]*string) {
	if err := j.validateSetLabelSelectorsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labelSelectors",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference)SetVolumeSnapshotEnabled(val interface{}) {
	if err := j.validateSetVolumeSnapshotEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"volumeSnapshotEnabled",
		val,
	)
}

func (d *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) ResetClusterScopedResourcesEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetClusterScopedResourcesEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) ResetExcludedNamespaces() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludedNamespaces",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) ResetExcludedResourceTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetExcludedResourceTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) ResetIncludedNamespaces() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludedNamespaces",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) ResetIncludedResourceTypes() {
	_jsii_.InvokeVoid(
		d,
		"resetIncludedResourceTypes",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) ResetLabelSelectors() {
	_jsii_.InvokeVoid(
		d,
		"resetLabelSelectors",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) ResetVolumeSnapshotEnabled() {
	_jsii_.InvokeVoid(
		d,
		"resetVolumeSnapshotEnabled",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataProtectionBackupInstanceKubernetesClusterBackupDatasourceParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


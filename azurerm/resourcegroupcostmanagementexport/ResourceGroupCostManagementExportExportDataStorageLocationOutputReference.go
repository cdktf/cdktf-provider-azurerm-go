// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcegroupcostmanagementexport

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/resourcegroupcostmanagementexport/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ResourceGroupCostManagementExportExportDataStorageLocationOutputReference interface {
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
	ContainerId() *string
	SetContainerId(val *string)
	ContainerIdInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *ResourceGroupCostManagementExportExportDataStorageLocation
	SetInternalValue(val *ResourceGroupCostManagementExportExportDataStorageLocation)
	RootFolderPath() *string
	SetRootFolderPath(val *string)
	RootFolderPathInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ResourceGroupCostManagementExportExportDataStorageLocationOutputReference
type jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference) ContainerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference) ContainerIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"containerIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference) InternalValue() *ResourceGroupCostManagementExportExportDataStorageLocation {
	var returns *ResourceGroupCostManagementExportExportDataStorageLocation
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference) RootFolderPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootFolderPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference) RootFolderPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rootFolderPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewResourceGroupCostManagementExportExportDataStorageLocationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ResourceGroupCostManagementExportExportDataStorageLocationOutputReference {
	_init_.Initialize()

	if err := validateNewResourceGroupCostManagementExportExportDataStorageLocationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.resourceGroupCostManagementExport.ResourceGroupCostManagementExportExportDataStorageLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewResourceGroupCostManagementExportExportDataStorageLocationOutputReference_Override(r ResourceGroupCostManagementExportExportDataStorageLocationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.resourceGroupCostManagementExport.ResourceGroupCostManagementExportExportDataStorageLocationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference)SetContainerId(val *string) {
	if err := j.validateSetContainerIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerId",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference)SetInternalValue(val *ResourceGroupCostManagementExportExportDataStorageLocation) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference)SetRootFolderPath(val *string) {
	if err := j.validateSetRootFolderPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootFolderPath",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := r.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		r,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := r.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := r.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		r,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := r.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		r,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := r.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		r,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := r.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		r,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := r.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		r,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := r.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		r,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := r.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		r,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := r.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := r.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		r,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementExportExportDataStorageLocationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


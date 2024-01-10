// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package resourcegroupcostmanagementview

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/resourcegroupcostmanagementview/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ResourceGroupCostManagementViewDatasetOutputReference interface {
	cdktf.ComplexObject
	Aggregation() ResourceGroupCostManagementViewDatasetAggregationList
	AggregationInput() interface{}
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
	Granularity() *string
	SetGranularity(val *string)
	GranularityInput() *string
	Grouping() ResourceGroupCostManagementViewDatasetGroupingList
	GroupingInput() interface{}
	InternalValue() *ResourceGroupCostManagementViewDataset
	SetInternalValue(val *ResourceGroupCostManagementViewDataset)
	Sorting() ResourceGroupCostManagementViewDatasetSortingList
	SortingInput() interface{}
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
	PutAggregation(value interface{})
	PutGrouping(value interface{})
	PutSorting(value interface{})
	ResetGrouping()
	ResetSorting()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ResourceGroupCostManagementViewDatasetOutputReference
type jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) Aggregation() ResourceGroupCostManagementViewDatasetAggregationList {
	var returns ResourceGroupCostManagementViewDatasetAggregationList
	_jsii_.Get(
		j,
		"aggregation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) AggregationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"aggregationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) Granularity() *string {
	var returns *string
	_jsii_.Get(
		j,
		"granularity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) GranularityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"granularityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) Grouping() ResourceGroupCostManagementViewDatasetGroupingList {
	var returns ResourceGroupCostManagementViewDatasetGroupingList
	_jsii_.Get(
		j,
		"grouping",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) GroupingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"groupingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) InternalValue() *ResourceGroupCostManagementViewDataset {
	var returns *ResourceGroupCostManagementViewDataset
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) Sorting() ResourceGroupCostManagementViewDatasetSortingList {
	var returns ResourceGroupCostManagementViewDatasetSortingList
	_jsii_.Get(
		j,
		"sorting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) SortingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sortingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewResourceGroupCostManagementViewDatasetOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ResourceGroupCostManagementViewDatasetOutputReference {
	_init_.Initialize()

	if err := validateNewResourceGroupCostManagementViewDatasetOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.resourceGroupCostManagementView.ResourceGroupCostManagementViewDatasetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewResourceGroupCostManagementViewDatasetOutputReference_Override(r ResourceGroupCostManagementViewDatasetOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.resourceGroupCostManagementView.ResourceGroupCostManagementViewDatasetOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		r,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference)SetGranularity(val *string) {
	if err := j.validateSetGranularityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"granularity",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference)SetInternalValue(val *ResourceGroupCostManagementViewDataset) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (r *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (r *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (r *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (r *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (r *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (r *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (r *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (r *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (r *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (r *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		r,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (r *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (r *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) PutAggregation(value interface{}) {
	if err := r.validatePutAggregationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putAggregation",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) PutGrouping(value interface{}) {
	if err := r.validatePutGroupingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putGrouping",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) PutSorting(value interface{}) {
	if err := r.validatePutSortingParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		r,
		"putSorting",
		[]interface{}{value},
	)
}

func (r *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) ResetGrouping() {
	_jsii_.InvokeVoid(
		r,
		"resetGrouping",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) ResetSorting() {
	_jsii_.InvokeVoid(
		r,
		"resetSorting",
		nil, // no parameters
	)
}

func (r *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (r *jsiiProxy_ResourceGroupCostManagementViewDatasetOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		r,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


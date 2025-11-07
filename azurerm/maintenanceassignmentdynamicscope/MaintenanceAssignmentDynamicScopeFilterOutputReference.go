// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package maintenanceassignmentdynamicscope

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/maintenanceassignmentdynamicscope/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MaintenanceAssignmentDynamicScopeFilterOutputReference interface {
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
	InternalValue() *MaintenanceAssignmentDynamicScopeFilter
	SetInternalValue(val *MaintenanceAssignmentDynamicScopeFilter)
	Locations() *[]*string
	SetLocations(val *[]*string)
	LocationsInput() *[]*string
	OsTypes() *[]*string
	SetOsTypes(val *[]*string)
	OsTypesInput() *[]*string
	ResourceGroups() *[]*string
	SetResourceGroups(val *[]*string)
	ResourceGroupsInput() *[]*string
	ResourceTypes() *[]*string
	SetResourceTypes(val *[]*string)
	ResourceTypesInput() *[]*string
	TagFilter() *string
	SetTagFilter(val *string)
	TagFilterInput() *string
	Tags() MaintenanceAssignmentDynamicScopeFilterTagsList
	TagsInput() interface{}
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutTags(value interface{})
	ResetLocations()
	ResetOsTypes()
	ResetResourceGroups()
	ResetResourceTypes()
	ResetTagFilter()
	ResetTags()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MaintenanceAssignmentDynamicScopeFilterOutputReference
type jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) InternalValue() *MaintenanceAssignmentDynamicScopeFilter {
	var returns *MaintenanceAssignmentDynamicScopeFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) Locations() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) LocationsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"locationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) OsTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"osTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) OsTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"osTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) ResourceGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) ResourceGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) ResourceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) ResourceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) TagFilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagFilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) TagFilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tagFilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) Tags() MaintenanceAssignmentDynamicScopeFilterTagsList {
	var returns MaintenanceAssignmentDynamicScopeFilterTagsList
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) TagsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMaintenanceAssignmentDynamicScopeFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MaintenanceAssignmentDynamicScopeFilterOutputReference {
	_init_.Initialize()

	if err := validateNewMaintenanceAssignmentDynamicScopeFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.maintenanceAssignmentDynamicScope.MaintenanceAssignmentDynamicScopeFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMaintenanceAssignmentDynamicScopeFilterOutputReference_Override(m MaintenanceAssignmentDynamicScopeFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.maintenanceAssignmentDynamicScope.MaintenanceAssignmentDynamicScopeFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference)SetInternalValue(val *MaintenanceAssignmentDynamicScopeFilter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference)SetLocations(val *[]*string) {
	if err := j.validateSetLocationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"locations",
		val,
	)
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference)SetOsTypes(val *[]*string) {
	if err := j.validateSetOsTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"osTypes",
		val,
	)
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference)SetResourceGroups(val *[]*string) {
	if err := j.validateSetResourceGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroups",
		val,
	)
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference)SetResourceTypes(val *[]*string) {
	if err := j.validateSetResourceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceTypes",
		val,
	)
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference)SetTagFilter(val *string) {
	if err := j.validateSetTagFilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tagFilter",
		val,
	)
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) PutTags(value interface{}) {
	if err := m.validatePutTagsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTags",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) ResetLocations() {
	_jsii_.InvokeVoid(
		m,
		"resetLocations",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) ResetOsTypes() {
	_jsii_.InvokeVoid(
		m,
		"resetOsTypes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) ResetResourceGroups() {
	_jsii_.InvokeVoid(
		m,
		"resetResourceGroups",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) ResetResourceTypes() {
	_jsii_.InvokeVoid(
		m,
		"resetResourceTypes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) ResetTagFilter() {
	_jsii_.InvokeVoid(
		m,
		"resetTagFilter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) ResetTags() {
	_jsii_.InvokeVoid(
		m,
		"resetTags",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MaintenanceAssignmentDynamicScopeFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


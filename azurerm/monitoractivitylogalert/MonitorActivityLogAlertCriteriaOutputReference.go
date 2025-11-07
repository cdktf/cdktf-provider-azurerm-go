// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitoractivitylogalert

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/monitoractivitylogalert/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorActivityLogAlertCriteriaOutputReference interface {
	cdktf.ComplexObject
	Caller() *string
	SetCaller(val *string)
	CallerInput() *string
	Category() *string
	SetCategory(val *string)
	CategoryInput() *string
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
	InternalValue() *MonitorActivityLogAlertCriteria
	SetInternalValue(val *MonitorActivityLogAlertCriteria)
	Level() *string
	SetLevel(val *string)
	LevelInput() *string
	Levels() *[]*string
	SetLevels(val *[]*string)
	LevelsInput() *[]*string
	OperationName() *string
	SetOperationName(val *string)
	OperationNameInput() *string
	RecommendationCategory() *string
	SetRecommendationCategory(val *string)
	RecommendationCategoryInput() *string
	RecommendationImpact() *string
	SetRecommendationImpact(val *string)
	RecommendationImpactInput() *string
	RecommendationType() *string
	SetRecommendationType(val *string)
	RecommendationTypeInput() *string
	ResourceGroup() *string
	SetResourceGroup(val *string)
	ResourceGroupInput() *string
	ResourceGroups() *[]*string
	SetResourceGroups(val *[]*string)
	ResourceGroupsInput() *[]*string
	ResourceHealth() MonitorActivityLogAlertCriteriaResourceHealthOutputReference
	ResourceHealthInput() *MonitorActivityLogAlertCriteriaResourceHealth
	ResourceId() *string
	SetResourceId(val *string)
	ResourceIdInput() *string
	ResourceIds() *[]*string
	SetResourceIds(val *[]*string)
	ResourceIdsInput() *[]*string
	ResourceProvider() *string
	SetResourceProvider(val *string)
	ResourceProviderInput() *string
	ResourceProviders() *[]*string
	SetResourceProviders(val *[]*string)
	ResourceProvidersInput() *[]*string
	ResourceType() *string
	SetResourceType(val *string)
	ResourceTypeInput() *string
	ResourceTypes() *[]*string
	SetResourceTypes(val *[]*string)
	ResourceTypesInput() *[]*string
	ServiceHealth() MonitorActivityLogAlertCriteriaServiceHealthOutputReference
	ServiceHealthInput() *MonitorActivityLogAlertCriteriaServiceHealth
	Status() *string
	SetStatus(val *string)
	Statuses() *[]*string
	SetStatuses(val *[]*string)
	StatusesInput() *[]*string
	StatusInput() *string
	SubStatus() *string
	SetSubStatus(val *string)
	SubStatuses() *[]*string
	SetSubStatuses(val *[]*string)
	SubStatusesInput() *[]*string
	SubStatusInput() *string
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
	PutResourceHealth(value *MonitorActivityLogAlertCriteriaResourceHealth)
	PutServiceHealth(value *MonitorActivityLogAlertCriteriaServiceHealth)
	ResetCaller()
	ResetLevel()
	ResetLevels()
	ResetOperationName()
	ResetRecommendationCategory()
	ResetRecommendationImpact()
	ResetRecommendationType()
	ResetResourceGroup()
	ResetResourceGroups()
	ResetResourceHealth()
	ResetResourceId()
	ResetResourceIds()
	ResetResourceProvider()
	ResetResourceProviders()
	ResetResourceType()
	ResetResourceTypes()
	ResetServiceHealth()
	ResetStatus()
	ResetStatuses()
	ResetSubStatus()
	ResetSubStatuses()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorActivityLogAlertCriteriaOutputReference
type jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) Caller() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caller",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) CallerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"callerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) Category() *string {
	var returns *string
	_jsii_.Get(
		j,
		"category",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) CategoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"categoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) InternalValue() *MonitorActivityLogAlertCriteria {
	var returns *MonitorActivityLogAlertCriteria
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) Level() *string {
	var returns *string
	_jsii_.Get(
		j,
		"level",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) LevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"levelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) Levels() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"levels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) LevelsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"levelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) OperationName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) OperationNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"operationNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) RecommendationCategory() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommendationCategory",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) RecommendationCategoryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommendationCategoryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) RecommendationImpact() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommendationImpact",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) RecommendationImpactInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommendationImpactInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) RecommendationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommendationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) RecommendationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"recommendationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceGroup() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceGroupInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceHealth() MonitorActivityLogAlertCriteriaResourceHealthOutputReference {
	var returns MonitorActivityLogAlertCriteriaResourceHealthOutputReference
	_jsii_.Get(
		j,
		"resourceHealth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceHealthInput() *MonitorActivityLogAlertCriteriaResourceHealth {
	var returns *MonitorActivityLogAlertCriteriaResourceHealth
	_jsii_.Get(
		j,
		"resourceHealthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceProvider() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceProvider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceProviderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceProviderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceProviders() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceProviders",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceProvidersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceProvidersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResourceTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"resourceTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ServiceHealth() MonitorActivityLogAlertCriteriaServiceHealthOutputReference {
	var returns MonitorActivityLogAlertCriteriaServiceHealthOutputReference
	_jsii_.Get(
		j,
		"serviceHealth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ServiceHealthInput() *MonitorActivityLogAlertCriteriaServiceHealth {
	var returns *MonitorActivityLogAlertCriteriaServiceHealth
	_jsii_.Get(
		j,
		"serviceHealthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) Status() *string {
	var returns *string
	_jsii_.Get(
		j,
		"status",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) Statuses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statuses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) StatusesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"statusesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) StatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"statusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) SubStatus() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subStatus",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) SubStatuses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subStatuses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) SubStatusesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"subStatusesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) SubStatusInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subStatusInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorActivityLogAlertCriteriaOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorActivityLogAlertCriteriaOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorActivityLogAlertCriteriaOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActivityLogAlert.MonitorActivityLogAlertCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorActivityLogAlertCriteriaOutputReference_Override(m MonitorActivityLogAlertCriteriaOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorActivityLogAlert.MonitorActivityLogAlertCriteriaOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference)SetCaller(val *string) {
	if err := j.validateSetCallerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caller",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference)SetCategory(val *string) {
	if err := j.validateSetCategoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"category",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference)SetInternalValue(val *MonitorActivityLogAlertCriteria) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference)SetLevel(val *string) {
	if err := j.validateSetLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"level",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference)SetLevels(val *[]*string) {
	if err := j.validateSetLevelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"levels",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference)SetOperationName(val *string) {
	if err := j.validateSetOperationNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationName",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference)SetRecommendationCategory(val *string) {
	if err := j.validateSetRecommendationCategoryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recommendationCategory",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference)SetRecommendationImpact(val *string) {
	if err := j.validateSetRecommendationImpactParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recommendationImpact",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference)SetRecommendationType(val *string) {
	if err := j.validateSetRecommendationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"recommendationType",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference)SetResourceGroup(val *string) {
	if err := j.validateSetResourceGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroup",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference)SetResourceGroups(val *[]*string) {
	if err := j.validateSetResourceGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroups",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference)SetResourceId(val *string) {
	if err := j.validateSetResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceId",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference)SetResourceIds(val *[]*string) {
	if err := j.validateSetResourceIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceIds",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference)SetResourceProvider(val *string) {
	if err := j.validateSetResourceProviderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceProvider",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference)SetResourceProviders(val *[]*string) {
	if err := j.validateSetResourceProvidersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceProviders",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference)SetResourceType(val *string) {
	if err := j.validateSetResourceTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceType",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference)SetResourceTypes(val *[]*string) {
	if err := j.validateSetResourceTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceTypes",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference)SetStatus(val *string) {
	if err := j.validateSetStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"status",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference)SetStatuses(val *[]*string) {
	if err := j.validateSetStatusesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"statuses",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference)SetSubStatus(val *string) {
	if err := j.validateSetSubStatusParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subStatus",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference)SetSubStatuses(val *[]*string) {
	if err := j.validateSetSubStatusesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subStatuses",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) PutResourceHealth(value *MonitorActivityLogAlertCriteriaResourceHealth) {
	if err := m.validatePutResourceHealthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putResourceHealth",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) PutServiceHealth(value *MonitorActivityLogAlertCriteriaServiceHealth) {
	if err := m.validatePutServiceHealthParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putServiceHealth",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetCaller() {
	_jsii_.InvokeVoid(
		m,
		"resetCaller",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetLevel() {
	_jsii_.InvokeVoid(
		m,
		"resetLevel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetLevels() {
	_jsii_.InvokeVoid(
		m,
		"resetLevels",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetOperationName() {
	_jsii_.InvokeVoid(
		m,
		"resetOperationName",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetRecommendationCategory() {
	_jsii_.InvokeVoid(
		m,
		"resetRecommendationCategory",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetRecommendationImpact() {
	_jsii_.InvokeVoid(
		m,
		"resetRecommendationImpact",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetRecommendationType() {
	_jsii_.InvokeVoid(
		m,
		"resetRecommendationType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetResourceGroup() {
	_jsii_.InvokeVoid(
		m,
		"resetResourceGroup",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetResourceGroups() {
	_jsii_.InvokeVoid(
		m,
		"resetResourceGroups",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetResourceHealth() {
	_jsii_.InvokeVoid(
		m,
		"resetResourceHealth",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetResourceId() {
	_jsii_.InvokeVoid(
		m,
		"resetResourceId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetResourceIds() {
	_jsii_.InvokeVoid(
		m,
		"resetResourceIds",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetResourceProvider() {
	_jsii_.InvokeVoid(
		m,
		"resetResourceProvider",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetResourceProviders() {
	_jsii_.InvokeVoid(
		m,
		"resetResourceProviders",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetResourceType() {
	_jsii_.InvokeVoid(
		m,
		"resetResourceType",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetResourceTypes() {
	_jsii_.InvokeVoid(
		m,
		"resetResourceTypes",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetServiceHealth() {
	_jsii_.InvokeVoid(
		m,
		"resetServiceHealth",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetStatus() {
	_jsii_.InvokeVoid(
		m,
		"resetStatus",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetStatuses() {
	_jsii_.InvokeVoid(
		m,
		"resetStatuses",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetSubStatus() {
	_jsii_.InvokeVoid(
		m,
		"resetSubStatus",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ResetSubStatuses() {
	_jsii_.InvokeVoid(
		m,
		"resetSubStatuses",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MonitorActivityLogAlertCriteriaOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


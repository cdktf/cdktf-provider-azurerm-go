// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermstoragemanagementpolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/dataazurermstoragemanagementpolicy/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference interface {
	cdktf.ComplexObject
	AutoTierToHotFromCoolEnabled() cdktf.IResolvable
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
	DeleteAfterDaysSinceCreationGreaterThan() *float64
	DeleteAfterDaysSinceLastAccessTimeGreaterThan() *float64
	DeleteAfterDaysSinceModificationGreaterThan() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *DataAzurermStorageManagementPolicyRuleActionsBaseBlob
	SetInternalValue(val *DataAzurermStorageManagementPolicyRuleActionsBaseBlob)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TierToArchiveAfterDaysSinceCreationGreaterThan() *float64
	TierToArchiveAfterDaysSinceLastAccessTimeGreaterThan() *float64
	TierToArchiveAfterDaysSinceLastTierChangeGreaterThan() *float64
	TierToArchiveAfterDaysSinceModificationGreaterThan() *float64
	TierToColdAfterDaysSinceCreationGreaterThan() *float64
	TierToColdAfterDaysSinceLastAccessTimeGreaterThan() *float64
	TierToColdAfterDaysSinceModificationGreaterThan() *float64
	TierToCoolAfterDaysSinceCreationGreaterThan() *float64
	TierToCoolAfterDaysSinceLastAccessTimeGreaterThan() *float64
	TierToCoolAfterDaysSinceModificationGreaterThan() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference
type jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) AutoTierToHotFromCoolEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"autoTierToHotFromCoolEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) DeleteAfterDaysSinceCreationGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteAfterDaysSinceCreationGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) DeleteAfterDaysSinceLastAccessTimeGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteAfterDaysSinceLastAccessTimeGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) DeleteAfterDaysSinceModificationGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"deleteAfterDaysSinceModificationGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) InternalValue() *DataAzurermStorageManagementPolicyRuleActionsBaseBlob {
	var returns *DataAzurermStorageManagementPolicyRuleActionsBaseBlob
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToArchiveAfterDaysSinceCreationGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToArchiveAfterDaysSinceCreationGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToArchiveAfterDaysSinceLastAccessTimeGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToArchiveAfterDaysSinceLastAccessTimeGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToArchiveAfterDaysSinceLastTierChangeGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToArchiveAfterDaysSinceLastTierChangeGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToArchiveAfterDaysSinceModificationGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToArchiveAfterDaysSinceModificationGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToColdAfterDaysSinceCreationGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToColdAfterDaysSinceCreationGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToColdAfterDaysSinceLastAccessTimeGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToColdAfterDaysSinceLastAccessTimeGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToColdAfterDaysSinceModificationGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToColdAfterDaysSinceModificationGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToCoolAfterDaysSinceCreationGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToCoolAfterDaysSinceCreationGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToCoolAfterDaysSinceLastAccessTimeGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToCoolAfterDaysSinceLastAccessTimeGreaterThan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) TierToCoolAfterDaysSinceModificationGreaterThan() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tierToCoolAfterDaysSinceModificationGreaterThan",
		&returns,
	)
	return returns
}


func NewDataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference {
	_init_.Initialize()

	if err := validateNewDataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermStorageManagementPolicy.DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference_Override(d DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermStorageManagementPolicy.DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference)SetInternalValue(val *DataAzurermStorageManagementPolicyRuleActionsBaseBlob) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermStorageManagementPolicyRuleActionsBaseBlobOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


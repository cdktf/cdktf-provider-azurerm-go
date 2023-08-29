// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cosmosdbsqlcontainer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/cosmosdbsqlcontainer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CosmosdbSqlContainerIndexingPolicyOutputReference interface {
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
	CompositeIndex() CosmosdbSqlContainerIndexingPolicyCompositeIndexList
	CompositeIndexInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExcludedPath() CosmosdbSqlContainerIndexingPolicyExcludedPathList
	ExcludedPathInput() interface{}
	// Experimental.
	Fqn() *string
	IncludedPath() CosmosdbSqlContainerIndexingPolicyIncludedPathList
	IncludedPathInput() interface{}
	IndexingMode() *string
	SetIndexingMode(val *string)
	IndexingModeInput() *string
	InternalValue() *CosmosdbSqlContainerIndexingPolicy
	SetInternalValue(val *CosmosdbSqlContainerIndexingPolicy)
	SpatialIndex() CosmosdbSqlContainerIndexingPolicySpatialIndexList
	SpatialIndexInput() interface{}
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
	PutCompositeIndex(value interface{})
	PutExcludedPath(value interface{})
	PutIncludedPath(value interface{})
	PutSpatialIndex(value interface{})
	ResetCompositeIndex()
	ResetExcludedPath()
	ResetIncludedPath()
	ResetIndexingMode()
	ResetSpatialIndex()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CosmosdbSqlContainerIndexingPolicyOutputReference
type jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) CompositeIndex() CosmosdbSqlContainerIndexingPolicyCompositeIndexList {
	var returns CosmosdbSqlContainerIndexingPolicyCompositeIndexList
	_jsii_.Get(
		j,
		"compositeIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) CompositeIndexInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compositeIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) ExcludedPath() CosmosdbSqlContainerIndexingPolicyExcludedPathList {
	var returns CosmosdbSqlContainerIndexingPolicyExcludedPathList
	_jsii_.Get(
		j,
		"excludedPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) ExcludedPathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludedPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) IncludedPath() CosmosdbSqlContainerIndexingPolicyIncludedPathList {
	var returns CosmosdbSqlContainerIndexingPolicyIncludedPathList
	_jsii_.Get(
		j,
		"includedPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) IncludedPathInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"includedPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) IndexingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) IndexingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) InternalValue() *CosmosdbSqlContainerIndexingPolicy {
	var returns *CosmosdbSqlContainerIndexingPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) SpatialIndex() CosmosdbSqlContainerIndexingPolicySpatialIndexList {
	var returns CosmosdbSqlContainerIndexingPolicySpatialIndexList
	_jsii_.Get(
		j,
		"spatialIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) SpatialIndexInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spatialIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCosmosdbSqlContainerIndexingPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CosmosdbSqlContainerIndexingPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewCosmosdbSqlContainerIndexingPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbSqlContainer.CosmosdbSqlContainerIndexingPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCosmosdbSqlContainerIndexingPolicyOutputReference_Override(c CosmosdbSqlContainerIndexingPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbSqlContainer.CosmosdbSqlContainerIndexingPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference)SetIndexingMode(val *string) {
	if err := j.validateSetIndexingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexingMode",
		val,
	)
}

func (j *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference)SetInternalValue(val *CosmosdbSqlContainerIndexingPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) PutCompositeIndex(value interface{}) {
	if err := c.validatePutCompositeIndexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCompositeIndex",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) PutExcludedPath(value interface{}) {
	if err := c.validatePutExcludedPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putExcludedPath",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) PutIncludedPath(value interface{}) {
	if err := c.validatePutIncludedPathParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIncludedPath",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) PutSpatialIndex(value interface{}) {
	if err := c.validatePutSpatialIndexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSpatialIndex",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) ResetCompositeIndex() {
	_jsii_.InvokeVoid(
		c,
		"resetCompositeIndex",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) ResetExcludedPath() {
	_jsii_.InvokeVoid(
		c,
		"resetExcludedPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) ResetIncludedPath() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludedPath",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) ResetIndexingMode() {
	_jsii_.InvokeVoid(
		c,
		"resetIndexingMode",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) ResetSpatialIndex() {
	_jsii_.InvokeVoid(
		c,
		"resetSpatialIndex",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbSqlContainerIndexingPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


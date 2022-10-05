package cosmosdbgremlingraph

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/cosmosdbgremlingraph/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CosmosdbGremlinGraphIndexPolicyOutputReference interface {
	cdktf.ComplexObject
	Automatic() interface{}
	SetAutomatic(val interface{})
	AutomaticInput() interface{}
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
	CompositeIndex() CosmosdbGremlinGraphIndexPolicyCompositeIndexList
	CompositeIndexInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	ExcludedPaths() *[]*string
	SetExcludedPaths(val *[]*string)
	ExcludedPathsInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludedPaths() *[]*string
	SetIncludedPaths(val *[]*string)
	IncludedPathsInput() *[]*string
	IndexingMode() *string
	SetIndexingMode(val *string)
	IndexingModeInput() *string
	InternalValue() *CosmosdbGremlinGraphIndexPolicy
	SetInternalValue(val *CosmosdbGremlinGraphIndexPolicy)
	SpatialIndex() CosmosdbGremlinGraphIndexPolicySpatialIndexList
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
	PutSpatialIndex(value interface{})
	ResetAutomatic()
	ResetCompositeIndex()
	ResetExcludedPaths()
	ResetIncludedPaths()
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

// The jsii proxy struct for CosmosdbGremlinGraphIndexPolicyOutputReference
type jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) Automatic() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automatic",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) AutomaticInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"automaticInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) CompositeIndex() CosmosdbGremlinGraphIndexPolicyCompositeIndexList {
	var returns CosmosdbGremlinGraphIndexPolicyCompositeIndexList
	_jsii_.Get(
		j,
		"compositeIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) CompositeIndexInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compositeIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) ExcludedPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) ExcludedPathsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludedPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) IncludedPaths() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) IncludedPathsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"includedPathsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) IndexingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) IndexingModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"indexingModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) InternalValue() *CosmosdbGremlinGraphIndexPolicy {
	var returns *CosmosdbGremlinGraphIndexPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) SpatialIndex() CosmosdbGremlinGraphIndexPolicySpatialIndexList {
	var returns CosmosdbGremlinGraphIndexPolicySpatialIndexList
	_jsii_.Get(
		j,
		"spatialIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) SpatialIndexInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"spatialIndexInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCosmosdbGremlinGraphIndexPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CosmosdbGremlinGraphIndexPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewCosmosdbGremlinGraphIndexPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbGremlinGraph.CosmosdbGremlinGraphIndexPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCosmosdbGremlinGraphIndexPolicyOutputReference_Override(c CosmosdbGremlinGraphIndexPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cosmosdbGremlinGraph.CosmosdbGremlinGraphIndexPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference)SetAutomatic(val interface{}) {
	if err := j.validateSetAutomaticParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"automatic",
		val,
	)
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference)SetExcludedPaths(val *[]*string) {
	if err := j.validateSetExcludedPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludedPaths",
		val,
	)
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference)SetIncludedPaths(val *[]*string) {
	if err := j.validateSetIncludedPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includedPaths",
		val,
	)
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference)SetIndexingMode(val *string) {
	if err := j.validateSetIndexingModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"indexingMode",
		val,
	)
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference)SetInternalValue(val *CosmosdbGremlinGraphIndexPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) PutCompositeIndex(value interface{}) {
	if err := c.validatePutCompositeIndexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCompositeIndex",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) PutSpatialIndex(value interface{}) {
	if err := c.validatePutSpatialIndexParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSpatialIndex",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) ResetAutomatic() {
	_jsii_.InvokeVoid(
		c,
		"resetAutomatic",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) ResetCompositeIndex() {
	_jsii_.InvokeVoid(
		c,
		"resetCompositeIndex",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) ResetExcludedPaths() {
	_jsii_.InvokeVoid(
		c,
		"resetExcludedPaths",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) ResetIncludedPaths() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludedPaths",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) ResetSpatialIndex() {
	_jsii_.InvokeVoid(
		c,
		"resetSpatialIndex",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CosmosdbGremlinGraphIndexPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package cdnfrontdoorroute

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/cdnfrontdoorroute/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CdnFrontdoorRouteCacheOutputReference interface {
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
	CompressionEnabled() interface{}
	SetCompressionEnabled(val interface{})
	CompressionEnabledInput() interface{}
	ContentTypesToCompress() *[]*string
	SetContentTypesToCompress(val *[]*string)
	ContentTypesToCompressInput() *[]*string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *CdnFrontdoorRouteCache
	SetInternalValue(val *CdnFrontdoorRouteCache)
	QueryStringCachingBehavior() *string
	SetQueryStringCachingBehavior(val *string)
	QueryStringCachingBehaviorInput() *string
	QueryStrings() *[]*string
	SetQueryStrings(val *[]*string)
	QueryStringsInput() *[]*string
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
	ResetCompressionEnabled()
	ResetContentTypesToCompress()
	ResetQueryStringCachingBehavior()
	ResetQueryStrings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnFrontdoorRouteCacheOutputReference
type jsiiProxy_CdnFrontdoorRouteCacheOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) CompressionEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressionEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) CompressionEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"compressionEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) ContentTypesToCompress() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"contentTypesToCompress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) ContentTypesToCompressInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"contentTypesToCompressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) InternalValue() *CdnFrontdoorRouteCache {
	var returns *CdnFrontdoorRouteCache
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) QueryStringCachingBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringCachingBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) QueryStringCachingBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"queryStringCachingBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) QueryStrings() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStrings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) QueryStringsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"queryStringsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCdnFrontdoorRouteCacheOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CdnFrontdoorRouteCacheOutputReference {
	_init_.Initialize()

	if err := validateNewCdnFrontdoorRouteCacheOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdnFrontdoorRouteCacheOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnFrontdoorRoute.CdnFrontdoorRouteCacheOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnFrontdoorRouteCacheOutputReference_Override(c CdnFrontdoorRouteCacheOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnFrontdoorRoute.CdnFrontdoorRouteCacheOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference)SetCompressionEnabled(val interface{}) {
	if err := j.validateSetCompressionEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"compressionEnabled",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference)SetContentTypesToCompress(val *[]*string) {
	if err := j.validateSetContentTypesToCompressParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"contentTypesToCompress",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference)SetInternalValue(val *CdnFrontdoorRouteCache) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference)SetQueryStringCachingBehavior(val *string) {
	if err := j.validateSetQueryStringCachingBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryStringCachingBehavior",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference)SetQueryStrings(val *[]*string) {
	if err := j.validateSetQueryStringsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"queryStrings",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRouteCacheOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) ResetCompressionEnabled() {
	_jsii_.InvokeVoid(
		c,
		"resetCompressionEnabled",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) ResetContentTypesToCompress() {
	_jsii_.InvokeVoid(
		c,
		"resetContentTypesToCompress",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) ResetQueryStringCachingBehavior() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryStringCachingBehavior",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) ResetQueryStrings() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryStrings",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CdnFrontdoorRouteCacheOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


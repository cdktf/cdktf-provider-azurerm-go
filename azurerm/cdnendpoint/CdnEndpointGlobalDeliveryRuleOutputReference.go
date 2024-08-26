// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/cdnendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CdnEndpointGlobalDeliveryRuleOutputReference interface {
	cdktf.ComplexObject
	CacheExpirationAction() CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference
	CacheExpirationActionInput() *CdnEndpointGlobalDeliveryRuleCacheExpirationAction
	CacheKeyQueryStringAction() CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference
	CacheKeyQueryStringActionInput() *CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringAction
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
	InternalValue() *CdnEndpointGlobalDeliveryRule
	SetInternalValue(val *CdnEndpointGlobalDeliveryRule)
	ModifyRequestHeaderAction() CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList
	ModifyRequestHeaderActionInput() interface{}
	ModifyResponseHeaderAction() CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList
	ModifyResponseHeaderActionInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UrlRedirectAction() CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference
	UrlRedirectActionInput() *CdnEndpointGlobalDeliveryRuleUrlRedirectAction
	UrlRewriteAction() CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference
	UrlRewriteActionInput() *CdnEndpointGlobalDeliveryRuleUrlRewriteAction
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
	PutCacheExpirationAction(value *CdnEndpointGlobalDeliveryRuleCacheExpirationAction)
	PutCacheKeyQueryStringAction(value *CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringAction)
	PutModifyRequestHeaderAction(value interface{})
	PutModifyResponseHeaderAction(value interface{})
	PutUrlRedirectAction(value *CdnEndpointGlobalDeliveryRuleUrlRedirectAction)
	PutUrlRewriteAction(value *CdnEndpointGlobalDeliveryRuleUrlRewriteAction)
	ResetCacheExpirationAction()
	ResetCacheKeyQueryStringAction()
	ResetModifyRequestHeaderAction()
	ResetModifyResponseHeaderAction()
	ResetUrlRedirectAction()
	ResetUrlRewriteAction()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnEndpointGlobalDeliveryRuleOutputReference
type jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) CacheExpirationAction() CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference {
	var returns CdnEndpointGlobalDeliveryRuleCacheExpirationActionOutputReference
	_jsii_.Get(
		j,
		"cacheExpirationAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) CacheExpirationActionInput() *CdnEndpointGlobalDeliveryRuleCacheExpirationAction {
	var returns *CdnEndpointGlobalDeliveryRuleCacheExpirationAction
	_jsii_.Get(
		j,
		"cacheExpirationActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) CacheKeyQueryStringAction() CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference {
	var returns CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringActionOutputReference
	_jsii_.Get(
		j,
		"cacheKeyQueryStringAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) CacheKeyQueryStringActionInput() *CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringAction {
	var returns *CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringAction
	_jsii_.Get(
		j,
		"cacheKeyQueryStringActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) InternalValue() *CdnEndpointGlobalDeliveryRule {
	var returns *CdnEndpointGlobalDeliveryRule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ModifyRequestHeaderAction() CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList {
	var returns CdnEndpointGlobalDeliveryRuleModifyRequestHeaderActionList
	_jsii_.Get(
		j,
		"modifyRequestHeaderAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ModifyRequestHeaderActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modifyRequestHeaderActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ModifyResponseHeaderAction() CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList {
	var returns CdnEndpointGlobalDeliveryRuleModifyResponseHeaderActionList
	_jsii_.Get(
		j,
		"modifyResponseHeaderAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ModifyResponseHeaderActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modifyResponseHeaderActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) UrlRedirectAction() CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference {
	var returns CdnEndpointGlobalDeliveryRuleUrlRedirectActionOutputReference
	_jsii_.Get(
		j,
		"urlRedirectAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) UrlRedirectActionInput() *CdnEndpointGlobalDeliveryRuleUrlRedirectAction {
	var returns *CdnEndpointGlobalDeliveryRuleUrlRedirectAction
	_jsii_.Get(
		j,
		"urlRedirectActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) UrlRewriteAction() CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference {
	var returns CdnEndpointGlobalDeliveryRuleUrlRewriteActionOutputReference
	_jsii_.Get(
		j,
		"urlRewriteAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) UrlRewriteActionInput() *CdnEndpointGlobalDeliveryRuleUrlRewriteAction {
	var returns *CdnEndpointGlobalDeliveryRuleUrlRewriteAction
	_jsii_.Get(
		j,
		"urlRewriteActionInput",
		&returns,
	)
	return returns
}


func NewCdnEndpointGlobalDeliveryRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CdnEndpointGlobalDeliveryRuleOutputReference {
	_init_.Initialize()

	if err := validateNewCdnEndpointGlobalDeliveryRuleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointGlobalDeliveryRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnEndpointGlobalDeliveryRuleOutputReference_Override(c CdnEndpointGlobalDeliveryRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointGlobalDeliveryRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference)SetInternalValue(val *CdnEndpointGlobalDeliveryRule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) PutCacheExpirationAction(value *CdnEndpointGlobalDeliveryRuleCacheExpirationAction) {
	if err := c.validatePutCacheExpirationActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCacheExpirationAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) PutCacheKeyQueryStringAction(value *CdnEndpointGlobalDeliveryRuleCacheKeyQueryStringAction) {
	if err := c.validatePutCacheKeyQueryStringActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCacheKeyQueryStringAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) PutModifyRequestHeaderAction(value interface{}) {
	if err := c.validatePutModifyRequestHeaderActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putModifyRequestHeaderAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) PutModifyResponseHeaderAction(value interface{}) {
	if err := c.validatePutModifyResponseHeaderActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putModifyResponseHeaderAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) PutUrlRedirectAction(value *CdnEndpointGlobalDeliveryRuleUrlRedirectAction) {
	if err := c.validatePutUrlRedirectActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUrlRedirectAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) PutUrlRewriteAction(value *CdnEndpointGlobalDeliveryRuleUrlRewriteAction) {
	if err := c.validatePutUrlRewriteActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUrlRewriteAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ResetCacheExpirationAction() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheExpirationAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ResetCacheKeyQueryStringAction() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheKeyQueryStringAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ResetModifyRequestHeaderAction() {
	_jsii_.InvokeVoid(
		c,
		"resetModifyRequestHeaderAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ResetModifyResponseHeaderAction() {
	_jsii_.InvokeVoid(
		c,
		"resetModifyResponseHeaderAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ResetUrlRedirectAction() {
	_jsii_.InvokeVoid(
		c,
		"resetUrlRedirectAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ResetUrlRewriteAction() {
	_jsii_.InvokeVoid(
		c,
		"resetUrlRewriteAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CdnEndpointGlobalDeliveryRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


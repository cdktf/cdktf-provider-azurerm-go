package cdnendpoint

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v6/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v6/cdnendpoint/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CdnEndpointDeliveryRuleOutputReference interface {
	cdktf.ComplexObject
	CacheExpirationAction() CdnEndpointDeliveryRuleCacheExpirationActionOutputReference
	CacheExpirationActionInput() *CdnEndpointDeliveryRuleCacheExpirationAction
	CacheKeyQueryStringAction() CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference
	CacheKeyQueryStringActionInput() *CdnEndpointDeliveryRuleCacheKeyQueryStringAction
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
	CookiesCondition() CdnEndpointDeliveryRuleCookiesConditionList
	CookiesConditionInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DeviceCondition() CdnEndpointDeliveryRuleDeviceConditionOutputReference
	DeviceConditionInput() *CdnEndpointDeliveryRuleDeviceCondition
	// Experimental.
	Fqn() *string
	HttpVersionCondition() CdnEndpointDeliveryRuleHttpVersionConditionList
	HttpVersionConditionInput() interface{}
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ModifyRequestHeaderAction() CdnEndpointDeliveryRuleModifyRequestHeaderActionList
	ModifyRequestHeaderActionInput() interface{}
	ModifyResponseHeaderAction() CdnEndpointDeliveryRuleModifyResponseHeaderActionList
	ModifyResponseHeaderActionInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	Order() *float64
	SetOrder(val *float64)
	OrderInput() *float64
	PostArgCondition() CdnEndpointDeliveryRulePostArgConditionList
	PostArgConditionInput() interface{}
	QueryStringCondition() CdnEndpointDeliveryRuleQueryStringConditionList
	QueryStringConditionInput() interface{}
	RemoteAddressCondition() CdnEndpointDeliveryRuleRemoteAddressConditionList
	RemoteAddressConditionInput() interface{}
	RequestBodyCondition() CdnEndpointDeliveryRuleRequestBodyConditionList
	RequestBodyConditionInput() interface{}
	RequestHeaderCondition() CdnEndpointDeliveryRuleRequestHeaderConditionList
	RequestHeaderConditionInput() interface{}
	RequestMethodCondition() CdnEndpointDeliveryRuleRequestMethodConditionOutputReference
	RequestMethodConditionInput() *CdnEndpointDeliveryRuleRequestMethodCondition
	RequestSchemeCondition() CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference
	RequestSchemeConditionInput() *CdnEndpointDeliveryRuleRequestSchemeCondition
	RequestUriCondition() CdnEndpointDeliveryRuleRequestUriConditionList
	RequestUriConditionInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UrlFileExtensionCondition() CdnEndpointDeliveryRuleUrlFileExtensionConditionList
	UrlFileExtensionConditionInput() interface{}
	UrlFileNameCondition() CdnEndpointDeliveryRuleUrlFileNameConditionList
	UrlFileNameConditionInput() interface{}
	UrlPathCondition() CdnEndpointDeliveryRuleUrlPathConditionList
	UrlPathConditionInput() interface{}
	UrlRedirectAction() CdnEndpointDeliveryRuleUrlRedirectActionOutputReference
	UrlRedirectActionInput() *CdnEndpointDeliveryRuleUrlRedirectAction
	UrlRewriteAction() CdnEndpointDeliveryRuleUrlRewriteActionOutputReference
	UrlRewriteActionInput() *CdnEndpointDeliveryRuleUrlRewriteAction
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
	PutCacheExpirationAction(value *CdnEndpointDeliveryRuleCacheExpirationAction)
	PutCacheKeyQueryStringAction(value *CdnEndpointDeliveryRuleCacheKeyQueryStringAction)
	PutCookiesCondition(value interface{})
	PutDeviceCondition(value *CdnEndpointDeliveryRuleDeviceCondition)
	PutHttpVersionCondition(value interface{})
	PutModifyRequestHeaderAction(value interface{})
	PutModifyResponseHeaderAction(value interface{})
	PutPostArgCondition(value interface{})
	PutQueryStringCondition(value interface{})
	PutRemoteAddressCondition(value interface{})
	PutRequestBodyCondition(value interface{})
	PutRequestHeaderCondition(value interface{})
	PutRequestMethodCondition(value *CdnEndpointDeliveryRuleRequestMethodCondition)
	PutRequestSchemeCondition(value *CdnEndpointDeliveryRuleRequestSchemeCondition)
	PutRequestUriCondition(value interface{})
	PutUrlFileExtensionCondition(value interface{})
	PutUrlFileNameCondition(value interface{})
	PutUrlPathCondition(value interface{})
	PutUrlRedirectAction(value *CdnEndpointDeliveryRuleUrlRedirectAction)
	PutUrlRewriteAction(value *CdnEndpointDeliveryRuleUrlRewriteAction)
	ResetCacheExpirationAction()
	ResetCacheKeyQueryStringAction()
	ResetCookiesCondition()
	ResetDeviceCondition()
	ResetHttpVersionCondition()
	ResetModifyRequestHeaderAction()
	ResetModifyResponseHeaderAction()
	ResetPostArgCondition()
	ResetQueryStringCondition()
	ResetRemoteAddressCondition()
	ResetRequestBodyCondition()
	ResetRequestHeaderCondition()
	ResetRequestMethodCondition()
	ResetRequestSchemeCondition()
	ResetRequestUriCondition()
	ResetUrlFileExtensionCondition()
	ResetUrlFileNameCondition()
	ResetUrlPathCondition()
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

// The jsii proxy struct for CdnEndpointDeliveryRuleOutputReference
type jsiiProxy_CdnEndpointDeliveryRuleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) CacheExpirationAction() CdnEndpointDeliveryRuleCacheExpirationActionOutputReference {
	var returns CdnEndpointDeliveryRuleCacheExpirationActionOutputReference
	_jsii_.Get(
		j,
		"cacheExpirationAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) CacheExpirationActionInput() *CdnEndpointDeliveryRuleCacheExpirationAction {
	var returns *CdnEndpointDeliveryRuleCacheExpirationAction
	_jsii_.Get(
		j,
		"cacheExpirationActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) CacheKeyQueryStringAction() CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference {
	var returns CdnEndpointDeliveryRuleCacheKeyQueryStringActionOutputReference
	_jsii_.Get(
		j,
		"cacheKeyQueryStringAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) CacheKeyQueryStringActionInput() *CdnEndpointDeliveryRuleCacheKeyQueryStringAction {
	var returns *CdnEndpointDeliveryRuleCacheKeyQueryStringAction
	_jsii_.Get(
		j,
		"cacheKeyQueryStringActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) CookiesCondition() CdnEndpointDeliveryRuleCookiesConditionList {
	var returns CdnEndpointDeliveryRuleCookiesConditionList
	_jsii_.Get(
		j,
		"cookiesCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) CookiesConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cookiesConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) DeviceCondition() CdnEndpointDeliveryRuleDeviceConditionOutputReference {
	var returns CdnEndpointDeliveryRuleDeviceConditionOutputReference
	_jsii_.Get(
		j,
		"deviceCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) DeviceConditionInput() *CdnEndpointDeliveryRuleDeviceCondition {
	var returns *CdnEndpointDeliveryRuleDeviceCondition
	_jsii_.Get(
		j,
		"deviceConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) HttpVersionCondition() CdnEndpointDeliveryRuleHttpVersionConditionList {
	var returns CdnEndpointDeliveryRuleHttpVersionConditionList
	_jsii_.Get(
		j,
		"httpVersionCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) HttpVersionConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpVersionConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ModifyRequestHeaderAction() CdnEndpointDeliveryRuleModifyRequestHeaderActionList {
	var returns CdnEndpointDeliveryRuleModifyRequestHeaderActionList
	_jsii_.Get(
		j,
		"modifyRequestHeaderAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ModifyRequestHeaderActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modifyRequestHeaderActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ModifyResponseHeaderAction() CdnEndpointDeliveryRuleModifyResponseHeaderActionList {
	var returns CdnEndpointDeliveryRuleModifyResponseHeaderActionList
	_jsii_.Get(
		j,
		"modifyResponseHeaderAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ModifyResponseHeaderActionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"modifyResponseHeaderActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) Order() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"order",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) OrderInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"orderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PostArgCondition() CdnEndpointDeliveryRulePostArgConditionList {
	var returns CdnEndpointDeliveryRulePostArgConditionList
	_jsii_.Get(
		j,
		"postArgCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PostArgConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"postArgConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) QueryStringCondition() CdnEndpointDeliveryRuleQueryStringConditionList {
	var returns CdnEndpointDeliveryRuleQueryStringConditionList
	_jsii_.Get(
		j,
		"queryStringCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) QueryStringConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) RemoteAddressCondition() CdnEndpointDeliveryRuleRemoteAddressConditionList {
	var returns CdnEndpointDeliveryRuleRemoteAddressConditionList
	_jsii_.Get(
		j,
		"remoteAddressCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) RemoteAddressConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remoteAddressConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) RequestBodyCondition() CdnEndpointDeliveryRuleRequestBodyConditionList {
	var returns CdnEndpointDeliveryRuleRequestBodyConditionList
	_jsii_.Get(
		j,
		"requestBodyCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) RequestBodyConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestBodyConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) RequestHeaderCondition() CdnEndpointDeliveryRuleRequestHeaderConditionList {
	var returns CdnEndpointDeliveryRuleRequestHeaderConditionList
	_jsii_.Get(
		j,
		"requestHeaderCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) RequestHeaderConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestHeaderConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) RequestMethodCondition() CdnEndpointDeliveryRuleRequestMethodConditionOutputReference {
	var returns CdnEndpointDeliveryRuleRequestMethodConditionOutputReference
	_jsii_.Get(
		j,
		"requestMethodCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) RequestMethodConditionInput() *CdnEndpointDeliveryRuleRequestMethodCondition {
	var returns *CdnEndpointDeliveryRuleRequestMethodCondition
	_jsii_.Get(
		j,
		"requestMethodConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) RequestSchemeCondition() CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference {
	var returns CdnEndpointDeliveryRuleRequestSchemeConditionOutputReference
	_jsii_.Get(
		j,
		"requestSchemeCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) RequestSchemeConditionInput() *CdnEndpointDeliveryRuleRequestSchemeCondition {
	var returns *CdnEndpointDeliveryRuleRequestSchemeCondition
	_jsii_.Get(
		j,
		"requestSchemeConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) RequestUriCondition() CdnEndpointDeliveryRuleRequestUriConditionList {
	var returns CdnEndpointDeliveryRuleRequestUriConditionList
	_jsii_.Get(
		j,
		"requestUriCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) RequestUriConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestUriConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) UrlFileExtensionCondition() CdnEndpointDeliveryRuleUrlFileExtensionConditionList {
	var returns CdnEndpointDeliveryRuleUrlFileExtensionConditionList
	_jsii_.Get(
		j,
		"urlFileExtensionCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) UrlFileExtensionConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlFileExtensionConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) UrlFileNameCondition() CdnEndpointDeliveryRuleUrlFileNameConditionList {
	var returns CdnEndpointDeliveryRuleUrlFileNameConditionList
	_jsii_.Get(
		j,
		"urlFileNameCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) UrlFileNameConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlFileNameConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) UrlPathCondition() CdnEndpointDeliveryRuleUrlPathConditionList {
	var returns CdnEndpointDeliveryRuleUrlPathConditionList
	_jsii_.Get(
		j,
		"urlPathCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) UrlPathConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlPathConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) UrlRedirectAction() CdnEndpointDeliveryRuleUrlRedirectActionOutputReference {
	var returns CdnEndpointDeliveryRuleUrlRedirectActionOutputReference
	_jsii_.Get(
		j,
		"urlRedirectAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) UrlRedirectActionInput() *CdnEndpointDeliveryRuleUrlRedirectAction {
	var returns *CdnEndpointDeliveryRuleUrlRedirectAction
	_jsii_.Get(
		j,
		"urlRedirectActionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) UrlRewriteAction() CdnEndpointDeliveryRuleUrlRewriteActionOutputReference {
	var returns CdnEndpointDeliveryRuleUrlRewriteActionOutputReference
	_jsii_.Get(
		j,
		"urlRewriteAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) UrlRewriteActionInput() *CdnEndpointDeliveryRuleUrlRewriteAction {
	var returns *CdnEndpointDeliveryRuleUrlRewriteAction
	_jsii_.Get(
		j,
		"urlRewriteActionInput",
		&returns,
	)
	return returns
}


func NewCdnEndpointDeliveryRuleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CdnEndpointDeliveryRuleOutputReference {
	_init_.Initialize()

	if err := validateNewCdnEndpointDeliveryRuleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdnEndpointDeliveryRuleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCdnEndpointDeliveryRuleOutputReference_Override(c CdnEndpointDeliveryRuleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnEndpoint.CdnEndpointDeliveryRuleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference)SetOrder(val *float64) {
	if err := j.validateSetOrderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"order",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnEndpointDeliveryRuleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutCacheExpirationAction(value *CdnEndpointDeliveryRuleCacheExpirationAction) {
	if err := c.validatePutCacheExpirationActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCacheExpirationAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutCacheKeyQueryStringAction(value *CdnEndpointDeliveryRuleCacheKeyQueryStringAction) {
	if err := c.validatePutCacheKeyQueryStringActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCacheKeyQueryStringAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutCookiesCondition(value interface{}) {
	if err := c.validatePutCookiesConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCookiesCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutDeviceCondition(value *CdnEndpointDeliveryRuleDeviceCondition) {
	if err := c.validatePutDeviceConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDeviceCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutHttpVersionCondition(value interface{}) {
	if err := c.validatePutHttpVersionConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttpVersionCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutModifyRequestHeaderAction(value interface{}) {
	if err := c.validatePutModifyRequestHeaderActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putModifyRequestHeaderAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutModifyResponseHeaderAction(value interface{}) {
	if err := c.validatePutModifyResponseHeaderActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putModifyResponseHeaderAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutPostArgCondition(value interface{}) {
	if err := c.validatePutPostArgConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPostArgCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutQueryStringCondition(value interface{}) {
	if err := c.validatePutQueryStringConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putQueryStringCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutRemoteAddressCondition(value interface{}) {
	if err := c.validatePutRemoteAddressConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRemoteAddressCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutRequestBodyCondition(value interface{}) {
	if err := c.validatePutRequestBodyConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestBodyCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutRequestHeaderCondition(value interface{}) {
	if err := c.validatePutRequestHeaderConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestHeaderCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutRequestMethodCondition(value *CdnEndpointDeliveryRuleRequestMethodCondition) {
	if err := c.validatePutRequestMethodConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestMethodCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutRequestSchemeCondition(value *CdnEndpointDeliveryRuleRequestSchemeCondition) {
	if err := c.validatePutRequestSchemeConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestSchemeCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutRequestUriCondition(value interface{}) {
	if err := c.validatePutRequestUriConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestUriCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutUrlFileExtensionCondition(value interface{}) {
	if err := c.validatePutUrlFileExtensionConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUrlFileExtensionCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutUrlFileNameCondition(value interface{}) {
	if err := c.validatePutUrlFileNameConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUrlFileNameCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutUrlPathCondition(value interface{}) {
	if err := c.validatePutUrlPathConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUrlPathCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutUrlRedirectAction(value *CdnEndpointDeliveryRuleUrlRedirectAction) {
	if err := c.validatePutUrlRedirectActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUrlRedirectAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) PutUrlRewriteAction(value *CdnEndpointDeliveryRuleUrlRewriteAction) {
	if err := c.validatePutUrlRewriteActionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUrlRewriteAction",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetCacheExpirationAction() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheExpirationAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetCacheKeyQueryStringAction() {
	_jsii_.InvokeVoid(
		c,
		"resetCacheKeyQueryStringAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetCookiesCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetCookiesCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetDeviceCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetDeviceCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetHttpVersionCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpVersionCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetModifyRequestHeaderAction() {
	_jsii_.InvokeVoid(
		c,
		"resetModifyRequestHeaderAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetModifyResponseHeaderAction() {
	_jsii_.InvokeVoid(
		c,
		"resetModifyResponseHeaderAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetPostArgCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetPostArgCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetQueryStringCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryStringCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetRemoteAddressCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetRemoteAddressCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetRequestBodyCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestBodyCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetRequestHeaderCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestHeaderCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetRequestMethodCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestMethodCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetRequestSchemeCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestSchemeCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetRequestUriCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestUriCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetUrlFileExtensionCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetUrlFileExtensionCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetUrlFileNameCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetUrlFileNameCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetUrlPathCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetUrlPathCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetUrlRedirectAction() {
	_jsii_.InvokeVoid(
		c,
		"resetUrlRedirectAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ResetUrlRewriteAction() {
	_jsii_.InvokeVoid(
		c,
		"resetUrlRewriteAction",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CdnEndpointDeliveryRuleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package cdnfrontdoorrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/cdnfrontdoorrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type CdnFrontdoorRuleConditionsOutputReference interface {
	cdktf.ComplexObject
	ClientPortCondition() CdnFrontdoorRuleConditionsClientPortConditionList
	ClientPortConditionInput() interface{}
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
	CookiesCondition() CdnFrontdoorRuleConditionsCookiesConditionList
	CookiesConditionInput() interface{}
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	HostNameCondition() CdnFrontdoorRuleConditionsHostNameConditionList
	HostNameConditionInput() interface{}
	HttpVersionCondition() CdnFrontdoorRuleConditionsHttpVersionConditionList
	HttpVersionConditionInput() interface{}
	InternalValue() *CdnFrontdoorRuleConditions
	SetInternalValue(val *CdnFrontdoorRuleConditions)
	IsDeviceCondition() CdnFrontdoorRuleConditionsIsDeviceConditionList
	IsDeviceConditionInput() interface{}
	PostArgsCondition() CdnFrontdoorRuleConditionsPostArgsConditionList
	PostArgsConditionInput() interface{}
	QueryStringCondition() CdnFrontdoorRuleConditionsQueryStringConditionList
	QueryStringConditionInput() interface{}
	RemoteAddressCondition() CdnFrontdoorRuleConditionsRemoteAddressConditionList
	RemoteAddressConditionInput() interface{}
	RequestBodyCondition() CdnFrontdoorRuleConditionsRequestBodyConditionList
	RequestBodyConditionInput() interface{}
	RequestHeaderCondition() CdnFrontdoorRuleConditionsRequestHeaderConditionList
	RequestHeaderConditionInput() interface{}
	RequestMethodCondition() CdnFrontdoorRuleConditionsRequestMethodConditionList
	RequestMethodConditionInput() interface{}
	RequestSchemeCondition() CdnFrontdoorRuleConditionsRequestSchemeConditionList
	RequestSchemeConditionInput() interface{}
	RequestUriCondition() CdnFrontdoorRuleConditionsRequestUriConditionList
	RequestUriConditionInput() interface{}
	ServerPortCondition() CdnFrontdoorRuleConditionsServerPortConditionList
	ServerPortConditionInput() interface{}
	SocketAddressCondition() CdnFrontdoorRuleConditionsSocketAddressConditionList
	SocketAddressConditionInput() interface{}
	SslProtocolCondition() CdnFrontdoorRuleConditionsSslProtocolConditionList
	SslProtocolConditionInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UrlFileExtensionCondition() CdnFrontdoorRuleConditionsUrlFileExtensionConditionList
	UrlFileExtensionConditionInput() interface{}
	UrlFilenameCondition() CdnFrontdoorRuleConditionsUrlFilenameConditionList
	UrlFilenameConditionInput() interface{}
	UrlPathCondition() CdnFrontdoorRuleConditionsUrlPathConditionList
	UrlPathConditionInput() interface{}
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
	PutClientPortCondition(value interface{})
	PutCookiesCondition(value interface{})
	PutHostNameCondition(value interface{})
	PutHttpVersionCondition(value interface{})
	PutIsDeviceCondition(value interface{})
	PutPostArgsCondition(value interface{})
	PutQueryStringCondition(value interface{})
	PutRemoteAddressCondition(value interface{})
	PutRequestBodyCondition(value interface{})
	PutRequestHeaderCondition(value interface{})
	PutRequestMethodCondition(value interface{})
	PutRequestSchemeCondition(value interface{})
	PutRequestUriCondition(value interface{})
	PutServerPortCondition(value interface{})
	PutSocketAddressCondition(value interface{})
	PutSslProtocolCondition(value interface{})
	PutUrlFileExtensionCondition(value interface{})
	PutUrlFilenameCondition(value interface{})
	PutUrlPathCondition(value interface{})
	ResetClientPortCondition()
	ResetCookiesCondition()
	ResetHostNameCondition()
	ResetHttpVersionCondition()
	ResetIsDeviceCondition()
	ResetPostArgsCondition()
	ResetQueryStringCondition()
	ResetRemoteAddressCondition()
	ResetRequestBodyCondition()
	ResetRequestHeaderCondition()
	ResetRequestMethodCondition()
	ResetRequestSchemeCondition()
	ResetRequestUriCondition()
	ResetServerPortCondition()
	ResetSocketAddressCondition()
	ResetSslProtocolCondition()
	ResetUrlFileExtensionCondition()
	ResetUrlFilenameCondition()
	ResetUrlPathCondition()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CdnFrontdoorRuleConditionsOutputReference
type jsiiProxy_CdnFrontdoorRuleConditionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ClientPortCondition() CdnFrontdoorRuleConditionsClientPortConditionList {
	var returns CdnFrontdoorRuleConditionsClientPortConditionList
	_jsii_.Get(
		j,
		"clientPortCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ClientPortConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientPortConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) CookiesCondition() CdnFrontdoorRuleConditionsCookiesConditionList {
	var returns CdnFrontdoorRuleConditionsCookiesConditionList
	_jsii_.Get(
		j,
		"cookiesCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) CookiesConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cookiesConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) HostNameCondition() CdnFrontdoorRuleConditionsHostNameConditionList {
	var returns CdnFrontdoorRuleConditionsHostNameConditionList
	_jsii_.Get(
		j,
		"hostNameCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) HostNameConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hostNameConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) HttpVersionCondition() CdnFrontdoorRuleConditionsHttpVersionConditionList {
	var returns CdnFrontdoorRuleConditionsHttpVersionConditionList
	_jsii_.Get(
		j,
		"httpVersionCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) HttpVersionConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"httpVersionConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) InternalValue() *CdnFrontdoorRuleConditions {
	var returns *CdnFrontdoorRuleConditions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) IsDeviceCondition() CdnFrontdoorRuleConditionsIsDeviceConditionList {
	var returns CdnFrontdoorRuleConditionsIsDeviceConditionList
	_jsii_.Get(
		j,
		"isDeviceCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) IsDeviceConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDeviceConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PostArgsCondition() CdnFrontdoorRuleConditionsPostArgsConditionList {
	var returns CdnFrontdoorRuleConditionsPostArgsConditionList
	_jsii_.Get(
		j,
		"postArgsCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PostArgsConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"postArgsConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) QueryStringCondition() CdnFrontdoorRuleConditionsQueryStringConditionList {
	var returns CdnFrontdoorRuleConditionsQueryStringConditionList
	_jsii_.Get(
		j,
		"queryStringCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) QueryStringConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"queryStringConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RemoteAddressCondition() CdnFrontdoorRuleConditionsRemoteAddressConditionList {
	var returns CdnFrontdoorRuleConditionsRemoteAddressConditionList
	_jsii_.Get(
		j,
		"remoteAddressCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RemoteAddressConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"remoteAddressConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestBodyCondition() CdnFrontdoorRuleConditionsRequestBodyConditionList {
	var returns CdnFrontdoorRuleConditionsRequestBodyConditionList
	_jsii_.Get(
		j,
		"requestBodyCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestBodyConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestBodyConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestHeaderCondition() CdnFrontdoorRuleConditionsRequestHeaderConditionList {
	var returns CdnFrontdoorRuleConditionsRequestHeaderConditionList
	_jsii_.Get(
		j,
		"requestHeaderCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestHeaderConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestHeaderConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestMethodCondition() CdnFrontdoorRuleConditionsRequestMethodConditionList {
	var returns CdnFrontdoorRuleConditionsRequestMethodConditionList
	_jsii_.Get(
		j,
		"requestMethodCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestMethodConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestMethodConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestSchemeCondition() CdnFrontdoorRuleConditionsRequestSchemeConditionList {
	var returns CdnFrontdoorRuleConditionsRequestSchemeConditionList
	_jsii_.Get(
		j,
		"requestSchemeCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestSchemeConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestSchemeConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestUriCondition() CdnFrontdoorRuleConditionsRequestUriConditionList {
	var returns CdnFrontdoorRuleConditionsRequestUriConditionList
	_jsii_.Get(
		j,
		"requestUriCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) RequestUriConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"requestUriConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ServerPortCondition() CdnFrontdoorRuleConditionsServerPortConditionList {
	var returns CdnFrontdoorRuleConditionsServerPortConditionList
	_jsii_.Get(
		j,
		"serverPortCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ServerPortConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"serverPortConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) SocketAddressCondition() CdnFrontdoorRuleConditionsSocketAddressConditionList {
	var returns CdnFrontdoorRuleConditionsSocketAddressConditionList
	_jsii_.Get(
		j,
		"socketAddressCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) SocketAddressConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"socketAddressConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) SslProtocolCondition() CdnFrontdoorRuleConditionsSslProtocolConditionList {
	var returns CdnFrontdoorRuleConditionsSslProtocolConditionList
	_jsii_.Get(
		j,
		"sslProtocolCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) SslProtocolConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sslProtocolConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) UrlFileExtensionCondition() CdnFrontdoorRuleConditionsUrlFileExtensionConditionList {
	var returns CdnFrontdoorRuleConditionsUrlFileExtensionConditionList
	_jsii_.Get(
		j,
		"urlFileExtensionCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) UrlFileExtensionConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlFileExtensionConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) UrlFilenameCondition() CdnFrontdoorRuleConditionsUrlFilenameConditionList {
	var returns CdnFrontdoorRuleConditionsUrlFilenameConditionList
	_jsii_.Get(
		j,
		"urlFilenameCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) UrlFilenameConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlFilenameConditionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) UrlPathCondition() CdnFrontdoorRuleConditionsUrlPathConditionList {
	var returns CdnFrontdoorRuleConditionsUrlPathConditionList
	_jsii_.Get(
		j,
		"urlPathCondition",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) UrlPathConditionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"urlPathConditionInput",
		&returns,
	)
	return returns
}


func NewCdnFrontdoorRuleConditionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CdnFrontdoorRuleConditionsOutputReference {
	_init_.Initialize()

	if err := validateNewCdnFrontdoorRuleConditionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CdnFrontdoorRuleConditionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnFrontdoorRule.CdnFrontdoorRuleConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCdnFrontdoorRuleConditionsOutputReference_Override(c CdnFrontdoorRuleConditionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.cdnFrontdoorRule.CdnFrontdoorRuleConditionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference)SetInternalValue(val *CdnFrontdoorRuleConditions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutClientPortCondition(value interface{}) {
	if err := c.validatePutClientPortConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putClientPortCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutCookiesCondition(value interface{}) {
	if err := c.validatePutCookiesConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putCookiesCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutHostNameCondition(value interface{}) {
	if err := c.validatePutHostNameConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHostNameCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutHttpVersionCondition(value interface{}) {
	if err := c.validatePutHttpVersionConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putHttpVersionCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutIsDeviceCondition(value interface{}) {
	if err := c.validatePutIsDeviceConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putIsDeviceCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutPostArgsCondition(value interface{}) {
	if err := c.validatePutPostArgsConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPostArgsCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutQueryStringCondition(value interface{}) {
	if err := c.validatePutQueryStringConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putQueryStringCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutRemoteAddressCondition(value interface{}) {
	if err := c.validatePutRemoteAddressConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRemoteAddressCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutRequestBodyCondition(value interface{}) {
	if err := c.validatePutRequestBodyConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestBodyCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutRequestHeaderCondition(value interface{}) {
	if err := c.validatePutRequestHeaderConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestHeaderCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutRequestMethodCondition(value interface{}) {
	if err := c.validatePutRequestMethodConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestMethodCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutRequestSchemeCondition(value interface{}) {
	if err := c.validatePutRequestSchemeConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestSchemeCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutRequestUriCondition(value interface{}) {
	if err := c.validatePutRequestUriConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRequestUriCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutServerPortCondition(value interface{}) {
	if err := c.validatePutServerPortConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putServerPortCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutSocketAddressCondition(value interface{}) {
	if err := c.validatePutSocketAddressConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSocketAddressCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutSslProtocolCondition(value interface{}) {
	if err := c.validatePutSslProtocolConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSslProtocolCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutUrlFileExtensionCondition(value interface{}) {
	if err := c.validatePutUrlFileExtensionConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUrlFileExtensionCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutUrlFilenameCondition(value interface{}) {
	if err := c.validatePutUrlFilenameConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUrlFilenameCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) PutUrlPathCondition(value interface{}) {
	if err := c.validatePutUrlPathConditionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putUrlPathCondition",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetClientPortCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetClientPortCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetCookiesCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetCookiesCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetHostNameCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetHostNameCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetHttpVersionCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetHttpVersionCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetIsDeviceCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetIsDeviceCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetPostArgsCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetPostArgsCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetQueryStringCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetQueryStringCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetRemoteAddressCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetRemoteAddressCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetRequestBodyCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestBodyCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetRequestHeaderCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestHeaderCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetRequestMethodCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestMethodCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetRequestSchemeCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestSchemeCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetRequestUriCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetRequestUriCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetServerPortCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetServerPortCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetSocketAddressCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetSocketAddressCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetSslProtocolCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetSslProtocolCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetUrlFileExtensionCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetUrlFileExtensionCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetUrlFilenameCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetUrlFilenameCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ResetUrlPathCondition() {
	_jsii_.InvokeVoid(
		c,
		"resetUrlPathCondition",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (c *jsiiProxy_CdnFrontdoorRuleConditionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


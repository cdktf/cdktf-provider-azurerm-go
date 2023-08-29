// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package functionapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/functionapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type FunctionAppSiteConfigOutputReference interface {
	cdktf.ComplexObject
	AlwaysOn() interface{}
	SetAlwaysOn(val interface{})
	AlwaysOnInput() interface{}
	AppScaleLimit() *float64
	SetAppScaleLimit(val *float64)
	AppScaleLimitInput() *float64
	AutoSwapSlotName() *string
	SetAutoSwapSlotName(val *string)
	AutoSwapSlotNameInput() *string
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
	Cors() FunctionAppSiteConfigCorsOutputReference
	CorsInput() *FunctionAppSiteConfigCors
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DotnetFrameworkVersion() *string
	SetDotnetFrameworkVersion(val *string)
	DotnetFrameworkVersionInput() *string
	ElasticInstanceMinimum() *float64
	SetElasticInstanceMinimum(val *float64)
	ElasticInstanceMinimumInput() *float64
	// Experimental.
	Fqn() *string
	FtpsState() *string
	SetFtpsState(val *string)
	FtpsStateInput() *string
	HealthCheckPath() *string
	SetHealthCheckPath(val *string)
	HealthCheckPathInput() *string
	Http2Enabled() interface{}
	SetHttp2Enabled(val interface{})
	Http2EnabledInput() interface{}
	InternalValue() *FunctionAppSiteConfig
	SetInternalValue(val *FunctionAppSiteConfig)
	IpRestriction() FunctionAppSiteConfigIpRestrictionList
	IpRestrictionInput() interface{}
	JavaVersion() *string
	SetJavaVersion(val *string)
	JavaVersionInput() *string
	LinuxFxVersion() *string
	SetLinuxFxVersion(val *string)
	LinuxFxVersionInput() *string
	MinTlsVersion() *string
	SetMinTlsVersion(val *string)
	MinTlsVersionInput() *string
	PreWarmedInstanceCount() *float64
	SetPreWarmedInstanceCount(val *float64)
	PreWarmedInstanceCountInput() *float64
	RuntimeScaleMonitoringEnabled() interface{}
	SetRuntimeScaleMonitoringEnabled(val interface{})
	RuntimeScaleMonitoringEnabledInput() interface{}
	ScmIpRestriction() FunctionAppSiteConfigScmIpRestrictionList
	ScmIpRestrictionInput() interface{}
	ScmType() *string
	SetScmType(val *string)
	ScmTypeInput() *string
	ScmUseMainIpRestriction() interface{}
	SetScmUseMainIpRestriction(val interface{})
	ScmUseMainIpRestrictionInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Use32BitWorkerProcess() interface{}
	SetUse32BitWorkerProcess(val interface{})
	Use32BitWorkerProcessInput() interface{}
	VnetRouteAllEnabled() interface{}
	SetVnetRouteAllEnabled(val interface{})
	VnetRouteAllEnabledInput() interface{}
	WebsocketsEnabled() interface{}
	SetWebsocketsEnabled(val interface{})
	WebsocketsEnabledInput() interface{}
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
	PutCors(value *FunctionAppSiteConfigCors)
	PutIpRestriction(value interface{})
	PutScmIpRestriction(value interface{})
	ResetAlwaysOn()
	ResetAppScaleLimit()
	ResetAutoSwapSlotName()
	ResetCors()
	ResetDotnetFrameworkVersion()
	ResetElasticInstanceMinimum()
	ResetFtpsState()
	ResetHealthCheckPath()
	ResetHttp2Enabled()
	ResetIpRestriction()
	ResetJavaVersion()
	ResetLinuxFxVersion()
	ResetMinTlsVersion()
	ResetPreWarmedInstanceCount()
	ResetRuntimeScaleMonitoringEnabled()
	ResetScmIpRestriction()
	ResetScmType()
	ResetScmUseMainIpRestriction()
	ResetUse32BitWorkerProcess()
	ResetVnetRouteAllEnabled()
	ResetWebsocketsEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for FunctionAppSiteConfigOutputReference
type jsiiProxy_FunctionAppSiteConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) AlwaysOn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) AlwaysOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) AppScaleLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"appScaleLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) AppScaleLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"appScaleLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) AutoSwapSlotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoSwapSlotName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) AutoSwapSlotNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoSwapSlotNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) Cors() FunctionAppSiteConfigCorsOutputReference {
	var returns FunctionAppSiteConfigCorsOutputReference
	_jsii_.Get(
		j,
		"cors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) CorsInput() *FunctionAppSiteConfigCors {
	var returns *FunctionAppSiteConfigCors
	_jsii_.Get(
		j,
		"corsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) DotnetFrameworkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dotnetFrameworkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) DotnetFrameworkVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dotnetFrameworkVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) ElasticInstanceMinimum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"elasticInstanceMinimum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) ElasticInstanceMinimumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"elasticInstanceMinimumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) FtpsState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ftpsState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) FtpsStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ftpsStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) HealthCheckPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) HealthCheckPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) Http2Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) Http2EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) InternalValue() *FunctionAppSiteConfig {
	var returns *FunctionAppSiteConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) IpRestriction() FunctionAppSiteConfigIpRestrictionList {
	var returns FunctionAppSiteConfigIpRestrictionList
	_jsii_.Get(
		j,
		"ipRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) IpRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) JavaVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) JavaVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"javaVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) LinuxFxVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linuxFxVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) LinuxFxVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linuxFxVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) MinTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) MinTlsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTlsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) PreWarmedInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preWarmedInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) PreWarmedInstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preWarmedInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) RuntimeScaleMonitoringEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runtimeScaleMonitoringEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) RuntimeScaleMonitoringEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runtimeScaleMonitoringEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) ScmIpRestriction() FunctionAppSiteConfigScmIpRestrictionList {
	var returns FunctionAppSiteConfigScmIpRestrictionList
	_jsii_.Get(
		j,
		"scmIpRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) ScmIpRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scmIpRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) ScmType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) ScmTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) ScmUseMainIpRestriction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scmUseMainIpRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) ScmUseMainIpRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scmUseMainIpRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) Use32BitWorkerProcess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"use32BitWorkerProcess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) Use32BitWorkerProcessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"use32BitWorkerProcessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) VnetRouteAllEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vnetRouteAllEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) VnetRouteAllEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vnetRouteAllEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) WebsocketsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"websocketsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference) WebsocketsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"websocketsEnabledInput",
		&returns,
	)
	return returns
}


func NewFunctionAppSiteConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) FunctionAppSiteConfigOutputReference {
	_init_.Initialize()

	if err := validateNewFunctionAppSiteConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_FunctionAppSiteConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.functionApp.FunctionAppSiteConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewFunctionAppSiteConfigOutputReference_Override(f FunctionAppSiteConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.functionApp.FunctionAppSiteConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		f,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference)SetAlwaysOn(val interface{}) {
	if err := j.validateSetAlwaysOnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alwaysOn",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference)SetAppScaleLimit(val *float64) {
	if err := j.validateSetAppScaleLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appScaleLimit",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference)SetAutoSwapSlotName(val *string) {
	if err := j.validateSetAutoSwapSlotNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoSwapSlotName",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference)SetDotnetFrameworkVersion(val *string) {
	if err := j.validateSetDotnetFrameworkVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dotnetFrameworkVersion",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference)SetElasticInstanceMinimum(val *float64) {
	if err := j.validateSetElasticInstanceMinimumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticInstanceMinimum",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference)SetFtpsState(val *string) {
	if err := j.validateSetFtpsStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ftpsState",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference)SetHealthCheckPath(val *string) {
	if err := j.validateSetHealthCheckPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckPath",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference)SetHttp2Enabled(val interface{}) {
	if err := j.validateSetHttp2EnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"http2Enabled",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference)SetInternalValue(val *FunctionAppSiteConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference)SetJavaVersion(val *string) {
	if err := j.validateSetJavaVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"javaVersion",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference)SetLinuxFxVersion(val *string) {
	if err := j.validateSetLinuxFxVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linuxFxVersion",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference)SetMinTlsVersion(val *string) {
	if err := j.validateSetMinTlsVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTlsVersion",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference)SetPreWarmedInstanceCount(val *float64) {
	if err := j.validateSetPreWarmedInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preWarmedInstanceCount",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference)SetRuntimeScaleMonitoringEnabled(val interface{}) {
	if err := j.validateSetRuntimeScaleMonitoringEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeScaleMonitoringEnabled",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference)SetScmType(val *string) {
	if err := j.validateSetScmTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scmType",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference)SetScmUseMainIpRestriction(val interface{}) {
	if err := j.validateSetScmUseMainIpRestrictionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scmUseMainIpRestriction",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference)SetUse32BitWorkerProcess(val interface{}) {
	if err := j.validateSetUse32BitWorkerProcessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"use32BitWorkerProcess",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference)SetVnetRouteAllEnabled(val interface{}) {
	if err := j.validateSetVnetRouteAllEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vnetRouteAllEnabled",
		val,
	)
}

func (j *jsiiProxy_FunctionAppSiteConfigOutputReference)SetWebsocketsEnabled(val interface{}) {
	if err := j.validateSetWebsocketsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"websocketsEnabled",
		val,
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := f.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		f,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := f.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := f.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		f,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := f.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		f,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := f.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		f,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := f.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		f,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := f.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		f,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := f.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		f,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := f.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		f,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := f.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		f,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) PutCors(value *FunctionAppSiteConfigCors) {
	if err := f.validatePutCorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putCors",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) PutIpRestriction(value interface{}) {
	if err := f.validatePutIpRestrictionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putIpRestriction",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) PutScmIpRestriction(value interface{}) {
	if err := f.validatePutScmIpRestrictionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		f,
		"putScmIpRestriction",
		[]interface{}{value},
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) ResetAlwaysOn() {
	_jsii_.InvokeVoid(
		f,
		"resetAlwaysOn",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) ResetAppScaleLimit() {
	_jsii_.InvokeVoid(
		f,
		"resetAppScaleLimit",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) ResetAutoSwapSlotName() {
	_jsii_.InvokeVoid(
		f,
		"resetAutoSwapSlotName",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) ResetCors() {
	_jsii_.InvokeVoid(
		f,
		"resetCors",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) ResetDotnetFrameworkVersion() {
	_jsii_.InvokeVoid(
		f,
		"resetDotnetFrameworkVersion",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) ResetElasticInstanceMinimum() {
	_jsii_.InvokeVoid(
		f,
		"resetElasticInstanceMinimum",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) ResetFtpsState() {
	_jsii_.InvokeVoid(
		f,
		"resetFtpsState",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) ResetHealthCheckPath() {
	_jsii_.InvokeVoid(
		f,
		"resetHealthCheckPath",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) ResetHttp2Enabled() {
	_jsii_.InvokeVoid(
		f,
		"resetHttp2Enabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) ResetIpRestriction() {
	_jsii_.InvokeVoid(
		f,
		"resetIpRestriction",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) ResetJavaVersion() {
	_jsii_.InvokeVoid(
		f,
		"resetJavaVersion",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) ResetLinuxFxVersion() {
	_jsii_.InvokeVoid(
		f,
		"resetLinuxFxVersion",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) ResetMinTlsVersion() {
	_jsii_.InvokeVoid(
		f,
		"resetMinTlsVersion",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) ResetPreWarmedInstanceCount() {
	_jsii_.InvokeVoid(
		f,
		"resetPreWarmedInstanceCount",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) ResetRuntimeScaleMonitoringEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetRuntimeScaleMonitoringEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) ResetScmIpRestriction() {
	_jsii_.InvokeVoid(
		f,
		"resetScmIpRestriction",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) ResetScmType() {
	_jsii_.InvokeVoid(
		f,
		"resetScmType",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) ResetScmUseMainIpRestriction() {
	_jsii_.InvokeVoid(
		f,
		"resetScmUseMainIpRestriction",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) ResetUse32BitWorkerProcess() {
	_jsii_.InvokeVoid(
		f,
		"resetUse32BitWorkerProcess",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) ResetVnetRouteAllEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetVnetRouteAllEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) ResetWebsocketsEnabled() {
	_jsii_.InvokeVoid(
		f,
		"resetWebsocketsEnabled",
		nil, // no parameters
	)
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := f.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		f,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (f *jsiiProxy_FunctionAppSiteConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		f,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


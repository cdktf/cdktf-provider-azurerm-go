// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package logicappstandard

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/logicappstandard/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LogicAppStandardSiteConfigOutputReference interface {
	cdktf.ComplexObject
	AlwaysOn() interface{}
	SetAlwaysOn(val interface{})
	AlwaysOnInput() interface{}
	AppScaleLimit() *float64
	SetAppScaleLimit(val *float64)
	AppScaleLimitInput() *float64
	AutoSwapSlotName() *string
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
	Cors() LogicAppStandardSiteConfigCorsOutputReference
	CorsInput() *LogicAppStandardSiteConfigCors
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
	InternalValue() *LogicAppStandardSiteConfig
	SetInternalValue(val *LogicAppStandardSiteConfig)
	IpRestriction() LogicAppStandardSiteConfigIpRestrictionList
	IpRestrictionInput() interface{}
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
	ScmIpRestriction() LogicAppStandardSiteConfigScmIpRestrictionList
	ScmIpRestrictionInput() interface{}
	ScmMinTlsVersion() *string
	SetScmMinTlsVersion(val *string)
	ScmMinTlsVersionInput() *string
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
	PutCors(value *LogicAppStandardSiteConfigCors)
	PutIpRestriction(value interface{})
	PutScmIpRestriction(value interface{})
	ResetAlwaysOn()
	ResetAppScaleLimit()
	ResetCors()
	ResetDotnetFrameworkVersion()
	ResetElasticInstanceMinimum()
	ResetFtpsState()
	ResetHealthCheckPath()
	ResetHttp2Enabled()
	ResetIpRestriction()
	ResetLinuxFxVersion()
	ResetMinTlsVersion()
	ResetPreWarmedInstanceCount()
	ResetRuntimeScaleMonitoringEnabled()
	ResetScmIpRestriction()
	ResetScmMinTlsVersion()
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

// The jsii proxy struct for LogicAppStandardSiteConfigOutputReference
type jsiiProxy_LogicAppStandardSiteConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) AlwaysOn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) AlwaysOnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"alwaysOnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) AppScaleLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"appScaleLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) AppScaleLimitInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"appScaleLimitInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) AutoSwapSlotName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoSwapSlotName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) Cors() LogicAppStandardSiteConfigCorsOutputReference {
	var returns LogicAppStandardSiteConfigCorsOutputReference
	_jsii_.Get(
		j,
		"cors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) CorsInput() *LogicAppStandardSiteConfigCors {
	var returns *LogicAppStandardSiteConfigCors
	_jsii_.Get(
		j,
		"corsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) DotnetFrameworkVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dotnetFrameworkVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) DotnetFrameworkVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dotnetFrameworkVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ElasticInstanceMinimum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"elasticInstanceMinimum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ElasticInstanceMinimumInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"elasticInstanceMinimumInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) FtpsState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ftpsState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) FtpsStateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ftpsStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) HealthCheckPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) HealthCheckPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) Http2Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) Http2EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"http2EnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) InternalValue() *LogicAppStandardSiteConfig {
	var returns *LogicAppStandardSiteConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) IpRestriction() LogicAppStandardSiteConfigIpRestrictionList {
	var returns LogicAppStandardSiteConfigIpRestrictionList
	_jsii_.Get(
		j,
		"ipRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) IpRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ipRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) LinuxFxVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linuxFxVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) LinuxFxVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"linuxFxVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) MinTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) MinTlsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minTlsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) PreWarmedInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preWarmedInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) PreWarmedInstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preWarmedInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) RuntimeScaleMonitoringEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runtimeScaleMonitoringEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) RuntimeScaleMonitoringEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"runtimeScaleMonitoringEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ScmIpRestriction() LogicAppStandardSiteConfigScmIpRestrictionList {
	var returns LogicAppStandardSiteConfigScmIpRestrictionList
	_jsii_.Get(
		j,
		"scmIpRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ScmIpRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scmIpRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ScmMinTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmMinTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ScmMinTlsVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmMinTlsVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ScmType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ScmTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ScmUseMainIpRestriction() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scmUseMainIpRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ScmUseMainIpRestrictionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"scmUseMainIpRestrictionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) Use32BitWorkerProcess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"use32BitWorkerProcess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) Use32BitWorkerProcessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"use32BitWorkerProcessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) VnetRouteAllEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vnetRouteAllEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) VnetRouteAllEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vnetRouteAllEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) WebsocketsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"websocketsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference) WebsocketsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"websocketsEnabledInput",
		&returns,
	)
	return returns
}


func NewLogicAppStandardSiteConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LogicAppStandardSiteConfigOutputReference {
	_init_.Initialize()

	if err := validateNewLogicAppStandardSiteConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LogicAppStandardSiteConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.logicAppStandard.LogicAppStandardSiteConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLogicAppStandardSiteConfigOutputReference_Override(l LogicAppStandardSiteConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.logicAppStandard.LogicAppStandardSiteConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference)SetAlwaysOn(val interface{}) {
	if err := j.validateSetAlwaysOnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"alwaysOn",
		val,
	)
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference)SetAppScaleLimit(val *float64) {
	if err := j.validateSetAppScaleLimitParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"appScaleLimit",
		val,
	)
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference)SetDotnetFrameworkVersion(val *string) {
	if err := j.validateSetDotnetFrameworkVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dotnetFrameworkVersion",
		val,
	)
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference)SetElasticInstanceMinimum(val *float64) {
	if err := j.validateSetElasticInstanceMinimumParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"elasticInstanceMinimum",
		val,
	)
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference)SetFtpsState(val *string) {
	if err := j.validateSetFtpsStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ftpsState",
		val,
	)
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference)SetHealthCheckPath(val *string) {
	if err := j.validateSetHealthCheckPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthCheckPath",
		val,
	)
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference)SetHttp2Enabled(val interface{}) {
	if err := j.validateSetHttp2EnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"http2Enabled",
		val,
	)
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference)SetInternalValue(val *LogicAppStandardSiteConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference)SetLinuxFxVersion(val *string) {
	if err := j.validateSetLinuxFxVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"linuxFxVersion",
		val,
	)
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference)SetMinTlsVersion(val *string) {
	if err := j.validateSetMinTlsVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"minTlsVersion",
		val,
	)
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference)SetPreWarmedInstanceCount(val *float64) {
	if err := j.validateSetPreWarmedInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preWarmedInstanceCount",
		val,
	)
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference)SetRuntimeScaleMonitoringEnabled(val interface{}) {
	if err := j.validateSetRuntimeScaleMonitoringEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"runtimeScaleMonitoringEnabled",
		val,
	)
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference)SetScmMinTlsVersion(val *string) {
	if err := j.validateSetScmMinTlsVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scmMinTlsVersion",
		val,
	)
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference)SetScmType(val *string) {
	if err := j.validateSetScmTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scmType",
		val,
	)
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference)SetScmUseMainIpRestriction(val interface{}) {
	if err := j.validateSetScmUseMainIpRestrictionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scmUseMainIpRestriction",
		val,
	)
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference)SetUse32BitWorkerProcess(val interface{}) {
	if err := j.validateSetUse32BitWorkerProcessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"use32BitWorkerProcess",
		val,
	)
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference)SetVnetRouteAllEnabled(val interface{}) {
	if err := j.validateSetVnetRouteAllEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vnetRouteAllEnabled",
		val,
	)
}

func (j *jsiiProxy_LogicAppStandardSiteConfigOutputReference)SetWebsocketsEnabled(val interface{}) {
	if err := j.validateSetWebsocketsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"websocketsEnabled",
		val,
	)
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) PutCors(value *LogicAppStandardSiteConfigCors) {
	if err := l.validatePutCorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putCors",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) PutIpRestriction(value interface{}) {
	if err := l.validatePutIpRestrictionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putIpRestriction",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) PutScmIpRestriction(value interface{}) {
	if err := l.validatePutScmIpRestrictionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putScmIpRestriction",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ResetAlwaysOn() {
	_jsii_.InvokeVoid(
		l,
		"resetAlwaysOn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ResetAppScaleLimit() {
	_jsii_.InvokeVoid(
		l,
		"resetAppScaleLimit",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ResetCors() {
	_jsii_.InvokeVoid(
		l,
		"resetCors",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ResetDotnetFrameworkVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetDotnetFrameworkVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ResetElasticInstanceMinimum() {
	_jsii_.InvokeVoid(
		l,
		"resetElasticInstanceMinimum",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ResetFtpsState() {
	_jsii_.InvokeVoid(
		l,
		"resetFtpsState",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ResetHealthCheckPath() {
	_jsii_.InvokeVoid(
		l,
		"resetHealthCheckPath",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ResetHttp2Enabled() {
	_jsii_.InvokeVoid(
		l,
		"resetHttp2Enabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ResetIpRestriction() {
	_jsii_.InvokeVoid(
		l,
		"resetIpRestriction",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ResetLinuxFxVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetLinuxFxVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ResetMinTlsVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetMinTlsVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ResetPreWarmedInstanceCount() {
	_jsii_.InvokeVoid(
		l,
		"resetPreWarmedInstanceCount",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ResetRuntimeScaleMonitoringEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetRuntimeScaleMonitoringEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ResetScmIpRestriction() {
	_jsii_.InvokeVoid(
		l,
		"resetScmIpRestriction",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ResetScmMinTlsVersion() {
	_jsii_.InvokeVoid(
		l,
		"resetScmMinTlsVersion",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ResetScmType() {
	_jsii_.InvokeVoid(
		l,
		"resetScmType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ResetScmUseMainIpRestriction() {
	_jsii_.InvokeVoid(
		l,
		"resetScmUseMainIpRestriction",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ResetUse32BitWorkerProcess() {
	_jsii_.InvokeVoid(
		l,
		"resetUse32BitWorkerProcess",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ResetVnetRouteAllEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetVnetRouteAllEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ResetWebsocketsEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetWebsocketsEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		l,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LogicAppStandardSiteConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


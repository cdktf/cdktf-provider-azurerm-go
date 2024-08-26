// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package dataazurermwindowsfunctionapp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v13/dataazurermwindowsfunctionapp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermWindowsFunctionAppSiteConfigOutputReference interface {
	cdktf.ComplexObject
	AlwaysOn() cdktf.IResolvable
	ApiDefinitionUrl() *string
	ApiManagementApiId() *string
	AppCommandLine() *string
	ApplicationInsightsConnectionString() *string
	ApplicationInsightsKey() *string
	ApplicationStack() DataAzurermWindowsFunctionAppSiteConfigApplicationStackList
	AppScaleLimit() *float64
	AppServiceLogs() DataAzurermWindowsFunctionAppSiteConfigAppServiceLogsList
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
	Cors() DataAzurermWindowsFunctionAppSiteConfigCorsList
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DefaultDocuments() *[]*string
	DetailedErrorLoggingEnabled() cdktf.IResolvable
	ElasticInstanceMinimum() *float64
	// Experimental.
	Fqn() *string
	FtpsState() *string
	HealthCheckEvictionTimeInMin() *float64
	HealthCheckPath() *string
	Http2Enabled() cdktf.IResolvable
	InternalValue() *DataAzurermWindowsFunctionAppSiteConfig
	SetInternalValue(val *DataAzurermWindowsFunctionAppSiteConfig)
	IpRestriction() DataAzurermWindowsFunctionAppSiteConfigIpRestrictionList
	IpRestrictionDefaultAction() *string
	LoadBalancingMode() *string
	ManagedPipelineMode() *string
	MinimumTlsVersion() *string
	PreWarmedInstanceCount() *float64
	RemoteDebuggingEnabled() cdktf.IResolvable
	RemoteDebuggingVersion() *string
	RuntimeScaleMonitoringEnabled() cdktf.IResolvable
	ScmIpRestriction() DataAzurermWindowsFunctionAppSiteConfigScmIpRestrictionList
	ScmIpRestrictionDefaultAction() *string
	ScmMinimumTlsVersion() *string
	ScmType() *string
	ScmUseMainIpRestriction() cdktf.IResolvable
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Use32BitWorker() cdktf.IResolvable
	VnetRouteAllEnabled() cdktf.IResolvable
	WebsocketsEnabled() cdktf.IResolvable
	WindowsFxVersion() *string
	WorkerCount() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DataAzurermWindowsFunctionAppSiteConfigOutputReference
type jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) AlwaysOn() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"alwaysOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) ApiDefinitionUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiDefinitionUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) ApiManagementApiId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiManagementApiId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) AppCommandLine() *string {
	var returns *string
	_jsii_.Get(
		j,
		"appCommandLine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) ApplicationInsightsConnectionString() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationInsightsConnectionString",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) ApplicationInsightsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationInsightsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) ApplicationStack() DataAzurermWindowsFunctionAppSiteConfigApplicationStackList {
	var returns DataAzurermWindowsFunctionAppSiteConfigApplicationStackList
	_jsii_.Get(
		j,
		"applicationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) AppScaleLimit() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"appScaleLimit",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) AppServiceLogs() DataAzurermWindowsFunctionAppSiteConfigAppServiceLogsList {
	var returns DataAzurermWindowsFunctionAppSiteConfigAppServiceLogsList
	_jsii_.Get(
		j,
		"appServiceLogs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) Cors() DataAzurermWindowsFunctionAppSiteConfigCorsList {
	var returns DataAzurermWindowsFunctionAppSiteConfigCorsList
	_jsii_.Get(
		j,
		"cors",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) DefaultDocuments() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"defaultDocuments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) DetailedErrorLoggingEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"detailedErrorLoggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) ElasticInstanceMinimum() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"elasticInstanceMinimum",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) FtpsState() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ftpsState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) HealthCheckEvictionTimeInMin() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"healthCheckEvictionTimeInMin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) HealthCheckPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthCheckPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) Http2Enabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"http2Enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) InternalValue() *DataAzurermWindowsFunctionAppSiteConfig {
	var returns *DataAzurermWindowsFunctionAppSiteConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) IpRestriction() DataAzurermWindowsFunctionAppSiteConfigIpRestrictionList {
	var returns DataAzurermWindowsFunctionAppSiteConfigIpRestrictionList
	_jsii_.Get(
		j,
		"ipRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) IpRestrictionDefaultAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipRestrictionDefaultAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) LoadBalancingMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"loadBalancingMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) ManagedPipelineMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"managedPipelineMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) MinimumTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"minimumTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) PreWarmedInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"preWarmedInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) RemoteDebuggingEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"remoteDebuggingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) RemoteDebuggingVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteDebuggingVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) RuntimeScaleMonitoringEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"runtimeScaleMonitoringEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) ScmIpRestriction() DataAzurermWindowsFunctionAppSiteConfigScmIpRestrictionList {
	var returns DataAzurermWindowsFunctionAppSiteConfigScmIpRestrictionList
	_jsii_.Get(
		j,
		"scmIpRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) ScmIpRestrictionDefaultAction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmIpRestrictionDefaultAction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) ScmMinimumTlsVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmMinimumTlsVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) ScmType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scmType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) ScmUseMainIpRestriction() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"scmUseMainIpRestriction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) Use32BitWorker() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"use32BitWorker",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) VnetRouteAllEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"vnetRouteAllEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) WebsocketsEnabled() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"websocketsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) WindowsFxVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"windowsFxVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) WorkerCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"workerCount",
		&returns,
	)
	return returns
}


func NewDataAzurermWindowsFunctionAppSiteConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAzurermWindowsFunctionAppSiteConfigOutputReference {
	_init_.Initialize()

	if err := validateNewDataAzurermWindowsFunctionAppSiteConfigOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermWindowsFunctionApp.DataAzurermWindowsFunctionAppSiteConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAzurermWindowsFunctionAppSiteConfigOutputReference_Override(d DataAzurermWindowsFunctionAppSiteConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermWindowsFunctionApp.DataAzurermWindowsFunctionAppSiteConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference)SetInternalValue(val *DataAzurermWindowsFunctionAppSiteConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermWindowsFunctionAppSiteConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


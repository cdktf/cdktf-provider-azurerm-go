// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitordatacollectionrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v10/monitordatacollectionrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorDataCollectionRuleDataSourcesOutputReference interface {
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
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DataImport() MonitorDataCollectionRuleDataSourcesDataImportOutputReference
	DataImportInput() *MonitorDataCollectionRuleDataSourcesDataImport
	Extension() MonitorDataCollectionRuleDataSourcesExtensionList
	ExtensionInput() interface{}
	// Experimental.
	Fqn() *string
	IisLog() MonitorDataCollectionRuleDataSourcesIisLogList
	IisLogInput() interface{}
	InternalValue() *MonitorDataCollectionRuleDataSources
	SetInternalValue(val *MonitorDataCollectionRuleDataSources)
	LogFile() MonitorDataCollectionRuleDataSourcesLogFileList
	LogFileInput() interface{}
	PerformanceCounter() MonitorDataCollectionRuleDataSourcesPerformanceCounterList
	PerformanceCounterInput() interface{}
	PlatformTelemetry() MonitorDataCollectionRuleDataSourcesPlatformTelemetryList
	PlatformTelemetryInput() interface{}
	PrometheusForwarder() MonitorDataCollectionRuleDataSourcesPrometheusForwarderList
	PrometheusForwarderInput() interface{}
	Syslog() MonitorDataCollectionRuleDataSourcesSyslogList
	SyslogInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WindowsEventLog() MonitorDataCollectionRuleDataSourcesWindowsEventLogList
	WindowsEventLogInput() interface{}
	WindowsFirewallLog() MonitorDataCollectionRuleDataSourcesWindowsFirewallLogList
	WindowsFirewallLogInput() interface{}
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
	PutDataImport(value *MonitorDataCollectionRuleDataSourcesDataImport)
	PutExtension(value interface{})
	PutIisLog(value interface{})
	PutLogFile(value interface{})
	PutPerformanceCounter(value interface{})
	PutPlatformTelemetry(value interface{})
	PutPrometheusForwarder(value interface{})
	PutSyslog(value interface{})
	PutWindowsEventLog(value interface{})
	PutWindowsFirewallLog(value interface{})
	ResetDataImport()
	ResetExtension()
	ResetIisLog()
	ResetLogFile()
	ResetPerformanceCounter()
	ResetPlatformTelemetry()
	ResetPrometheusForwarder()
	ResetSyslog()
	ResetWindowsEventLog()
	ResetWindowsFirewallLog()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorDataCollectionRuleDataSourcesOutputReference
type jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) DataImport() MonitorDataCollectionRuleDataSourcesDataImportOutputReference {
	var returns MonitorDataCollectionRuleDataSourcesDataImportOutputReference
	_jsii_.Get(
		j,
		"dataImport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) DataImportInput() *MonitorDataCollectionRuleDataSourcesDataImport {
	var returns *MonitorDataCollectionRuleDataSourcesDataImport
	_jsii_.Get(
		j,
		"dataImportInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) Extension() MonitorDataCollectionRuleDataSourcesExtensionList {
	var returns MonitorDataCollectionRuleDataSourcesExtensionList
	_jsii_.Get(
		j,
		"extension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) ExtensionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) IisLog() MonitorDataCollectionRuleDataSourcesIisLogList {
	var returns MonitorDataCollectionRuleDataSourcesIisLogList
	_jsii_.Get(
		j,
		"iisLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) IisLogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"iisLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) InternalValue() *MonitorDataCollectionRuleDataSources {
	var returns *MonitorDataCollectionRuleDataSources
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) LogFile() MonitorDataCollectionRuleDataSourcesLogFileList {
	var returns MonitorDataCollectionRuleDataSourcesLogFileList
	_jsii_.Get(
		j,
		"logFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) LogFileInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) PerformanceCounter() MonitorDataCollectionRuleDataSourcesPerformanceCounterList {
	var returns MonitorDataCollectionRuleDataSourcesPerformanceCounterList
	_jsii_.Get(
		j,
		"performanceCounter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) PerformanceCounterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"performanceCounterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) PlatformTelemetry() MonitorDataCollectionRuleDataSourcesPlatformTelemetryList {
	var returns MonitorDataCollectionRuleDataSourcesPlatformTelemetryList
	_jsii_.Get(
		j,
		"platformTelemetry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) PlatformTelemetryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"platformTelemetryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) PrometheusForwarder() MonitorDataCollectionRuleDataSourcesPrometheusForwarderList {
	var returns MonitorDataCollectionRuleDataSourcesPrometheusForwarderList
	_jsii_.Get(
		j,
		"prometheusForwarder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) PrometheusForwarderInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"prometheusForwarderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) Syslog() MonitorDataCollectionRuleDataSourcesSyslogList {
	var returns MonitorDataCollectionRuleDataSourcesSyslogList
	_jsii_.Get(
		j,
		"syslog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) SyslogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"syslogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) WindowsEventLog() MonitorDataCollectionRuleDataSourcesWindowsEventLogList {
	var returns MonitorDataCollectionRuleDataSourcesWindowsEventLogList
	_jsii_.Get(
		j,
		"windowsEventLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) WindowsEventLogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"windowsEventLogInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) WindowsFirewallLog() MonitorDataCollectionRuleDataSourcesWindowsFirewallLogList {
	var returns MonitorDataCollectionRuleDataSourcesWindowsFirewallLogList
	_jsii_.Get(
		j,
		"windowsFirewallLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) WindowsFirewallLogInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"windowsFirewallLogInput",
		&returns,
	)
	return returns
}


func NewMonitorDataCollectionRuleDataSourcesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorDataCollectionRuleDataSourcesOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorDataCollectionRuleDataSourcesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorDataCollectionRule.MonitorDataCollectionRuleDataSourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorDataCollectionRuleDataSourcesOutputReference_Override(m MonitorDataCollectionRuleDataSourcesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorDataCollectionRule.MonitorDataCollectionRuleDataSourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference)SetInternalValue(val *MonitorDataCollectionRuleDataSources) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := m.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := m.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := m.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := m.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := m.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := m.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := m.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := m.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) PutDataImport(value *MonitorDataCollectionRuleDataSourcesDataImport) {
	if err := m.validatePutDataImportParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDataImport",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) PutExtension(value interface{}) {
	if err := m.validatePutExtensionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putExtension",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) PutIisLog(value interface{}) {
	if err := m.validatePutIisLogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putIisLog",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) PutLogFile(value interface{}) {
	if err := m.validatePutLogFileParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putLogFile",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) PutPerformanceCounter(value interface{}) {
	if err := m.validatePutPerformanceCounterParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putPerformanceCounter",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) PutPlatformTelemetry(value interface{}) {
	if err := m.validatePutPlatformTelemetryParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putPlatformTelemetry",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) PutPrometheusForwarder(value interface{}) {
	if err := m.validatePutPrometheusForwarderParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putPrometheusForwarder",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) PutSyslog(value interface{}) {
	if err := m.validatePutSyslogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putSyslog",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) PutWindowsEventLog(value interface{}) {
	if err := m.validatePutWindowsEventLogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putWindowsEventLog",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) PutWindowsFirewallLog(value interface{}) {
	if err := m.validatePutWindowsFirewallLogParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putWindowsFirewallLog",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) ResetDataImport() {
	_jsii_.InvokeVoid(
		m,
		"resetDataImport",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) ResetExtension() {
	_jsii_.InvokeVoid(
		m,
		"resetExtension",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) ResetIisLog() {
	_jsii_.InvokeVoid(
		m,
		"resetIisLog",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) ResetLogFile() {
	_jsii_.InvokeVoid(
		m,
		"resetLogFile",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) ResetPerformanceCounter() {
	_jsii_.InvokeVoid(
		m,
		"resetPerformanceCounter",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) ResetPlatformTelemetry() {
	_jsii_.InvokeVoid(
		m,
		"resetPlatformTelemetry",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) ResetPrometheusForwarder() {
	_jsii_.InvokeVoid(
		m,
		"resetPrometheusForwarder",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) ResetSyslog() {
	_jsii_.InvokeVoid(
		m,
		"resetSyslog",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) ResetWindowsEventLog() {
	_jsii_.InvokeVoid(
		m,
		"resetWindowsEventLog",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) ResetWindowsFirewallLog() {
	_jsii_.InvokeVoid(
		m,
		"resetWindowsFirewallLog",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


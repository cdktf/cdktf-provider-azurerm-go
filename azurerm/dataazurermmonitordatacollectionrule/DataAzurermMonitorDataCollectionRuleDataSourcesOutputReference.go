package dataazurermmonitordatacollectionrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/dataazurermmonitordatacollectionrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference interface {
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
	DataImport() DataAzurermMonitorDataCollectionRuleDataSourcesDataImportList
	Extension() DataAzurermMonitorDataCollectionRuleDataSourcesExtensionList
	// Experimental.
	Fqn() *string
	IisLog() DataAzurermMonitorDataCollectionRuleDataSourcesIisLogList
	InternalValue() *DataAzurermMonitorDataCollectionRuleDataSources
	SetInternalValue(val *DataAzurermMonitorDataCollectionRuleDataSources)
	LogFile() DataAzurermMonitorDataCollectionRuleDataSourcesLogFileList
	PerformanceCounter() DataAzurermMonitorDataCollectionRuleDataSourcesPerformanceCounterList
	PlatformTelemetry() DataAzurermMonitorDataCollectionRuleDataSourcesPlatformTelemetryList
	PrometheusForwarder() DataAzurermMonitorDataCollectionRuleDataSourcesPrometheusForwarderList
	Syslog() DataAzurermMonitorDataCollectionRuleDataSourcesSyslogList
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WindowsEventLog() DataAzurermMonitorDataCollectionRuleDataSourcesWindowsEventLogList
	WindowsFirewallLog() DataAzurermMonitorDataCollectionRuleDataSourcesWindowsFirewallLogList
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

// The jsii proxy struct for DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference
type jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) DataImport() DataAzurermMonitorDataCollectionRuleDataSourcesDataImportList {
	var returns DataAzurermMonitorDataCollectionRuleDataSourcesDataImportList
	_jsii_.Get(
		j,
		"dataImport",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) Extension() DataAzurermMonitorDataCollectionRuleDataSourcesExtensionList {
	var returns DataAzurermMonitorDataCollectionRuleDataSourcesExtensionList
	_jsii_.Get(
		j,
		"extension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) IisLog() DataAzurermMonitorDataCollectionRuleDataSourcesIisLogList {
	var returns DataAzurermMonitorDataCollectionRuleDataSourcesIisLogList
	_jsii_.Get(
		j,
		"iisLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) InternalValue() *DataAzurermMonitorDataCollectionRuleDataSources {
	var returns *DataAzurermMonitorDataCollectionRuleDataSources
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) LogFile() DataAzurermMonitorDataCollectionRuleDataSourcesLogFileList {
	var returns DataAzurermMonitorDataCollectionRuleDataSourcesLogFileList
	_jsii_.Get(
		j,
		"logFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) PerformanceCounter() DataAzurermMonitorDataCollectionRuleDataSourcesPerformanceCounterList {
	var returns DataAzurermMonitorDataCollectionRuleDataSourcesPerformanceCounterList
	_jsii_.Get(
		j,
		"performanceCounter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) PlatformTelemetry() DataAzurermMonitorDataCollectionRuleDataSourcesPlatformTelemetryList {
	var returns DataAzurermMonitorDataCollectionRuleDataSourcesPlatformTelemetryList
	_jsii_.Get(
		j,
		"platformTelemetry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) PrometheusForwarder() DataAzurermMonitorDataCollectionRuleDataSourcesPrometheusForwarderList {
	var returns DataAzurermMonitorDataCollectionRuleDataSourcesPrometheusForwarderList
	_jsii_.Get(
		j,
		"prometheusForwarder",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) Syslog() DataAzurermMonitorDataCollectionRuleDataSourcesSyslogList {
	var returns DataAzurermMonitorDataCollectionRuleDataSourcesSyslogList
	_jsii_.Get(
		j,
		"syslog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) WindowsEventLog() DataAzurermMonitorDataCollectionRuleDataSourcesWindowsEventLogList {
	var returns DataAzurermMonitorDataCollectionRuleDataSourcesWindowsEventLogList
	_jsii_.Get(
		j,
		"windowsEventLog",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) WindowsFirewallLog() DataAzurermMonitorDataCollectionRuleDataSourcesWindowsFirewallLogList {
	var returns DataAzurermMonitorDataCollectionRuleDataSourcesWindowsFirewallLogList
	_jsii_.Get(
		j,
		"windowsFirewallLog",
		&returns,
	)
	return returns
}


func NewDataAzurermMonitorDataCollectionRuleDataSourcesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference {
	_init_.Initialize()

	if err := validateNewDataAzurermMonitorDataCollectionRuleDataSourcesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermMonitorDataCollectionRule.DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDataAzurermMonitorDataCollectionRuleDataSourcesOutputReference_Override(d DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.dataAzurermMonitorDataCollectionRule.DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference)SetInternalValue(val *DataAzurermMonitorDataCollectionRuleDataSources) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DataAzurermMonitorDataCollectionRuleDataSourcesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


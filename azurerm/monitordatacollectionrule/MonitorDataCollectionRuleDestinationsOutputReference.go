// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package monitordatacollectionrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/monitordatacollectionrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorDataCollectionRuleDestinationsOutputReference interface {
	cdktf.ComplexObject
	AzureMonitorMetrics() MonitorDataCollectionRuleDestinationsAzureMonitorMetricsOutputReference
	AzureMonitorMetricsInput() *MonitorDataCollectionRuleDestinationsAzureMonitorMetrics
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
	EventHub() MonitorDataCollectionRuleDestinationsEventHubOutputReference
	EventHubDirect() MonitorDataCollectionRuleDestinationsEventHubDirectOutputReference
	EventHubDirectInput() *MonitorDataCollectionRuleDestinationsEventHubDirect
	EventHubInput() *MonitorDataCollectionRuleDestinationsEventHub
	// Experimental.
	Fqn() *string
	InternalValue() *MonitorDataCollectionRuleDestinations
	SetInternalValue(val *MonitorDataCollectionRuleDestinations)
	LogAnalytics() MonitorDataCollectionRuleDestinationsLogAnalyticsList
	LogAnalyticsInput() interface{}
	MonitorAccount() MonitorDataCollectionRuleDestinationsMonitorAccountList
	MonitorAccountInput() interface{}
	StorageBlob() MonitorDataCollectionRuleDestinationsStorageBlobList
	StorageBlobDirect() MonitorDataCollectionRuleDestinationsStorageBlobDirectList
	StorageBlobDirectInput() interface{}
	StorageBlobInput() interface{}
	StorageTableDirect() MonitorDataCollectionRuleDestinationsStorageTableDirectList
	StorageTableDirectInput() interface{}
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
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	PutAzureMonitorMetrics(value *MonitorDataCollectionRuleDestinationsAzureMonitorMetrics)
	PutEventHub(value *MonitorDataCollectionRuleDestinationsEventHub)
	PutEventHubDirect(value *MonitorDataCollectionRuleDestinationsEventHubDirect)
	PutLogAnalytics(value interface{})
	PutMonitorAccount(value interface{})
	PutStorageBlob(value interface{})
	PutStorageBlobDirect(value interface{})
	PutStorageTableDirect(value interface{})
	ResetAzureMonitorMetrics()
	ResetEventHub()
	ResetEventHubDirect()
	ResetLogAnalytics()
	ResetMonitorAccount()
	ResetStorageBlob()
	ResetStorageBlobDirect()
	ResetStorageTableDirect()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorDataCollectionRuleDestinationsOutputReference
type jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) AzureMonitorMetrics() MonitorDataCollectionRuleDestinationsAzureMonitorMetricsOutputReference {
	var returns MonitorDataCollectionRuleDestinationsAzureMonitorMetricsOutputReference
	_jsii_.Get(
		j,
		"azureMonitorMetrics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) AzureMonitorMetricsInput() *MonitorDataCollectionRuleDestinationsAzureMonitorMetrics {
	var returns *MonitorDataCollectionRuleDestinationsAzureMonitorMetrics
	_jsii_.Get(
		j,
		"azureMonitorMetricsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) EventHub() MonitorDataCollectionRuleDestinationsEventHubOutputReference {
	var returns MonitorDataCollectionRuleDestinationsEventHubOutputReference
	_jsii_.Get(
		j,
		"eventHub",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) EventHubDirect() MonitorDataCollectionRuleDestinationsEventHubDirectOutputReference {
	var returns MonitorDataCollectionRuleDestinationsEventHubDirectOutputReference
	_jsii_.Get(
		j,
		"eventHubDirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) EventHubDirectInput() *MonitorDataCollectionRuleDestinationsEventHubDirect {
	var returns *MonitorDataCollectionRuleDestinationsEventHubDirect
	_jsii_.Get(
		j,
		"eventHubDirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) EventHubInput() *MonitorDataCollectionRuleDestinationsEventHub {
	var returns *MonitorDataCollectionRuleDestinationsEventHub
	_jsii_.Get(
		j,
		"eventHubInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) InternalValue() *MonitorDataCollectionRuleDestinations {
	var returns *MonitorDataCollectionRuleDestinations
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) LogAnalytics() MonitorDataCollectionRuleDestinationsLogAnalyticsList {
	var returns MonitorDataCollectionRuleDestinationsLogAnalyticsList
	_jsii_.Get(
		j,
		"logAnalytics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) LogAnalyticsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"logAnalyticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) MonitorAccount() MonitorDataCollectionRuleDestinationsMonitorAccountList {
	var returns MonitorDataCollectionRuleDestinationsMonitorAccountList
	_jsii_.Get(
		j,
		"monitorAccount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) MonitorAccountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"monitorAccountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) StorageBlob() MonitorDataCollectionRuleDestinationsStorageBlobList {
	var returns MonitorDataCollectionRuleDestinationsStorageBlobList
	_jsii_.Get(
		j,
		"storageBlob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) StorageBlobDirect() MonitorDataCollectionRuleDestinationsStorageBlobDirectList {
	var returns MonitorDataCollectionRuleDestinationsStorageBlobDirectList
	_jsii_.Get(
		j,
		"storageBlobDirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) StorageBlobDirectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageBlobDirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) StorageBlobInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageBlobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) StorageTableDirect() MonitorDataCollectionRuleDestinationsStorageTableDirectList {
	var returns MonitorDataCollectionRuleDestinationsStorageTableDirectList
	_jsii_.Get(
		j,
		"storageTableDirect",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) StorageTableDirectInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageTableDirectInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorDataCollectionRuleDestinationsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MonitorDataCollectionRuleDestinationsOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorDataCollectionRuleDestinationsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorDataCollectionRule.MonitorDataCollectionRuleDestinationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMonitorDataCollectionRuleDestinationsOutputReference_Override(m MonitorDataCollectionRuleDestinationsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorDataCollectionRule.MonitorDataCollectionRuleDestinationsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference)SetInternalValue(val *MonitorDataCollectionRuleDestinations) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := m.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) PutAzureMonitorMetrics(value *MonitorDataCollectionRuleDestinationsAzureMonitorMetrics) {
	if err := m.validatePutAzureMonitorMetricsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putAzureMonitorMetrics",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) PutEventHub(value *MonitorDataCollectionRuleDestinationsEventHub) {
	if err := m.validatePutEventHubParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEventHub",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) PutEventHubDirect(value *MonitorDataCollectionRuleDestinationsEventHubDirect) {
	if err := m.validatePutEventHubDirectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putEventHubDirect",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) PutLogAnalytics(value interface{}) {
	if err := m.validatePutLogAnalyticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putLogAnalytics",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) PutMonitorAccount(value interface{}) {
	if err := m.validatePutMonitorAccountParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMonitorAccount",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) PutStorageBlob(value interface{}) {
	if err := m.validatePutStorageBlobParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putStorageBlob",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) PutStorageBlobDirect(value interface{}) {
	if err := m.validatePutStorageBlobDirectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putStorageBlobDirect",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) PutStorageTableDirect(value interface{}) {
	if err := m.validatePutStorageTableDirectParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putStorageTableDirect",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) ResetAzureMonitorMetrics() {
	_jsii_.InvokeVoid(
		m,
		"resetAzureMonitorMetrics",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) ResetEventHub() {
	_jsii_.InvokeVoid(
		m,
		"resetEventHub",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) ResetEventHubDirect() {
	_jsii_.InvokeVoid(
		m,
		"resetEventHubDirect",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) ResetLogAnalytics() {
	_jsii_.InvokeVoid(
		m,
		"resetLogAnalytics",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) ResetMonitorAccount() {
	_jsii_.InvokeVoid(
		m,
		"resetMonitorAccount",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) ResetStorageBlob() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageBlob",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) ResetStorageBlobDirect() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageBlobDirect",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) ResetStorageTableDirect() {
	_jsii_.InvokeVoid(
		m,
		"resetStorageTableDirect",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) Resolve(context cdktf.IResolveContext) interface{} {
	if err := m.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		m,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRuleDestinationsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


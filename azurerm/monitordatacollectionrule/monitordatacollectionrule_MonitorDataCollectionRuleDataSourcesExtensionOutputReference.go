package monitordatacollectionrule

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v4/monitordatacollectionrule/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MonitorDataCollectionRuleDataSourcesExtensionOutputReference interface {
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
	ExtensionJson() *string
	SetExtensionJson(val *string)
	ExtensionJsonInput() *string
	ExtensionName() *string
	SetExtensionName(val *string)
	ExtensionNameInput() *string
	// Experimental.
	Fqn() *string
	InputDataSources() *[]*string
	SetInputDataSources(val *[]*string)
	InputDataSourcesInput() *[]*string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	Streams() *[]*string
	SetStreams(val *[]*string)
	StreamsInput() *[]*string
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
	ResetExtensionJson()
	ResetInputDataSources()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MonitorDataCollectionRuleDataSourcesExtensionOutputReference
type jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) ExtensionJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) ExtensionJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) ExtensionName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) ExtensionNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) InputDataSources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputDataSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) InputDataSourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"inputDataSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) Streams() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"streams",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) StreamsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"streamsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMonitorDataCollectionRuleDataSourcesExtensionOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) MonitorDataCollectionRuleDataSourcesExtensionOutputReference {
	_init_.Initialize()

	if err := validateNewMonitorDataCollectionRuleDataSourcesExtensionOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorDataCollectionRule.MonitorDataCollectionRuleDataSourcesExtensionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewMonitorDataCollectionRuleDataSourcesExtensionOutputReference_Override(m MonitorDataCollectionRuleDataSourcesExtensionOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.monitorDataCollectionRule.MonitorDataCollectionRuleDataSourcesExtensionOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		m,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference)SetExtensionJson(val *string) {
	if err := j.validateSetExtensionJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extensionJson",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference)SetExtensionName(val *string) {
	if err := j.validateSetExtensionNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extensionName",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference)SetInputDataSources(val *[]*string) {
	if err := j.validateSetInputDataSourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"inputDataSources",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference)SetStreams(val *[]*string) {
	if err := j.validateSetStreamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"streams",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) ResetExtensionJson() {
	_jsii_.InvokeVoid(
		m,
		"resetExtensionJson",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) ResetInputDataSources() {
	_jsii_.InvokeVoid(
		m,
		"resetInputDataSources",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MonitorDataCollectionRuleDataSourcesExtensionOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


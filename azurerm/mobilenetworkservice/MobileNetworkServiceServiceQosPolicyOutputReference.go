package mobilenetworkservice

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/mobilenetworkservice/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MobileNetworkServiceServiceQosPolicyOutputReference interface {
	cdktf.ComplexObject
	AllocationAndRetentionPriorityLevel() *float64
	SetAllocationAndRetentionPriorityLevel(val *float64)
	AllocationAndRetentionPriorityLevelInput() *float64
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
	InternalValue() *MobileNetworkServiceServiceQosPolicy
	SetInternalValue(val *MobileNetworkServiceServiceQosPolicy)
	MaximumBitRate() MobileNetworkServiceServiceQosPolicyMaximumBitRateOutputReference
	MaximumBitRateInput() *MobileNetworkServiceServiceQosPolicyMaximumBitRate
	PreemptionCapability() *string
	SetPreemptionCapability(val *string)
	PreemptionCapabilityInput() *string
	PreemptionVulnerability() *string
	SetPreemptionVulnerability(val *string)
	PreemptionVulnerabilityInput() *string
	QosIndicator() *float64
	SetQosIndicator(val *float64)
	QosIndicatorInput() *float64
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
	PutMaximumBitRate(value *MobileNetworkServiceServiceQosPolicyMaximumBitRate)
	ResetAllocationAndRetentionPriorityLevel()
	ResetPreemptionCapability()
	ResetPreemptionVulnerability()
	ResetQosIndicator()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MobileNetworkServiceServiceQosPolicyOutputReference
type jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) AllocationAndRetentionPriorityLevel() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationAndRetentionPriorityLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) AllocationAndRetentionPriorityLevelInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"allocationAndRetentionPriorityLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) InternalValue() *MobileNetworkServiceServiceQosPolicy {
	var returns *MobileNetworkServiceServiceQosPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) MaximumBitRate() MobileNetworkServiceServiceQosPolicyMaximumBitRateOutputReference {
	var returns MobileNetworkServiceServiceQosPolicyMaximumBitRateOutputReference
	_jsii_.Get(
		j,
		"maximumBitRate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) MaximumBitRateInput() *MobileNetworkServiceServiceQosPolicyMaximumBitRate {
	var returns *MobileNetworkServiceServiceQosPolicyMaximumBitRate
	_jsii_.Get(
		j,
		"maximumBitRateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) PreemptionCapability() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preemptionCapability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) PreemptionCapabilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preemptionCapabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) PreemptionVulnerability() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preemptionVulnerability",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) PreemptionVulnerabilityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"preemptionVulnerabilityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) QosIndicator() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"qosIndicator",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) QosIndicatorInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"qosIndicatorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMobileNetworkServiceServiceQosPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MobileNetworkServiceServiceQosPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewMobileNetworkServiceServiceQosPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mobileNetworkService.MobileNetworkServiceServiceQosPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMobileNetworkServiceServiceQosPolicyOutputReference_Override(m MobileNetworkServiceServiceQosPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mobileNetworkService.MobileNetworkServiceServiceQosPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference)SetAllocationAndRetentionPriorityLevel(val *float64) {
	if err := j.validateSetAllocationAndRetentionPriorityLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allocationAndRetentionPriorityLevel",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference)SetInternalValue(val *MobileNetworkServiceServiceQosPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference)SetPreemptionCapability(val *string) {
	if err := j.validateSetPreemptionCapabilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preemptionCapability",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference)SetPreemptionVulnerability(val *string) {
	if err := j.validateSetPreemptionVulnerabilityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preemptionVulnerability",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference)SetQosIndicator(val *float64) {
	if err := j.validateSetQosIndicatorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qosIndicator",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) PutMaximumBitRate(value *MobileNetworkServiceServiceQosPolicyMaximumBitRate) {
	if err := m.validatePutMaximumBitRateParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putMaximumBitRate",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) ResetAllocationAndRetentionPriorityLevel() {
	_jsii_.InvokeVoid(
		m,
		"resetAllocationAndRetentionPriorityLevel",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) ResetPreemptionCapability() {
	_jsii_.InvokeVoid(
		m,
		"resetPreemptionCapability",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) ResetPreemptionVulnerability() {
	_jsii_.InvokeVoid(
		m,
		"resetPreemptionVulnerability",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) ResetQosIndicator() {
	_jsii_.InvokeVoid(
		m,
		"resetQosIndicator",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MobileNetworkServiceServiceQosPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package mssqlvirtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v3/mssqlvirtualmachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MssqlVirtualMachineAutoBackupManualScheduleOutputReference interface {
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
	// Experimental.
	Fqn() *string
	FullBackupFrequency() *string
	SetFullBackupFrequency(val *string)
	FullBackupFrequencyInput() *string
	FullBackupStartHour() *float64
	SetFullBackupStartHour(val *float64)
	FullBackupStartHourInput() *float64
	FullBackupWindowInHours() *float64
	SetFullBackupWindowInHours(val *float64)
	FullBackupWindowInHoursInput() *float64
	InternalValue() *MssqlVirtualMachineAutoBackupManualSchedule
	SetInternalValue(val *MssqlVirtualMachineAutoBackupManualSchedule)
	LogBackupFrequencyInMinutes() *float64
	SetLogBackupFrequencyInMinutes(val *float64)
	LogBackupFrequencyInMinutesInput() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MssqlVirtualMachineAutoBackupManualScheduleOutputReference
type jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) FullBackupFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullBackupFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) FullBackupFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fullBackupFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) FullBackupStartHour() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fullBackupStartHour",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) FullBackupStartHourInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fullBackupStartHourInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) FullBackupWindowInHours() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fullBackupWindowInHours",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) FullBackupWindowInHoursInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"fullBackupWindowInHoursInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) InternalValue() *MssqlVirtualMachineAutoBackupManualSchedule {
	var returns *MssqlVirtualMachineAutoBackupManualSchedule
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) LogBackupFrequencyInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logBackupFrequencyInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) LogBackupFrequencyInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logBackupFrequencyInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMssqlVirtualMachineAutoBackupManualScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MssqlVirtualMachineAutoBackupManualScheduleOutputReference {
	_init_.Initialize()

	if err := validateNewMssqlVirtualMachineAutoBackupManualScheduleOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineAutoBackupManualScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMssqlVirtualMachineAutoBackupManualScheduleOutputReference_Override(m MssqlVirtualMachineAutoBackupManualScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineAutoBackupManualScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference)SetFullBackupFrequency(val *string) {
	if err := j.validateSetFullBackupFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullBackupFrequency",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference)SetFullBackupStartHour(val *float64) {
	if err := j.validateSetFullBackupStartHourParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullBackupStartHour",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference)SetFullBackupWindowInHours(val *float64) {
	if err := j.validateSetFullBackupWindowInHoursParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fullBackupWindowInHours",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference)SetInternalValue(val *MssqlVirtualMachineAutoBackupManualSchedule) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference)SetLogBackupFrequencyInMinutes(val *float64) {
	if err := j.validateSetLogBackupFrequencyInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logBackupFrequencyInMinutes",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MssqlVirtualMachineAutoBackupManualScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


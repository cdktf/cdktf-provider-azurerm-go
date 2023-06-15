package mssqlvirtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/mssqlvirtualmachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference interface {
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
	DataFileCount() *float64
	SetDataFileCount(val *float64)
	DataFileCountInput() *float64
	DataFileGrowthInMb() *float64
	SetDataFileGrowthInMb(val *float64)
	DataFileGrowthInMbInput() *float64
	DataFileSizeMb() *float64
	SetDataFileSizeMb(val *float64)
	DataFileSizeMbInput() *float64
	DefaultFilePath() *string
	SetDefaultFilePath(val *string)
	DefaultFilePathInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MssqlVirtualMachineStorageConfigurationTempDbSettings
	SetInternalValue(val *MssqlVirtualMachineStorageConfigurationTempDbSettings)
	LogFileGrowthMb() *float64
	SetLogFileGrowthMb(val *float64)
	LogFileGrowthMbInput() *float64
	LogFileSizeMb() *float64
	SetLogFileSizeMb(val *float64)
	LogFileSizeMbInput() *float64
	Luns() *[]*float64
	SetLuns(val *[]*float64)
	LunsInput() *[]*float64
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
	ResetDataFileCount()
	ResetDataFileGrowthInMb()
	ResetDataFileSizeMb()
	ResetLogFileGrowthMb()
	ResetLogFileSizeMb()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference
type jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) DataFileCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataFileCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) DataFileCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataFileCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) DataFileGrowthInMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataFileGrowthInMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) DataFileGrowthInMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataFileGrowthInMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) DataFileSizeMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataFileSizeMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) DataFileSizeMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataFileSizeMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) DefaultFilePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultFilePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) DefaultFilePathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultFilePathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) InternalValue() *MssqlVirtualMachineStorageConfigurationTempDbSettings {
	var returns *MssqlVirtualMachineStorageConfigurationTempDbSettings
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) LogFileGrowthMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logFileGrowthMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) LogFileGrowthMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logFileGrowthMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) LogFileSizeMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logFileSizeMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) LogFileSizeMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"logFileSizeMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) Luns() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"luns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) LunsInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"lunsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference {
	_init_.Initialize()

	if err := validateNewMssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference_Override(m MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference)SetDataFileCount(val *float64) {
	if err := j.validateSetDataFileCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataFileCount",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference)SetDataFileGrowthInMb(val *float64) {
	if err := j.validateSetDataFileGrowthInMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataFileGrowthInMb",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference)SetDataFileSizeMb(val *float64) {
	if err := j.validateSetDataFileSizeMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataFileSizeMb",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference)SetDefaultFilePath(val *string) {
	if err := j.validateSetDefaultFilePathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultFilePath",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference)SetInternalValue(val *MssqlVirtualMachineStorageConfigurationTempDbSettings) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference)SetLogFileGrowthMb(val *float64) {
	if err := j.validateSetLogFileGrowthMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logFileGrowthMb",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference)SetLogFileSizeMb(val *float64) {
	if err := j.validateSetLogFileSizeMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logFileSizeMb",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference)SetLuns(val *[]*float64) {
	if err := j.validateSetLunsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"luns",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) ResetDataFileCount() {
	_jsii_.InvokeVoid(
		m,
		"resetDataFileCount",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) ResetDataFileGrowthInMb() {
	_jsii_.InvokeVoid(
		m,
		"resetDataFileGrowthInMb",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) ResetDataFileSizeMb() {
	_jsii_.InvokeVoid(
		m,
		"resetDataFileSizeMb",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) ResetLogFileGrowthMb() {
	_jsii_.InvokeVoid(
		m,
		"resetLogFileGrowthMb",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) ResetLogFileSizeMb() {
	_jsii_.InvokeVoid(
		m,
		"resetLogFileSizeMb",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


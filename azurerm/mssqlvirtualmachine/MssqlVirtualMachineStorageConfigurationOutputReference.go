package mssqlvirtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/mssqlvirtualmachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type MssqlVirtualMachineStorageConfigurationOutputReference interface {
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
	DataSettings() MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference
	DataSettingsInput() *MssqlVirtualMachineStorageConfigurationDataSettings
	DiskType() *string
	SetDiskType(val *string)
	DiskTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *MssqlVirtualMachineStorageConfiguration
	SetInternalValue(val *MssqlVirtualMachineStorageConfiguration)
	LogSettings() MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference
	LogSettingsInput() *MssqlVirtualMachineStorageConfigurationLogSettings
	StorageWorkloadType() *string
	SetStorageWorkloadType(val *string)
	StorageWorkloadTypeInput() *string
	SystemDbOnDataDiskEnabled() interface{}
	SetSystemDbOnDataDiskEnabled(val interface{})
	SystemDbOnDataDiskEnabledInput() interface{}
	TempDbSettings() MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference
	TempDbSettingsInput() *MssqlVirtualMachineStorageConfigurationTempDbSettings
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
	PutDataSettings(value *MssqlVirtualMachineStorageConfigurationDataSettings)
	PutLogSettings(value *MssqlVirtualMachineStorageConfigurationLogSettings)
	PutTempDbSettings(value *MssqlVirtualMachineStorageConfigurationTempDbSettings)
	ResetDataSettings()
	ResetLogSettings()
	ResetSystemDbOnDataDiskEnabled()
	ResetTempDbSettings()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for MssqlVirtualMachineStorageConfigurationOutputReference
type jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) DataSettings() MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference {
	var returns MssqlVirtualMachineStorageConfigurationDataSettingsOutputReference
	_jsii_.Get(
		j,
		"dataSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) DataSettingsInput() *MssqlVirtualMachineStorageConfigurationDataSettings {
	var returns *MssqlVirtualMachineStorageConfigurationDataSettings
	_jsii_.Get(
		j,
		"dataSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) DiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) DiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) InternalValue() *MssqlVirtualMachineStorageConfiguration {
	var returns *MssqlVirtualMachineStorageConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) LogSettings() MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference {
	var returns MssqlVirtualMachineStorageConfigurationLogSettingsOutputReference
	_jsii_.Get(
		j,
		"logSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) LogSettingsInput() *MssqlVirtualMachineStorageConfigurationLogSettings {
	var returns *MssqlVirtualMachineStorageConfigurationLogSettings
	_jsii_.Get(
		j,
		"logSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) StorageWorkloadType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageWorkloadType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) StorageWorkloadTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageWorkloadTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) SystemDbOnDataDiskEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"systemDbOnDataDiskEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) SystemDbOnDataDiskEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"systemDbOnDataDiskEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) TempDbSettings() MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference {
	var returns MssqlVirtualMachineStorageConfigurationTempDbSettingsOutputReference
	_jsii_.Get(
		j,
		"tempDbSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) TempDbSettingsInput() *MssqlVirtualMachineStorageConfigurationTempDbSettings {
	var returns *MssqlVirtualMachineStorageConfigurationTempDbSettings
	_jsii_.Get(
		j,
		"tempDbSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewMssqlVirtualMachineStorageConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) MssqlVirtualMachineStorageConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewMssqlVirtualMachineStorageConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineStorageConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewMssqlVirtualMachineStorageConfigurationOutputReference_Override(m MssqlVirtualMachineStorageConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.mssqlVirtualMachine.MssqlVirtualMachineStorageConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		m,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference)SetDiskType(val *string) {
	if err := j.validateSetDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskType",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference)SetInternalValue(val *MssqlVirtualMachineStorageConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference)SetStorageWorkloadType(val *string) {
	if err := j.validateSetStorageWorkloadTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageWorkloadType",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference)SetSystemDbOnDataDiskEnabled(val interface{}) {
	if err := j.validateSetSystemDbOnDataDiskEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"systemDbOnDataDiskEnabled",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) PutDataSettings(value *MssqlVirtualMachineStorageConfigurationDataSettings) {
	if err := m.validatePutDataSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putDataSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) PutLogSettings(value *MssqlVirtualMachineStorageConfigurationLogSettings) {
	if err := m.validatePutLogSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putLogSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) PutTempDbSettings(value *MssqlVirtualMachineStorageConfigurationTempDbSettings) {
	if err := m.validatePutTempDbSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		m,
		"putTempDbSettings",
		[]interface{}{value},
	)
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) ResetDataSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetDataSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) ResetLogSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetLogSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) ResetSystemDbOnDataDiskEnabled() {
	_jsii_.InvokeVoid(
		m,
		"resetSystemDbOnDataDiskEnabled",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) ResetTempDbSettings() {
	_jsii_.InvokeVoid(
		m,
		"resetTempDbSettings",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (m *jsiiProxy_MssqlVirtualMachineStorageConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


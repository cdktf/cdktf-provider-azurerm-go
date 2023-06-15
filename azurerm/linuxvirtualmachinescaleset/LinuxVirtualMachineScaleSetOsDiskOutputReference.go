package linuxvirtualmachinescaleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/linuxvirtualmachinescaleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type LinuxVirtualMachineScaleSetOsDiskOutputReference interface {
	cdktf.ComplexObject
	Caching() *string
	SetCaching(val *string)
	CachingInput() *string
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
	DiffDiskSettings() LinuxVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference
	DiffDiskSettingsInput() *LinuxVirtualMachineScaleSetOsDiskDiffDiskSettings
	DiskEncryptionSetId() *string
	SetDiskEncryptionSetId(val *string)
	DiskEncryptionSetIdInput() *string
	DiskSizeGb() *float64
	SetDiskSizeGb(val *float64)
	DiskSizeGbInput() *float64
	// Experimental.
	Fqn() *string
	InternalValue() *LinuxVirtualMachineScaleSetOsDisk
	SetInternalValue(val *LinuxVirtualMachineScaleSetOsDisk)
	SecureVmDiskEncryptionSetId() *string
	SetSecureVmDiskEncryptionSetId(val *string)
	SecureVmDiskEncryptionSetIdInput() *string
	SecurityEncryptionType() *string
	SetSecurityEncryptionType(val *string)
	SecurityEncryptionTypeInput() *string
	StorageAccountType() *string
	SetStorageAccountType(val *string)
	StorageAccountTypeInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WriteAcceleratorEnabled() interface{}
	SetWriteAcceleratorEnabled(val interface{})
	WriteAcceleratorEnabledInput() interface{}
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
	PutDiffDiskSettings(value *LinuxVirtualMachineScaleSetOsDiskDiffDiskSettings)
	ResetDiffDiskSettings()
	ResetDiskEncryptionSetId()
	ResetDiskSizeGb()
	ResetSecureVmDiskEncryptionSetId()
	ResetSecurityEncryptionType()
	ResetWriteAcceleratorEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LinuxVirtualMachineScaleSetOsDiskOutputReference
type jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) Caching() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caching",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) CachingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cachingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) DiffDiskSettings() LinuxVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference {
	var returns LinuxVirtualMachineScaleSetOsDiskDiffDiskSettingsOutputReference
	_jsii_.Get(
		j,
		"diffDiskSettings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) DiffDiskSettingsInput() *LinuxVirtualMachineScaleSetOsDiskDiffDiskSettings {
	var returns *LinuxVirtualMachineScaleSetOsDiskDiffDiskSettings
	_jsii_.Get(
		j,
		"diffDiskSettingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) DiskEncryptionSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) DiskEncryptionSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskEncryptionSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) DiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) DiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) InternalValue() *LinuxVirtualMachineScaleSetOsDisk {
	var returns *LinuxVirtualMachineScaleSetOsDisk
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) SecureVmDiskEncryptionSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secureVmDiskEncryptionSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) SecureVmDiskEncryptionSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secureVmDiskEncryptionSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) SecurityEncryptionType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityEncryptionType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) SecurityEncryptionTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"securityEncryptionTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) StorageAccountType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) StorageAccountTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) WriteAcceleratorEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeAcceleratorEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) WriteAcceleratorEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"writeAcceleratorEnabledInput",
		&returns,
	)
	return returns
}


func NewLinuxVirtualMachineScaleSetOsDiskOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LinuxVirtualMachineScaleSetOsDiskOutputReference {
	_init_.Initialize()

	if err := validateNewLinuxVirtualMachineScaleSetOsDiskOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxVirtualMachineScaleSet.LinuxVirtualMachineScaleSetOsDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLinuxVirtualMachineScaleSetOsDiskOutputReference_Override(l LinuxVirtualMachineScaleSetOsDiskOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxVirtualMachineScaleSet.LinuxVirtualMachineScaleSetOsDiskOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference)SetCaching(val *string) {
	if err := j.validateSetCachingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caching",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference)SetDiskEncryptionSetId(val *string) {
	if err := j.validateSetDiskEncryptionSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskEncryptionSetId",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference)SetDiskSizeGb(val *float64) {
	if err := j.validateSetDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskSizeGb",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference)SetInternalValue(val *LinuxVirtualMachineScaleSetOsDisk) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference)SetSecureVmDiskEncryptionSetId(val *string) {
	if err := j.validateSetSecureVmDiskEncryptionSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secureVmDiskEncryptionSetId",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference)SetSecurityEncryptionType(val *string) {
	if err := j.validateSetSecurityEncryptionTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"securityEncryptionType",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference)SetStorageAccountType(val *string) {
	if err := j.validateSetStorageAccountTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountType",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference)SetWriteAcceleratorEnabled(val interface{}) {
	if err := j.validateSetWriteAcceleratorEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeAcceleratorEnabled",
		val,
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (l *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (l *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (l *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (l *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (l *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (l *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (l *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (l *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (l *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) PutDiffDiskSettings(value *LinuxVirtualMachineScaleSetOsDiskDiffDiskSettings) {
	if err := l.validatePutDiffDiskSettingsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDiffDiskSettings",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) ResetDiffDiskSettings() {
	_jsii_.InvokeVoid(
		l,
		"resetDiffDiskSettings",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) ResetDiskEncryptionSetId() {
	_jsii_.InvokeVoid(
		l,
		"resetDiskEncryptionSetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) ResetDiskSizeGb() {
	_jsii_.InvokeVoid(
		l,
		"resetDiskSizeGb",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) ResetSecureVmDiskEncryptionSetId() {
	_jsii_.InvokeVoid(
		l,
		"resetSecureVmDiskEncryptionSetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) ResetSecurityEncryptionType() {
	_jsii_.InvokeVoid(
		l,
		"resetSecurityEncryptionType",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) ResetWriteAcceleratorEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetWriteAcceleratorEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (l *jsiiProxy_LinuxVirtualMachineScaleSetOsDiskOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


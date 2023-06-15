package orchestratedvirtualmachinescaleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/orchestratedvirtualmachinescaleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference interface {
	cdktf.ComplexObject
	AdminPassword() *string
	SetAdminPassword(val *string)
	AdminPasswordInput() *string
	AdminSshKey() OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationAdminSshKeyList
	AdminSshKeyInput() interface{}
	AdminUsername() *string
	SetAdminUsername(val *string)
	AdminUsernameInput() *string
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
	ComputerNamePrefix() *string
	SetComputerNamePrefix(val *string)
	ComputerNamePrefixInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	DisablePasswordAuthentication() interface{}
	SetDisablePasswordAuthentication(val interface{})
	DisablePasswordAuthenticationInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *OrchestratedVirtualMachineScaleSetOsProfileLinuxConfiguration
	SetInternalValue(val *OrchestratedVirtualMachineScaleSetOsProfileLinuxConfiguration)
	PatchAssessmentMode() *string
	SetPatchAssessmentMode(val *string)
	PatchAssessmentModeInput() *string
	PatchMode() *string
	SetPatchMode(val *string)
	PatchModeInput() *string
	ProvisionVmAgent() interface{}
	SetProvisionVmAgent(val interface{})
	ProvisionVmAgentInput() interface{}
	Secret() OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationSecretList
	SecretInput() interface{}
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
	PutAdminSshKey(value interface{})
	PutSecret(value interface{})
	ResetAdminPassword()
	ResetAdminSshKey()
	ResetComputerNamePrefix()
	ResetDisablePasswordAuthentication()
	ResetPatchAssessmentMode()
	ResetPatchMode()
	ResetProvisionVmAgent()
	ResetSecret()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference
type jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) AdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) AdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) AdminSshKey() OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationAdminSshKeyList {
	var returns OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationAdminSshKeyList
	_jsii_.Get(
		j,
		"adminSshKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) AdminSshKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminSshKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) AdminUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) AdminUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) ComputerNamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerNamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) ComputerNamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerNamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) DisablePasswordAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePasswordAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) DisablePasswordAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePasswordAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) InternalValue() *OrchestratedVirtualMachineScaleSetOsProfileLinuxConfiguration {
	var returns *OrchestratedVirtualMachineScaleSetOsProfileLinuxConfiguration
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) PatchAssessmentMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchAssessmentMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) PatchAssessmentModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchAssessmentModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) PatchMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) PatchModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) ProvisionVmAgent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionVmAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) ProvisionVmAgentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionVmAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) Secret() OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationSecretList {
	var returns OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationSecretList
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference {
	_init_.Initialize()

	if err := validateNewOrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference_Override(o OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.orchestratedVirtualMachineScaleSet.OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference)SetAdminPassword(val *string) {
	if err := j.validateSetAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminPassword",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference)SetAdminUsername(val *string) {
	if err := j.validateSetAdminUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminUsername",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference)SetComputerNamePrefix(val *string) {
	if err := j.validateSetComputerNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computerNamePrefix",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference)SetDisablePasswordAuthentication(val interface{}) {
	if err := j.validateSetDisablePasswordAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disablePasswordAuthentication",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference)SetInternalValue(val *OrchestratedVirtualMachineScaleSetOsProfileLinuxConfiguration) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference)SetPatchAssessmentMode(val *string) {
	if err := j.validateSetPatchAssessmentModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patchAssessmentMode",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference)SetPatchMode(val *string) {
	if err := j.validateSetPatchModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patchMode",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference)SetProvisionVmAgent(val interface{}) {
	if err := j.validateSetProvisionVmAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionVmAgent",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) PutAdminSshKey(value interface{}) {
	if err := o.validatePutAdminSshKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putAdminSshKey",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) PutSecret(value interface{}) {
	if err := o.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		o,
		"putSecret",
		[]interface{}{value},
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) ResetAdminPassword() {
	_jsii_.InvokeVoid(
		o,
		"resetAdminPassword",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) ResetAdminSshKey() {
	_jsii_.InvokeVoid(
		o,
		"resetAdminSshKey",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) ResetComputerNamePrefix() {
	_jsii_.InvokeVoid(
		o,
		"resetComputerNamePrefix",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) ResetDisablePasswordAuthentication() {
	_jsii_.InvokeVoid(
		o,
		"resetDisablePasswordAuthentication",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) ResetPatchAssessmentMode() {
	_jsii_.InvokeVoid(
		o,
		"resetPatchAssessmentMode",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) ResetPatchMode() {
	_jsii_.InvokeVoid(
		o,
		"resetPatchMode",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) ResetProvisionVmAgent() {
	_jsii_.InvokeVoid(
		o,
		"resetProvisionVmAgent",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) ResetSecret() {
	_jsii_.InvokeVoid(
		o,
		"resetSecret",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OrchestratedVirtualMachineScaleSetOsProfileLinuxConfigurationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


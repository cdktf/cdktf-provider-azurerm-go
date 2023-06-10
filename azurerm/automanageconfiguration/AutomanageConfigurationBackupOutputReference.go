package automanageconfiguration

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/automanageconfiguration/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type AutomanageConfigurationBackupOutputReference interface {
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
	InstantRpRetentionRangeInDays() *float64
	SetInstantRpRetentionRangeInDays(val *float64)
	InstantRpRetentionRangeInDaysInput() *float64
	InternalValue() *AutomanageConfigurationBackup
	SetInternalValue(val *AutomanageConfigurationBackup)
	PolicyName() *string
	SetPolicyName(val *string)
	PolicyNameInput() *string
	RetentionPolicy() AutomanageConfigurationBackupRetentionPolicyOutputReference
	RetentionPolicyInput() *AutomanageConfigurationBackupRetentionPolicy
	SchedulePolicy() AutomanageConfigurationBackupSchedulePolicyOutputReference
	SchedulePolicyInput() *AutomanageConfigurationBackupSchedulePolicy
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TimeZone() *string
	SetTimeZone(val *string)
	TimeZoneInput() *string
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
	PutRetentionPolicy(value *AutomanageConfigurationBackupRetentionPolicy)
	PutSchedulePolicy(value *AutomanageConfigurationBackupSchedulePolicy)
	ResetInstantRpRetentionRangeInDays()
	ResetPolicyName()
	ResetRetentionPolicy()
	ResetSchedulePolicy()
	ResetTimeZone()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for AutomanageConfigurationBackupOutputReference
type jsiiProxy_AutomanageConfigurationBackupOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_AutomanageConfigurationBackupOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupOutputReference) InstantRpRetentionRangeInDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instantRpRetentionRangeInDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupOutputReference) InstantRpRetentionRangeInDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instantRpRetentionRangeInDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupOutputReference) InternalValue() *AutomanageConfigurationBackup {
	var returns *AutomanageConfigurationBackup
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupOutputReference) PolicyName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupOutputReference) PolicyNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupOutputReference) RetentionPolicy() AutomanageConfigurationBackupRetentionPolicyOutputReference {
	var returns AutomanageConfigurationBackupRetentionPolicyOutputReference
	_jsii_.Get(
		j,
		"retentionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupOutputReference) RetentionPolicyInput() *AutomanageConfigurationBackupRetentionPolicy {
	var returns *AutomanageConfigurationBackupRetentionPolicy
	_jsii_.Get(
		j,
		"retentionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupOutputReference) SchedulePolicy() AutomanageConfigurationBackupSchedulePolicyOutputReference {
	var returns AutomanageConfigurationBackupSchedulePolicyOutputReference
	_jsii_.Get(
		j,
		"schedulePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupOutputReference) SchedulePolicyInput() *AutomanageConfigurationBackupSchedulePolicy {
	var returns *AutomanageConfigurationBackupSchedulePolicy
	_jsii_.Get(
		j,
		"schedulePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupOutputReference) TimeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AutomanageConfigurationBackupOutputReference) TimeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timeZoneInput",
		&returns,
	)
	return returns
}


func NewAutomanageConfigurationBackupOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) AutomanageConfigurationBackupOutputReference {
	_init_.Initialize()

	if err := validateNewAutomanageConfigurationBackupOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_AutomanageConfigurationBackupOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.automanageConfiguration.AutomanageConfigurationBackupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewAutomanageConfigurationBackupOutputReference_Override(a AutomanageConfigurationBackupOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.automanageConfiguration.AutomanageConfigurationBackupOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		a,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupOutputReference)SetInstantRpRetentionRangeInDays(val *float64) {
	if err := j.validateSetInstantRpRetentionRangeInDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instantRpRetentionRangeInDays",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupOutputReference)SetInternalValue(val *AutomanageConfigurationBackup) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupOutputReference)SetPolicyName(val *string) {
	if err := j.validateSetPolicyNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyName",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_AutomanageConfigurationBackupOutputReference)SetTimeZone(val *string) {
	if err := j.validateSetTimeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timeZone",
		val,
	)
}

func (a *jsiiProxy_AutomanageConfigurationBackupOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfigurationBackupOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfigurationBackupOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfigurationBackupOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfigurationBackupOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfigurationBackupOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfigurationBackupOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfigurationBackupOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfigurationBackupOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfigurationBackupOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfigurationBackupOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfigurationBackupOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfigurationBackupOutputReference) PutRetentionPolicy(value *AutomanageConfigurationBackupRetentionPolicy) {
	if err := a.validatePutRetentionPolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putRetentionPolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomanageConfigurationBackupOutputReference) PutSchedulePolicy(value *AutomanageConfigurationBackupSchedulePolicy) {
	if err := a.validatePutSchedulePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"putSchedulePolicy",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AutomanageConfigurationBackupOutputReference) ResetInstantRpRetentionRangeInDays() {
	_jsii_.InvokeVoid(
		a,
		"resetInstantRpRetentionRangeInDays",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfigurationBackupOutputReference) ResetPolicyName() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicyName",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfigurationBackupOutputReference) ResetRetentionPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetRetentionPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfigurationBackupOutputReference) ResetSchedulePolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetSchedulePolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfigurationBackupOutputReference) ResetTimeZone() {
	_jsii_.InvokeVoid(
		a,
		"resetTimeZone",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AutomanageConfigurationBackupOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := a.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		a,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AutomanageConfigurationBackupOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


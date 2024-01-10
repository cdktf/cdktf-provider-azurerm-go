// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package backuppolicyvmworkload

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/backuppolicyvmworkload/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BackupPolicyVmWorkloadProtectionPolicyOutputReference interface {
	cdktf.ComplexObject
	Backup() BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference
	BackupInput() *BackupPolicyVmWorkloadProtectionPolicyBackup
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
	InternalValue() interface{}
	SetInternalValue(val interface{})
	PolicyType() *string
	SetPolicyType(val *string)
	PolicyTypeInput() *string
	RetentionDaily() BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference
	RetentionDailyInput() *BackupPolicyVmWorkloadProtectionPolicyRetentionDaily
	RetentionMonthly() BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference
	RetentionMonthlyInput() *BackupPolicyVmWorkloadProtectionPolicyRetentionMonthly
	RetentionWeekly() BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference
	RetentionWeeklyInput() *BackupPolicyVmWorkloadProtectionPolicyRetentionWeekly
	RetentionYearly() BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference
	RetentionYearlyInput() *BackupPolicyVmWorkloadProtectionPolicyRetentionYearly
	SimpleRetention() BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference
	SimpleRetentionInput() *BackupPolicyVmWorkloadProtectionPolicySimpleRetention
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
	PutBackup(value *BackupPolicyVmWorkloadProtectionPolicyBackup)
	PutRetentionDaily(value *BackupPolicyVmWorkloadProtectionPolicyRetentionDaily)
	PutRetentionMonthly(value *BackupPolicyVmWorkloadProtectionPolicyRetentionMonthly)
	PutRetentionWeekly(value *BackupPolicyVmWorkloadProtectionPolicyRetentionWeekly)
	PutRetentionYearly(value *BackupPolicyVmWorkloadProtectionPolicyRetentionYearly)
	PutSimpleRetention(value *BackupPolicyVmWorkloadProtectionPolicySimpleRetention)
	ResetRetentionDaily()
	ResetRetentionMonthly()
	ResetRetentionWeekly()
	ResetRetentionYearly()
	ResetSimpleRetention()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for BackupPolicyVmWorkloadProtectionPolicyOutputReference
type jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) Backup() BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference {
	var returns BackupPolicyVmWorkloadProtectionPolicyBackupOutputReference
	_jsii_.Get(
		j,
		"backup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) BackupInput() *BackupPolicyVmWorkloadProtectionPolicyBackup {
	var returns *BackupPolicyVmWorkloadProtectionPolicyBackup
	_jsii_.Get(
		j,
		"backupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) PolicyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) PolicyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) RetentionDaily() BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference {
	var returns BackupPolicyVmWorkloadProtectionPolicyRetentionDailyOutputReference
	_jsii_.Get(
		j,
		"retentionDaily",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) RetentionDailyInput() *BackupPolicyVmWorkloadProtectionPolicyRetentionDaily {
	var returns *BackupPolicyVmWorkloadProtectionPolicyRetentionDaily
	_jsii_.Get(
		j,
		"retentionDailyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) RetentionMonthly() BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference {
	var returns BackupPolicyVmWorkloadProtectionPolicyRetentionMonthlyOutputReference
	_jsii_.Get(
		j,
		"retentionMonthly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) RetentionMonthlyInput() *BackupPolicyVmWorkloadProtectionPolicyRetentionMonthly {
	var returns *BackupPolicyVmWorkloadProtectionPolicyRetentionMonthly
	_jsii_.Get(
		j,
		"retentionMonthlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) RetentionWeekly() BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference {
	var returns BackupPolicyVmWorkloadProtectionPolicyRetentionWeeklyOutputReference
	_jsii_.Get(
		j,
		"retentionWeekly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) RetentionWeeklyInput() *BackupPolicyVmWorkloadProtectionPolicyRetentionWeekly {
	var returns *BackupPolicyVmWorkloadProtectionPolicyRetentionWeekly
	_jsii_.Get(
		j,
		"retentionWeeklyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) RetentionYearly() BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference {
	var returns BackupPolicyVmWorkloadProtectionPolicyRetentionYearlyOutputReference
	_jsii_.Get(
		j,
		"retentionYearly",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) RetentionYearlyInput() *BackupPolicyVmWorkloadProtectionPolicyRetentionYearly {
	var returns *BackupPolicyVmWorkloadProtectionPolicyRetentionYearly
	_jsii_.Get(
		j,
		"retentionYearlyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) SimpleRetention() BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference {
	var returns BackupPolicyVmWorkloadProtectionPolicySimpleRetentionOutputReference
	_jsii_.Get(
		j,
		"simpleRetention",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) SimpleRetentionInput() *BackupPolicyVmWorkloadProtectionPolicySimpleRetention {
	var returns *BackupPolicyVmWorkloadProtectionPolicySimpleRetention
	_jsii_.Get(
		j,
		"simpleRetentionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewBackupPolicyVmWorkloadProtectionPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) BackupPolicyVmWorkloadProtectionPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewBackupPolicyVmWorkloadProtectionPolicyOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyVmWorkload.BackupPolicyVmWorkloadProtectionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewBackupPolicyVmWorkloadProtectionPolicyOutputReference_Override(b BackupPolicyVmWorkloadProtectionPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.backupPolicyVmWorkload.BackupPolicyVmWorkloadProtectionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		b,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference)SetPolicyType(val *string) {
	if err := j.validateSetPolicyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyType",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := b.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		b,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := b.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := b.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		b,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := b.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		b,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := b.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		b,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := b.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		b,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := b.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		b,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := b.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		b,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := b.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		b,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := b.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		b,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) PutBackup(value *BackupPolicyVmWorkloadProtectionPolicyBackup) {
	if err := b.validatePutBackupParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putBackup",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) PutRetentionDaily(value *BackupPolicyVmWorkloadProtectionPolicyRetentionDaily) {
	if err := b.validatePutRetentionDailyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRetentionDaily",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) PutRetentionMonthly(value *BackupPolicyVmWorkloadProtectionPolicyRetentionMonthly) {
	if err := b.validatePutRetentionMonthlyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRetentionMonthly",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) PutRetentionWeekly(value *BackupPolicyVmWorkloadProtectionPolicyRetentionWeekly) {
	if err := b.validatePutRetentionWeeklyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRetentionWeekly",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) PutRetentionYearly(value *BackupPolicyVmWorkloadProtectionPolicyRetentionYearly) {
	if err := b.validatePutRetentionYearlyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putRetentionYearly",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) PutSimpleRetention(value *BackupPolicyVmWorkloadProtectionPolicySimpleRetention) {
	if err := b.validatePutSimpleRetentionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		b,
		"putSimpleRetention",
		[]interface{}{value},
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) ResetRetentionDaily() {
	_jsii_.InvokeVoid(
		b,
		"resetRetentionDaily",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) ResetRetentionMonthly() {
	_jsii_.InvokeVoid(
		b,
		"resetRetentionMonthly",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) ResetRetentionWeekly() {
	_jsii_.InvokeVoid(
		b,
		"resetRetentionWeekly",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) ResetRetentionYearly() {
	_jsii_.InvokeVoid(
		b,
		"resetRetentionYearly",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) ResetSimpleRetention() {
	_jsii_.InvokeVoid(
		b,
		"resetSimpleRetention",
		nil, // no parameters
	)
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := b.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		b,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (b *jsiiProxy_BackupPolicyVmWorkloadProtectionPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		b,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package virtualdesktopscalingplan

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/virtualdesktopscalingplan/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type VirtualDesktopScalingPlanScheduleOutputReference interface {
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
	DaysOfWeek() *[]*string
	SetDaysOfWeek(val *[]*string)
	DaysOfWeekInput() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	Name() *string
	SetName(val *string)
	NameInput() *string
	OffPeakLoadBalancingAlgorithm() *string
	SetOffPeakLoadBalancingAlgorithm(val *string)
	OffPeakLoadBalancingAlgorithmInput() *string
	OffPeakStartTime() *string
	SetOffPeakStartTime(val *string)
	OffPeakStartTimeInput() *string
	PeakLoadBalancingAlgorithm() *string
	SetPeakLoadBalancingAlgorithm(val *string)
	PeakLoadBalancingAlgorithmInput() *string
	PeakStartTime() *string
	SetPeakStartTime(val *string)
	PeakStartTimeInput() *string
	RampDownCapacityThresholdPercent() *float64
	SetRampDownCapacityThresholdPercent(val *float64)
	RampDownCapacityThresholdPercentInput() *float64
	RampDownForceLogoffUsers() interface{}
	SetRampDownForceLogoffUsers(val interface{})
	RampDownForceLogoffUsersInput() interface{}
	RampDownLoadBalancingAlgorithm() *string
	SetRampDownLoadBalancingAlgorithm(val *string)
	RampDownLoadBalancingAlgorithmInput() *string
	RampDownMinimumHostsPercent() *float64
	SetRampDownMinimumHostsPercent(val *float64)
	RampDownMinimumHostsPercentInput() *float64
	RampDownNotificationMessage() *string
	SetRampDownNotificationMessage(val *string)
	RampDownNotificationMessageInput() *string
	RampDownStartTime() *string
	SetRampDownStartTime(val *string)
	RampDownStartTimeInput() *string
	RampDownStopHostsWhen() *string
	SetRampDownStopHostsWhen(val *string)
	RampDownStopHostsWhenInput() *string
	RampDownWaitTimeMinutes() *float64
	SetRampDownWaitTimeMinutes(val *float64)
	RampDownWaitTimeMinutesInput() *float64
	RampUpCapacityThresholdPercent() *float64
	SetRampUpCapacityThresholdPercent(val *float64)
	RampUpCapacityThresholdPercentInput() *float64
	RampUpLoadBalancingAlgorithm() *string
	SetRampUpLoadBalancingAlgorithm(val *string)
	RampUpLoadBalancingAlgorithmInput() *string
	RampUpMinimumHostsPercent() *float64
	SetRampUpMinimumHostsPercent(val *float64)
	RampUpMinimumHostsPercentInput() *float64
	RampUpStartTime() *string
	SetRampUpStartTime(val *string)
	RampUpStartTimeInput() *string
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
	ResetRampUpCapacityThresholdPercent()
	ResetRampUpMinimumHostsPercent()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for VirtualDesktopScalingPlanScheduleOutputReference
type jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) DaysOfWeek() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfWeek",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) DaysOfWeekInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"daysOfWeekInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) OffPeakLoadBalancingAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"offPeakLoadBalancingAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) OffPeakLoadBalancingAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"offPeakLoadBalancingAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) OffPeakStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"offPeakStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) OffPeakStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"offPeakStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) PeakLoadBalancingAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peakLoadBalancingAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) PeakLoadBalancingAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peakLoadBalancingAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) PeakStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peakStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) PeakStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"peakStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownCapacityThresholdPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rampDownCapacityThresholdPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownCapacityThresholdPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rampDownCapacityThresholdPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownForceLogoffUsers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rampDownForceLogoffUsers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownForceLogoffUsersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rampDownForceLogoffUsersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownLoadBalancingAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rampDownLoadBalancingAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownLoadBalancingAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rampDownLoadBalancingAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownMinimumHostsPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rampDownMinimumHostsPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownMinimumHostsPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rampDownMinimumHostsPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownNotificationMessage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rampDownNotificationMessage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownNotificationMessageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rampDownNotificationMessageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rampDownStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rampDownStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownStopHostsWhen() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rampDownStopHostsWhen",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownStopHostsWhenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rampDownStopHostsWhenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownWaitTimeMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rampDownWaitTimeMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampDownWaitTimeMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rampDownWaitTimeMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampUpCapacityThresholdPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rampUpCapacityThresholdPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampUpCapacityThresholdPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rampUpCapacityThresholdPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampUpLoadBalancingAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rampUpLoadBalancingAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampUpLoadBalancingAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rampUpLoadBalancingAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampUpMinimumHostsPercent() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rampUpMinimumHostsPercent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampUpMinimumHostsPercentInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rampUpMinimumHostsPercentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampUpStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rampUpStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) RampUpStartTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rampUpStartTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewVirtualDesktopScalingPlanScheduleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) VirtualDesktopScalingPlanScheduleOutputReference {
	_init_.Initialize()

	if err := validateNewVirtualDesktopScalingPlanScheduleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualDesktopScalingPlan.VirtualDesktopScalingPlanScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewVirtualDesktopScalingPlanScheduleOutputReference_Override(v VirtualDesktopScalingPlanScheduleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.virtualDesktopScalingPlan.VirtualDesktopScalingPlanScheduleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		v,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference)SetDaysOfWeek(val *[]*string) {
	if err := j.validateSetDaysOfWeekParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"daysOfWeek",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference)SetOffPeakLoadBalancingAlgorithm(val *string) {
	if err := j.validateSetOffPeakLoadBalancingAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"offPeakLoadBalancingAlgorithm",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference)SetOffPeakStartTime(val *string) {
	if err := j.validateSetOffPeakStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"offPeakStartTime",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference)SetPeakLoadBalancingAlgorithm(val *string) {
	if err := j.validateSetPeakLoadBalancingAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peakLoadBalancingAlgorithm",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference)SetPeakStartTime(val *string) {
	if err := j.validateSetPeakStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"peakStartTime",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference)SetRampDownCapacityThresholdPercent(val *float64) {
	if err := j.validateSetRampDownCapacityThresholdPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rampDownCapacityThresholdPercent",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference)SetRampDownForceLogoffUsers(val interface{}) {
	if err := j.validateSetRampDownForceLogoffUsersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rampDownForceLogoffUsers",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference)SetRampDownLoadBalancingAlgorithm(val *string) {
	if err := j.validateSetRampDownLoadBalancingAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rampDownLoadBalancingAlgorithm",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference)SetRampDownMinimumHostsPercent(val *float64) {
	if err := j.validateSetRampDownMinimumHostsPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rampDownMinimumHostsPercent",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference)SetRampDownNotificationMessage(val *string) {
	if err := j.validateSetRampDownNotificationMessageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rampDownNotificationMessage",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference)SetRampDownStartTime(val *string) {
	if err := j.validateSetRampDownStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rampDownStartTime",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference)SetRampDownStopHostsWhen(val *string) {
	if err := j.validateSetRampDownStopHostsWhenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rampDownStopHostsWhen",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference)SetRampDownWaitTimeMinutes(val *float64) {
	if err := j.validateSetRampDownWaitTimeMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rampDownWaitTimeMinutes",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference)SetRampUpCapacityThresholdPercent(val *float64) {
	if err := j.validateSetRampUpCapacityThresholdPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rampUpCapacityThresholdPercent",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference)SetRampUpLoadBalancingAlgorithm(val *string) {
	if err := j.validateSetRampUpLoadBalancingAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rampUpLoadBalancingAlgorithm",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference)SetRampUpMinimumHostsPercent(val *float64) {
	if err := j.validateSetRampUpMinimumHostsPercentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rampUpMinimumHostsPercent",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference)SetRampUpStartTime(val *string) {
	if err := j.validateSetRampUpStartTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rampUpStartTime",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := v.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := v.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := v.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		v,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := v.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		v,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := v.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		v,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := v.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		v,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := v.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		v,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := v.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		v,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := v.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		v,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := v.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		v,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) ResetRampUpCapacityThresholdPercent() {
	_jsii_.InvokeVoid(
		v,
		"resetRampUpCapacityThresholdPercent",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) ResetRampUpMinimumHostsPercent() {
	_jsii_.InvokeVoid(
		v,
		"resetRampUpMinimumHostsPercent",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := v.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		v,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VirtualDesktopScalingPlanScheduleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


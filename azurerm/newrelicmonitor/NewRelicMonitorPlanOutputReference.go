package newrelicmonitor

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/newrelicmonitor/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NewRelicMonitorPlanOutputReference interface {
	cdktf.ComplexObject
	BillingCycle() *string
	SetBillingCycle(val *string)
	BillingCycleInput() *string
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
	EffectiveDate() *string
	SetEffectiveDate(val *string)
	EffectiveDateInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *NewRelicMonitorPlan
	SetInternalValue(val *NewRelicMonitorPlan)
	PlanId() *string
	SetPlanId(val *string)
	PlanIdInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UsageType() *string
	SetUsageType(val *string)
	UsageTypeInput() *string
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
	ResetBillingCycle()
	ResetPlanId()
	ResetUsageType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NewRelicMonitorPlanOutputReference
type jsiiProxy_NewRelicMonitorPlanOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NewRelicMonitorPlanOutputReference) BillingCycle() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingCycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicMonitorPlanOutputReference) BillingCycleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"billingCycleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicMonitorPlanOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicMonitorPlanOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicMonitorPlanOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicMonitorPlanOutputReference) EffectiveDate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveDate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicMonitorPlanOutputReference) EffectiveDateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"effectiveDateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicMonitorPlanOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicMonitorPlanOutputReference) InternalValue() *NewRelicMonitorPlan {
	var returns *NewRelicMonitorPlan
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicMonitorPlanOutputReference) PlanId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"planId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicMonitorPlanOutputReference) PlanIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"planIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicMonitorPlanOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicMonitorPlanOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicMonitorPlanOutputReference) UsageType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NewRelicMonitorPlanOutputReference) UsageTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageTypeInput",
		&returns,
	)
	return returns
}


func NewNewRelicMonitorPlanOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NewRelicMonitorPlanOutputReference {
	_init_.Initialize()

	if err := validateNewNewRelicMonitorPlanOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NewRelicMonitorPlanOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.newRelicMonitor.NewRelicMonitorPlanOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNewRelicMonitorPlanOutputReference_Override(n NewRelicMonitorPlanOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.newRelicMonitor.NewRelicMonitorPlanOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NewRelicMonitorPlanOutputReference)SetBillingCycle(val *string) {
	if err := j.validateSetBillingCycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"billingCycle",
		val,
	)
}

func (j *jsiiProxy_NewRelicMonitorPlanOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NewRelicMonitorPlanOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NewRelicMonitorPlanOutputReference)SetEffectiveDate(val *string) {
	if err := j.validateSetEffectiveDateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"effectiveDate",
		val,
	)
}

func (j *jsiiProxy_NewRelicMonitorPlanOutputReference)SetInternalValue(val *NewRelicMonitorPlan) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NewRelicMonitorPlanOutputReference)SetPlanId(val *string) {
	if err := j.validateSetPlanIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"planId",
		val,
	)
}

func (j *jsiiProxy_NewRelicMonitorPlanOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NewRelicMonitorPlanOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NewRelicMonitorPlanOutputReference)SetUsageType(val *string) {
	if err := j.validateSetUsageTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usageType",
		val,
	)
}

func (n *jsiiProxy_NewRelicMonitorPlanOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicMonitorPlanOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := n.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		n,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicMonitorPlanOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := n.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicMonitorPlanOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := n.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		n,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicMonitorPlanOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := n.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		n,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicMonitorPlanOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := n.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		n,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicMonitorPlanOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := n.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		n,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicMonitorPlanOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := n.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		n,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicMonitorPlanOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := n.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		n,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicMonitorPlanOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := n.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		n,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicMonitorPlanOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicMonitorPlanOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := n.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicMonitorPlanOutputReference) ResetBillingCycle() {
	_jsii_.InvokeVoid(
		n,
		"resetBillingCycle",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NewRelicMonitorPlanOutputReference) ResetPlanId() {
	_jsii_.InvokeVoid(
		n,
		"resetPlanId",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NewRelicMonitorPlanOutputReference) ResetUsageType() {
	_jsii_.InvokeVoid(
		n,
		"resetUsageType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NewRelicMonitorPlanOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := n.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		n,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NewRelicMonitorPlanOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


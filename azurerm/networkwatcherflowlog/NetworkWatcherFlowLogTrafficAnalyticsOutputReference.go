package networkwatcherflowlog

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v6/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v6/networkwatcherflowlog/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetworkWatcherFlowLogTrafficAnalyticsOutputReference interface {
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
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	// Experimental.
	Fqn() *string
	InternalValue() *NetworkWatcherFlowLogTrafficAnalytics
	SetInternalValue(val *NetworkWatcherFlowLogTrafficAnalytics)
	IntervalInMinutes() *float64
	SetIntervalInMinutes(val *float64)
	IntervalInMinutesInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WorkspaceId() *string
	SetWorkspaceId(val *string)
	WorkspaceIdInput() *string
	WorkspaceRegion() *string
	SetWorkspaceRegion(val *string)
	WorkspaceRegionInput() *string
	WorkspaceResourceId() *string
	SetWorkspaceResourceId(val *string)
	WorkspaceResourceIdInput() *string
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
	ResetIntervalInMinutes()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetworkWatcherFlowLogTrafficAnalyticsOutputReference
type jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) InternalValue() *NetworkWatcherFlowLogTrafficAnalytics {
	var returns *NetworkWatcherFlowLogTrafficAnalytics
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) IntervalInMinutes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInMinutes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) IntervalInMinutesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"intervalInMinutesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) WorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) WorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) WorkspaceRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) WorkspaceRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) WorkspaceResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) WorkspaceResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workspaceResourceIdInput",
		&returns,
	)
	return returns
}


func NewNetworkWatcherFlowLogTrafficAnalyticsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetworkWatcherFlowLogTrafficAnalyticsOutputReference {
	_init_.Initialize()

	if err := validateNewNetworkWatcherFlowLogTrafficAnalyticsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.networkWatcherFlowLog.NetworkWatcherFlowLogTrafficAnalyticsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetworkWatcherFlowLogTrafficAnalyticsOutputReference_Override(n NetworkWatcherFlowLogTrafficAnalyticsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.networkWatcherFlowLog.NetworkWatcherFlowLogTrafficAnalyticsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference)SetInternalValue(val *NetworkWatcherFlowLogTrafficAnalytics) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference)SetIntervalInMinutes(val *float64) {
	if err := j.validateSetIntervalInMinutesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"intervalInMinutes",
		val,
	)
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference)SetWorkspaceId(val *string) {
	if err := j.validateSetWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceId",
		val,
	)
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference)SetWorkspaceRegion(val *string) {
	if err := j.validateSetWorkspaceRegionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceRegion",
		val,
	)
}

func (j *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference)SetWorkspaceResourceId(val *string) {
	if err := j.validateSetWorkspaceResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workspaceResourceId",
		val,
	)
}

func (n *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) ResetIntervalInMinutes() {
	_jsii_.InvokeVoid(
		n,
		"resetIntervalInMinutes",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetworkWatcherFlowLogTrafficAnalyticsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


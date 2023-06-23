package kubernetescluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v9/kubernetescluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesClusterAutoScalerProfileOutputReference interface {
	cdktf.ComplexObject
	BalanceSimilarNodeGroups() interface{}
	SetBalanceSimilarNodeGroups(val interface{})
	BalanceSimilarNodeGroupsInput() interface{}
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
	EmptyBulkDeleteMax() *string
	SetEmptyBulkDeleteMax(val *string)
	EmptyBulkDeleteMaxInput() *string
	Expander() *string
	SetExpander(val *string)
	ExpanderInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *KubernetesClusterAutoScalerProfile
	SetInternalValue(val *KubernetesClusterAutoScalerProfile)
	MaxGracefulTerminationSec() *string
	SetMaxGracefulTerminationSec(val *string)
	MaxGracefulTerminationSecInput() *string
	MaxNodeProvisioningTime() *string
	SetMaxNodeProvisioningTime(val *string)
	MaxNodeProvisioningTimeInput() *string
	MaxUnreadyNodes() *float64
	SetMaxUnreadyNodes(val *float64)
	MaxUnreadyNodesInput() *float64
	MaxUnreadyPercentage() *float64
	SetMaxUnreadyPercentage(val *float64)
	MaxUnreadyPercentageInput() *float64
	NewPodScaleUpDelay() *string
	SetNewPodScaleUpDelay(val *string)
	NewPodScaleUpDelayInput() *string
	ScaleDownDelayAfterAdd() *string
	SetScaleDownDelayAfterAdd(val *string)
	ScaleDownDelayAfterAddInput() *string
	ScaleDownDelayAfterDelete() *string
	SetScaleDownDelayAfterDelete(val *string)
	ScaleDownDelayAfterDeleteInput() *string
	ScaleDownDelayAfterFailure() *string
	SetScaleDownDelayAfterFailure(val *string)
	ScaleDownDelayAfterFailureInput() *string
	ScaleDownUnneeded() *string
	SetScaleDownUnneeded(val *string)
	ScaleDownUnneededInput() *string
	ScaleDownUnready() *string
	SetScaleDownUnready(val *string)
	ScaleDownUnreadyInput() *string
	ScaleDownUtilizationThreshold() *string
	SetScaleDownUtilizationThreshold(val *string)
	ScaleDownUtilizationThresholdInput() *string
	ScanInterval() *string
	SetScanInterval(val *string)
	ScanIntervalInput() *string
	SkipNodesWithLocalStorage() interface{}
	SetSkipNodesWithLocalStorage(val interface{})
	SkipNodesWithLocalStorageInput() interface{}
	SkipNodesWithSystemPods() interface{}
	SetSkipNodesWithSystemPods(val interface{})
	SkipNodesWithSystemPodsInput() interface{}
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
	ResetBalanceSimilarNodeGroups()
	ResetEmptyBulkDeleteMax()
	ResetExpander()
	ResetMaxGracefulTerminationSec()
	ResetMaxNodeProvisioningTime()
	ResetMaxUnreadyNodes()
	ResetMaxUnreadyPercentage()
	ResetNewPodScaleUpDelay()
	ResetScaleDownDelayAfterAdd()
	ResetScaleDownDelayAfterDelete()
	ResetScaleDownDelayAfterFailure()
	ResetScaleDownUnneeded()
	ResetScaleDownUnready()
	ResetScaleDownUtilizationThreshold()
	ResetScanInterval()
	ResetSkipNodesWithLocalStorage()
	ResetSkipNodesWithSystemPods()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterAutoScalerProfileOutputReference
type jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) BalanceSimilarNodeGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"balanceSimilarNodeGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) BalanceSimilarNodeGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"balanceSimilarNodeGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) EmptyBulkDeleteMax() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emptyBulkDeleteMax",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) EmptyBulkDeleteMaxInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"emptyBulkDeleteMaxInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) Expander() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expander",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ExpanderInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expanderInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) InternalValue() *KubernetesClusterAutoScalerProfile {
	var returns *KubernetesClusterAutoScalerProfile
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) MaxGracefulTerminationSec() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxGracefulTerminationSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) MaxGracefulTerminationSecInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxGracefulTerminationSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) MaxNodeProvisioningTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxNodeProvisioningTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) MaxNodeProvisioningTimeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxNodeProvisioningTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) MaxUnreadyNodes() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnreadyNodes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) MaxUnreadyNodesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnreadyNodesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) MaxUnreadyPercentage() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnreadyPercentage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) MaxUnreadyPercentageInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxUnreadyPercentageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) NewPodScaleUpDelay() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newPodScaleUpDelay",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) NewPodScaleUpDelayInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"newPodScaleUpDelayInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScaleDownDelayAfterAdd() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownDelayAfterAdd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScaleDownDelayAfterAddInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownDelayAfterAddInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScaleDownDelayAfterDelete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownDelayAfterDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScaleDownDelayAfterDeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownDelayAfterDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScaleDownDelayAfterFailure() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownDelayAfterFailure",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScaleDownDelayAfterFailureInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownDelayAfterFailureInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScaleDownUnneeded() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownUnneeded",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScaleDownUnneededInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownUnneededInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScaleDownUnready() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownUnready",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScaleDownUnreadyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownUnreadyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScaleDownUtilizationThreshold() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownUtilizationThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScaleDownUtilizationThresholdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleDownUtilizationThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScanInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scanInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ScanIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scanIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SkipNodesWithLocalStorage() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipNodesWithLocalStorage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SkipNodesWithLocalStorageInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipNodesWithLocalStorageInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SkipNodesWithSystemPods() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipNodesWithSystemPods",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) SkipNodesWithSystemPodsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipNodesWithSystemPodsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewKubernetesClusterAutoScalerProfileOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterAutoScalerProfileOutputReference {
	_init_.Initialize()

	if err := validateNewKubernetesClusterAutoScalerProfileOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterAutoScalerProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterAutoScalerProfileOutputReference_Override(k KubernetesClusterAutoScalerProfileOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterAutoScalerProfileOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference)SetBalanceSimilarNodeGroups(val interface{}) {
	if err := j.validateSetBalanceSimilarNodeGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"balanceSimilarNodeGroups",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference)SetEmptyBulkDeleteMax(val *string) {
	if err := j.validateSetEmptyBulkDeleteMaxParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"emptyBulkDeleteMax",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference)SetExpander(val *string) {
	if err := j.validateSetExpanderParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expander",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference)SetInternalValue(val *KubernetesClusterAutoScalerProfile) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference)SetMaxGracefulTerminationSec(val *string) {
	if err := j.validateSetMaxGracefulTerminationSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxGracefulTerminationSec",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference)SetMaxNodeProvisioningTime(val *string) {
	if err := j.validateSetMaxNodeProvisioningTimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxNodeProvisioningTime",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference)SetMaxUnreadyNodes(val *float64) {
	if err := j.validateSetMaxUnreadyNodesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUnreadyNodes",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference)SetMaxUnreadyPercentage(val *float64) {
	if err := j.validateSetMaxUnreadyPercentageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxUnreadyPercentage",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference)SetNewPodScaleUpDelay(val *string) {
	if err := j.validateSetNewPodScaleUpDelayParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"newPodScaleUpDelay",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference)SetScaleDownDelayAfterAdd(val *string) {
	if err := j.validateSetScaleDownDelayAfterAddParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleDownDelayAfterAdd",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference)SetScaleDownDelayAfterDelete(val *string) {
	if err := j.validateSetScaleDownDelayAfterDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleDownDelayAfterDelete",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference)SetScaleDownDelayAfterFailure(val *string) {
	if err := j.validateSetScaleDownDelayAfterFailureParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleDownDelayAfterFailure",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference)SetScaleDownUnneeded(val *string) {
	if err := j.validateSetScaleDownUnneededParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleDownUnneeded",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference)SetScaleDownUnready(val *string) {
	if err := j.validateSetScaleDownUnreadyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleDownUnready",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference)SetScaleDownUtilizationThreshold(val *string) {
	if err := j.validateSetScaleDownUtilizationThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleDownUtilizationThreshold",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference)SetScanInterval(val *string) {
	if err := j.validateSetScanIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scanInterval",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference)SetSkipNodesWithLocalStorage(val interface{}) {
	if err := j.validateSetSkipNodesWithLocalStorageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipNodesWithLocalStorage",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference)SetSkipNodesWithSystemPods(val interface{}) {
	if err := j.validateSetSkipNodesWithSystemPodsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipNodesWithSystemPods",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetBalanceSimilarNodeGroups() {
	_jsii_.InvokeVoid(
		k,
		"resetBalanceSimilarNodeGroups",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetEmptyBulkDeleteMax() {
	_jsii_.InvokeVoid(
		k,
		"resetEmptyBulkDeleteMax",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetExpander() {
	_jsii_.InvokeVoid(
		k,
		"resetExpander",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetMaxGracefulTerminationSec() {
	_jsii_.InvokeVoid(
		k,
		"resetMaxGracefulTerminationSec",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetMaxNodeProvisioningTime() {
	_jsii_.InvokeVoid(
		k,
		"resetMaxNodeProvisioningTime",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetMaxUnreadyNodes() {
	_jsii_.InvokeVoid(
		k,
		"resetMaxUnreadyNodes",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetMaxUnreadyPercentage() {
	_jsii_.InvokeVoid(
		k,
		"resetMaxUnreadyPercentage",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetNewPodScaleUpDelay() {
	_jsii_.InvokeVoid(
		k,
		"resetNewPodScaleUpDelay",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetScaleDownDelayAfterAdd() {
	_jsii_.InvokeVoid(
		k,
		"resetScaleDownDelayAfterAdd",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetScaleDownDelayAfterDelete() {
	_jsii_.InvokeVoid(
		k,
		"resetScaleDownDelayAfterDelete",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetScaleDownDelayAfterFailure() {
	_jsii_.InvokeVoid(
		k,
		"resetScaleDownDelayAfterFailure",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetScaleDownUnneeded() {
	_jsii_.InvokeVoid(
		k,
		"resetScaleDownUnneeded",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetScaleDownUnready() {
	_jsii_.InvokeVoid(
		k,
		"resetScaleDownUnready",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetScaleDownUtilizationThreshold() {
	_jsii_.InvokeVoid(
		k,
		"resetScaleDownUtilizationThreshold",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetScanInterval() {
	_jsii_.InvokeVoid(
		k,
		"resetScanInterval",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetSkipNodesWithLocalStorage() {
	_jsii_.InvokeVoid(
		k,
		"resetSkipNodesWithLocalStorage",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ResetSkipNodesWithSystemPods() {
	_jsii_.InvokeVoid(
		k,
		"resetSkipNodesWithSystemPods",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := k.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		k,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterAutoScalerProfileOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


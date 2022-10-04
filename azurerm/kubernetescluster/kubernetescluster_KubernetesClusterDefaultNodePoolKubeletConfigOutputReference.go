package kubernetescluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/jsii"

	"github.com/hashicorp/cdktf-provider-azurerm-go/azurerm/v3/kubernetescluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type KubernetesClusterDefaultNodePoolKubeletConfigOutputReference interface {
	cdktf.ComplexObject
	AllowedUnsafeSysctls() *[]*string
	SetAllowedUnsafeSysctls(val *[]*string)
	AllowedUnsafeSysctlsInput() *[]*string
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
	ContainerLogMaxLine() *float64
	SetContainerLogMaxLine(val *float64)
	ContainerLogMaxLineInput() *float64
	ContainerLogMaxSizeMb() *float64
	SetContainerLogMaxSizeMb(val *float64)
	ContainerLogMaxSizeMbInput() *float64
	CpuCfsQuotaEnabled() interface{}
	SetCpuCfsQuotaEnabled(val interface{})
	CpuCfsQuotaEnabledInput() interface{}
	CpuCfsQuotaPeriod() *string
	SetCpuCfsQuotaPeriod(val *string)
	CpuCfsQuotaPeriodInput() *string
	CpuManagerPolicy() *string
	SetCpuManagerPolicy(val *string)
	CpuManagerPolicyInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	ImageGcHighThreshold() *float64
	SetImageGcHighThreshold(val *float64)
	ImageGcHighThresholdInput() *float64
	ImageGcLowThreshold() *float64
	SetImageGcLowThreshold(val *float64)
	ImageGcLowThresholdInput() *float64
	InternalValue() *KubernetesClusterDefaultNodePoolKubeletConfig
	SetInternalValue(val *KubernetesClusterDefaultNodePoolKubeletConfig)
	PodMaxPid() *float64
	SetPodMaxPid(val *float64)
	PodMaxPidInput() *float64
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TopologyManagerPolicy() *string
	SetTopologyManagerPolicy(val *string)
	TopologyManagerPolicyInput() *string
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
	ResetAllowedUnsafeSysctls()
	ResetContainerLogMaxLine()
	ResetContainerLogMaxSizeMb()
	ResetCpuCfsQuotaEnabled()
	ResetCpuCfsQuotaPeriod()
	ResetCpuManagerPolicy()
	ResetImageGcHighThreshold()
	ResetImageGcLowThreshold()
	ResetPodMaxPid()
	ResetTopologyManagerPolicy()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for KubernetesClusterDefaultNodePoolKubeletConfigOutputReference
type jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) AllowedUnsafeSysctls() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUnsafeSysctls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) AllowedUnsafeSysctlsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedUnsafeSysctlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ContainerLogMaxLine() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerLogMaxLine",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ContainerLogMaxLineInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerLogMaxLineInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ContainerLogMaxSizeMb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerLogMaxSizeMb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ContainerLogMaxSizeMbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"containerLogMaxSizeMbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) CpuCfsQuotaEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuCfsQuotaEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) CpuCfsQuotaEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cpuCfsQuotaEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) CpuCfsQuotaPeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuCfsQuotaPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) CpuCfsQuotaPeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuCfsQuotaPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) CpuManagerPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuManagerPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) CpuManagerPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cpuManagerPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ImageGcHighThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"imageGcHighThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ImageGcHighThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"imageGcHighThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ImageGcLowThreshold() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"imageGcLowThreshold",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ImageGcLowThresholdInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"imageGcLowThresholdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) InternalValue() *KubernetesClusterDefaultNodePoolKubeletConfig {
	var returns *KubernetesClusterDefaultNodePoolKubeletConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) PodMaxPid() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"podMaxPid",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) PodMaxPidInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"podMaxPidInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) TopologyManagerPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topologyManagerPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) TopologyManagerPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"topologyManagerPolicyInput",
		&returns,
	)
	return returns
}


func NewKubernetesClusterDefaultNodePoolKubeletConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) KubernetesClusterDefaultNodePoolKubeletConfigOutputReference {
	_init_.Initialize()

	if err := validateNewKubernetesClusterDefaultNodePoolKubeletConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterDefaultNodePoolKubeletConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewKubernetesClusterDefaultNodePoolKubeletConfigOutputReference_Override(k KubernetesClusterDefaultNodePoolKubeletConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.kubernetesCluster.KubernetesClusterDefaultNodePoolKubeletConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		k,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference)SetAllowedUnsafeSysctls(val *[]*string) {
	if err := j.validateSetAllowedUnsafeSysctlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedUnsafeSysctls",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference)SetContainerLogMaxLine(val *float64) {
	if err := j.validateSetContainerLogMaxLineParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerLogMaxLine",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference)SetContainerLogMaxSizeMb(val *float64) {
	if err := j.validateSetContainerLogMaxSizeMbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"containerLogMaxSizeMb",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference)SetCpuCfsQuotaEnabled(val interface{}) {
	if err := j.validateSetCpuCfsQuotaEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuCfsQuotaEnabled",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference)SetCpuCfsQuotaPeriod(val *string) {
	if err := j.validateSetCpuCfsQuotaPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuCfsQuotaPeriod",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference)SetCpuManagerPolicy(val *string) {
	if err := j.validateSetCpuManagerPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cpuManagerPolicy",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference)SetImageGcHighThreshold(val *float64) {
	if err := j.validateSetImageGcHighThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageGcHighThreshold",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference)SetImageGcLowThreshold(val *float64) {
	if err := j.validateSetImageGcLowThresholdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"imageGcLowThreshold",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference)SetInternalValue(val *KubernetesClusterDefaultNodePoolKubeletConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference)SetPodMaxPid(val *float64) {
	if err := j.validateSetPodMaxPidParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"podMaxPid",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference)SetTopologyManagerPolicy(val *string) {
	if err := j.validateSetTopologyManagerPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"topologyManagerPolicy",
		val,
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ResetAllowedUnsafeSysctls() {
	_jsii_.InvokeVoid(
		k,
		"resetAllowedUnsafeSysctls",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ResetContainerLogMaxLine() {
	_jsii_.InvokeVoid(
		k,
		"resetContainerLogMaxLine",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ResetContainerLogMaxSizeMb() {
	_jsii_.InvokeVoid(
		k,
		"resetContainerLogMaxSizeMb",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ResetCpuCfsQuotaEnabled() {
	_jsii_.InvokeVoid(
		k,
		"resetCpuCfsQuotaEnabled",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ResetCpuCfsQuotaPeriod() {
	_jsii_.InvokeVoid(
		k,
		"resetCpuCfsQuotaPeriod",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ResetCpuManagerPolicy() {
	_jsii_.InvokeVoid(
		k,
		"resetCpuManagerPolicy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ResetImageGcHighThreshold() {
	_jsii_.InvokeVoid(
		k,
		"resetImageGcHighThreshold",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ResetImageGcLowThreshold() {
	_jsii_.InvokeVoid(
		k,
		"resetImageGcLowThreshold",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ResetPodMaxPid() {
	_jsii_.InvokeVoid(
		k,
		"resetPodMaxPid",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ResetTopologyManagerPolicy() {
	_jsii_.InvokeVoid(
		k,
		"resetTopologyManagerPolicy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (k *jsiiProxy_KubernetesClusterDefaultNodePoolKubeletConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


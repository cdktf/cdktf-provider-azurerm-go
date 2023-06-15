package hdinsightkafkacluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/hdinsightkafkacluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type HdinsightKafkaClusterRolesOutputReference interface {
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
	HeadNode() HdinsightKafkaClusterRolesHeadNodeOutputReference
	HeadNodeInput() *HdinsightKafkaClusterRolesHeadNode
	InternalValue() *HdinsightKafkaClusterRoles
	SetInternalValue(val *HdinsightKafkaClusterRoles)
	KafkaManagementNode() HdinsightKafkaClusterRolesKafkaManagementNodeOutputReference
	KafkaManagementNodeInput() *HdinsightKafkaClusterRolesKafkaManagementNode
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	WorkerNode() HdinsightKafkaClusterRolesWorkerNodeOutputReference
	WorkerNodeInput() *HdinsightKafkaClusterRolesWorkerNode
	ZookeeperNode() HdinsightKafkaClusterRolesZookeeperNodeOutputReference
	ZookeeperNodeInput() *HdinsightKafkaClusterRolesZookeeperNode
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
	PutHeadNode(value *HdinsightKafkaClusterRolesHeadNode)
	PutKafkaManagementNode(value *HdinsightKafkaClusterRolesKafkaManagementNode)
	PutWorkerNode(value *HdinsightKafkaClusterRolesWorkerNode)
	PutZookeeperNode(value *HdinsightKafkaClusterRolesZookeeperNode)
	ResetKafkaManagementNode()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for HdinsightKafkaClusterRolesOutputReference
type jsiiProxy_HdinsightKafkaClusterRolesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) HeadNode() HdinsightKafkaClusterRolesHeadNodeOutputReference {
	var returns HdinsightKafkaClusterRolesHeadNodeOutputReference
	_jsii_.Get(
		j,
		"headNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) HeadNodeInput() *HdinsightKafkaClusterRolesHeadNode {
	var returns *HdinsightKafkaClusterRolesHeadNode
	_jsii_.Get(
		j,
		"headNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) InternalValue() *HdinsightKafkaClusterRoles {
	var returns *HdinsightKafkaClusterRoles
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) KafkaManagementNode() HdinsightKafkaClusterRolesKafkaManagementNodeOutputReference {
	var returns HdinsightKafkaClusterRolesKafkaManagementNodeOutputReference
	_jsii_.Get(
		j,
		"kafkaManagementNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) KafkaManagementNodeInput() *HdinsightKafkaClusterRolesKafkaManagementNode {
	var returns *HdinsightKafkaClusterRolesKafkaManagementNode
	_jsii_.Get(
		j,
		"kafkaManagementNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) WorkerNode() HdinsightKafkaClusterRolesWorkerNodeOutputReference {
	var returns HdinsightKafkaClusterRolesWorkerNodeOutputReference
	_jsii_.Get(
		j,
		"workerNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) WorkerNodeInput() *HdinsightKafkaClusterRolesWorkerNode {
	var returns *HdinsightKafkaClusterRolesWorkerNode
	_jsii_.Get(
		j,
		"workerNodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) ZookeeperNode() HdinsightKafkaClusterRolesZookeeperNodeOutputReference {
	var returns HdinsightKafkaClusterRolesZookeeperNodeOutputReference
	_jsii_.Get(
		j,
		"zookeeperNode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) ZookeeperNodeInput() *HdinsightKafkaClusterRolesZookeeperNode {
	var returns *HdinsightKafkaClusterRolesZookeeperNode
	_jsii_.Get(
		j,
		"zookeeperNodeInput",
		&returns,
	)
	return returns
}


func NewHdinsightKafkaClusterRolesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) HdinsightKafkaClusterRolesOutputReference {
	_init_.Initialize()

	if err := validateNewHdinsightKafkaClusterRolesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_HdinsightKafkaClusterRolesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightKafkaCluster.HdinsightKafkaClusterRolesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewHdinsightKafkaClusterRolesOutputReference_Override(h HdinsightKafkaClusterRolesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.hdinsightKafkaCluster.HdinsightKafkaClusterRolesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		h,
	)
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesOutputReference)SetInternalValue(val *HdinsightKafkaClusterRoles) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_HdinsightKafkaClusterRolesOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := h.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		h,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := h.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := h.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		h,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := h.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		h,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := h.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		h,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := h.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		h,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := h.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		h,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := h.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		h,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := h.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		h,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := h.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		h,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) PutHeadNode(value *HdinsightKafkaClusterRolesHeadNode) {
	if err := h.validatePutHeadNodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putHeadNode",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) PutKafkaManagementNode(value *HdinsightKafkaClusterRolesKafkaManagementNode) {
	if err := h.validatePutKafkaManagementNodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putKafkaManagementNode",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) PutWorkerNode(value *HdinsightKafkaClusterRolesWorkerNode) {
	if err := h.validatePutWorkerNodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putWorkerNode",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) PutZookeeperNode(value *HdinsightKafkaClusterRolesZookeeperNode) {
	if err := h.validatePutZookeeperNodeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		h,
		"putZookeeperNode",
		[]interface{}{value},
	)
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) ResetKafkaManagementNode() {
	_jsii_.InvokeVoid(
		h,
		"resetKafkaManagementNode",
		nil, // no parameters
	)
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := h.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		h,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (h *jsiiProxy_HdinsightKafkaClusterRolesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		h,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


package netappvolumegroupsaphana

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/netappvolumegroupsaphana/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference interface {
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
	EndpointType() *string
	SetEndpointType(val *string)
	EndpointTypeInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *NetappVolumeGroupSapHanaVolumeDataProtectionReplication
	SetInternalValue(val *NetappVolumeGroupSapHanaVolumeDataProtectionReplication)
	RemoteVolumeLocation() *string
	SetRemoteVolumeLocation(val *string)
	RemoteVolumeLocationInput() *string
	RemoteVolumeResourceId() *string
	SetRemoteVolumeResourceId(val *string)
	RemoteVolumeResourceIdInput() *string
	ReplicationFrequency() *string
	SetReplicationFrequency(val *string)
	ReplicationFrequencyInput() *string
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
	ResetEndpointType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference
type jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) EndpointType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) EndpointTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"endpointTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) InternalValue() *NetappVolumeGroupSapHanaVolumeDataProtectionReplication {
	var returns *NetappVolumeGroupSapHanaVolumeDataProtectionReplication
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) RemoteVolumeLocation() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteVolumeLocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) RemoteVolumeLocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteVolumeLocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) RemoteVolumeResourceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteVolumeResourceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) RemoteVolumeResourceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"remoteVolumeResourceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) ReplicationFrequency() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationFrequency",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) ReplicationFrequencyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"replicationFrequencyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewNetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference {
	_init_.Initialize()

	if err := validateNewNetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.netappVolumeGroupSapHana.NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewNetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference_Override(n NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.netappVolumeGroupSapHana.NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		n,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference)SetEndpointType(val *string) {
	if err := j.validateSetEndpointTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"endpointType",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference)SetInternalValue(val *NetappVolumeGroupSapHanaVolumeDataProtectionReplication) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference)SetRemoteVolumeLocation(val *string) {
	if err := j.validateSetRemoteVolumeLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteVolumeLocation",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference)SetRemoteVolumeResourceId(val *string) {
	if err := j.validateSetRemoteVolumeResourceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"remoteVolumeResourceId",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference)SetReplicationFrequency(val *string) {
	if err := j.validateSetReplicationFrequencyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"replicationFrequency",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		n,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) ResetEndpointType() {
	_jsii_.InvokeVoid(
		n,
		"resetEndpointType",
		nil, // no parameters
	)
}

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (n *jsiiProxy_NetappVolumeGroupSapHanaVolumeDataProtectionReplicationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		n,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


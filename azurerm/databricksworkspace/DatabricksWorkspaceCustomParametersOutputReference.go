package databricksworkspace

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v7/databricksworkspace/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabricksWorkspaceCustomParametersOutputReference interface {
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
	InternalValue() *DatabricksWorkspaceCustomParameters
	SetInternalValue(val *DatabricksWorkspaceCustomParameters)
	MachineLearningWorkspaceId() *string
	SetMachineLearningWorkspaceId(val *string)
	MachineLearningWorkspaceIdInput() *string
	NatGatewayName() *string
	SetNatGatewayName(val *string)
	NatGatewayNameInput() *string
	NoPublicIp() interface{}
	SetNoPublicIp(val interface{})
	NoPublicIpInput() interface{}
	PrivateSubnetName() *string
	SetPrivateSubnetName(val *string)
	PrivateSubnetNameInput() *string
	PrivateSubnetNetworkSecurityGroupAssociationId() *string
	SetPrivateSubnetNetworkSecurityGroupAssociationId(val *string)
	PrivateSubnetNetworkSecurityGroupAssociationIdInput() *string
	PublicIpName() *string
	SetPublicIpName(val *string)
	PublicIpNameInput() *string
	PublicSubnetName() *string
	SetPublicSubnetName(val *string)
	PublicSubnetNameInput() *string
	PublicSubnetNetworkSecurityGroupAssociationId() *string
	SetPublicSubnetNetworkSecurityGroupAssociationId(val *string)
	PublicSubnetNetworkSecurityGroupAssociationIdInput() *string
	StorageAccountName() *string
	SetStorageAccountName(val *string)
	StorageAccountNameInput() *string
	StorageAccountSkuName() *string
	SetStorageAccountSkuName(val *string)
	StorageAccountSkuNameInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VirtualNetworkId() *string
	SetVirtualNetworkId(val *string)
	VirtualNetworkIdInput() *string
	VnetAddressPrefix() *string
	SetVnetAddressPrefix(val *string)
	VnetAddressPrefixInput() *string
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
	ResetMachineLearningWorkspaceId()
	ResetNatGatewayName()
	ResetNoPublicIp()
	ResetPrivateSubnetName()
	ResetPrivateSubnetNetworkSecurityGroupAssociationId()
	ResetPublicIpName()
	ResetPublicSubnetName()
	ResetPublicSubnetNetworkSecurityGroupAssociationId()
	ResetStorageAccountName()
	ResetStorageAccountSkuName()
	ResetVirtualNetworkId()
	ResetVnetAddressPrefix()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatabricksWorkspaceCustomParametersOutputReference
type jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) InternalValue() *DatabricksWorkspaceCustomParameters {
	var returns *DatabricksWorkspaceCustomParameters
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) MachineLearningWorkspaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineLearningWorkspaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) MachineLearningWorkspaceIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineLearningWorkspaceIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) NatGatewayName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natGatewayName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) NatGatewayNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"natGatewayNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) NoPublicIp() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) NoPublicIpInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"noPublicIpInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) PrivateSubnetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateSubnetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) PrivateSubnetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateSubnetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) PrivateSubnetNetworkSecurityGroupAssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateSubnetNetworkSecurityGroupAssociationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) PrivateSubnetNetworkSecurityGroupAssociationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateSubnetNetworkSecurityGroupAssociationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) PublicIpName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) PublicIpNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) PublicSubnetName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicSubnetName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) PublicSubnetNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicSubnetNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) PublicSubnetNetworkSecurityGroupAssociationId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicSubnetNetworkSecurityGroupAssociationId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) PublicSubnetNetworkSecurityGroupAssociationIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicSubnetNetworkSecurityGroupAssociationIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) StorageAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) StorageAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) StorageAccountSkuName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountSkuName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) StorageAccountSkuNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"storageAccountSkuNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) VirtualNetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) VirtualNetworkIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualNetworkIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) VnetAddressPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnetAddressPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) VnetAddressPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vnetAddressPrefixInput",
		&returns,
	)
	return returns
}


func NewDatabricksWorkspaceCustomParametersOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatabricksWorkspaceCustomParametersOutputReference {
	_init_.Initialize()

	if err := validateNewDatabricksWorkspaceCustomParametersOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.databricksWorkspace.DatabricksWorkspaceCustomParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatabricksWorkspaceCustomParametersOutputReference_Override(d DatabricksWorkspaceCustomParametersOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.databricksWorkspace.DatabricksWorkspaceCustomParametersOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference)SetInternalValue(val *DatabricksWorkspaceCustomParameters) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference)SetMachineLearningWorkspaceId(val *string) {
	if err := j.validateSetMachineLearningWorkspaceIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineLearningWorkspaceId",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference)SetNatGatewayName(val *string) {
	if err := j.validateSetNatGatewayNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"natGatewayName",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference)SetNoPublicIp(val interface{}) {
	if err := j.validateSetNoPublicIpParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"noPublicIp",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference)SetPrivateSubnetName(val *string) {
	if err := j.validateSetPrivateSubnetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateSubnetName",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference)SetPrivateSubnetNetworkSecurityGroupAssociationId(val *string) {
	if err := j.validateSetPrivateSubnetNetworkSecurityGroupAssociationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"privateSubnetNetworkSecurityGroupAssociationId",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference)SetPublicIpName(val *string) {
	if err := j.validateSetPublicIpNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicIpName",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference)SetPublicSubnetName(val *string) {
	if err := j.validateSetPublicSubnetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicSubnetName",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference)SetPublicSubnetNetworkSecurityGroupAssociationId(val *string) {
	if err := j.validateSetPublicSubnetNetworkSecurityGroupAssociationIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"publicSubnetNetworkSecurityGroupAssociationId",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference)SetStorageAccountName(val *string) {
	if err := j.validateSetStorageAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountName",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference)SetStorageAccountSkuName(val *string) {
	if err := j.validateSetStorageAccountSkuNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"storageAccountSkuName",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference)SetVirtualNetworkId(val *string) {
	if err := j.validateSetVirtualNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualNetworkId",
		val,
	)
}

func (j *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference)SetVnetAddressPrefix(val *string) {
	if err := j.validateSetVnetAddressPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vnetAddressPrefix",
		val,
	)
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ResetMachineLearningWorkspaceId() {
	_jsii_.InvokeVoid(
		d,
		"resetMachineLearningWorkspaceId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ResetNatGatewayName() {
	_jsii_.InvokeVoid(
		d,
		"resetNatGatewayName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ResetNoPublicIp() {
	_jsii_.InvokeVoid(
		d,
		"resetNoPublicIp",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ResetPrivateSubnetName() {
	_jsii_.InvokeVoid(
		d,
		"resetPrivateSubnetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ResetPrivateSubnetNetworkSecurityGroupAssociationId() {
	_jsii_.InvokeVoid(
		d,
		"resetPrivateSubnetNetworkSecurityGroupAssociationId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ResetPublicIpName() {
	_jsii_.InvokeVoid(
		d,
		"resetPublicIpName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ResetPublicSubnetName() {
	_jsii_.InvokeVoid(
		d,
		"resetPublicSubnetName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ResetPublicSubnetNetworkSecurityGroupAssociationId() {
	_jsii_.InvokeVoid(
		d,
		"resetPublicSubnetNetworkSecurityGroupAssociationId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ResetStorageAccountName() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageAccountName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ResetStorageAccountSkuName() {
	_jsii_.InvokeVoid(
		d,
		"resetStorageAccountSkuName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ResetVirtualNetworkId() {
	_jsii_.InvokeVoid(
		d,
		"resetVirtualNetworkId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ResetVnetAddressPrefix() {
	_jsii_.InvokeVoid(
		d,
		"resetVnetAddressPrefix",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabricksWorkspaceCustomParametersOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


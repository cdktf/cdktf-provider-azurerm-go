package servicefabricmanagedcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/jsii"

	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v8/servicefabricmanagedcluster/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ServiceFabricManagedClusterNodeTypeOutputReference interface {
	cdktf.ComplexObject
	ApplicationPortRange() *string
	SetApplicationPortRange(val *string)
	ApplicationPortRangeInput() *string
	Capacities() *map[string]*string
	SetCapacities(val *map[string]*string)
	CapacitiesInput() *map[string]*string
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
	DataDiskSizeGb() *float64
	SetDataDiskSizeGb(val *float64)
	DataDiskSizeGbInput() *float64
	DataDiskType() *string
	SetDataDiskType(val *string)
	DataDiskTypeInput() *string
	EphemeralPortRange() *string
	SetEphemeralPortRange(val *string)
	EphemeralPortRangeInput() *string
	// Experimental.
	Fqn() *string
	Id() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MultiplePlacementGroupsEnabled() interface{}
	SetMultiplePlacementGroupsEnabled(val interface{})
	MultiplePlacementGroupsEnabledInput() interface{}
	Name() *string
	SetName(val *string)
	NameInput() *string
	PlacementProperties() *map[string]*string
	SetPlacementProperties(val *map[string]*string)
	PlacementPropertiesInput() *map[string]*string
	Primary() interface{}
	SetPrimary(val interface{})
	PrimaryInput() interface{}
	Stateless() interface{}
	SetStateless(val interface{})
	StatelessInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VmImageOffer() *string
	SetVmImageOffer(val *string)
	VmImageOfferInput() *string
	VmImagePublisher() *string
	SetVmImagePublisher(val *string)
	VmImagePublisherInput() *string
	VmImageSku() *string
	SetVmImageSku(val *string)
	VmImageSkuInput() *string
	VmImageVersion() *string
	SetVmImageVersion(val *string)
	VmImageVersionInput() *string
	VmInstanceCount() *float64
	SetVmInstanceCount(val *float64)
	VmInstanceCountInput() *float64
	VmSecrets() ServiceFabricManagedClusterNodeTypeVmSecretsList
	VmSecretsInput() interface{}
	VmSize() *string
	SetVmSize(val *string)
	VmSizeInput() *string
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
	PutVmSecrets(value interface{})
	ResetCapacities()
	ResetDataDiskType()
	ResetMultiplePlacementGroupsEnabled()
	ResetPlacementProperties()
	ResetPrimary()
	ResetStateless()
	ResetVmSecrets()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ServiceFabricManagedClusterNodeTypeOutputReference
type jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) ApplicationPortRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationPortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) ApplicationPortRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationPortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) Capacities() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"capacities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) CapacitiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"capacitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) DataDiskSizeGb() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataDiskSizeGb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) DataDiskSizeGbInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"dataDiskSizeGbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) DataDiskType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataDiskType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) DataDiskTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataDiskTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) EphemeralPortRange() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ephemeralPortRange",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) EphemeralPortRangeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ephemeralPortRangeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) MultiplePlacementGroupsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiplePlacementGroupsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) MultiplePlacementGroupsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"multiplePlacementGroupsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) PlacementProperties() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"placementProperties",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) PlacementPropertiesInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"placementPropertiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) Primary() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primary",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) PrimaryInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primaryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) Stateless() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"stateless",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) StatelessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"statelessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmImageOffer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmImageOffer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmImageOfferInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmImageOfferInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmImagePublisher() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmImagePublisher",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmImagePublisherInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmImagePublisherInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmImageSku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmImageSku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmImageSkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmImageSkuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmImageVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmImageVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmImageVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmImageVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmInstanceCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmInstanceCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmInstanceCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"vmInstanceCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmSecrets() ServiceFabricManagedClusterNodeTypeVmSecretsList {
	var returns ServiceFabricManagedClusterNodeTypeVmSecretsList
	_jsii_.Get(
		j,
		"vmSecrets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmSecretsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vmSecretsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmSize() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) VmSizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"vmSizeInput",
		&returns,
	)
	return returns
}


func NewServiceFabricManagedClusterNodeTypeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) ServiceFabricManagedClusterNodeTypeOutputReference {
	_init_.Initialize()

	if err := validateNewServiceFabricManagedClusterNodeTypeOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterNodeTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewServiceFabricManagedClusterNodeTypeOutputReference_Override(s ServiceFabricManagedClusterNodeTypeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.serviceFabricManagedCluster.ServiceFabricManagedClusterNodeTypeOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		s,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference)SetApplicationPortRange(val *string) {
	if err := j.validateSetApplicationPortRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"applicationPortRange",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference)SetCapacities(val *map[string]*string) {
	if err := j.validateSetCapacitiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacities",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference)SetDataDiskSizeGb(val *float64) {
	if err := j.validateSetDataDiskSizeGbParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataDiskSizeGb",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference)SetDataDiskType(val *string) {
	if err := j.validateSetDataDiskTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dataDiskType",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference)SetEphemeralPortRange(val *string) {
	if err := j.validateSetEphemeralPortRangeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ephemeralPortRange",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference)SetMultiplePlacementGroupsEnabled(val interface{}) {
	if err := j.validateSetMultiplePlacementGroupsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"multiplePlacementGroupsEnabled",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference)SetPlacementProperties(val *map[string]*string) {
	if err := j.validateSetPlacementPropertiesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"placementProperties",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference)SetPrimary(val interface{}) {
	if err := j.validateSetPrimaryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"primary",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference)SetStateless(val interface{}) {
	if err := j.validateSetStatelessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"stateless",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference)SetVmImageOffer(val *string) {
	if err := j.validateSetVmImageOfferParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmImageOffer",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference)SetVmImagePublisher(val *string) {
	if err := j.validateSetVmImagePublisherParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmImagePublisher",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference)SetVmImageSku(val *string) {
	if err := j.validateSetVmImageSkuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmImageSku",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference)SetVmImageVersion(val *string) {
	if err := j.validateSetVmImageVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmImageVersion",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference)SetVmInstanceCount(val *float64) {
	if err := j.validateSetVmInstanceCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmInstanceCount",
		val,
	)
}

func (j *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference)SetVmSize(val *string) {
	if err := j.validateSetVmSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmSize",
		val,
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := s.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		s,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := s.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := s.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		s,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := s.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		s,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := s.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		s,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := s.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		s,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := s.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		s,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := s.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		s,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := s.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		s,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := s.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		s,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) PutVmSecrets(value interface{}) {
	if err := s.validatePutVmSecretsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		s,
		"putVmSecrets",
		[]interface{}{value},
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) ResetCapacities() {
	_jsii_.InvokeVoid(
		s,
		"resetCapacities",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) ResetDataDiskType() {
	_jsii_.InvokeVoid(
		s,
		"resetDataDiskType",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) ResetMultiplePlacementGroupsEnabled() {
	_jsii_.InvokeVoid(
		s,
		"resetMultiplePlacementGroupsEnabled",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) ResetPlacementProperties() {
	_jsii_.InvokeVoid(
		s,
		"resetPlacementProperties",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) ResetPrimary() {
	_jsii_.InvokeVoid(
		s,
		"resetPrimary",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) ResetStateless() {
	_jsii_.InvokeVoid(
		s,
		"resetStateless",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) ResetVmSecrets() {
	_jsii_.InvokeVoid(
		s,
		"resetVmSecrets",
		nil, // no parameters
	)
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := s.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		s,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (s *jsiiProxy_ServiceFabricManagedClusterNodeTypeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		s,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package linuxvirtualmachinescaleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v14/linuxvirtualmachinescaleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/linux_virtual_machine_scale_set azurerm_linux_virtual_machine_scale_set}.
type LinuxVirtualMachineScaleSet interface {
	cdktf.TerraformResource
	AdditionalCapabilities() LinuxVirtualMachineScaleSetAdditionalCapabilitiesOutputReference
	AdditionalCapabilitiesInput() *LinuxVirtualMachineScaleSetAdditionalCapabilities
	AdminPassword() *string
	SetAdminPassword(val *string)
	AdminPasswordInput() *string
	AdminSshKey() LinuxVirtualMachineScaleSetAdminSshKeyList
	AdminSshKeyInput() interface{}
	AdminUsername() *string
	SetAdminUsername(val *string)
	AdminUsernameInput() *string
	AutomaticInstanceRepair() LinuxVirtualMachineScaleSetAutomaticInstanceRepairOutputReference
	AutomaticInstanceRepairInput() *LinuxVirtualMachineScaleSetAutomaticInstanceRepair
	AutomaticOsUpgradePolicy() LinuxVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference
	AutomaticOsUpgradePolicyInput() *LinuxVirtualMachineScaleSetAutomaticOsUpgradePolicy
	BootDiagnostics() LinuxVirtualMachineScaleSetBootDiagnosticsOutputReference
	BootDiagnosticsInput() *LinuxVirtualMachineScaleSetBootDiagnostics
	CapacityReservationGroupId() *string
	SetCapacityReservationGroupId(val *string)
	CapacityReservationGroupIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ComputerNamePrefix() *string
	SetComputerNamePrefix(val *string)
	ComputerNamePrefixInput() *string
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	CustomData() *string
	SetCustomData(val *string)
	CustomDataInput() *string
	DataDisk() LinuxVirtualMachineScaleSetDataDiskList
	DataDiskInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DisablePasswordAuthentication() interface{}
	SetDisablePasswordAuthentication(val interface{})
	DisablePasswordAuthenticationInput() interface{}
	DoNotRunExtensionsOnOverprovisionedMachines() interface{}
	SetDoNotRunExtensionsOnOverprovisionedMachines(val interface{})
	DoNotRunExtensionsOnOverprovisionedMachinesInput() interface{}
	EdgeZone() *string
	SetEdgeZone(val *string)
	EdgeZoneInput() *string
	EncryptionAtHostEnabled() interface{}
	SetEncryptionAtHostEnabled(val interface{})
	EncryptionAtHostEnabledInput() interface{}
	EvictionPolicy() *string
	SetEvictionPolicy(val *string)
	EvictionPolicyInput() *string
	Extension() LinuxVirtualMachineScaleSetExtensionList
	ExtensionInput() interface{}
	ExtensionOperationsEnabled() interface{}
	SetExtensionOperationsEnabled(val interface{})
	ExtensionOperationsEnabledInput() interface{}
	ExtensionsTimeBudget() *string
	SetExtensionsTimeBudget(val *string)
	ExtensionsTimeBudgetInput() *string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GalleryApplication() LinuxVirtualMachineScaleSetGalleryApplicationList
	GalleryApplicationInput() interface{}
	HealthProbeId() *string
	SetHealthProbeId(val *string)
	HealthProbeIdInput() *string
	HostGroupId() *string
	SetHostGroupId(val *string)
	HostGroupIdInput() *string
	Id() *string
	SetId(val *string)
	Identity() LinuxVirtualMachineScaleSetIdentityOutputReference
	IdentityInput() *LinuxVirtualMachineScaleSetIdentity
	IdInput() *string
	Instances() *float64
	SetInstances(val *float64)
	InstancesInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Location() *string
	SetLocation(val *string)
	LocationInput() *string
	MaxBidPrice() *float64
	SetMaxBidPrice(val *float64)
	MaxBidPriceInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	NetworkInterface() LinuxVirtualMachineScaleSetNetworkInterfaceList
	NetworkInterfaceInput() interface{}
	// The tree node.
	Node() constructs.Node
	OsDisk() LinuxVirtualMachineScaleSetOsDiskOutputReference
	OsDiskInput() *LinuxVirtualMachineScaleSetOsDisk
	Overprovision() interface{}
	SetOverprovision(val interface{})
	OverprovisionInput() interface{}
	Plan() LinuxVirtualMachineScaleSetPlanOutputReference
	PlanInput() *LinuxVirtualMachineScaleSetPlan
	PlatformFaultDomainCount() *float64
	SetPlatformFaultDomainCount(val *float64)
	PlatformFaultDomainCountInput() *float64
	Priority() *string
	SetPriority(val *string)
	PriorityInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	ProvisionVmAgent() interface{}
	SetProvisionVmAgent(val interface{})
	ProvisionVmAgentInput() interface{}
	ProximityPlacementGroupId() *string
	SetProximityPlacementGroupId(val *string)
	ProximityPlacementGroupIdInput() *string
	// Experimental.
	RawOverrides() interface{}
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	RollingUpgradePolicy() LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference
	RollingUpgradePolicyInput() *LinuxVirtualMachineScaleSetRollingUpgradePolicy
	ScaleIn() LinuxVirtualMachineScaleSetScaleInOutputReference
	ScaleInInput() *LinuxVirtualMachineScaleSetScaleIn
	Secret() LinuxVirtualMachineScaleSetSecretList
	SecretInput() interface{}
	SecureBootEnabled() interface{}
	SetSecureBootEnabled(val interface{})
	SecureBootEnabledInput() interface{}
	SinglePlacementGroup() interface{}
	SetSinglePlacementGroup(val interface{})
	SinglePlacementGroupInput() interface{}
	Sku() *string
	SetSku(val *string)
	SkuInput() *string
	SourceImageId() *string
	SetSourceImageId(val *string)
	SourceImageIdInput() *string
	SourceImageReference() LinuxVirtualMachineScaleSetSourceImageReferenceOutputReference
	SourceImageReferenceInput() *LinuxVirtualMachineScaleSetSourceImageReference
	SpotRestore() LinuxVirtualMachineScaleSetSpotRestoreOutputReference
	SpotRestoreInput() *LinuxVirtualMachineScaleSetSpotRestore
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TerminationNotification() LinuxVirtualMachineScaleSetTerminationNotificationOutputReference
	TerminationNotificationInput() *LinuxVirtualMachineScaleSetTerminationNotification
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() LinuxVirtualMachineScaleSetTimeoutsOutputReference
	TimeoutsInput() interface{}
	UniqueId() *string
	UpgradeMode() *string
	SetUpgradeMode(val *string)
	UpgradeModeInput() *string
	UserData() *string
	SetUserData(val *string)
	UserDataInput() *string
	VtpmEnabled() interface{}
	SetVtpmEnabled(val interface{})
	VtpmEnabledInput() interface{}
	ZoneBalance() interface{}
	SetZoneBalance(val interface{})
	ZoneBalanceInput() interface{}
	Zones() *[]*string
	SetZones(val *[]*string)
	ZonesInput() *[]*string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
	// Experimental.
	AddOverride(path *string, value interface{})
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAdditionalCapabilities(value *LinuxVirtualMachineScaleSetAdditionalCapabilities)
	PutAdminSshKey(value interface{})
	PutAutomaticInstanceRepair(value *LinuxVirtualMachineScaleSetAutomaticInstanceRepair)
	PutAutomaticOsUpgradePolicy(value *LinuxVirtualMachineScaleSetAutomaticOsUpgradePolicy)
	PutBootDiagnostics(value *LinuxVirtualMachineScaleSetBootDiagnostics)
	PutDataDisk(value interface{})
	PutExtension(value interface{})
	PutGalleryApplication(value interface{})
	PutIdentity(value *LinuxVirtualMachineScaleSetIdentity)
	PutNetworkInterface(value interface{})
	PutOsDisk(value *LinuxVirtualMachineScaleSetOsDisk)
	PutPlan(value *LinuxVirtualMachineScaleSetPlan)
	PutRollingUpgradePolicy(value *LinuxVirtualMachineScaleSetRollingUpgradePolicy)
	PutScaleIn(value *LinuxVirtualMachineScaleSetScaleIn)
	PutSecret(value interface{})
	PutSourceImageReference(value *LinuxVirtualMachineScaleSetSourceImageReference)
	PutSpotRestore(value *LinuxVirtualMachineScaleSetSpotRestore)
	PutTerminationNotification(value *LinuxVirtualMachineScaleSetTerminationNotification)
	PutTimeouts(value *LinuxVirtualMachineScaleSetTimeouts)
	ResetAdditionalCapabilities()
	ResetAdminPassword()
	ResetAdminSshKey()
	ResetAutomaticInstanceRepair()
	ResetAutomaticOsUpgradePolicy()
	ResetBootDiagnostics()
	ResetCapacityReservationGroupId()
	ResetComputerNamePrefix()
	ResetCustomData()
	ResetDataDisk()
	ResetDisablePasswordAuthentication()
	ResetDoNotRunExtensionsOnOverprovisionedMachines()
	ResetEdgeZone()
	ResetEncryptionAtHostEnabled()
	ResetEvictionPolicy()
	ResetExtension()
	ResetExtensionOperationsEnabled()
	ResetExtensionsTimeBudget()
	ResetGalleryApplication()
	ResetHealthProbeId()
	ResetHostGroupId()
	ResetId()
	ResetIdentity()
	ResetInstances()
	ResetMaxBidPrice()
	ResetOverprovision()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPlan()
	ResetPlatformFaultDomainCount()
	ResetPriority()
	ResetProvisionVmAgent()
	ResetProximityPlacementGroupId()
	ResetRollingUpgradePolicy()
	ResetScaleIn()
	ResetSecret()
	ResetSecureBootEnabled()
	ResetSinglePlacementGroup()
	ResetSourceImageId()
	ResetSourceImageReference()
	ResetSpotRestore()
	ResetTags()
	ResetTerminationNotification()
	ResetTimeouts()
	ResetUpgradeMode()
	ResetUserData()
	ResetVtpmEnabled()
	ResetZoneBalance()
	ResetZones()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for LinuxVirtualMachineScaleSet
type jsiiProxy_LinuxVirtualMachineScaleSet struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) AdditionalCapabilities() LinuxVirtualMachineScaleSetAdditionalCapabilitiesOutputReference {
	var returns LinuxVirtualMachineScaleSetAdditionalCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"additionalCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) AdditionalCapabilitiesInput() *LinuxVirtualMachineScaleSetAdditionalCapabilities {
	var returns *LinuxVirtualMachineScaleSetAdditionalCapabilities
	_jsii_.Get(
		j,
		"additionalCapabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) AdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) AdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) AdminSshKey() LinuxVirtualMachineScaleSetAdminSshKeyList {
	var returns LinuxVirtualMachineScaleSetAdminSshKeyList
	_jsii_.Get(
		j,
		"adminSshKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) AdminSshKeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"adminSshKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) AdminUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) AdminUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) AutomaticInstanceRepair() LinuxVirtualMachineScaleSetAutomaticInstanceRepairOutputReference {
	var returns LinuxVirtualMachineScaleSetAutomaticInstanceRepairOutputReference
	_jsii_.Get(
		j,
		"automaticInstanceRepair",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) AutomaticInstanceRepairInput() *LinuxVirtualMachineScaleSetAutomaticInstanceRepair {
	var returns *LinuxVirtualMachineScaleSetAutomaticInstanceRepair
	_jsii_.Get(
		j,
		"automaticInstanceRepairInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) AutomaticOsUpgradePolicy() LinuxVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference {
	var returns LinuxVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference
	_jsii_.Get(
		j,
		"automaticOsUpgradePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) AutomaticOsUpgradePolicyInput() *LinuxVirtualMachineScaleSetAutomaticOsUpgradePolicy {
	var returns *LinuxVirtualMachineScaleSetAutomaticOsUpgradePolicy
	_jsii_.Get(
		j,
		"automaticOsUpgradePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) BootDiagnostics() LinuxVirtualMachineScaleSetBootDiagnosticsOutputReference {
	var returns LinuxVirtualMachineScaleSetBootDiagnosticsOutputReference
	_jsii_.Get(
		j,
		"bootDiagnostics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) BootDiagnosticsInput() *LinuxVirtualMachineScaleSetBootDiagnostics {
	var returns *LinuxVirtualMachineScaleSetBootDiagnostics
	_jsii_.Get(
		j,
		"bootDiagnosticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) CapacityReservationGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) CapacityReservationGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) ComputerNamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerNamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) ComputerNamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerNamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) CustomData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) CustomDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) DataDisk() LinuxVirtualMachineScaleSetDataDiskList {
	var returns LinuxVirtualMachineScaleSetDataDiskList
	_jsii_.Get(
		j,
		"dataDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) DataDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) DisablePasswordAuthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePasswordAuthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) DisablePasswordAuthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disablePasswordAuthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) DoNotRunExtensionsOnOverprovisionedMachines() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"doNotRunExtensionsOnOverprovisionedMachines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) DoNotRunExtensionsOnOverprovisionedMachinesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"doNotRunExtensionsOnOverprovisionedMachinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) EdgeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) EdgeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) EncryptionAtHostEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionAtHostEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) EncryptionAtHostEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionAtHostEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) EvictionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evictionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) EvictionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evictionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) Extension() LinuxVirtualMachineScaleSetExtensionList {
	var returns LinuxVirtualMachineScaleSetExtensionList
	_jsii_.Get(
		j,
		"extension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) ExtensionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) ExtensionOperationsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extensionOperationsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) ExtensionOperationsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extensionOperationsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) ExtensionsTimeBudget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionsTimeBudget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) ExtensionsTimeBudgetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionsTimeBudgetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) GalleryApplication() LinuxVirtualMachineScaleSetGalleryApplicationList {
	var returns LinuxVirtualMachineScaleSetGalleryApplicationList
	_jsii_.Get(
		j,
		"galleryApplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) GalleryApplicationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"galleryApplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) HealthProbeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthProbeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) HealthProbeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthProbeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) HostGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) HostGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) Identity() LinuxVirtualMachineScaleSetIdentityOutputReference {
	var returns LinuxVirtualMachineScaleSetIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) IdentityInput() *LinuxVirtualMachineScaleSetIdentity {
	var returns *LinuxVirtualMachineScaleSetIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) Instances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) InstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) MaxBidPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBidPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) MaxBidPriceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBidPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) NetworkInterface() LinuxVirtualMachineScaleSetNetworkInterfaceList {
	var returns LinuxVirtualMachineScaleSetNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) NetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) OsDisk() LinuxVirtualMachineScaleSetOsDiskOutputReference {
	var returns LinuxVirtualMachineScaleSetOsDiskOutputReference
	_jsii_.Get(
		j,
		"osDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) OsDiskInput() *LinuxVirtualMachineScaleSetOsDisk {
	var returns *LinuxVirtualMachineScaleSetOsDisk
	_jsii_.Get(
		j,
		"osDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) Overprovision() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overprovision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) OverprovisionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overprovisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) Plan() LinuxVirtualMachineScaleSetPlanOutputReference {
	var returns LinuxVirtualMachineScaleSetPlanOutputReference
	_jsii_.Get(
		j,
		"plan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) PlanInput() *LinuxVirtualMachineScaleSetPlan {
	var returns *LinuxVirtualMachineScaleSetPlan
	_jsii_.Get(
		j,
		"planInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) PlatformFaultDomainCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"platformFaultDomainCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) PlatformFaultDomainCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"platformFaultDomainCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) Priority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) PriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) ProvisionVmAgent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionVmAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) ProvisionVmAgentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionVmAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) ProximityPlacementGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proximityPlacementGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) ProximityPlacementGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proximityPlacementGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) RollingUpgradePolicy() LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference {
	var returns LinuxVirtualMachineScaleSetRollingUpgradePolicyOutputReference
	_jsii_.Get(
		j,
		"rollingUpgradePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) RollingUpgradePolicyInput() *LinuxVirtualMachineScaleSetRollingUpgradePolicy {
	var returns *LinuxVirtualMachineScaleSetRollingUpgradePolicy
	_jsii_.Get(
		j,
		"rollingUpgradePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) ScaleIn() LinuxVirtualMachineScaleSetScaleInOutputReference {
	var returns LinuxVirtualMachineScaleSetScaleInOutputReference
	_jsii_.Get(
		j,
		"scaleIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) ScaleInInput() *LinuxVirtualMachineScaleSetScaleIn {
	var returns *LinuxVirtualMachineScaleSetScaleIn
	_jsii_.Get(
		j,
		"scaleInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) Secret() LinuxVirtualMachineScaleSetSecretList {
	var returns LinuxVirtualMachineScaleSetSecretList
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) SecureBootEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secureBootEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) SecureBootEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secureBootEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) SinglePlacementGroup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singlePlacementGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) SinglePlacementGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singlePlacementGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) Sku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) SkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) SourceImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) SourceImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceImageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) SourceImageReference() LinuxVirtualMachineScaleSetSourceImageReferenceOutputReference {
	var returns LinuxVirtualMachineScaleSetSourceImageReferenceOutputReference
	_jsii_.Get(
		j,
		"sourceImageReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) SourceImageReferenceInput() *LinuxVirtualMachineScaleSetSourceImageReference {
	var returns *LinuxVirtualMachineScaleSetSourceImageReference
	_jsii_.Get(
		j,
		"sourceImageReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) SpotRestore() LinuxVirtualMachineScaleSetSpotRestoreOutputReference {
	var returns LinuxVirtualMachineScaleSetSpotRestoreOutputReference
	_jsii_.Get(
		j,
		"spotRestore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) SpotRestoreInput() *LinuxVirtualMachineScaleSetSpotRestore {
	var returns *LinuxVirtualMachineScaleSetSpotRestore
	_jsii_.Get(
		j,
		"spotRestoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) TerminationNotification() LinuxVirtualMachineScaleSetTerminationNotificationOutputReference {
	var returns LinuxVirtualMachineScaleSetTerminationNotificationOutputReference
	_jsii_.Get(
		j,
		"terminationNotification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) TerminationNotificationInput() *LinuxVirtualMachineScaleSetTerminationNotification {
	var returns *LinuxVirtualMachineScaleSetTerminationNotification
	_jsii_.Get(
		j,
		"terminationNotificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) Timeouts() LinuxVirtualMachineScaleSetTimeoutsOutputReference {
	var returns LinuxVirtualMachineScaleSetTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) UniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) UpgradeMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) UpgradeModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) VtpmEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vtpmEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) VtpmEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vtpmEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) ZoneBalance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneBalance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) ZoneBalanceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneBalanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) Zones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet) ZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zonesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/linux_virtual_machine_scale_set azurerm_linux_virtual_machine_scale_set} Resource.
func NewLinuxVirtualMachineScaleSet(scope constructs.Construct, id *string, config *LinuxVirtualMachineScaleSetConfig) LinuxVirtualMachineScaleSet {
	_init_.Initialize()

	if err := validateNewLinuxVirtualMachineScaleSetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_LinuxVirtualMachineScaleSet{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxVirtualMachineScaleSet.LinuxVirtualMachineScaleSet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/4.45.0/docs/resources/linux_virtual_machine_scale_set azurerm_linux_virtual_machine_scale_set} Resource.
func NewLinuxVirtualMachineScaleSet_Override(l LinuxVirtualMachineScaleSet, scope constructs.Construct, id *string, config *LinuxVirtualMachineScaleSetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.linuxVirtualMachineScaleSet.LinuxVirtualMachineScaleSet",
		[]interface{}{scope, id, config},
		l,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetAdminPassword(val *string) {
	if err := j.validateSetAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminPassword",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetAdminUsername(val *string) {
	if err := j.validateSetAdminUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminUsername",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetCapacityReservationGroupId(val *string) {
	if err := j.validateSetCapacityReservationGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityReservationGroupId",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetComputerNamePrefix(val *string) {
	if err := j.validateSetComputerNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computerNamePrefix",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetCustomData(val *string) {
	if err := j.validateSetCustomDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customData",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetDisablePasswordAuthentication(val interface{}) {
	if err := j.validateSetDisablePasswordAuthenticationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disablePasswordAuthentication",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetDoNotRunExtensionsOnOverprovisionedMachines(val interface{}) {
	if err := j.validateSetDoNotRunExtensionsOnOverprovisionedMachinesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"doNotRunExtensionsOnOverprovisionedMachines",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetEdgeZone(val *string) {
	if err := j.validateSetEdgeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeZone",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetEncryptionAtHostEnabled(val interface{}) {
	if err := j.validateSetEncryptionAtHostEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionAtHostEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetEvictionPolicy(val *string) {
	if err := j.validateSetEvictionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evictionPolicy",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetExtensionOperationsEnabled(val interface{}) {
	if err := j.validateSetExtensionOperationsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extensionOperationsEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetExtensionsTimeBudget(val *string) {
	if err := j.validateSetExtensionsTimeBudgetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extensionsTimeBudget",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetHealthProbeId(val *string) {
	if err := j.validateSetHealthProbeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthProbeId",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetHostGroupId(val *string) {
	if err := j.validateSetHostGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostGroupId",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetInstances(val *float64) {
	if err := j.validateSetInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instances",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetMaxBidPrice(val *float64) {
	if err := j.validateSetMaxBidPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBidPrice",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetOverprovision(val interface{}) {
	if err := j.validateSetOverprovisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overprovision",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetPlatformFaultDomainCount(val *float64) {
	if err := j.validateSetPlatformFaultDomainCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platformFaultDomainCount",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetPriority(val *string) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetProvisionVmAgent(val interface{}) {
	if err := j.validateSetProvisionVmAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionVmAgent",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetProximityPlacementGroupId(val *string) {
	if err := j.validateSetProximityPlacementGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proximityPlacementGroupId",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetSecureBootEnabled(val interface{}) {
	if err := j.validateSetSecureBootEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secureBootEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetSinglePlacementGroup(val interface{}) {
	if err := j.validateSetSinglePlacementGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singlePlacementGroup",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetSku(val *string) {
	if err := j.validateSetSkuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sku",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetSourceImageId(val *string) {
	if err := j.validateSetSourceImageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceImageId",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetUpgradeMode(val *string) {
	if err := j.validateSetUpgradeModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upgradeMode",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetVtpmEnabled(val interface{}) {
	if err := j.validateSetVtpmEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vtpmEnabled",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetZoneBalance(val interface{}) {
	if err := j.validateSetZoneBalanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneBalance",
		val,
	)
}

func (j *jsiiProxy_LinuxVirtualMachineScaleSet)SetZones(val *[]*string) {
	if err := j.validateSetZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zones",
		val,
	)
}

// Generates CDKTF code for importing a LinuxVirtualMachineScaleSet resource upon running "cdktf plan <stack-name>".
func LinuxVirtualMachineScaleSet_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateLinuxVirtualMachineScaleSet_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.linuxVirtualMachineScaleSet.LinuxVirtualMachineScaleSet",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func LinuxVirtualMachineScaleSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLinuxVirtualMachineScaleSet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.linuxVirtualMachineScaleSet.LinuxVirtualMachineScaleSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LinuxVirtualMachineScaleSet_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLinuxVirtualMachineScaleSet_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.linuxVirtualMachineScaleSet.LinuxVirtualMachineScaleSet",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func LinuxVirtualMachineScaleSet_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateLinuxVirtualMachineScaleSet_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.linuxVirtualMachineScaleSet.LinuxVirtualMachineScaleSet",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func LinuxVirtualMachineScaleSet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.linuxVirtualMachineScaleSet.LinuxVirtualMachineScaleSet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) AddMoveTarget(moveTarget *string) {
	if err := l.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) AddOverride(path *string, value interface{}) {
	if err := l.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := l.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) MoveFromId(id *string) {
	if err := l.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveFromId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) MoveTo(moveTarget *string, index interface{}) {
	if err := l.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) MoveToId(id *string) {
	if err := l.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"moveToId",
		[]interface{}{id},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) OverrideLogicalId(newLogicalId *string) {
	if err := l.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) PutAdditionalCapabilities(value *LinuxVirtualMachineScaleSetAdditionalCapabilities) {
	if err := l.validatePutAdditionalCapabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAdditionalCapabilities",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) PutAdminSshKey(value interface{}) {
	if err := l.validatePutAdminSshKeyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAdminSshKey",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) PutAutomaticInstanceRepair(value *LinuxVirtualMachineScaleSetAutomaticInstanceRepair) {
	if err := l.validatePutAutomaticInstanceRepairParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAutomaticInstanceRepair",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) PutAutomaticOsUpgradePolicy(value *LinuxVirtualMachineScaleSetAutomaticOsUpgradePolicy) {
	if err := l.validatePutAutomaticOsUpgradePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putAutomaticOsUpgradePolicy",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) PutBootDiagnostics(value *LinuxVirtualMachineScaleSetBootDiagnostics) {
	if err := l.validatePutBootDiagnosticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putBootDiagnostics",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) PutDataDisk(value interface{}) {
	if err := l.validatePutDataDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putDataDisk",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) PutExtension(value interface{}) {
	if err := l.validatePutExtensionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putExtension",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) PutGalleryApplication(value interface{}) {
	if err := l.validatePutGalleryApplicationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putGalleryApplication",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) PutIdentity(value *LinuxVirtualMachineScaleSetIdentity) {
	if err := l.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putIdentity",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) PutNetworkInterface(value interface{}) {
	if err := l.validatePutNetworkInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putNetworkInterface",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) PutOsDisk(value *LinuxVirtualMachineScaleSetOsDisk) {
	if err := l.validatePutOsDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putOsDisk",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) PutPlan(value *LinuxVirtualMachineScaleSetPlan) {
	if err := l.validatePutPlanParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putPlan",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) PutRollingUpgradePolicy(value *LinuxVirtualMachineScaleSetRollingUpgradePolicy) {
	if err := l.validatePutRollingUpgradePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putRollingUpgradePolicy",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) PutScaleIn(value *LinuxVirtualMachineScaleSetScaleIn) {
	if err := l.validatePutScaleInParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putScaleIn",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) PutSecret(value interface{}) {
	if err := l.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSecret",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) PutSourceImageReference(value *LinuxVirtualMachineScaleSetSourceImageReference) {
	if err := l.validatePutSourceImageReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSourceImageReference",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) PutSpotRestore(value *LinuxVirtualMachineScaleSetSpotRestore) {
	if err := l.validatePutSpotRestoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putSpotRestore",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) PutTerminationNotification(value *LinuxVirtualMachineScaleSetTerminationNotification) {
	if err := l.validatePutTerminationNotificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTerminationNotification",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) PutTimeouts(value *LinuxVirtualMachineScaleSetTimeouts) {
	if err := l.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetAdditionalCapabilities() {
	_jsii_.InvokeVoid(
		l,
		"resetAdditionalCapabilities",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetAdminPassword() {
	_jsii_.InvokeVoid(
		l,
		"resetAdminPassword",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetAdminSshKey() {
	_jsii_.InvokeVoid(
		l,
		"resetAdminSshKey",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetAutomaticInstanceRepair() {
	_jsii_.InvokeVoid(
		l,
		"resetAutomaticInstanceRepair",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetAutomaticOsUpgradePolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetAutomaticOsUpgradePolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetBootDiagnostics() {
	_jsii_.InvokeVoid(
		l,
		"resetBootDiagnostics",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetCapacityReservationGroupId() {
	_jsii_.InvokeVoid(
		l,
		"resetCapacityReservationGroupId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetComputerNamePrefix() {
	_jsii_.InvokeVoid(
		l,
		"resetComputerNamePrefix",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetCustomData() {
	_jsii_.InvokeVoid(
		l,
		"resetCustomData",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetDataDisk() {
	_jsii_.InvokeVoid(
		l,
		"resetDataDisk",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetDisablePasswordAuthentication() {
	_jsii_.InvokeVoid(
		l,
		"resetDisablePasswordAuthentication",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetDoNotRunExtensionsOnOverprovisionedMachines() {
	_jsii_.InvokeVoid(
		l,
		"resetDoNotRunExtensionsOnOverprovisionedMachines",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetEdgeZone() {
	_jsii_.InvokeVoid(
		l,
		"resetEdgeZone",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetEncryptionAtHostEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetEncryptionAtHostEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetEvictionPolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetEvictionPolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetExtension() {
	_jsii_.InvokeVoid(
		l,
		"resetExtension",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetExtensionOperationsEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetExtensionOperationsEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetExtensionsTimeBudget() {
	_jsii_.InvokeVoid(
		l,
		"resetExtensionsTimeBudget",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetGalleryApplication() {
	_jsii_.InvokeVoid(
		l,
		"resetGalleryApplication",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetHealthProbeId() {
	_jsii_.InvokeVoid(
		l,
		"resetHealthProbeId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetHostGroupId() {
	_jsii_.InvokeVoid(
		l,
		"resetHostGroupId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetId() {
	_jsii_.InvokeVoid(
		l,
		"resetId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetIdentity() {
	_jsii_.InvokeVoid(
		l,
		"resetIdentity",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetInstances() {
	_jsii_.InvokeVoid(
		l,
		"resetInstances",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetMaxBidPrice() {
	_jsii_.InvokeVoid(
		l,
		"resetMaxBidPrice",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetOverprovision() {
	_jsii_.InvokeVoid(
		l,
		"resetOverprovision",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		l,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetPlan() {
	_jsii_.InvokeVoid(
		l,
		"resetPlan",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetPlatformFaultDomainCount() {
	_jsii_.InvokeVoid(
		l,
		"resetPlatformFaultDomainCount",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetPriority() {
	_jsii_.InvokeVoid(
		l,
		"resetPriority",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetProvisionVmAgent() {
	_jsii_.InvokeVoid(
		l,
		"resetProvisionVmAgent",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetProximityPlacementGroupId() {
	_jsii_.InvokeVoid(
		l,
		"resetProximityPlacementGroupId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetRollingUpgradePolicy() {
	_jsii_.InvokeVoid(
		l,
		"resetRollingUpgradePolicy",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetScaleIn() {
	_jsii_.InvokeVoid(
		l,
		"resetScaleIn",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetSecret() {
	_jsii_.InvokeVoid(
		l,
		"resetSecret",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetSecureBootEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetSecureBootEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetSinglePlacementGroup() {
	_jsii_.InvokeVoid(
		l,
		"resetSinglePlacementGroup",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetSourceImageId() {
	_jsii_.InvokeVoid(
		l,
		"resetSourceImageId",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetSourceImageReference() {
	_jsii_.InvokeVoid(
		l,
		"resetSourceImageReference",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetSpotRestore() {
	_jsii_.InvokeVoid(
		l,
		"resetSpotRestore",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetTags() {
	_jsii_.InvokeVoid(
		l,
		"resetTags",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetTerminationNotification() {
	_jsii_.InvokeVoid(
		l,
		"resetTerminationNotification",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetTimeouts() {
	_jsii_.InvokeVoid(
		l,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetUpgradeMode() {
	_jsii_.InvokeVoid(
		l,
		"resetUpgradeMode",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetUserData() {
	_jsii_.InvokeVoid(
		l,
		"resetUserData",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetVtpmEnabled() {
	_jsii_.InvokeVoid(
		l,
		"resetVtpmEnabled",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetZoneBalance() {
	_jsii_.InvokeVoid(
		l,
		"resetZoneBalance",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ResetZones() {
	_jsii_.InvokeVoid(
		l,
		"resetZones",
		nil, // no parameters
	)
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		l,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LinuxVirtualMachineScaleSet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		l,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


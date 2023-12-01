// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsvirtualmachinescaleset

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v11/windowsvirtualmachinescaleset/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/windows_virtual_machine_scale_set azurerm_windows_virtual_machine_scale_set}.
type WindowsVirtualMachineScaleSet interface {
	cdktf.TerraformResource
	AdditionalCapabilities() WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference
	AdditionalCapabilitiesInput() *WindowsVirtualMachineScaleSetAdditionalCapabilities
	AdditionalUnattendContent() WindowsVirtualMachineScaleSetAdditionalUnattendContentList
	AdditionalUnattendContentInput() interface{}
	AdminPassword() *string
	SetAdminPassword(val *string)
	AdminPasswordInput() *string
	AdminUsername() *string
	SetAdminUsername(val *string)
	AdminUsernameInput() *string
	AutomaticInstanceRepair() WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference
	AutomaticInstanceRepairInput() *WindowsVirtualMachineScaleSetAutomaticInstanceRepair
	AutomaticOsUpgradePolicy() WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference
	AutomaticOsUpgradePolicyInput() *WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicy
	BootDiagnostics() WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference
	BootDiagnosticsInput() *WindowsVirtualMachineScaleSetBootDiagnostics
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
	DataDisk() WindowsVirtualMachineScaleSetDataDiskList
	DataDiskInput() interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DoNotRunExtensionsOnOverprovisionedMachines() interface{}
	SetDoNotRunExtensionsOnOverprovisionedMachines(val interface{})
	DoNotRunExtensionsOnOverprovisionedMachinesInput() interface{}
	EdgeZone() *string
	SetEdgeZone(val *string)
	EdgeZoneInput() *string
	EnableAutomaticUpdates() interface{}
	SetEnableAutomaticUpdates(val interface{})
	EnableAutomaticUpdatesInput() interface{}
	EncryptionAtHostEnabled() interface{}
	SetEncryptionAtHostEnabled(val interface{})
	EncryptionAtHostEnabledInput() interface{}
	EvictionPolicy() *string
	SetEvictionPolicy(val *string)
	EvictionPolicyInput() *string
	Extension() WindowsVirtualMachineScaleSetExtensionList
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
	GalleryApplication() WindowsVirtualMachineScaleSetGalleryApplicationList
	GalleryApplicationInput() interface{}
	GalleryApplications() WindowsVirtualMachineScaleSetGalleryApplicationsList
	GalleryApplicationsInput() interface{}
	HealthProbeId() *string
	SetHealthProbeId(val *string)
	HealthProbeIdInput() *string
	HostGroupId() *string
	SetHostGroupId(val *string)
	HostGroupIdInput() *string
	Id() *string
	SetId(val *string)
	Identity() WindowsVirtualMachineScaleSetIdentityOutputReference
	IdentityInput() *WindowsVirtualMachineScaleSetIdentity
	IdInput() *string
	Instances() *float64
	SetInstances(val *float64)
	InstancesInput() *float64
	LicenseType() *string
	SetLicenseType(val *string)
	LicenseTypeInput() *string
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
	NetworkInterface() WindowsVirtualMachineScaleSetNetworkInterfaceList
	NetworkInterfaceInput() interface{}
	// The tree node.
	Node() constructs.Node
	OsDisk() WindowsVirtualMachineScaleSetOsDiskOutputReference
	OsDiskInput() *WindowsVirtualMachineScaleSetOsDisk
	Overprovision() interface{}
	SetOverprovision(val interface{})
	OverprovisionInput() interface{}
	Plan() WindowsVirtualMachineScaleSetPlanOutputReference
	PlanInput() *WindowsVirtualMachineScaleSetPlan
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
	RollingUpgradePolicy() WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference
	RollingUpgradePolicyInput() *WindowsVirtualMachineScaleSetRollingUpgradePolicy
	ScaleIn() WindowsVirtualMachineScaleSetScaleInOutputReference
	ScaleInInput() *WindowsVirtualMachineScaleSetScaleIn
	ScaleInPolicy() *string
	SetScaleInPolicy(val *string)
	ScaleInPolicyInput() *string
	Secret() WindowsVirtualMachineScaleSetSecretList
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
	SourceImageReference() WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference
	SourceImageReferenceInput() *WindowsVirtualMachineScaleSetSourceImageReference
	SpotRestore() WindowsVirtualMachineScaleSetSpotRestoreOutputReference
	SpotRestoreInput() *WindowsVirtualMachineScaleSetSpotRestore
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TerminateNotification() WindowsVirtualMachineScaleSetTerminateNotificationOutputReference
	TerminateNotificationInput() *WindowsVirtualMachineScaleSetTerminateNotification
	TerminationNotification() WindowsVirtualMachineScaleSetTerminationNotificationOutputReference
	TerminationNotificationInput() *WindowsVirtualMachineScaleSetTerminationNotification
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() WindowsVirtualMachineScaleSetTimeoutsOutputReference
	TimeoutsInput() interface{}
	Timezone() *string
	SetTimezone(val *string)
	TimezoneInput() *string
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
	WinrmListener() WindowsVirtualMachineScaleSetWinrmListenerList
	WinrmListenerInput() interface{}
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutAdditionalCapabilities(value *WindowsVirtualMachineScaleSetAdditionalCapabilities)
	PutAdditionalUnattendContent(value interface{})
	PutAutomaticInstanceRepair(value *WindowsVirtualMachineScaleSetAutomaticInstanceRepair)
	PutAutomaticOsUpgradePolicy(value *WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicy)
	PutBootDiagnostics(value *WindowsVirtualMachineScaleSetBootDiagnostics)
	PutDataDisk(value interface{})
	PutExtension(value interface{})
	PutGalleryApplication(value interface{})
	PutGalleryApplications(value interface{})
	PutIdentity(value *WindowsVirtualMachineScaleSetIdentity)
	PutNetworkInterface(value interface{})
	PutOsDisk(value *WindowsVirtualMachineScaleSetOsDisk)
	PutPlan(value *WindowsVirtualMachineScaleSetPlan)
	PutRollingUpgradePolicy(value *WindowsVirtualMachineScaleSetRollingUpgradePolicy)
	PutScaleIn(value *WindowsVirtualMachineScaleSetScaleIn)
	PutSecret(value interface{})
	PutSourceImageReference(value *WindowsVirtualMachineScaleSetSourceImageReference)
	PutSpotRestore(value *WindowsVirtualMachineScaleSetSpotRestore)
	PutTerminateNotification(value *WindowsVirtualMachineScaleSetTerminateNotification)
	PutTerminationNotification(value *WindowsVirtualMachineScaleSetTerminationNotification)
	PutTimeouts(value *WindowsVirtualMachineScaleSetTimeouts)
	PutWinrmListener(value interface{})
	ResetAdditionalCapabilities()
	ResetAdditionalUnattendContent()
	ResetAutomaticInstanceRepair()
	ResetAutomaticOsUpgradePolicy()
	ResetBootDiagnostics()
	ResetCapacityReservationGroupId()
	ResetComputerNamePrefix()
	ResetCustomData()
	ResetDataDisk()
	ResetDoNotRunExtensionsOnOverprovisionedMachines()
	ResetEdgeZone()
	ResetEnableAutomaticUpdates()
	ResetEncryptionAtHostEnabled()
	ResetEvictionPolicy()
	ResetExtension()
	ResetExtensionOperationsEnabled()
	ResetExtensionsTimeBudget()
	ResetGalleryApplication()
	ResetGalleryApplications()
	ResetHealthProbeId()
	ResetHostGroupId()
	ResetId()
	ResetIdentity()
	ResetLicenseType()
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
	ResetScaleInPolicy()
	ResetSecret()
	ResetSecureBootEnabled()
	ResetSinglePlacementGroup()
	ResetSourceImageId()
	ResetSourceImageReference()
	ResetSpotRestore()
	ResetTags()
	ResetTerminateNotification()
	ResetTerminationNotification()
	ResetTimeouts()
	ResetTimezone()
	ResetUpgradeMode()
	ResetUserData()
	ResetVtpmEnabled()
	ResetWinrmListener()
	ResetZoneBalance()
	ResetZones()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for WindowsVirtualMachineScaleSet
type jsiiProxy_WindowsVirtualMachineScaleSet struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) AdditionalCapabilities() WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference {
	var returns WindowsVirtualMachineScaleSetAdditionalCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"additionalCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) AdditionalCapabilitiesInput() *WindowsVirtualMachineScaleSetAdditionalCapabilities {
	var returns *WindowsVirtualMachineScaleSetAdditionalCapabilities
	_jsii_.Get(
		j,
		"additionalCapabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) AdditionalUnattendContent() WindowsVirtualMachineScaleSetAdditionalUnattendContentList {
	var returns WindowsVirtualMachineScaleSetAdditionalUnattendContentList
	_jsii_.Get(
		j,
		"additionalUnattendContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) AdditionalUnattendContentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalUnattendContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) AdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) AdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) AdminUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) AdminUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) AutomaticInstanceRepair() WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference {
	var returns WindowsVirtualMachineScaleSetAutomaticInstanceRepairOutputReference
	_jsii_.Get(
		j,
		"automaticInstanceRepair",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) AutomaticInstanceRepairInput() *WindowsVirtualMachineScaleSetAutomaticInstanceRepair {
	var returns *WindowsVirtualMachineScaleSetAutomaticInstanceRepair
	_jsii_.Get(
		j,
		"automaticInstanceRepairInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) AutomaticOsUpgradePolicy() WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference {
	var returns WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicyOutputReference
	_jsii_.Get(
		j,
		"automaticOsUpgradePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) AutomaticOsUpgradePolicyInput() *WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicy {
	var returns *WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicy
	_jsii_.Get(
		j,
		"automaticOsUpgradePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) BootDiagnostics() WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference {
	var returns WindowsVirtualMachineScaleSetBootDiagnosticsOutputReference
	_jsii_.Get(
		j,
		"bootDiagnostics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) BootDiagnosticsInput() *WindowsVirtualMachineScaleSetBootDiagnostics {
	var returns *WindowsVirtualMachineScaleSetBootDiagnostics
	_jsii_.Get(
		j,
		"bootDiagnosticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) CapacityReservationGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) CapacityReservationGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ComputerNamePrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerNamePrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ComputerNamePrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerNamePrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) CustomData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) CustomDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) DataDisk() WindowsVirtualMachineScaleSetDataDiskList {
	var returns WindowsVirtualMachineScaleSetDataDiskList
	_jsii_.Get(
		j,
		"dataDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) DataDiskInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"dataDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) DoNotRunExtensionsOnOverprovisionedMachines() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"doNotRunExtensionsOnOverprovisionedMachines",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) DoNotRunExtensionsOnOverprovisionedMachinesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"doNotRunExtensionsOnOverprovisionedMachinesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) EdgeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) EdgeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) EnableAutomaticUpdates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticUpdates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) EnableAutomaticUpdatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticUpdatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) EncryptionAtHostEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionAtHostEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) EncryptionAtHostEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionAtHostEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) EvictionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evictionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) EvictionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evictionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Extension() WindowsVirtualMachineScaleSetExtensionList {
	var returns WindowsVirtualMachineScaleSetExtensionList
	_jsii_.Get(
		j,
		"extension",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ExtensionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extensionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ExtensionOperationsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extensionOperationsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ExtensionOperationsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"extensionOperationsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ExtensionsTimeBudget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionsTimeBudget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ExtensionsTimeBudgetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionsTimeBudgetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) GalleryApplication() WindowsVirtualMachineScaleSetGalleryApplicationList {
	var returns WindowsVirtualMachineScaleSetGalleryApplicationList
	_jsii_.Get(
		j,
		"galleryApplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) GalleryApplicationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"galleryApplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) GalleryApplications() WindowsVirtualMachineScaleSetGalleryApplicationsList {
	var returns WindowsVirtualMachineScaleSetGalleryApplicationsList
	_jsii_.Get(
		j,
		"galleryApplications",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) GalleryApplicationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"galleryApplicationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) HealthProbeId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthProbeId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) HealthProbeIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"healthProbeIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) HostGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) HostGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"hostGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Identity() WindowsVirtualMachineScaleSetIdentityOutputReference {
	var returns WindowsVirtualMachineScaleSetIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) IdentityInput() *WindowsVirtualMachineScaleSetIdentity {
	var returns *WindowsVirtualMachineScaleSetIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Instances() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) InstancesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"instancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) LicenseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) LicenseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) MaxBidPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBidPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) MaxBidPriceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBidPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) NetworkInterface() WindowsVirtualMachineScaleSetNetworkInterfaceList {
	var returns WindowsVirtualMachineScaleSetNetworkInterfaceList
	_jsii_.Get(
		j,
		"networkInterface",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) NetworkInterfaceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"networkInterfaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) OsDisk() WindowsVirtualMachineScaleSetOsDiskOutputReference {
	var returns WindowsVirtualMachineScaleSetOsDiskOutputReference
	_jsii_.Get(
		j,
		"osDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) OsDiskInput() *WindowsVirtualMachineScaleSetOsDisk {
	var returns *WindowsVirtualMachineScaleSetOsDisk
	_jsii_.Get(
		j,
		"osDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Overprovision() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overprovision",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) OverprovisionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"overprovisionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Plan() WindowsVirtualMachineScaleSetPlanOutputReference {
	var returns WindowsVirtualMachineScaleSetPlanOutputReference
	_jsii_.Get(
		j,
		"plan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) PlanInput() *WindowsVirtualMachineScaleSetPlan {
	var returns *WindowsVirtualMachineScaleSetPlan
	_jsii_.Get(
		j,
		"planInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) PlatformFaultDomainCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"platformFaultDomainCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) PlatformFaultDomainCountInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"platformFaultDomainCountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Priority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) PriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ProvisionVmAgent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionVmAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ProvisionVmAgentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionVmAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ProximityPlacementGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proximityPlacementGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ProximityPlacementGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proximityPlacementGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) RollingUpgradePolicy() WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference {
	var returns WindowsVirtualMachineScaleSetRollingUpgradePolicyOutputReference
	_jsii_.Get(
		j,
		"rollingUpgradePolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) RollingUpgradePolicyInput() *WindowsVirtualMachineScaleSetRollingUpgradePolicy {
	var returns *WindowsVirtualMachineScaleSetRollingUpgradePolicy
	_jsii_.Get(
		j,
		"rollingUpgradePolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ScaleIn() WindowsVirtualMachineScaleSetScaleInOutputReference {
	var returns WindowsVirtualMachineScaleSetScaleInOutputReference
	_jsii_.Get(
		j,
		"scaleIn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ScaleInInput() *WindowsVirtualMachineScaleSetScaleIn {
	var returns *WindowsVirtualMachineScaleSetScaleIn
	_jsii_.Get(
		j,
		"scaleInInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ScaleInPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleInPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ScaleInPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scaleInPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Secret() WindowsVirtualMachineScaleSetSecretList {
	var returns WindowsVirtualMachineScaleSetSecretList
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SecureBootEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secureBootEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SecureBootEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secureBootEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SinglePlacementGroup() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singlePlacementGroup",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SinglePlacementGroupInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"singlePlacementGroupInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Sku() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sku",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SkuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"skuInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SourceImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SourceImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceImageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SourceImageReference() WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference {
	var returns WindowsVirtualMachineScaleSetSourceImageReferenceOutputReference
	_jsii_.Get(
		j,
		"sourceImageReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SourceImageReferenceInput() *WindowsVirtualMachineScaleSetSourceImageReference {
	var returns *WindowsVirtualMachineScaleSetSourceImageReference
	_jsii_.Get(
		j,
		"sourceImageReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SpotRestore() WindowsVirtualMachineScaleSetSpotRestoreOutputReference {
	var returns WindowsVirtualMachineScaleSetSpotRestoreOutputReference
	_jsii_.Get(
		j,
		"spotRestore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) SpotRestoreInput() *WindowsVirtualMachineScaleSetSpotRestore {
	var returns *WindowsVirtualMachineScaleSetSpotRestore
	_jsii_.Get(
		j,
		"spotRestoreInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) TerminateNotification() WindowsVirtualMachineScaleSetTerminateNotificationOutputReference {
	var returns WindowsVirtualMachineScaleSetTerminateNotificationOutputReference
	_jsii_.Get(
		j,
		"terminateNotification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) TerminateNotificationInput() *WindowsVirtualMachineScaleSetTerminateNotification {
	var returns *WindowsVirtualMachineScaleSetTerminateNotification
	_jsii_.Get(
		j,
		"terminateNotificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) TerminationNotification() WindowsVirtualMachineScaleSetTerminationNotificationOutputReference {
	var returns WindowsVirtualMachineScaleSetTerminationNotificationOutputReference
	_jsii_.Get(
		j,
		"terminationNotification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) TerminationNotificationInput() *WindowsVirtualMachineScaleSetTerminationNotification {
	var returns *WindowsVirtualMachineScaleSetTerminationNotification
	_jsii_.Get(
		j,
		"terminationNotificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Timeouts() WindowsVirtualMachineScaleSetTimeoutsOutputReference {
	var returns WindowsVirtualMachineScaleSetTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) UniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) UpgradeMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) UpgradeModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upgradeModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) VtpmEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vtpmEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) VtpmEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vtpmEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) WinrmListener() WindowsVirtualMachineScaleSetWinrmListenerList {
	var returns WindowsVirtualMachineScaleSetWinrmListenerList
	_jsii_.Get(
		j,
		"winrmListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) WinrmListenerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"winrmListenerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ZoneBalance() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneBalance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ZoneBalanceInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"zoneBalanceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) Zones() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zones",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet) ZonesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"zonesInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/windows_virtual_machine_scale_set azurerm_windows_virtual_machine_scale_set} Resource.
func NewWindowsVirtualMachineScaleSet(scope constructs.Construct, id *string, config *WindowsVirtualMachineScaleSetConfig) WindowsVirtualMachineScaleSet {
	_init_.Initialize()

	if err := validateNewWindowsVirtualMachineScaleSetParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_WindowsVirtualMachineScaleSet{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSet",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.83.0/docs/resources/windows_virtual_machine_scale_set azurerm_windows_virtual_machine_scale_set} Resource.
func NewWindowsVirtualMachineScaleSet_Override(w WindowsVirtualMachineScaleSet, scope constructs.Construct, id *string, config *WindowsVirtualMachineScaleSetConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSet",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetAdminPassword(val *string) {
	if err := j.validateSetAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminPassword",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetAdminUsername(val *string) {
	if err := j.validateSetAdminUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminUsername",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetCapacityReservationGroupId(val *string) {
	if err := j.validateSetCapacityReservationGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityReservationGroupId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetComputerNamePrefix(val *string) {
	if err := j.validateSetComputerNamePrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computerNamePrefix",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetCustomData(val *string) {
	if err := j.validateSetCustomDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customData",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetDoNotRunExtensionsOnOverprovisionedMachines(val interface{}) {
	if err := j.validateSetDoNotRunExtensionsOnOverprovisionedMachinesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"doNotRunExtensionsOnOverprovisionedMachines",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetEdgeZone(val *string) {
	if err := j.validateSetEdgeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeZone",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetEnableAutomaticUpdates(val interface{}) {
	if err := j.validateSetEnableAutomaticUpdatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAutomaticUpdates",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetEncryptionAtHostEnabled(val interface{}) {
	if err := j.validateSetEncryptionAtHostEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionAtHostEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetEvictionPolicy(val *string) {
	if err := j.validateSetEvictionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evictionPolicy",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetExtensionOperationsEnabled(val interface{}) {
	if err := j.validateSetExtensionOperationsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extensionOperationsEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetExtensionsTimeBudget(val *string) {
	if err := j.validateSetExtensionsTimeBudgetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extensionsTimeBudget",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetHealthProbeId(val *string) {
	if err := j.validateSetHealthProbeIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"healthProbeId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetHostGroupId(val *string) {
	if err := j.validateSetHostGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hostGroupId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetInstances(val *float64) {
	if err := j.validateSetInstancesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instances",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetLicenseType(val *string) {
	if err := j.validateSetLicenseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseType",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetMaxBidPrice(val *float64) {
	if err := j.validateSetMaxBidPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBidPrice",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetOverprovision(val interface{}) {
	if err := j.validateSetOverprovisionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"overprovision",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetPlatformFaultDomainCount(val *float64) {
	if err := j.validateSetPlatformFaultDomainCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platformFaultDomainCount",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetPriority(val *string) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetProvisionVmAgent(val interface{}) {
	if err := j.validateSetProvisionVmAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionVmAgent",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetProximityPlacementGroupId(val *string) {
	if err := j.validateSetProximityPlacementGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proximityPlacementGroupId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetScaleInPolicy(val *string) {
	if err := j.validateSetScaleInPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scaleInPolicy",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetSecureBootEnabled(val interface{}) {
	if err := j.validateSetSecureBootEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secureBootEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetSinglePlacementGroup(val interface{}) {
	if err := j.validateSetSinglePlacementGroupParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"singlePlacementGroup",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetSku(val *string) {
	if err := j.validateSetSkuParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sku",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetSourceImageId(val *string) {
	if err := j.validateSetSourceImageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceImageId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetUpgradeMode(val *string) {
	if err := j.validateSetUpgradeModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upgradeMode",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetVtpmEnabled(val interface{}) {
	if err := j.validateSetVtpmEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vtpmEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetZoneBalance(val interface{}) {
	if err := j.validateSetZoneBalanceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zoneBalance",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachineScaleSet)SetZones(val *[]*string) {
	if err := j.validateSetZonesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zones",
		val,
	)
}

// Generates CDKTF code for importing a WindowsVirtualMachineScaleSet resource upon running "cdktf plan <stack-name>".
func WindowsVirtualMachineScaleSet_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateWindowsVirtualMachineScaleSet_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSet",
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
func WindowsVirtualMachineScaleSet_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWindowsVirtualMachineScaleSet_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSet",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WindowsVirtualMachineScaleSet_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWindowsVirtualMachineScaleSet_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSet",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WindowsVirtualMachineScaleSet_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWindowsVirtualMachineScaleSet_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSet",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WindowsVirtualMachineScaleSet_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.windowsVirtualMachineScaleSet.WindowsVirtualMachineScaleSet",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) AddMoveTarget(moveTarget *string) {
	if err := w.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := w.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := w.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		w,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := w.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		w,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := w.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		w,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := w.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		w,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := w.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		w,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) GetStringAttribute(terraformAttribute *string) *string {
	if err := w.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		w,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := w.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		w,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := w.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := w.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		w,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) MoveTo(moveTarget *string, index interface{}) {
	if err := w.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutAdditionalCapabilities(value *WindowsVirtualMachineScaleSetAdditionalCapabilities) {
	if err := w.validatePutAdditionalCapabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAdditionalCapabilities",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutAdditionalUnattendContent(value interface{}) {
	if err := w.validatePutAdditionalUnattendContentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAdditionalUnattendContent",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutAutomaticInstanceRepair(value *WindowsVirtualMachineScaleSetAutomaticInstanceRepair) {
	if err := w.validatePutAutomaticInstanceRepairParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAutomaticInstanceRepair",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutAutomaticOsUpgradePolicy(value *WindowsVirtualMachineScaleSetAutomaticOsUpgradePolicy) {
	if err := w.validatePutAutomaticOsUpgradePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAutomaticOsUpgradePolicy",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutBootDiagnostics(value *WindowsVirtualMachineScaleSetBootDiagnostics) {
	if err := w.validatePutBootDiagnosticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBootDiagnostics",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutDataDisk(value interface{}) {
	if err := w.validatePutDataDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putDataDisk",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutExtension(value interface{}) {
	if err := w.validatePutExtensionParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putExtension",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutGalleryApplication(value interface{}) {
	if err := w.validatePutGalleryApplicationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putGalleryApplication",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutGalleryApplications(value interface{}) {
	if err := w.validatePutGalleryApplicationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putGalleryApplications",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutIdentity(value *WindowsVirtualMachineScaleSetIdentity) {
	if err := w.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putIdentity",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutNetworkInterface(value interface{}) {
	if err := w.validatePutNetworkInterfaceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putNetworkInterface",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutOsDisk(value *WindowsVirtualMachineScaleSetOsDisk) {
	if err := w.validatePutOsDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putOsDisk",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutPlan(value *WindowsVirtualMachineScaleSetPlan) {
	if err := w.validatePutPlanParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putPlan",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutRollingUpgradePolicy(value *WindowsVirtualMachineScaleSetRollingUpgradePolicy) {
	if err := w.validatePutRollingUpgradePolicyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putRollingUpgradePolicy",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutScaleIn(value *WindowsVirtualMachineScaleSetScaleIn) {
	if err := w.validatePutScaleInParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putScaleIn",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutSecret(value interface{}) {
	if err := w.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSecret",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutSourceImageReference(value *WindowsVirtualMachineScaleSetSourceImageReference) {
	if err := w.validatePutSourceImageReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSourceImageReference",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutSpotRestore(value *WindowsVirtualMachineScaleSetSpotRestore) {
	if err := w.validatePutSpotRestoreParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSpotRestore",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutTerminateNotification(value *WindowsVirtualMachineScaleSetTerminateNotification) {
	if err := w.validatePutTerminateNotificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTerminateNotification",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutTerminationNotification(value *WindowsVirtualMachineScaleSetTerminationNotification) {
	if err := w.validatePutTerminationNotificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTerminationNotification",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutTimeouts(value *WindowsVirtualMachineScaleSetTimeouts) {
	if err := w.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) PutWinrmListener(value interface{}) {
	if err := w.validatePutWinrmListenerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putWinrmListener",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetAdditionalCapabilities() {
	_jsii_.InvokeVoid(
		w,
		"resetAdditionalCapabilities",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetAdditionalUnattendContent() {
	_jsii_.InvokeVoid(
		w,
		"resetAdditionalUnattendContent",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetAutomaticInstanceRepair() {
	_jsii_.InvokeVoid(
		w,
		"resetAutomaticInstanceRepair",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetAutomaticOsUpgradePolicy() {
	_jsii_.InvokeVoid(
		w,
		"resetAutomaticOsUpgradePolicy",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetBootDiagnostics() {
	_jsii_.InvokeVoid(
		w,
		"resetBootDiagnostics",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetCapacityReservationGroupId() {
	_jsii_.InvokeVoid(
		w,
		"resetCapacityReservationGroupId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetComputerNamePrefix() {
	_jsii_.InvokeVoid(
		w,
		"resetComputerNamePrefix",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetCustomData() {
	_jsii_.InvokeVoid(
		w,
		"resetCustomData",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetDataDisk() {
	_jsii_.InvokeVoid(
		w,
		"resetDataDisk",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetDoNotRunExtensionsOnOverprovisionedMachines() {
	_jsii_.InvokeVoid(
		w,
		"resetDoNotRunExtensionsOnOverprovisionedMachines",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetEdgeZone() {
	_jsii_.InvokeVoid(
		w,
		"resetEdgeZone",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetEnableAutomaticUpdates() {
	_jsii_.InvokeVoid(
		w,
		"resetEnableAutomaticUpdates",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetEncryptionAtHostEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetEncryptionAtHostEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetEvictionPolicy() {
	_jsii_.InvokeVoid(
		w,
		"resetEvictionPolicy",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetExtension() {
	_jsii_.InvokeVoid(
		w,
		"resetExtension",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetExtensionOperationsEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetExtensionOperationsEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetExtensionsTimeBudget() {
	_jsii_.InvokeVoid(
		w,
		"resetExtensionsTimeBudget",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetGalleryApplication() {
	_jsii_.InvokeVoid(
		w,
		"resetGalleryApplication",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetGalleryApplications() {
	_jsii_.InvokeVoid(
		w,
		"resetGalleryApplications",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetHealthProbeId() {
	_jsii_.InvokeVoid(
		w,
		"resetHealthProbeId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetHostGroupId() {
	_jsii_.InvokeVoid(
		w,
		"resetHostGroupId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetId() {
	_jsii_.InvokeVoid(
		w,
		"resetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetIdentity() {
	_jsii_.InvokeVoid(
		w,
		"resetIdentity",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetLicenseType() {
	_jsii_.InvokeVoid(
		w,
		"resetLicenseType",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetMaxBidPrice() {
	_jsii_.InvokeVoid(
		w,
		"resetMaxBidPrice",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetOverprovision() {
	_jsii_.InvokeVoid(
		w,
		"resetOverprovision",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetPlan() {
	_jsii_.InvokeVoid(
		w,
		"resetPlan",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetPlatformFaultDomainCount() {
	_jsii_.InvokeVoid(
		w,
		"resetPlatformFaultDomainCount",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetPriority() {
	_jsii_.InvokeVoid(
		w,
		"resetPriority",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetProvisionVmAgent() {
	_jsii_.InvokeVoid(
		w,
		"resetProvisionVmAgent",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetProximityPlacementGroupId() {
	_jsii_.InvokeVoid(
		w,
		"resetProximityPlacementGroupId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetRollingUpgradePolicy() {
	_jsii_.InvokeVoid(
		w,
		"resetRollingUpgradePolicy",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetScaleIn() {
	_jsii_.InvokeVoid(
		w,
		"resetScaleIn",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetScaleInPolicy() {
	_jsii_.InvokeVoid(
		w,
		"resetScaleInPolicy",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetSecret() {
	_jsii_.InvokeVoid(
		w,
		"resetSecret",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetSecureBootEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetSecureBootEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetSinglePlacementGroup() {
	_jsii_.InvokeVoid(
		w,
		"resetSinglePlacementGroup",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetSourceImageId() {
	_jsii_.InvokeVoid(
		w,
		"resetSourceImageId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetSourceImageReference() {
	_jsii_.InvokeVoid(
		w,
		"resetSourceImageReference",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetSpotRestore() {
	_jsii_.InvokeVoid(
		w,
		"resetSpotRestore",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetTags() {
	_jsii_.InvokeVoid(
		w,
		"resetTags",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetTerminateNotification() {
	_jsii_.InvokeVoid(
		w,
		"resetTerminateNotification",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetTerminationNotification() {
	_jsii_.InvokeVoid(
		w,
		"resetTerminationNotification",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetTimeouts() {
	_jsii_.InvokeVoid(
		w,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetTimezone() {
	_jsii_.InvokeVoid(
		w,
		"resetTimezone",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetUpgradeMode() {
	_jsii_.InvokeVoid(
		w,
		"resetUpgradeMode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetUserData() {
	_jsii_.InvokeVoid(
		w,
		"resetUserData",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetVtpmEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetVtpmEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetWinrmListener() {
	_jsii_.InvokeVoid(
		w,
		"resetWinrmListener",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetZoneBalance() {
	_jsii_.InvokeVoid(
		w,
		"resetZoneBalance",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ResetZones() {
	_jsii_.InvokeVoid(
		w,
		"resetZones",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachineScaleSet) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package windowsvirtualmachine

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-azurerm-go/azurerm/v12/windowsvirtualmachine/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/windows_virtual_machine azurerm_windows_virtual_machine}.
type WindowsVirtualMachine interface {
	cdktf.TerraformResource
	AdditionalCapabilities() WindowsVirtualMachineAdditionalCapabilitiesOutputReference
	AdditionalCapabilitiesInput() *WindowsVirtualMachineAdditionalCapabilities
	AdditionalUnattendContent() WindowsVirtualMachineAdditionalUnattendContentList
	AdditionalUnattendContentInput() interface{}
	AdminPassword() *string
	SetAdminPassword(val *string)
	AdminPasswordInput() *string
	AdminUsername() *string
	SetAdminUsername(val *string)
	AdminUsernameInput() *string
	AllowExtensionOperations() interface{}
	SetAllowExtensionOperations(val interface{})
	AllowExtensionOperationsInput() interface{}
	AvailabilitySetId() *string
	SetAvailabilitySetId(val *string)
	AvailabilitySetIdInput() *string
	BootDiagnostics() WindowsVirtualMachineBootDiagnosticsOutputReference
	BootDiagnosticsInput() *WindowsVirtualMachineBootDiagnostics
	BypassPlatformSafetyChecksOnUserScheduleEnabled() interface{}
	SetBypassPlatformSafetyChecksOnUserScheduleEnabled(val interface{})
	BypassPlatformSafetyChecksOnUserScheduleEnabledInput() interface{}
	CapacityReservationGroupId() *string
	SetCapacityReservationGroupId(val *string)
	CapacityReservationGroupIdInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ComputerName() *string
	SetComputerName(val *string)
	ComputerNameInput() *string
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
	DedicatedHostGroupId() *string
	SetDedicatedHostGroupId(val *string)
	DedicatedHostGroupIdInput() *string
	DedicatedHostId() *string
	SetDedicatedHostId(val *string)
	DedicatedHostIdInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DiskControllerType() *string
	SetDiskControllerType(val *string)
	DiskControllerTypeInput() *string
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
	GalleryApplication() WindowsVirtualMachineGalleryApplicationList
	GalleryApplicationInput() interface{}
	HotpatchingEnabled() interface{}
	SetHotpatchingEnabled(val interface{})
	HotpatchingEnabledInput() interface{}
	Id() *string
	SetId(val *string)
	Identity() WindowsVirtualMachineIdentityOutputReference
	IdentityInput() *WindowsVirtualMachineIdentity
	IdInput() *string
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
	NetworkInterfaceIds() *[]*string
	SetNetworkInterfaceIds(val *[]*string)
	NetworkInterfaceIdsInput() *[]*string
	// The tree node.
	Node() constructs.Node
	OsDisk() WindowsVirtualMachineOsDiskOutputReference
	OsDiskInput() *WindowsVirtualMachineOsDisk
	OsImageNotification() WindowsVirtualMachineOsImageNotificationOutputReference
	OsImageNotificationInput() *WindowsVirtualMachineOsImageNotification
	PatchAssessmentMode() *string
	SetPatchAssessmentMode(val *string)
	PatchAssessmentModeInput() *string
	PatchMode() *string
	SetPatchMode(val *string)
	PatchModeInput() *string
	Plan() WindowsVirtualMachinePlanOutputReference
	PlanInput() *WindowsVirtualMachinePlan
	PlatformFaultDomain() *float64
	SetPlatformFaultDomain(val *float64)
	PlatformFaultDomainInput() *float64
	Priority() *string
	SetPriority(val *string)
	PriorityInput() *string
	PrivateIpAddress() *string
	PrivateIpAddresses() *[]*string
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
	PublicIpAddress() *string
	PublicIpAddresses() *[]*string
	// Experimental.
	RawOverrides() interface{}
	RebootSetting() *string
	SetRebootSetting(val *string)
	RebootSettingInput() *string
	ResourceGroupName() *string
	SetResourceGroupName(val *string)
	ResourceGroupNameInput() *string
	Secret() WindowsVirtualMachineSecretList
	SecretInput() interface{}
	SecureBootEnabled() interface{}
	SetSecureBootEnabled(val interface{})
	SecureBootEnabledInput() interface{}
	Size() *string
	SetSize(val *string)
	SizeInput() *string
	SourceImageId() *string
	SetSourceImageId(val *string)
	SourceImageIdInput() *string
	SourceImageReference() WindowsVirtualMachineSourceImageReferenceOutputReference
	SourceImageReferenceInput() *WindowsVirtualMachineSourceImageReference
	Tags() *map[string]*string
	SetTags(val *map[string]*string)
	TagsInput() *map[string]*string
	TerminationNotification() WindowsVirtualMachineTerminationNotificationOutputReference
	TerminationNotificationInput() *WindowsVirtualMachineTerminationNotification
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Timeouts() WindowsVirtualMachineTimeoutsOutputReference
	TimeoutsInput() interface{}
	Timezone() *string
	SetTimezone(val *string)
	TimezoneInput() *string
	UserData() *string
	SetUserData(val *string)
	UserDataInput() *string
	VirtualMachineId() *string
	VirtualMachineScaleSetId() *string
	SetVirtualMachineScaleSetId(val *string)
	VirtualMachineScaleSetIdInput() *string
	VmAgentPlatformUpdatesEnabled() interface{}
	SetVmAgentPlatformUpdatesEnabled(val interface{})
	VmAgentPlatformUpdatesEnabledInput() interface{}
	VtpmEnabled() interface{}
	SetVtpmEnabled(val interface{})
	VtpmEnabledInput() interface{}
	WinrmListener() WindowsVirtualMachineWinrmListenerList
	WinrmListenerInput() interface{}
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
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
	PutAdditionalCapabilities(value *WindowsVirtualMachineAdditionalCapabilities)
	PutAdditionalUnattendContent(value interface{})
	PutBootDiagnostics(value *WindowsVirtualMachineBootDiagnostics)
	PutGalleryApplication(value interface{})
	PutIdentity(value *WindowsVirtualMachineIdentity)
	PutOsDisk(value *WindowsVirtualMachineOsDisk)
	PutOsImageNotification(value *WindowsVirtualMachineOsImageNotification)
	PutPlan(value *WindowsVirtualMachinePlan)
	PutSecret(value interface{})
	PutSourceImageReference(value *WindowsVirtualMachineSourceImageReference)
	PutTerminationNotification(value *WindowsVirtualMachineTerminationNotification)
	PutTimeouts(value *WindowsVirtualMachineTimeouts)
	PutWinrmListener(value interface{})
	ResetAdditionalCapabilities()
	ResetAdditionalUnattendContent()
	ResetAllowExtensionOperations()
	ResetAvailabilitySetId()
	ResetBootDiagnostics()
	ResetBypassPlatformSafetyChecksOnUserScheduleEnabled()
	ResetCapacityReservationGroupId()
	ResetComputerName()
	ResetCustomData()
	ResetDedicatedHostGroupId()
	ResetDedicatedHostId()
	ResetDiskControllerType()
	ResetEdgeZone()
	ResetEnableAutomaticUpdates()
	ResetEncryptionAtHostEnabled()
	ResetEvictionPolicy()
	ResetExtensionsTimeBudget()
	ResetGalleryApplication()
	ResetHotpatchingEnabled()
	ResetId()
	ResetIdentity()
	ResetLicenseType()
	ResetMaxBidPrice()
	ResetOsImageNotification()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPatchAssessmentMode()
	ResetPatchMode()
	ResetPlan()
	ResetPlatformFaultDomain()
	ResetPriority()
	ResetProvisionVmAgent()
	ResetProximityPlacementGroupId()
	ResetRebootSetting()
	ResetSecret()
	ResetSecureBootEnabled()
	ResetSourceImageId()
	ResetSourceImageReference()
	ResetTags()
	ResetTerminationNotification()
	ResetTimeouts()
	ResetTimezone()
	ResetUserData()
	ResetVirtualMachineScaleSetId()
	ResetVmAgentPlatformUpdatesEnabled()
	ResetVtpmEnabled()
	ResetWinrmListener()
	ResetZone()
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

// The jsii proxy struct for WindowsVirtualMachine
type jsiiProxy_WindowsVirtualMachine struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_WindowsVirtualMachine) AdditionalCapabilities() WindowsVirtualMachineAdditionalCapabilitiesOutputReference {
	var returns WindowsVirtualMachineAdditionalCapabilitiesOutputReference
	_jsii_.Get(
		j,
		"additionalCapabilities",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) AdditionalCapabilitiesInput() *WindowsVirtualMachineAdditionalCapabilities {
	var returns *WindowsVirtualMachineAdditionalCapabilities
	_jsii_.Get(
		j,
		"additionalCapabilitiesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) AdditionalUnattendContent() WindowsVirtualMachineAdditionalUnattendContentList {
	var returns WindowsVirtualMachineAdditionalUnattendContentList
	_jsii_.Get(
		j,
		"additionalUnattendContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) AdditionalUnattendContentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"additionalUnattendContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) AdminPassword() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPassword",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) AdminPasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminPasswordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) AdminUsername() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsername",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) AdminUsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"adminUsernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) AllowExtensionOperations() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowExtensionOperations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) AllowExtensionOperationsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowExtensionOperationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) AvailabilitySetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilitySetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) AvailabilitySetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"availabilitySetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) BootDiagnostics() WindowsVirtualMachineBootDiagnosticsOutputReference {
	var returns WindowsVirtualMachineBootDiagnosticsOutputReference
	_jsii_.Get(
		j,
		"bootDiagnostics",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) BootDiagnosticsInput() *WindowsVirtualMachineBootDiagnostics {
	var returns *WindowsVirtualMachineBootDiagnostics
	_jsii_.Get(
		j,
		"bootDiagnosticsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) BypassPlatformSafetyChecksOnUserScheduleEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bypassPlatformSafetyChecksOnUserScheduleEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) BypassPlatformSafetyChecksOnUserScheduleEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"bypassPlatformSafetyChecksOnUserScheduleEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) CapacityReservationGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) CapacityReservationGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"capacityReservationGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) ComputerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) ComputerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"computerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) CustomData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) CustomDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"customDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) DedicatedHostGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedHostGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) DedicatedHostGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedHostGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) DedicatedHostId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedHostId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) DedicatedHostIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dedicatedHostIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) DiskControllerType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskControllerType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) DiskControllerTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"diskControllerTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) EdgeZone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeZone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) EdgeZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"edgeZoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) EnableAutomaticUpdates() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticUpdates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) EnableAutomaticUpdatesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAutomaticUpdatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) EncryptionAtHostEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionAtHostEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) EncryptionAtHostEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"encryptionAtHostEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) EvictionPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evictionPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) EvictionPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"evictionPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) ExtensionsTimeBudget() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionsTimeBudget",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) ExtensionsTimeBudgetInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"extensionsTimeBudgetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) GalleryApplication() WindowsVirtualMachineGalleryApplicationList {
	var returns WindowsVirtualMachineGalleryApplicationList
	_jsii_.Get(
		j,
		"galleryApplication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) GalleryApplicationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"galleryApplicationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) HotpatchingEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hotpatchingEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) HotpatchingEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hotpatchingEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) Identity() WindowsVirtualMachineIdentityOutputReference {
	var returns WindowsVirtualMachineIdentityOutputReference
	_jsii_.Get(
		j,
		"identity",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) IdentityInput() *WindowsVirtualMachineIdentity {
	var returns *WindowsVirtualMachineIdentity
	_jsii_.Get(
		j,
		"identityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) LicenseType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) LicenseTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"licenseTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) Location() *string {
	var returns *string
	_jsii_.Get(
		j,
		"location",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) LocationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) MaxBidPrice() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBidPrice",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) MaxBidPriceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxBidPriceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) NetworkInterfaceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkInterfaceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) NetworkInterfaceIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"networkInterfaceIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) OsDisk() WindowsVirtualMachineOsDiskOutputReference {
	var returns WindowsVirtualMachineOsDiskOutputReference
	_jsii_.Get(
		j,
		"osDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) OsDiskInput() *WindowsVirtualMachineOsDisk {
	var returns *WindowsVirtualMachineOsDisk
	_jsii_.Get(
		j,
		"osDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) OsImageNotification() WindowsVirtualMachineOsImageNotificationOutputReference {
	var returns WindowsVirtualMachineOsImageNotificationOutputReference
	_jsii_.Get(
		j,
		"osImageNotification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) OsImageNotificationInput() *WindowsVirtualMachineOsImageNotification {
	var returns *WindowsVirtualMachineOsImageNotification
	_jsii_.Get(
		j,
		"osImageNotificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) PatchAssessmentMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchAssessmentMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) PatchAssessmentModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchAssessmentModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) PatchMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) PatchModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"patchModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) Plan() WindowsVirtualMachinePlanOutputReference {
	var returns WindowsVirtualMachinePlanOutputReference
	_jsii_.Get(
		j,
		"plan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) PlanInput() *WindowsVirtualMachinePlan {
	var returns *WindowsVirtualMachinePlan
	_jsii_.Get(
		j,
		"planInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) PlatformFaultDomain() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"platformFaultDomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) PlatformFaultDomainInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"platformFaultDomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) Priority() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priority",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) PriorityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"priorityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) PrivateIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) PrivateIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"privateIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) ProvisionVmAgent() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionVmAgent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) ProvisionVmAgentInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"provisionVmAgentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) ProximityPlacementGroupId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proximityPlacementGroupId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) ProximityPlacementGroupIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"proximityPlacementGroupIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) PublicIpAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"publicIpAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) PublicIpAddresses() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"publicIpAddresses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) RebootSetting() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rebootSetting",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) RebootSettingInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"rebootSettingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) ResourceGroupName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) ResourceGroupNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"resourceGroupNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) Secret() WindowsVirtualMachineSecretList {
	var returns WindowsVirtualMachineSecretList
	_jsii_.Get(
		j,
		"secret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) SecretInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) SecureBootEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secureBootEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) SecureBootEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"secureBootEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) Size() *string {
	var returns *string
	_jsii_.Get(
		j,
		"size",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) SizeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) SourceImageId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceImageId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) SourceImageIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"sourceImageIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) SourceImageReference() WindowsVirtualMachineSourceImageReferenceOutputReference {
	var returns WindowsVirtualMachineSourceImageReferenceOutputReference
	_jsii_.Get(
		j,
		"sourceImageReference",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) SourceImageReferenceInput() *WindowsVirtualMachineSourceImageReference {
	var returns *WindowsVirtualMachineSourceImageReference
	_jsii_.Get(
		j,
		"sourceImageReferenceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) Tags() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tags",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) TagsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"tagsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) TerminationNotification() WindowsVirtualMachineTerminationNotificationOutputReference {
	var returns WindowsVirtualMachineTerminationNotificationOutputReference
	_jsii_.Get(
		j,
		"terminationNotification",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) TerminationNotificationInput() *WindowsVirtualMachineTerminationNotification {
	var returns *WindowsVirtualMachineTerminationNotification
	_jsii_.Get(
		j,
		"terminationNotificationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) Timeouts() WindowsVirtualMachineTimeoutsOutputReference {
	var returns WindowsVirtualMachineTimeoutsOutputReference
	_jsii_.Get(
		j,
		"timeouts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) TimeoutsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"timeoutsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) Timezone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) TimezoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"timezoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) UserData() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userData",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) UserDataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userDataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) VirtualMachineId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualMachineId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) VirtualMachineScaleSetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualMachineScaleSetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) VirtualMachineScaleSetIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"virtualMachineScaleSetIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) VmAgentPlatformUpdatesEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vmAgentPlatformUpdatesEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) VmAgentPlatformUpdatesEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vmAgentPlatformUpdatesEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) VtpmEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vtpmEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) VtpmEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vtpmEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) WinrmListener() WindowsVirtualMachineWinrmListenerList {
	var returns WindowsVirtualMachineWinrmListenerList
	_jsii_.Get(
		j,
		"winrmListener",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) WinrmListenerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"winrmListenerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_WindowsVirtualMachine) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/windows_virtual_machine azurerm_windows_virtual_machine} Resource.
func NewWindowsVirtualMachine(scope constructs.Construct, id *string, config *WindowsVirtualMachineConfig) WindowsVirtualMachine {
	_init_.Initialize()

	if err := validateNewWindowsVirtualMachineParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_WindowsVirtualMachine{}

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachine.WindowsVirtualMachine",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/azurerm/3.105.0/docs/resources/windows_virtual_machine azurerm_windows_virtual_machine} Resource.
func NewWindowsVirtualMachine_Override(w WindowsVirtualMachine, scope constructs.Construct, id *string, config *WindowsVirtualMachineConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-azurerm.windowsVirtualMachine.WindowsVirtualMachine",
		[]interface{}{scope, id, config},
		w,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetAdminPassword(val *string) {
	if err := j.validateSetAdminPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminPassword",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetAdminUsername(val *string) {
	if err := j.validateSetAdminUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"adminUsername",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetAllowExtensionOperations(val interface{}) {
	if err := j.validateSetAllowExtensionOperationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowExtensionOperations",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetAvailabilitySetId(val *string) {
	if err := j.validateSetAvailabilitySetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"availabilitySetId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetBypassPlatformSafetyChecksOnUserScheduleEnabled(val interface{}) {
	if err := j.validateSetBypassPlatformSafetyChecksOnUserScheduleEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bypassPlatformSafetyChecksOnUserScheduleEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetCapacityReservationGroupId(val *string) {
	if err := j.validateSetCapacityReservationGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"capacityReservationGroupId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetComputerName(val *string) {
	if err := j.validateSetComputerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"computerName",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetCustomData(val *string) {
	if err := j.validateSetCustomDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"customData",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetDedicatedHostGroupId(val *string) {
	if err := j.validateSetDedicatedHostGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicatedHostGroupId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetDedicatedHostId(val *string) {
	if err := j.validateSetDedicatedHostIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dedicatedHostId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetDiskControllerType(val *string) {
	if err := j.validateSetDiskControllerTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"diskControllerType",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetEdgeZone(val *string) {
	if err := j.validateSetEdgeZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"edgeZone",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetEnableAutomaticUpdates(val interface{}) {
	if err := j.validateSetEnableAutomaticUpdatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAutomaticUpdates",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetEncryptionAtHostEnabled(val interface{}) {
	if err := j.validateSetEncryptionAtHostEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"encryptionAtHostEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetEvictionPolicy(val *string) {
	if err := j.validateSetEvictionPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"evictionPolicy",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetExtensionsTimeBudget(val *string) {
	if err := j.validateSetExtensionsTimeBudgetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extensionsTimeBudget",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetHotpatchingEnabled(val interface{}) {
	if err := j.validateSetHotpatchingEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"hotpatchingEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetLicenseType(val *string) {
	if err := j.validateSetLicenseTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"licenseType",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetLocation(val *string) {
	if err := j.validateSetLocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"location",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetMaxBidPrice(val *float64) {
	if err := j.validateSetMaxBidPriceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxBidPrice",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetNetworkInterfaceIds(val *[]*string) {
	if err := j.validateSetNetworkInterfaceIdsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkInterfaceIds",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetPatchAssessmentMode(val *string) {
	if err := j.validateSetPatchAssessmentModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patchAssessmentMode",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetPatchMode(val *string) {
	if err := j.validateSetPatchModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"patchMode",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetPlatformFaultDomain(val *float64) {
	if err := j.validateSetPlatformFaultDomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"platformFaultDomain",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetPriority(val *string) {
	if err := j.validateSetPriorityParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"priority",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetProvisionVmAgent(val interface{}) {
	if err := j.validateSetProvisionVmAgentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisionVmAgent",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetProximityPlacementGroupId(val *string) {
	if err := j.validateSetProximityPlacementGroupIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"proximityPlacementGroupId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetRebootSetting(val *string) {
	if err := j.validateSetRebootSettingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rebootSetting",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetResourceGroupName(val *string) {
	if err := j.validateSetResourceGroupNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"resourceGroupName",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetSecureBootEnabled(val interface{}) {
	if err := j.validateSetSecureBootEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secureBootEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetSize(val *string) {
	if err := j.validateSetSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"size",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetSourceImageId(val *string) {
	if err := j.validateSetSourceImageIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"sourceImageId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetTags(val *map[string]*string) {
	if err := j.validateSetTagsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tags",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetTimezone(val *string) {
	if err := j.validateSetTimezoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"timezone",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetUserData(val *string) {
	if err := j.validateSetUserDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userData",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetVirtualMachineScaleSetId(val *string) {
	if err := j.validateSetVirtualMachineScaleSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"virtualMachineScaleSetId",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetVmAgentPlatformUpdatesEnabled(val interface{}) {
	if err := j.validateSetVmAgentPlatformUpdatesEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vmAgentPlatformUpdatesEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetVtpmEnabled(val interface{}) {
	if err := j.validateSetVtpmEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"vtpmEnabled",
		val,
	)
}

func (j *jsiiProxy_WindowsVirtualMachine)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

// Generates CDKTF code for importing a WindowsVirtualMachine resource upon running "cdktf plan <stack-name>".
func WindowsVirtualMachine_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateWindowsVirtualMachine_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.windowsVirtualMachine.WindowsVirtualMachine",
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
func WindowsVirtualMachine_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWindowsVirtualMachine_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.windowsVirtualMachine.WindowsVirtualMachine",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WindowsVirtualMachine_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWindowsVirtualMachine_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.windowsVirtualMachine.WindowsVirtualMachine",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func WindowsVirtualMachine_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateWindowsVirtualMachine_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-azurerm.windowsVirtualMachine.WindowsVirtualMachine",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func WindowsVirtualMachine_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-azurerm.windowsVirtualMachine.WindowsVirtualMachine",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (w *jsiiProxy_WindowsVirtualMachine) AddMoveTarget(moveTarget *string) {
	if err := w.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) AddOverride(path *string, value interface{}) {
	if err := w.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (w *jsiiProxy_WindowsVirtualMachine) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WindowsVirtualMachine) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (w *jsiiProxy_WindowsVirtualMachine) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (w *jsiiProxy_WindowsVirtualMachine) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (w *jsiiProxy_WindowsVirtualMachine) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (w *jsiiProxy_WindowsVirtualMachine) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (w *jsiiProxy_WindowsVirtualMachine) GetStringAttribute(terraformAttribute *string) *string {
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

func (w *jsiiProxy_WindowsVirtualMachine) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (w *jsiiProxy_WindowsVirtualMachine) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachine) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := w.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (w *jsiiProxy_WindowsVirtualMachine) MoveFromId(id *string) {
	if err := w.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveFromId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) MoveTo(moveTarget *string, index interface{}) {
	if err := w.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) MoveToId(id *string) {
	if err := w.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"moveToId",
		[]interface{}{id},
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) OverrideLogicalId(newLogicalId *string) {
	if err := w.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) PutAdditionalCapabilities(value *WindowsVirtualMachineAdditionalCapabilities) {
	if err := w.validatePutAdditionalCapabilitiesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAdditionalCapabilities",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) PutAdditionalUnattendContent(value interface{}) {
	if err := w.validatePutAdditionalUnattendContentParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putAdditionalUnattendContent",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) PutBootDiagnostics(value *WindowsVirtualMachineBootDiagnostics) {
	if err := w.validatePutBootDiagnosticsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putBootDiagnostics",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) PutGalleryApplication(value interface{}) {
	if err := w.validatePutGalleryApplicationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putGalleryApplication",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) PutIdentity(value *WindowsVirtualMachineIdentity) {
	if err := w.validatePutIdentityParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putIdentity",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) PutOsDisk(value *WindowsVirtualMachineOsDisk) {
	if err := w.validatePutOsDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putOsDisk",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) PutOsImageNotification(value *WindowsVirtualMachineOsImageNotification) {
	if err := w.validatePutOsImageNotificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putOsImageNotification",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) PutPlan(value *WindowsVirtualMachinePlan) {
	if err := w.validatePutPlanParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putPlan",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) PutSecret(value interface{}) {
	if err := w.validatePutSecretParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSecret",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) PutSourceImageReference(value *WindowsVirtualMachineSourceImageReference) {
	if err := w.validatePutSourceImageReferenceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putSourceImageReference",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) PutTerminationNotification(value *WindowsVirtualMachineTerminationNotification) {
	if err := w.validatePutTerminationNotificationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTerminationNotification",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) PutTimeouts(value *WindowsVirtualMachineTimeouts) {
	if err := w.validatePutTimeoutsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putTimeouts",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) PutWinrmListener(value interface{}) {
	if err := w.validatePutWinrmListenerParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		w,
		"putWinrmListener",
		[]interface{}{value},
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetAdditionalCapabilities() {
	_jsii_.InvokeVoid(
		w,
		"resetAdditionalCapabilities",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetAdditionalUnattendContent() {
	_jsii_.InvokeVoid(
		w,
		"resetAdditionalUnattendContent",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetAllowExtensionOperations() {
	_jsii_.InvokeVoid(
		w,
		"resetAllowExtensionOperations",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetAvailabilitySetId() {
	_jsii_.InvokeVoid(
		w,
		"resetAvailabilitySetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetBootDiagnostics() {
	_jsii_.InvokeVoid(
		w,
		"resetBootDiagnostics",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetBypassPlatformSafetyChecksOnUserScheduleEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetBypassPlatformSafetyChecksOnUserScheduleEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetCapacityReservationGroupId() {
	_jsii_.InvokeVoid(
		w,
		"resetCapacityReservationGroupId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetComputerName() {
	_jsii_.InvokeVoid(
		w,
		"resetComputerName",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetCustomData() {
	_jsii_.InvokeVoid(
		w,
		"resetCustomData",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetDedicatedHostGroupId() {
	_jsii_.InvokeVoid(
		w,
		"resetDedicatedHostGroupId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetDedicatedHostId() {
	_jsii_.InvokeVoid(
		w,
		"resetDedicatedHostId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetDiskControllerType() {
	_jsii_.InvokeVoid(
		w,
		"resetDiskControllerType",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetEdgeZone() {
	_jsii_.InvokeVoid(
		w,
		"resetEdgeZone",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetEnableAutomaticUpdates() {
	_jsii_.InvokeVoid(
		w,
		"resetEnableAutomaticUpdates",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetEncryptionAtHostEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetEncryptionAtHostEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetEvictionPolicy() {
	_jsii_.InvokeVoid(
		w,
		"resetEvictionPolicy",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetExtensionsTimeBudget() {
	_jsii_.InvokeVoid(
		w,
		"resetExtensionsTimeBudget",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetGalleryApplication() {
	_jsii_.InvokeVoid(
		w,
		"resetGalleryApplication",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetHotpatchingEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetHotpatchingEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetId() {
	_jsii_.InvokeVoid(
		w,
		"resetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetIdentity() {
	_jsii_.InvokeVoid(
		w,
		"resetIdentity",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetLicenseType() {
	_jsii_.InvokeVoid(
		w,
		"resetLicenseType",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetMaxBidPrice() {
	_jsii_.InvokeVoid(
		w,
		"resetMaxBidPrice",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetOsImageNotification() {
	_jsii_.InvokeVoid(
		w,
		"resetOsImageNotification",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		w,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetPatchAssessmentMode() {
	_jsii_.InvokeVoid(
		w,
		"resetPatchAssessmentMode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetPatchMode() {
	_jsii_.InvokeVoid(
		w,
		"resetPatchMode",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetPlan() {
	_jsii_.InvokeVoid(
		w,
		"resetPlan",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetPlatformFaultDomain() {
	_jsii_.InvokeVoid(
		w,
		"resetPlatformFaultDomain",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetPriority() {
	_jsii_.InvokeVoid(
		w,
		"resetPriority",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetProvisionVmAgent() {
	_jsii_.InvokeVoid(
		w,
		"resetProvisionVmAgent",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetProximityPlacementGroupId() {
	_jsii_.InvokeVoid(
		w,
		"resetProximityPlacementGroupId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetRebootSetting() {
	_jsii_.InvokeVoid(
		w,
		"resetRebootSetting",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetSecret() {
	_jsii_.InvokeVoid(
		w,
		"resetSecret",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetSecureBootEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetSecureBootEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetSourceImageId() {
	_jsii_.InvokeVoid(
		w,
		"resetSourceImageId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetSourceImageReference() {
	_jsii_.InvokeVoid(
		w,
		"resetSourceImageReference",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetTags() {
	_jsii_.InvokeVoid(
		w,
		"resetTags",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetTerminationNotification() {
	_jsii_.InvokeVoid(
		w,
		"resetTerminationNotification",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetTimeouts() {
	_jsii_.InvokeVoid(
		w,
		"resetTimeouts",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetTimezone() {
	_jsii_.InvokeVoid(
		w,
		"resetTimezone",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetUserData() {
	_jsii_.InvokeVoid(
		w,
		"resetUserData",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetVirtualMachineScaleSetId() {
	_jsii_.InvokeVoid(
		w,
		"resetVirtualMachineScaleSetId",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetVmAgentPlatformUpdatesEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetVmAgentPlatformUpdatesEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetVtpmEnabled() {
	_jsii_.InvokeVoid(
		w,
		"resetVtpmEnabled",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetWinrmListener() {
	_jsii_.InvokeVoid(
		w,
		"resetWinrmListener",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) ResetZone() {
	_jsii_.InvokeVoid(
		w,
		"resetZone",
		nil, // no parameters
	)
}

func (w *jsiiProxy_WindowsVirtualMachine) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachine) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		w,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachine) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachine) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachine) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		w,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (w *jsiiProxy_WindowsVirtualMachine) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		w,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}


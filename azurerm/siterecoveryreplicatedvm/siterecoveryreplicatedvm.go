package siterecoveryreplicatedvm

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVm",
		reflect.TypeOf((*SiteRecoveryReplicatedVm)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "managedDisk", GoGetter: "ManagedDisk"},
			_jsii_.MemberProperty{JsiiProperty: "managedDiskInput", GoGetter: "ManagedDiskInput"},
			_jsii_.MemberProperty{JsiiProperty: "multiVmGroupName", GoGetter: "MultiVmGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "multiVmGroupNameInput", GoGetter: "MultiVmGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterface", GoGetter: "NetworkInterface"},
			_jsii_.MemberProperty{JsiiProperty: "networkInterfaceInput", GoGetter: "NetworkInterfaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putManagedDisk", GoMethod: "PutManagedDisk"},
			_jsii_.MemberMethod{JsiiMethod: "putNetworkInterface", GoMethod: "PutNetworkInterface"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "putUnmanagedDisk", GoMethod: "PutUnmanagedDisk"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "recoveryReplicationPolicyId", GoGetter: "RecoveryReplicationPolicyId"},
			_jsii_.MemberProperty{JsiiProperty: "recoveryReplicationPolicyIdInput", GoGetter: "RecoveryReplicationPolicyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "recoveryVaultName", GoGetter: "RecoveryVaultName"},
			_jsii_.MemberProperty{JsiiProperty: "recoveryVaultNameInput", GoGetter: "RecoveryVaultNameInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetManagedDisk", GoMethod: "ResetManagedDisk"},
			_jsii_.MemberMethod{JsiiMethod: "resetMultiVmGroupName", GoMethod: "ResetMultiVmGroupName"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetworkInterface", GoMethod: "ResetNetworkInterface"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetAvailabilitySetId", GoMethod: "ResetTargetAvailabilitySetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetBootDiagnosticStorageAccountId", GoMethod: "ResetTargetBootDiagnosticStorageAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetCapacityReservationGroupId", GoMethod: "ResetTargetCapacityReservationGroupId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetEdgeZone", GoMethod: "ResetTargetEdgeZone"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetNetworkId", GoMethod: "ResetTargetNetworkId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetProximityPlacementGroupId", GoMethod: "ResetTargetProximityPlacementGroupId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetVirtualMachineScaleSetId", GoMethod: "ResetTargetVirtualMachineScaleSetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetZone", GoMethod: "ResetTargetZone"},
			_jsii_.MemberMethod{JsiiMethod: "resetTestNetworkId", GoMethod: "ResetTestNetworkId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "resetUnmanagedDisk", GoMethod: "ResetUnmanagedDisk"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupName", GoGetter: "ResourceGroupName"},
			_jsii_.MemberProperty{JsiiProperty: "resourceGroupNameInput", GoGetter: "ResourceGroupNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceRecoveryFabricName", GoGetter: "SourceRecoveryFabricName"},
			_jsii_.MemberProperty{JsiiProperty: "sourceRecoveryFabricNameInput", GoGetter: "SourceRecoveryFabricNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceRecoveryProtectionContainerName", GoGetter: "SourceRecoveryProtectionContainerName"},
			_jsii_.MemberProperty{JsiiProperty: "sourceRecoveryProtectionContainerNameInput", GoGetter: "SourceRecoveryProtectionContainerNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceVmId", GoGetter: "SourceVmId"},
			_jsii_.MemberProperty{JsiiProperty: "sourceVmIdInput", GoGetter: "SourceVmIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "targetAvailabilitySetId", GoGetter: "TargetAvailabilitySetId"},
			_jsii_.MemberProperty{JsiiProperty: "targetAvailabilitySetIdInput", GoGetter: "TargetAvailabilitySetIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetBootDiagnosticStorageAccountId", GoGetter: "TargetBootDiagnosticStorageAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "targetBootDiagnosticStorageAccountIdInput", GoGetter: "TargetBootDiagnosticStorageAccountIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetCapacityReservationGroupId", GoGetter: "TargetCapacityReservationGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "targetCapacityReservationGroupIdInput", GoGetter: "TargetCapacityReservationGroupIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetEdgeZone", GoGetter: "TargetEdgeZone"},
			_jsii_.MemberProperty{JsiiProperty: "targetEdgeZoneInput", GoGetter: "TargetEdgeZoneInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetNetworkId", GoGetter: "TargetNetworkId"},
			_jsii_.MemberProperty{JsiiProperty: "targetNetworkIdInput", GoGetter: "TargetNetworkIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetProximityPlacementGroupId", GoGetter: "TargetProximityPlacementGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "targetProximityPlacementGroupIdInput", GoGetter: "TargetProximityPlacementGroupIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetRecoveryFabricId", GoGetter: "TargetRecoveryFabricId"},
			_jsii_.MemberProperty{JsiiProperty: "targetRecoveryFabricIdInput", GoGetter: "TargetRecoveryFabricIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetRecoveryProtectionContainerId", GoGetter: "TargetRecoveryProtectionContainerId"},
			_jsii_.MemberProperty{JsiiProperty: "targetRecoveryProtectionContainerIdInput", GoGetter: "TargetRecoveryProtectionContainerIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetResourceGroupId", GoGetter: "TargetResourceGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "targetResourceGroupIdInput", GoGetter: "TargetResourceGroupIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetVirtualMachineScaleSetId", GoGetter: "TargetVirtualMachineScaleSetId"},
			_jsii_.MemberProperty{JsiiProperty: "targetVirtualMachineScaleSetIdInput", GoGetter: "TargetVirtualMachineScaleSetIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetZone", GoGetter: "TargetZone"},
			_jsii_.MemberProperty{JsiiProperty: "targetZoneInput", GoGetter: "TargetZoneInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "testNetworkId", GoGetter: "TestNetworkId"},
			_jsii_.MemberProperty{JsiiProperty: "testNetworkIdInput", GoGetter: "TestNetworkIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "unmanagedDisk", GoGetter: "UnmanagedDisk"},
			_jsii_.MemberProperty{JsiiProperty: "unmanagedDiskInput", GoGetter: "UnmanagedDiskInput"},
		},
		func() interface{} {
			j := jsiiProxy_SiteRecoveryReplicatedVm{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmConfig",
		reflect.TypeOf((*SiteRecoveryReplicatedVmConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmManagedDisk",
		reflect.TypeOf((*SiteRecoveryReplicatedVmManagedDisk)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmManagedDiskList",
		reflect.TypeOf((*SiteRecoveryReplicatedVmManagedDiskList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_SiteRecoveryReplicatedVmManagedDiskList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmManagedDiskOutputReference",
		reflect.TypeOf((*SiteRecoveryReplicatedVmManagedDiskOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "diskId", GoGetter: "DiskId"},
			_jsii_.MemberProperty{JsiiProperty: "diskIdInput", GoGetter: "DiskIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putTargetDiskEncryption", GoMethod: "PutTargetDiskEncryption"},
			_jsii_.MemberMethod{JsiiMethod: "resetDiskId", GoMethod: "ResetDiskId"},
			_jsii_.MemberMethod{JsiiMethod: "resetStagingStorageAccountId", GoMethod: "ResetStagingStorageAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetDiskEncryption", GoMethod: "ResetTargetDiskEncryption"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetDiskEncryptionSetId", GoMethod: "ResetTargetDiskEncryptionSetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetDiskType", GoMethod: "ResetTargetDiskType"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetReplicaDiskType", GoMethod: "ResetTargetReplicaDiskType"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetResourceGroupId", GoMethod: "ResetTargetResourceGroupId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "stagingStorageAccountId", GoGetter: "StagingStorageAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "stagingStorageAccountIdInput", GoGetter: "StagingStorageAccountIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetDiskEncryption", GoGetter: "TargetDiskEncryption"},
			_jsii_.MemberProperty{JsiiProperty: "targetDiskEncryptionInput", GoGetter: "TargetDiskEncryptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetDiskEncryptionSetId", GoGetter: "TargetDiskEncryptionSetId"},
			_jsii_.MemberProperty{JsiiProperty: "targetDiskEncryptionSetIdInput", GoGetter: "TargetDiskEncryptionSetIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetDiskType", GoGetter: "TargetDiskType"},
			_jsii_.MemberProperty{JsiiProperty: "targetDiskTypeInput", GoGetter: "TargetDiskTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetReplicaDiskType", GoGetter: "TargetReplicaDiskType"},
			_jsii_.MemberProperty{JsiiProperty: "targetReplicaDiskTypeInput", GoGetter: "TargetReplicaDiskTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetResourceGroupId", GoGetter: "TargetResourceGroupId"},
			_jsii_.MemberProperty{JsiiProperty: "targetResourceGroupIdInput", GoGetter: "TargetResourceGroupIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SiteRecoveryReplicatedVmManagedDiskOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryption",
		reflect.TypeOf((*SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryption)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionDiskEncryptionKey",
		reflect.TypeOf((*SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionDiskEncryptionKey)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionDiskEncryptionKeyList",
		reflect.TypeOf((*SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionDiskEncryptionKeyList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionDiskEncryptionKeyList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionDiskEncryptionKeyOutputReference",
		reflect.TypeOf((*SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionDiskEncryptionKeyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecretUrl", GoMethod: "ResetSecretUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetVaultId", GoMethod: "ResetVaultId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "secretUrl", GoGetter: "SecretUrl"},
			_jsii_.MemberProperty{JsiiProperty: "secretUrlInput", GoGetter: "SecretUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vaultId", GoGetter: "VaultId"},
			_jsii_.MemberProperty{JsiiProperty: "vaultIdInput", GoGetter: "VaultIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionDiskEncryptionKeyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionKeyEncryptionKey",
		reflect.TypeOf((*SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionKeyEncryptionKey)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionKeyEncryptionKeyList",
		reflect.TypeOf((*SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionKeyEncryptionKeyList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionKeyEncryptionKeyList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionKeyEncryptionKeyOutputReference",
		reflect.TypeOf((*SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionKeyEncryptionKeyOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keyUrl", GoGetter: "KeyUrl"},
			_jsii_.MemberProperty{JsiiProperty: "keyUrlInput", GoGetter: "KeyUrlInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyUrl", GoMethod: "ResetKeyUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetVaultId", GoMethod: "ResetVaultId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "vaultId", GoGetter: "VaultId"},
			_jsii_.MemberProperty{JsiiProperty: "vaultIdInput", GoGetter: "VaultIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionKeyEncryptionKeyOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionList",
		reflect.TypeOf((*SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionOutputReference",
		reflect.TypeOf((*SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "diskEncryptionKey", GoGetter: "DiskEncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "diskEncryptionKeyInput", GoGetter: "DiskEncryptionKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keyEncryptionKey", GoGetter: "KeyEncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "keyEncryptionKeyInput", GoGetter: "KeyEncryptionKeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "putDiskEncryptionKey", GoMethod: "PutDiskEncryptionKey"},
			_jsii_.MemberMethod{JsiiMethod: "putKeyEncryptionKey", GoMethod: "PutKeyEncryptionKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetDiskEncryptionKey", GoMethod: "ResetDiskEncryptionKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyEncryptionKey", GoMethod: "ResetKeyEncryptionKey"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SiteRecoveryReplicatedVmManagedDiskTargetDiskEncryptionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmNetworkInterface",
		reflect.TypeOf((*SiteRecoveryReplicatedVmNetworkInterface)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmNetworkInterfaceList",
		reflect.TypeOf((*SiteRecoveryReplicatedVmNetworkInterfaceList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmNetworkInterfaceOutputReference",
		reflect.TypeOf((*SiteRecoveryReplicatedVmNetworkInterfaceOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "failoverTestPublicIpAddressId", GoGetter: "FailoverTestPublicIpAddressId"},
			_jsii_.MemberProperty{JsiiProperty: "failoverTestPublicIpAddressIdInput", GoGetter: "FailoverTestPublicIpAddressIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "failoverTestStaticIp", GoGetter: "FailoverTestStaticIp"},
			_jsii_.MemberProperty{JsiiProperty: "failoverTestStaticIpInput", GoGetter: "FailoverTestStaticIpInput"},
			_jsii_.MemberProperty{JsiiProperty: "failoverTestSubnetName", GoGetter: "FailoverTestSubnetName"},
			_jsii_.MemberProperty{JsiiProperty: "failoverTestSubnetNameInput", GoGetter: "FailoverTestSubnetNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isPrimary", GoGetter: "IsPrimary"},
			_jsii_.MemberProperty{JsiiProperty: "isPrimaryInput", GoGetter: "IsPrimaryInput"},
			_jsii_.MemberProperty{JsiiProperty: "recoveryPublicIpAddressId", GoGetter: "RecoveryPublicIpAddressId"},
			_jsii_.MemberProperty{JsiiProperty: "recoveryPublicIpAddressIdInput", GoGetter: "RecoveryPublicIpAddressIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetFailoverTestPublicIpAddressId", GoMethod: "ResetFailoverTestPublicIpAddressId"},
			_jsii_.MemberMethod{JsiiMethod: "resetFailoverTestStaticIp", GoMethod: "ResetFailoverTestStaticIp"},
			_jsii_.MemberMethod{JsiiMethod: "resetFailoverTestSubnetName", GoMethod: "ResetFailoverTestSubnetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsPrimary", GoMethod: "ResetIsPrimary"},
			_jsii_.MemberMethod{JsiiMethod: "resetRecoveryPublicIpAddressId", GoMethod: "ResetRecoveryPublicIpAddressId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceNetworkInterfaceId", GoMethod: "ResetSourceNetworkInterfaceId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetStaticIp", GoMethod: "ResetTargetStaticIp"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetSubnetName", GoMethod: "ResetTargetSubnetName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sourceNetworkInterfaceId", GoGetter: "SourceNetworkInterfaceId"},
			_jsii_.MemberProperty{JsiiProperty: "sourceNetworkInterfaceIdInput", GoGetter: "SourceNetworkInterfaceIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetStaticIp", GoGetter: "TargetStaticIp"},
			_jsii_.MemberProperty{JsiiProperty: "targetStaticIpInput", GoGetter: "TargetStaticIpInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetSubnetName", GoGetter: "TargetSubnetName"},
			_jsii_.MemberProperty{JsiiProperty: "targetSubnetNameInput", GoGetter: "TargetSubnetNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SiteRecoveryReplicatedVmNetworkInterfaceOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmTimeouts",
		reflect.TypeOf((*SiteRecoveryReplicatedVmTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmTimeoutsOutputReference",
		reflect.TypeOf((*SiteRecoveryReplicatedVmTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "read", GoGetter: "Read"},
			_jsii_.MemberProperty{JsiiProperty: "readInput", GoGetter: "ReadInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetRead", GoMethod: "ResetRead"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_SiteRecoveryReplicatedVmTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmUnmanagedDisk",
		reflect.TypeOf((*SiteRecoveryReplicatedVmUnmanagedDisk)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmUnmanagedDiskList",
		reflect.TypeOf((*SiteRecoveryReplicatedVmUnmanagedDiskList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_SiteRecoveryReplicatedVmUnmanagedDiskList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-azurerm.siteRecoveryReplicatedVm.SiteRecoveryReplicatedVmUnmanagedDiskOutputReference",
		reflect.TypeOf((*SiteRecoveryReplicatedVmUnmanagedDiskOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "diskUri", GoGetter: "DiskUri"},
			_jsii_.MemberProperty{JsiiProperty: "diskUriInput", GoGetter: "DiskUriInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetDiskUri", GoMethod: "ResetDiskUri"},
			_jsii_.MemberMethod{JsiiMethod: "resetStagingStorageAccountId", GoMethod: "ResetStagingStorageAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetStorageAccountId", GoMethod: "ResetTargetStorageAccountId"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "stagingStorageAccountId", GoGetter: "StagingStorageAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "stagingStorageAccountIdInput", GoGetter: "StagingStorageAccountIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetStorageAccountId", GoGetter: "TargetStorageAccountId"},
			_jsii_.MemberProperty{JsiiProperty: "targetStorageAccountIdInput", GoGetter: "TargetStorageAccountIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SiteRecoveryReplicatedVmUnmanagedDiskOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}

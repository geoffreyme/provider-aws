/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

// Code generated by ack-generate. DO NOT EDIT.

package v1alpha1

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Hack to avoid import errors during build...
var (
	_ = &metav1.Time{}
)

// +kubebuilder:skipversion
type AccessPointDescription struct {
	AccessPointARN *string `json:"accessPointARN,omitempty"`

	AccessPointID *string `json:"accessPointID,omitempty"`

	ClientToken *string `json:"clientToken,omitempty"`

	FileSystemID *string `json:"fileSystemID,omitempty"`

	LifeCycleState *string `json:"lifeCycleState,omitempty"`

	Name *string `json:"name,omitempty"`

	OwnerID *string `json:"ownerID,omitempty"`
	// The full POSIX identity, including the user ID, group ID, and any secondary
	// group IDs, on the access point that is used for all file system operations
	// performed by NFS clients using the access point.
	PosixUser *PosixUser `json:"posixUser,omitempty"`
	// Specifies the directory on the Amazon EFS file system that the access point
	// provides access to. The access point exposes the specified file system path
	// as the root directory of your file system to applications using the access
	// point. NFS clients using the access point can only access data in the access
	// point's RootDirectory and it's subdirectories.
	RootDirectory *RootDirectory `json:"rootDirectory,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`
}

// +kubebuilder:skipversion
type CreationInfo struct {
	OwnerGid *int64 `json:"ownerGid,omitempty"`

	OwnerUid *int64 `json:"ownerUid,omitempty"`

	Permissions *string `json:"permissions,omitempty"`
}

// +kubebuilder:skipversion
type FileSystemDescription struct {
	AvailabilityZoneID *string `json:"availabilityZoneID,omitempty"`

	AvailabilityZoneName *string `json:"availabilityZoneName,omitempty"`

	CreationTime *metav1.Time `json:"creationTime,omitempty"`

	CreationToken *string `json:"creationToken,omitempty"`

	Encrypted *bool `json:"encrypted,omitempty"`

	FileSystemARN *string `json:"fileSystemARN,omitempty"`

	FileSystemID *string `json:"fileSystemID,omitempty"`

	KMSKeyID *string `json:"kmsKeyID,omitempty"`

	LifeCycleState *string `json:"lifeCycleState,omitempty"`

	Name *string `json:"name,omitempty"`

	NumberOfMountTargets *int64 `json:"numberOfMountTargets,omitempty"`

	OwnerID *string `json:"ownerID,omitempty"`

	PerformanceMode *string `json:"performanceMode,omitempty"`
	// The latest known metered size (in bytes) of data stored in the file system,
	// in its Value field, and the time at which that size was determined in its
	// Timestamp field. The value doesn't represent the size of a consistent snapshot
	// of the file system, but it is eventually consistent when there are no writes
	// to the file system. That is, the value represents the actual size only if
	// the file system is not modified for a period longer than a couple of hours.
	// Otherwise, the value is not necessarily the exact size the file system was
	// at any instant in time.
	SizeInBytes *FileSystemSize `json:"sizeInBytes,omitempty"`

	Tags []*Tag `json:"tags,omitempty"`

	ThroughputMode *string `json:"throughputMode,omitempty"`
}

// +kubebuilder:skipversion
type FileSystemSize struct {
	Timestamp *metav1.Time `json:"timestamp,omitempty"`

	Value *int64 `json:"value,omitempty"`

	ValueInIA *int64 `json:"valueInIA,omitempty"`

	ValueInStandard *int64 `json:"valueInStandard,omitempty"`
}

// +kubebuilder:skipversion
type MountTargetDescription struct {
	AvailabilityZoneID *string `json:"availabilityZoneID,omitempty"`

	AvailabilityZoneName *string `json:"availabilityZoneName,omitempty"`

	FileSystemID *string `json:"fileSystemID,omitempty"`

	IPAddress *string `json:"ipAddress,omitempty"`

	LifeCycleState *string `json:"lifeCycleState,omitempty"`

	MountTargetID *string `json:"mountTargetID,omitempty"`

	NetworkInterfaceID *string `json:"networkInterfaceID,omitempty"`

	OwnerID *string `json:"ownerID,omitempty"`

	SubnetID *string `json:"subnetID,omitempty"`

	VPCID *string `json:"vpcID,omitempty"`
}

// +kubebuilder:skipversion
type PosixUser struct {
	Gid *int64 `json:"gid,omitempty"`

	SecondaryGids []*int64 `json:"secondaryGids,omitempty"`

	Uid *int64 `json:"uid,omitempty"`
}

// +kubebuilder:skipversion
type RootDirectory struct {
	// Required if the RootDirectory > Path specified does not exist. Specifies
	// the POSIX IDs and permissions to apply to the access point's RootDirectory
	// > Path. If the access point root directory does not exist, EFS creates it
	// with these settings when a client connects to the access point. When specifying
	// CreationInfo, you must include values for all properties.
	//
	// Amazon EFS creates a root directory only if you have provided the CreationInfo:
	// OwnUid, OwnGID, and permissions for the directory. If you do not provide
	// this information, Amazon EFS does not create the root directory. If the root
	// directory does not exist, attempts to mount using the access point will fail.
	//
	// If you do not provide CreationInfo and the specified RootDirectory does not
	// exist, attempts to mount the file system using the access point will fail.
	CreationInfo *CreationInfo `json:"creationInfo,omitempty"`

	Path *string `json:"path,omitempty"`
}

// +kubebuilder:skipversion
type Tag struct {
	Key *string `json:"key,omitempty"`

	Value *string `json:"value,omitempty"`
}

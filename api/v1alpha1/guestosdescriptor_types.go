// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

//go:generate go run ./internal/generate_guestosidentifiers.go

package v1alpha1

// VirtualMachineGuestOSDescriptor is used to specify information about the
//
type VirtualMachineGuestOSDescriptor struct {
	// VirtualMachineGuestOSIdentifier is the guest OS identifier.
	//
	// If ommited the Family field is used to determine a default
	// VirtualMachineGuestOSIdentifier value.
	//
	// +optional
	ID VirtualMachineGuestOSIdentifier `json:"id,omitempty"`

	// VirtualMachineGuestOSFamily specifies the family of the guest OS.
	//
	// If omitted the VirtualMachineImage is used to attempt to determine
	// the guest's family and/or ID.
	//
	// +optional
	Family VirtualMachineGuestOSFamily `json:"family,omitempty"`
}

// VirtualMachineGuestOSFamily specifies the family of the guest OS.
type VirtualMachineGuestOSFamily string

const (
	VirtualMachineGuestOSFamilyWindows       VirtualMachineGuestOSFamily = "windowsGuest"
	VirtualMachineGuestOSFamilyLinux         VirtualMachineGuestOSFamily = "linuxGuest"
	VirtualMachineGuestOSFamilyNovellNetware VirtualMachineGuestOSFamily = "netwareGuest"
	VirtualMachineGuestOSFamilySolaris       VirtualMachineGuestOSFamily = "solarisGuest"
	VirtualMachineGuestOSFamilyDarwin        VirtualMachineGuestOSFamily = "darwinGuestFamily"
	VirtualMachineGuestOSFamilyOther         VirtualMachineGuestOSFamily = "otherGuestFamily"
)

// VirtualMachineGuestOSIdentifier is the guest operating system identifier.
type VirtualMachineGuestOSIdentifier string

// validVirtualMachineGuestOSIdentifiers is a map used to determine if a
// VirtualMachineGuestOSIdentifier value is valid.
var validVirtualMachineGuestOSIdentifiers map[VirtualMachineGuestOSIdentifier]struct{}

// IsValidVirtualMachineGuestOSIdentifier returns true if the provided value is
// valid.
func IsValidVirtualMachineGuestOSIdentifier(s string) bool {
	return VirtualMachineGuestOSIdentifier(s).IsValid()
}

// IsValidVirtualMachineGuestOSIdentifier returns true if the identifier is
// valid.
func (e VirtualMachineGuestOSIdentifier) IsValid() bool {
	_, ok := validVirtualMachineGuestOSIdentifiers[e]
	return ok
}

// Copyright (c) 2020 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

package v1alpha1_test

import (
	"testing"

	"github.com/vmware-tanzu/vm-operator-api/api/v1alpha1"
)

func TestIsValidVirtualMachineGuestOSIdentifier(t *testing.T) {

	t.Run("valid", func(t *testing.T) {
		testCases := []struct {
			name string
			id   string
		}{
			{
				name: "otherLinuxGuest",
				id:   "otherLinuxGuest",
			},
			{
				name: "genericLinuxGuest",
				id:   "genericLinuxGuest",
			},
		}
		for i := range testCases {
			tc := testCases[i]
			t.Run(tc.name, func(t *testing.T) {
				if !v1alpha1.IsValidVirtualMachineGuestOSIdentifier(tc.id) {
					t.Errorf("%s is invalid & should be valid", tc.id)
				}

			})
		}
	})

	t.Run("invalid", func(t *testing.T) {
		testCases := []struct {
			name string
			id   string
		}{
			{
				name: "windows",
				id:   "windows",
			},
			{
				name: "windowsGuest",
				id:   "windowsGuest",
			},
		}
		for i := range testCases {
			tc := testCases[i]
			t.Run(tc.name, func(t *testing.T) {
				if v1alpha1.IsValidVirtualMachineGuestOSIdentifier(tc.id) {
					t.Errorf("%s is valid & should be invalid", tc.id)
				}

			})
		}
	})

}

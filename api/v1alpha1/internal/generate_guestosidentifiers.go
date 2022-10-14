// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
// SPDX-License-Identifier: Apache-2.0

//go:build guestosidentifier_types
// +build guestosidentifier_types

package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"os/exec"
)

func main() {
	generate()
	format()
}

func generate() {
	out, err := os.Create("guestosidentifier_types.go")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
	defer out.Close()

	out.WriteString(`// Copyright (c) 2022 VMware, Inc. All Rights Reserved.
	// SPDX-License-Identifier: Apache-2.0

	//
	// !! GENERATED CONTENT -- DO NOT MODIFY DIRECTLY !!
	//
	
	package v1alpha1
	
	const (`)

	in, err := os.Open("./internal/guestosidentifiers.csv")
	if err != nil {
		os.Stderr.WriteString(err.Error())
		os.Exit(1)
	}
	defer in.Close()

	var guestOSIDConstantNames []string
	r := csv.NewReader(in)
	r.Comment = '`'

	for {
		data, err := r.Read()
		if err != nil {
			if err == io.EOF {
				break
			}
			os.Stderr.WriteString(err.Error())
			os.Exit(1)
		}

		guestOSIDConstName := "VirtualMachineGuestOSIdentifier" + data[1]
		guestOSIDConstantNames = append(guestOSIDConstantNames, guestOSIDConstName)
		fmt.Fprintf(out, "\t%s VirtualMachineGuestOSIdentifier = \"%s\"\n", guestOSIDConstName, data[0])
	}

	fmt.Fprintf(out, ")\n")
	out.WriteString(`
	func init() {
		validVirtualMachineGuestOSIdentifiers = map[VirtualMachineGuestOSIdentifier]struct{}{
	`)
	for _, name := range guestOSIDConstantNames {
		fmt.Fprintf(out, "%s: struct{}{},\n", name)
	}
	out.WriteString(`}}`)
}

func format() {
	cmd := exec.Command("go", "fmt", "guestosidentifier_types.go")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		os.Exit(1)
	}
}

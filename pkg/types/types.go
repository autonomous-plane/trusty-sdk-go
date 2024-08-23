//
// Copyright 2024 Stacklok, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package types is the collection of main data types used by the Trusty libraries
package types

// Ecosystem is an identifier of a packaging system supported by Trusty
type Ecosystem int32

// Dependency represents a generic dependency structure
type Dependency struct {
	Name      string
	Version   string
	Ecosystem Ecosystem
}

const (
	// EcosystemNpm identifies the NPM ecosystem
	EcosystemNpm Ecosystem = 1

	// EcosystemGo identifies the Go language
	EcosystemGo Ecosystem = 2

	// EcosystemPypi identifies the Python Package Index
	EcosystemPypi Ecosystem = 3

	// EcosystemMaven identifies the Python Package Index
	EcosystemMaven Ecosystem = 4

	// EcosystemCrates identifies the Python Package Index
	EcosystemCrates Ecosystem = 5
)

// Ecosystems enumerates the supported ecosystems
var Ecosystems = map[string]Ecosystem{
	"EcosystemNpm":  EcosystemNpm,
	"EcosystemGo":   EcosystemGo,
	"EcosystemPypi": EcosystemPypi,
	"EcosystemMaven": EcosystemMaven,
	"EcosystemCrates": EcosystemCrates,
}

// AsString returns the string representation of the DepEcosystem
func (ecosystem Ecosystem) AsString() string {
	switch ecosystem {
	case EcosystemNpm:
		return "npm"
	case EcosystemGo:
		return "Go"
	case EcosystemPypi:
		return "PyPI"
	case EcosystemMaven:
		return "Maven"
	case EcosystemCrates:
		return "crates"
	default:
		return ""
	}
}

// ConvertDepsToMap converts a slice of Dependency structs to a map for easier comparison
func ConvertDepsToMap(deps []Dependency) map[string]string {
	depMap := make(map[string]string)
	for _, dep := range deps {
		depMap[dep.Name] = dep.Version
	}
	return depMap
}

// DiffDependencies compares two sets of dependencies (represented as maps) and finds what's added in newDeps.
func DiffDependencies(oldDeps, newDeps map[string]string) map[string]string {
	addedDeps := make(map[string]string)
	for dep, version := range newDeps {
		if _, exists := oldDeps[dep]; !exists {
			addedDeps[dep] = version
		}
	}
	return addedDeps
}

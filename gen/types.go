// Copyright © 2019 The Samply Development Community
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

package gen

import "fmt"

// Object is like a JSON object
type Object map[string]interface{}

// Array is like a JSON array
type Array []interface{}

var (
	vitalSigns = codeableConcept(coding("http://terminology.hl7.org/CodeSystem/observation-category", "vital-signs"))
)

func coding(system string, code string) Object {
	return Object{"system": system, "code": code}
}

func codingWithVersion(system string, version string, code string) Object {
	return Object{"system": system, "version": version, "code": code}
}

func codeableConcept(coding Object) Object {
	return Object{"coding": Array{coding}}
}

func quantity(value float64, unit string) Object {
	return Object{"value": value, "unit": unit, "system": "http://unitsofmeasure.org", "code": unit}
}

func reference(resourceType string, id int) Object {
	return Object{"reference": fmt.Sprintf("%s/%d", resourceType, id)}
}

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

import (
	"fmt"
	"math/rand"
)

func CauseOfDeath(r *rand.Rand, patientIdx int) Object {
	return Object{
		"resourceType":         "Observation",
		"id":                   fmt.Sprintf("bbmri-%d-cause-of-death", patientIdx),
		"meta":                 meta("https://fhir.bbmri.de/StructureDefinition/CauseOfDeath"),
		"status":               "final",
		"code":                 codeableConcept(coding("http://loinc.org", "68343-3")),
		"subject":              patientReference(patientIdx),
		"valueCodeableConcept": codeableConcept(coding("http://hl7.org/fhir/sid/icd-10", randIcd10Code(r))),
	}
}

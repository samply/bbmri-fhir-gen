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
	"time"
)

func BodyWeight(patientIdx int, date time.Time, value float64) Object {
	return Object{
		"resourceType":      "Observation",
		"id":                fmt.Sprintf("%d-body-weight", patientIdx),
		"meta":              meta("https://fhir.bbmri.de/StructureDefinition/BodyWeight"),
		"status":            "final",
		"category":          Array{vitalSigns},
		"subject":           reference("Patient", patientIdx),
		"code":              codeableConcept(coding("http://loinc.org", "29463-7")),
		"effectiveDateTime": date.Format("2006-01-02"),
		"valueQuantity":     quantity(value, "kg"),
	}
}

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
	"time"
)

func TobaccoUse(r *rand.Rand, patientIdx int, time time.Time) Object {
	tobaccoUse := make(map[string]interface{})
	tobaccoUse["resourceType"] = "Observation"
	tobaccoUse["id"] = fmt.Sprintf("%d-tobacco-use", patientIdx)
	tobaccoUse["meta"] = meta("https://fhir.bbmri.de/StructureDefinition/TobaccoUse")
	tobaccoUse["status"] = "final"
	tobaccoUse["code"] = codeableConcept(coding("http://loinc.org", "72166-2"))
	tobaccoUse["subject"] = reference("Patient", patientIdx)
	tobaccoUse["effectiveDateTime"] = time.Format("2006-01-02")
	tobaccoUse["valueCodeableConcept"] = codeableConcept(coding("http://loinc.org", randSmokingStatus(r)))

	return tobaccoUse
}

func randSmokingStatus(r *rand.Rand) string {
	switch r.Intn(7) {
	case 0:
		return "LA18976-3"
	case 1:
		return "LA18977-1"
	case 2:
		return "LA15920-4"
	case 3:
		return "LA18978-9"
	case 4:
		return "LA18979-7"
	case 5:
		return "LA18980-5"
	case 6:
		return "LA18981-3"
	case 7:
		return "LA18982-1"
	default:
		panic("too many switch cases")
	}
}

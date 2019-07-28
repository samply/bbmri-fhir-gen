// Copyright © 2019 Samply Community
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

func Condition(r *rand.Rand, donorIdx int, date time.Time) Object {
	return Object{
		"resourceType": "Condition",
		"id":           fmt.Sprintf("%d-condition", donorIdx),
		"meta":         meta("https://fhir.bbmri.de/StructureDefinition/BbmriCondition"),
		"subject":      reference("Patient", donorIdx),
		"code":         codeableConcept(codingWithVersion("http://hl7.org/fhir/sid/icd-10", "2016", randIcd10Code(r))),
		//"onsetDateTime": date.Format("2006-01-02"),
	}
}

func randIcd10Code(r *rand.Rand) string {
	return fmt.Sprintf("%s%02d", string(65+r.Intn(26)), r.Intn(100))
}

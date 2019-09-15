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

func Patient(r *rand.Rand, idx int) Object {
	patient := make(map[string]interface{})
	patient["resourceType"] = "Patient"
	patient["id"] = fmt.Sprint(idx)
	patient["meta"] = meta("https://fhir.bbmri.de/StructureDefinition/Patient")
	patient["gender"] = randGender(r)

	birthDate := randDate(r, 1950, 2000)
	patient["birthDate"] = birthDate.Format("2006-01-02")

	deceasedDate := birthDate.Add(randAge(r))
	if deceasedDate.Before(time.Now()) {
		patient["deceasedDateTime"] = deceasedDate.Format("2006-01-02")
	}

	return patient
}

var genders = []string{"male", "female"}

func randGender(r *rand.Rand) string {
	return genders[r.Intn(len(genders))]
}

func randDate(r *rand.Rand, startYear int, endYear int) time.Time {
	start := time.Date(startYear, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()
	end := time.Date(endYear, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()
	return time.Unix(start+r.Int63n(end-start), 0)
}

func randAge(r *rand.Rand) time.Duration {
	return time.Duration(r.Intn(80*365)+10*365) * 24 * time.Hour
}

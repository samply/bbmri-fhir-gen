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
	"github.com/stretchr/testify/assert"
	"math/rand"
	"testing"
	"time"
)

func TestSpecimen(t *testing.T) {
	date, err := time.Parse("2006-01-02", "2019-07-30")
	if err != nil {
		t.Error(err)
		return
	}
	specimen := Specimen(rand.New(rand.NewSource(1)), 0, date)
	assert.Equal(t,
		"https://fhir.bbmri.de/StructureDefinition/Specimen",
		specimen["meta"].(Object)["profile"].(Array)[0])
	assert.Equal(t, "Specimen", specimen["resourceType"])
	assert.Equal(t, "0-specimen", specimen["id"])
	assert.Equal(t, "https://fhir.bbmri.de/StructureDefinition/StorageTemperature", specimen["extension"].(Array)[0].(Object)["url"])
	assert.Equal(t, "https://fhir.bbmri.de/CodeSystem/StorageTemperature", specimen["extension"].(Array)[0].(Object)["valueCodeableConcept"].(Object)["coding"].(Array)[0].(Object)["system"])
	assert.Equal(t, "temperatureRoom", specimen["extension"].(Array)[0].(Object)["valueCodeableConcept"].(Object)["coding"].(Array)[0].(Object)["code"])
	assert.Equal(t, "https://fhir.bbmri.de/StructureDefinition/SampleDiagnosis", specimen["extension"].(Array)[1].(Object)["url"])
	assert.Equal(t, "http://hl7.org/fhir/sid/icd-10", specimen["extension"].(Array)[1].(Object)["valueCodeableConcept"].(Object)["coding"].(Array)[0].(Object)["system"])
	assert.Equal(t, "V47.9", specimen["extension"].(Array)[1].(Object)["valueCodeableConcept"].(Object)["coding"].(Array)[0].(Object)["code"])
	assert.Equal(t, "https://fhir.bbmri.de/StructureDefinition/Custodian", specimen["extension"].(Array)[2].(Object)["url"])
	assert.Equal(t, "Organization/collection-1", specimen["extension"].(Array)[2].(Object)["valueReference"].(Object)["reference"])

	assert.Equal(t, "Patient/0", specimen["subject"].(Object)["reference"])

	assert.Equal(t, "2019-07-30", specimen["collection"].(Object)["collectedDateTime"])
	assert.Equal(t, "urn:oid:2.16.840.1.113883.6.43.1", specimen["collection"].(Object)["bodySite"].(Object)["coding"].(Array)[0].(Object)["system"])
	assert.Equal(t, "C18.5", specimen["collection"].(Object)["bodySite"].(Object)["coding"].(Array)[0].(Object)["code"])
	assert.Equal(t, "http://terminology.hl7.org/CodeSystem/v2-0916", specimen["collection"].(Object)["fastingStatusCodeableConcept"].(Object)["coding"].(Array)[0].(Object)["system"])
	assert.Equal(t, "NF", specimen["collection"].(Object)["fastingStatusCodeableConcept"].(Object)["coding"].(Array)[0].(Object)["code"])
}

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

func TestTobaccoUse(t *testing.T) {
	date, err := time.Parse("2006-01-02", "2019-07-30")
	if err != nil {
		t.Error(err)
		return
	}
	tobaccoUse := TobaccoUse(rand.New(rand.NewSource(1)), 0, date)
	assert.Equal(t,
		"https://fhir.bbmri.de/StructureDefinition/TobaccoUse",
		tobaccoUse["meta"].(Object)["profile"].(Array)[0])
	assert.Equal(t, "Observation", tobaccoUse["resourceType"])
	assert.Equal(t, "0-tobacco-use", tobaccoUse["id"])
	assert.Equal(t, "http://loinc.org", tobaccoUse["code"].(Object)["coding"].(Array)[0].(Object)["system"])
	assert.Equal(t, "72166-2", tobaccoUse["code"].(Object)["coding"].(Array)[0].(Object)["code"])
	assert.Equal(t, "http://loinc.org", tobaccoUse["valueCodeableConcept"].(Object)["coding"].(Array)[0].(Object)["system"])
	assert.Equal(t, "LA18977-1", tobaccoUse["valueCodeableConcept"].(Object)["coding"].(Array)[0].(Object)["code"])
	assert.Equal(t, "Patient/0", tobaccoUse["subject"].(Object)["reference"])
}

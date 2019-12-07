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
	"testing"
	"time"
)

func TestBodyWeight(t *testing.T) {
	date, err := time.Parse("2006-01-02", "2019-07-30")
	if err != nil {
		t.Error(err)
		return
	}
	weight := BodyWeight(0, date, 70.1)
	assert.Equal(t,
		"https://fhir.bbmri.de/StructureDefinition/BodyWeight",
		weight["meta"].(Object)["profile"].(Array)[0])
	assert.Equal(t, "Observation", weight["resourceType"])
	assert.Equal(t, "bbmri-0-body-weight", weight["id"])
	assert.Equal(t, "http://loinc.org", weight["code"].(Object)["coding"].(Array)[0].(Object)["system"])
	assert.Equal(t, "29463-7", weight["code"].(Object)["coding"].(Array)[0].(Object)["code"])
	assert.Equal(t, "2019-07-30", weight["effectiveDateTime"])
	assert.Equal(t, 70.0, weight["valueQuantity"].(Object)["value"])
	assert.Equal(t, "kg", weight["valueQuantity"].(Object)["unit"])
	assert.Equal(t, "kg", weight["valueQuantity"].(Object)["code"])
	assert.Equal(t, "http://unitsofmeasure.org", weight["valueQuantity"].(Object)["system"])
	assert.Equal(t, "Patient/bbmri-0", weight["subject"].(Object)["reference"])
}

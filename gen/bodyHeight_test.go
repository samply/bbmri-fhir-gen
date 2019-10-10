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

func TestBodyHeight(t *testing.T) {
	date, err := time.Parse("2006-01-02", "2019-07-30")
	if err != nil {
		t.Error(err)
		return
	}
	height := BodyHeight(0, date, 170.0)
	assert.Equal(t,
		"https://fhir.bbmri.de/StructureDefinition/BodyHeight",
		height["meta"].(Object)["profile"].(Array)[0])
	assert.Equal(t, "Observation", height["resourceType"])
	assert.Equal(t, "bbmri-0-body-height", height["id"])
	assert.Equal(t, "http://loinc.org", height["code"].(Object)["coding"].(Array)[0].(Object)["system"])
	assert.Equal(t, "8302-2", height["code"].(Object)["coding"].(Array)[0].(Object)["code"])
	assert.Equal(t, "2019-07-30", height["effectiveDateTime"])
	assert.Equal(t, 170.0, height["valueQuantity"].(Object)["value"])
	assert.Equal(t, "cm", height["valueQuantity"].(Object)["unit"])
	assert.Equal(t, "cm", height["valueQuantity"].(Object)["code"])
	assert.Equal(t, "http://unitsofmeasure.org", height["valueQuantity"].(Object)["system"])
	assert.Equal(t, "Patient/bbmri-0", height["subject"].(Object)["reference"])
}

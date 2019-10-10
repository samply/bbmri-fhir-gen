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

func TestCondition(t *testing.T) {
	date, err := time.Parse("2006-01-02", "2019-07-30")
	if err != nil {
		t.Error(err)
		return
	}
	condition := Condition(rand.New(rand.NewSource(5)), 0, 0, date)
	assert.Equal(t,
		"https://fhir.bbmri.de/StructureDefinition/Condition",
		condition["meta"].(Object)["profile"].(Array)[0])
	assert.Equal(t, "Condition", condition["resourceType"])
	assert.Equal(t, "bbmri-0-condition-0", condition["id"])
	assert.Equal(t, "2019-07-30", condition["onsetDateTime"])
	assert.Equal(t, "http://hl7.org/fhir/sid/icd-10", condition["code"].(Object)["coding"].(Array)[0].(Object)["system"])
	assert.Equal(t, "2016", condition["code"].(Object)["coding"].(Array)[0].(Object)["version"])
	assert.Equal(t, "Q36.9", condition["code"].(Object)["coding"].(Array)[0].(Object)["code"])
	assert.Equal(t, "Patient/bbmri-0", condition["subject"].(Object)["reference"])
}

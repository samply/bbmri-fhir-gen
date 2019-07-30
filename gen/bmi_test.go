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

func TestBmi(t *testing.T) {
	date, err := time.Parse("2006-01-02", "2019-07-30")
	if err != nil {
		t.Error(err)
		return
	}
	bmi := Bmi(0, date, 42.0)
	assert.Equal(t,
		"https://fhir.bbmri.de/StructureDefinition/BbmriBmi",
		bmi["meta"].(Object)["profile"].(Array)[0])
	assert.Equal(t, "0-bmi", bmi["id"])
	assert.Equal(t, "2019-07-30", bmi["effectiveDateTime"])
	assert.Equal(t, 42.0, bmi["valueQuantity"].(Object)["value"])
	assert.Equal(t, "kg/m2", bmi["valueQuantity"].(Object)["unit"])
	assert.Equal(t, "kg/m2", bmi["valueQuantity"].(Object)["code"])
	assert.Equal(t, "http://unitsofmeasure.org", bmi["valueQuantity"].(Object)["system"])
}

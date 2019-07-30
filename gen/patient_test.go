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
)

func TestPatient(t *testing.T) {
	patient := Patient(rand.New(rand.NewSource(0)), 0)
	assert.Equal(t,
		"https://fhir.bbmri.de/StructureDefinition/Patient",
		patient["meta"].(Object)["profile"].(Array)[0])
	assert.Equal(t, "0", patient["id"])
}

func TestPatient_0(t *testing.T) {
	patient := Patient(rand.New(rand.NewSource(0)), 0)
	assert.Equal(t, "male", patient["gender"])
	assert.Equal(t, "1989-11-18", patient["birthDate"])
	assert.Nil(t, patient["deceasedDateTime"])
}

func TestPatient_1(t *testing.T) {
	patient := Patient(rand.New(rand.NewSource(1)), 0)
	assert.Equal(t, "female", patient["gender"])
	assert.Equal(t, "1953-11-30", patient["birthDate"])
	assert.Equal(t, "1994-02-25", patient["deceasedDateTime"])
}

func TestPatient_2(t *testing.T) {
	patient := Patient(rand.New(rand.NewSource(2)), 0)
	assert.Equal(t, "male", patient["gender"])
	assert.Equal(t, "1969-08-10", patient["birthDate"])
	assert.Nil(t, patient["deceasedDateTime"])
}

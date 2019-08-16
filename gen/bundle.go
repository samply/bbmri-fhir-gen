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
	"github.com/google/uuid"
	"math"
	"math/rand"
)

func Bundle(r *rand.Rand, start int, n int) Object {
	entries := make(Array, 0, n)
	for i := start; i < start+n; i++ {

		patient := Patient(r, i)
		entries = append(entries, entry(patient))


		encounterDate := randDate(r, 2000, 2018)
		bmi := RandBmiValue(r)
		bodyHeight := RandBodyHeightValue(r)
		bodyWeight := bmi * math.Pow(bodyHeight/100, 2)
		entries = append(entries, entry(Bmi(i, encounterDate, bmi)))
		entries = append(entries, entry(BodyHeight(i, encounterDate, bodyHeight)))
		entries = append(entries, entry(BodyWeight(i, encounterDate, bodyWeight)))
		entries = append(entries, entry(Condition(r, i, encounterDate)))
		entries = append(entries, entry(TobaccoUse(r,i,encounterDate)))
		entries = append(entries, entry(Specimen(r,i,encounterDate)))

		if patient["deceasedDateTime"] != nil {
			entries = append(entries, entry(CauseOfDeath(r,i)))
		}

	}
	return Object{
		"resourceType": "Bundle",
		"id":           uuid.New().String(),
		"type":         "transaction",
		"entry":        entries,
	}
}

func entry(resource Object) Object {
	return Object{
		"fullUrl":  uuid.New().String(),
		"resource": resource,
		"request": Object{
			"method": "PUT",
			"url":    fmt.Sprintf("%s/%s", resource["resourceType"], resource["id"]),
		},
	}
}

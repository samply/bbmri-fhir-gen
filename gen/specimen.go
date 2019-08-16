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

func Specimen(r *rand.Rand, patientIdx int, date time.Time) Object {
	return Object{
		"resourceType": "Specimen",
		"id":           fmt.Sprintf("%d-specimen", patientIdx),
		"meta":         meta("https://fhir.bbmri.de/StructureDefinition/Specimen"),
		"extension":    Array{storageTemp(r), sampleDiagnosis(r), custodian(r)},
		"type":         codeableConcept(coding("https://fhir.bbmri.de/CodeSystem/SampleMaterialType", "whole-blood")),
		"subject":      reference("Patient", patientIdx),
		"collection":   collection(r, date),
	}
}

func storageTemp(r *rand.Rand) Object {
	coding := coding(
		"https://fhir.bbmri.de/CodeSystem/StorageTemperature",
		randStorageTemp(r))

	return bbmriExtensionCodeableConcept("StorageTemperature", codeableConcept(coding))
}

func randStorageTemp(r *rand.Rand) string {
	switch n := r.Intn(6); n {
	case 0:
		return "temperature2to10"
	case 1:
		return "temperature-18to-35"
	case 2:
		return "temperature-60to-85"
	case 3:
		return "temperatureGN"
	case 4:
		return "temperatureLN"
	case 5:
		return "temperatureRoom"
	case 6:
		return "temperatureOther"
	default:
		panic("too many switch cases")
	}
}

func sampleDiagnosis(r *rand.Rand) Object {
	coding := coding("http://hl7.org/fhir/sid/icd-10", randIcd10Code(r))

	return bbmriExtensionCodeableConcept("SampleDiagnosis", codeableConcept(coding))
}

func custodian(r *rand.Rand) Object {
	return bbmriExtensionReference(
		"Custodian",
		stringReference("Organization", fmt.Sprintf("collection-%d", r.Intn(10))))
}

func collection(r *rand.Rand, date time.Time) Object {
	return Object{
		"collectedDateTime":            date.Format("2006-01-02"),
		"bodySite":                     codeableConcept(coding("https://www.who.int/classifications/icd/adaptations/oncology/en/", randIcdOCode(r))),
		"fastingStatusCodeableConcept": codeableConcept(coding("http://terminology.hl7.org/CodeSystem/v2-0916", randFastingStatus(r))),
	}
}

func randFastingStatus(r *rand.Rand) string {
	switch n := r.Intn(3); n {
	case 0:
		return "F"
	case 1:
		return "FNA"
	case 2:
		return "NF"
	case 3:
		return "NG"
	default:
		panic("too many switch cases")
	}
}

func randIcdOCode(r *rand.Rand) string {
	return fmt.Sprintf("C%02d.%d", r.Intn(100), r.Intn(10))
}

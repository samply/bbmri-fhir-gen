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

func Specimen(r *rand.Rand, patientIdx int, specimenIdx int, date time.Time, ejprd bool) Object {
	specimen := Object{
		"resourceType": "Specimen",
		"id":           fmt.Sprintf("bbmri-%d-specimen-%d", patientIdx, specimenIdx),
		"meta":         meta("https://fhir.bbmri.de/StructureDefinition/Specimen"),
		"extension":    Array{storageTemp(r), sampleDiagnosis(r), custodian(r)},
		"type":         codeableConcept(materialTypeCoding(r)),
		"subject":      patientReference(patientIdx),
		"collection":   collection(r, date),
	}

	if ejprd {
		specimen["status"] = randStatus(r)
	}

	return specimen
}

func storageTemp(r *rand.Rand) Object {
	coding := coding(
		"https://fhir.bbmri.de/CodeSystem/StorageTemperature",
		randStorageTemp(r))

	return bbmriExtensionCodeableConcept("StorageTemperature", codeableConcept(coding))
}

var storageTemps = []string{
	"temperature2to10",
	"temperature-18to-35",
	"temperature-60to-85",
	"temperatureGN",
	"temperatureLN",
	"temperatureRoom",
	"temperatureOther",
}

func randStorageTemp(r *rand.Rand) string {
	return storageTemps[r.Intn(len(storageTemps))]
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
		"bodySite":                     codeableConcept(coding("urn:oid:2.16.840.1.113883.6.43.1", randIcdOCode(r))),
		"fastingStatusCodeableConcept": codeableConcept(coding("http://terminology.hl7.org/CodeSystem/v2-0916", randFastingStatus(r))),
	}
}

var fastingStatus = []string{"F", "FNA", "NF", "NG"}

func randFastingStatus(r *rand.Rand) string {
	return fastingStatus[r.Intn(len(fastingStatus))]
}

func randIcdOCode(r *rand.Rand) string {
	return fmt.Sprintf("C%02d.%d", r.Intn(100), r.Intn(10))
}

func materialTypeCoding(r *rand.Rand) Object {
	return coding("https://fhir.bbmri.de/CodeSystem/SampleMaterialType",
		randMaterialType(r))
}

var materialTypes = []string{
	"tissue",
	"tissue-formalin",
	"tissue-frozen",
	"tissue-paxgene-or-else",
	"tissue-other",
	"liquid",
	"whole-blood",
	"blood-plasma",
	"blood-serum",
	"peripheral-blood-cells-vital",
	"buffy-coat",
	"bone-marrow",
	"csf-liquor",
	"ascites",
	"urine",
	"saliva",
	"stool-faeces",
	"liquid-other",
	"derivative",
	"dna",
	"cf-dna",
	"rna",
	"derivative-other",
}

func randMaterialType(r *rand.Rand) string {
	return materialTypes[r.Intn(len(materialTypes))]
}

// Skew the distribution towards "available"
var statuses = []string{
	"available",
	"available",
	"available",
	"available",
	"available",
	"available",
	"available",
	"available",
	"unavailable",
	"available",
	"available",
	"available",
	"available",
	"available",
	"available",
	"available",
	"available",
	"unsatisfactory",
	"available",
	"available",
	"available",
	"available",
	"available",
	"available",
	"available",
	"available",
}

func randStatus(r *rand.Rand) string {
	return statuses[r.Intn(len(statuses))]
}

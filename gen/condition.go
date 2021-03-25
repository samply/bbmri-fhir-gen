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

func Condition(r *rand.Rand, patientIdx int, conditionIdx int, date time.Time, ejprd bool) Object {
	resource := make(map[string]interface{})
	resource["resourceType"] = "Condition"
	resource["id"] = fmt.Sprintf("bbmri-%d-condition-%d", patientIdx, conditionIdx)
	resource["meta"] = meta("https://fhir.bbmri.de/StructureDefinition/Condition")
	resource["subject"] = patientReference(patientIdx)
	resource["code"] = generateCodeableConceptCode(r, ejprd)
	resource["onsetDateTime"] = date.Format("2006-01-02")
	if ejprd{
		resource["clinicalStatus"] = generateCodeableConceptStatus(r, ejprd)
	}
	return resource
}

func randIcd10Code(r *rand.Rand) string {
	return fmt.Sprintf("%s%02d.%d", string(65+r.Intn(26)), r.Intn(100), r.Intn(10))
}

var orphaCodes = [][]string{
	{"423461","Mucolipidosis type III alpha/beta"},
	{"1440","Ring chromosome 14 syndrome"},
	{"1440","Ring chromosome 14 syndrome"},
	{"1440","Ring chromosome 14 syndrome"},
	{"778","Rett syndrome - classic rett"},
	{"98896","Duchenne muscular dystrophy"},
	{"267","Limb girdle muscular dystrophy type 2A"},
	{"97242","Congenital myopathy"},
	{"98895","Becker muscular dystrophy"},
	{"206549","HyperCKemia"},
	{"98896","Duchenne muscular dystrophy"},
	{"2322","Kabuki syndrome"},
	{"98895","Becker muscular dystrophy"},
	{"98896","Duchenne muscular dystrophy"},
	{"104","Leber hereditary optic atrophy (LHON); LHON14484C; MTND6"},
	{"314051","Combined oxidative phosphorylation deficiency 12; EARS2"},
	{"321","Multiple Ostechondromas"},
	{"43","X-linked adrenoleukodystrophy"},
	{"93545","Duplication of ureter"},
	{"93545","Congenital vesico-uretero-renal reflux"},
	{"2133","Haemoglobin variant"},
	{"848","Beta thalassemia carrier"},
	{"848","Beta thalassemia carrier"},
}

func randHpOrphaCode(r *rand.Rand) []string {
	return orphaCodes[r.Intn(len(orphaCodes))]
}

func generateCodings(r *rand.Rand, ejprd bool) []Object {
	icd10Coding := codingWithVersion("http://hl7.org/fhir/sid/icd-10", "2016", randIcd10Code(r))
	orphaCode := randHpOrphaCode(r)
	codingCount := 1
	if ejprd {
		codingCount = 2
	}
	codings := make([] Object, codingCount)
	codings[0] = icd10Coding
	if ejprd {
		orphaCoding := codingDisplay("http://www.orpha.net", orphaCode[0], orphaCode[1])
		codings[1] = orphaCoding
	}
	return codings
}

func generateCodeableConceptCode(r *rand.Rand, ejprd bool) Object {
	codings := generateCodings(r, ejprd)
	concept := codeableConcepts(codings)
	return concept
}

// Skew the distribution towards "active".
var clinicalStatus = []string{
	"active",
	"active",
	"active",
	"recurrence",
	"active",
	"active",
	"active",
	"relapse",
	"active",
	"active",
	"active",
	"inactive",
	"active",
	"active",
	"active",
	"remission",
	"active",
	"active",
	"active",
	"resolved",
	"active",
	"active",
	"active",
}

func randClinicalStatus(r *rand.Rand) string {
	return clinicalStatus[r.Intn(len(clinicalStatus))]
}

func generateCodeableConceptStatus(r *rand.Rand, ejprd bool) Object {
	concept := codeableConcept(coding("http://terminology.hl7.org/CodeSystem/condition-clinical", randClinicalStatus(r)))
	return concept
}

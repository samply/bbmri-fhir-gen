// Copyright © 2021 The Samply Development Community
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

func HpDiagnosis(r *rand.Rand, patientIdx int, time time.Time) Object {
	hpDiagnosis := make(map[string]interface{})
	codeableConcepts := diagnosticCodeableConcept(r)
	hpDiagnosis["resourceType"] = "Observation"
	hpDiagnosis["id"] = fmt.Sprintf("bbmri-%d-hp-diagnosis", patientIdx)
	hpDiagnosis["meta"] = meta("https://fhir.bbmri.de/StructureDefinition/Observation")
	hpDiagnosis["status"] = "final"
	hpDiagnosis["subject"] = patientReference(patientIdx)
	hpDiagnosis["effectiveDateTime"] = time.Format("2006-01-02")
	hpDiagnosis["code"] = codeableConcept(codeableConcepts)

	return hpDiagnosis
}

var diagnoses = [][]string{
	{"HP:0000007","Autosomal recessive inheritance"},
	{"HP:0000020","Urinary incontinence"},
	{"HP:0000076","Vesicoureteral reflux"},
	{"HP:0000175","Cleft palate"},
	{"HP:0000252","Microcephaly"},
	{"HP:0000365","Hearing impairment"},
	{"HP:0000479","Abnormality of the retina"},
	{"HP:0000505","Visual impairment"},
	{"HP:0000508","Ptosis"},
	{"HP:0000572","Visual loss"},
	{"HP:0000602","Ophthalmoplegia"},
	{"HP:0000924","Abnormality of the skeletal system"},
	{"HP:0001004","Lymphedema"},
	{"HP:0001249","Intellectual disability"},
	{"HP:0001250","Seizures"},
	{"HP:0001263","Global developmental delay"},
	{"HP:0001285","Spastic tetraparesis"},
	{"HP:0001319","Neonatal hypotonia"},
	{"HP:0001332","Dystonia"},
	{"HP:0001396","Cholestasis"},
	{"HP:0001403","Macrovesicular hepatic steatosis"},
	{"HP:0001508","Failure to thrive"},
	{"HP:0001510","Growth delay"},
	{"HP:0001640","Cardiomegaly"},
	{"HP:0002067","Bradykinesia"},
	{"HP:0002079","Hypoplasia of the corpus callosum"},
	{"HP:0002151","Increased serum lactate"},
	{"HP:0002240","Hepatomegaly"},
	{"HP:0002352","Leukoencephalopathy"},
	{"HP:0002355","difficulty in walking"},
	{"HP:0002376","Developmental regression"},
	{"HP:0002515","waddling gait"},
	{"HP:0002913","Myoglobinuria"},
	{"HP:0003128","Lactic acidosis"},
	{"HP:0003200","Ragged-red muscle fibers"},
	{"HP:0003236","hyperCKemia"},
	{"HP:0003325","limb girdle muscle weakness"},
	{"HP:0003593","Infantile onset"},
	{"HP:0003701","Muscle proximal weakness"},
	{"HP:0006989","Dysplastic corpus callosum"},
	{"HP:0007126","Proximal amyotrophy"},
	{"HP:0007340","Lower limb muscle weakness"},
	{"HP:0008347","Decreased activity of mitochondrial complex IV"},
	{"HP:0008981","calf muscle hypertrophy"},
	{"HP:0011342","Mild global developmental delay"},
	{"HP:0011923","Decreased activity of mitochondrial complex I"},
	{"HP:0011924","Decreased activity of mitochondrial complex III"},
	{"HP:0012572","Ureter duplex"},
}

func randHpDiagnosis(r *rand.Rand) []string {
	return diagnoses[r.Intn(len(diagnoses))]
}

func randHpDiagnoses(r *rand.Rand) [][]string {
	const maxDiagnoses int = 10
	diagnosisCount := r.Intn(maxDiagnoses) + 1
	entries := make([][]string, diagnosisCount)
	duplicateCount := 0
	for i := range entries {
		entries[i] = make([]string, 2)
		idx := r.Intn(len(diagnoses))
		hpCode := diagnoses[idx][0]
		description := diagnoses[idx][1]
		duplicate := false
		for j := 0; j < i; j++ {
			if entries[j][0] == hpCode {
				duplicate = true
				break
			}
		}
		if duplicate {
			duplicateCount++
			entries[i][0] = "ZZZZ"
			continue
		}
		entries[i][0] = hpCode
		entries[i][1] = description
	}
	uniqueCount := diagnosisCount - duplicateCount
	uniqueEntries := make([][]string, uniqueCount)
	uniqueNum := 0
	for i := range entries {
		if entries[i][0] != "ZZZZ" {
			uniqueEntries[uniqueNum] = entries[i]
			uniqueNum++
		}
	}
	return uniqueEntries
}

func diagnosisCodings(r *rand.Rand) []Object {
	diagnoses := randHpDiagnoses(r)
	codings := make([] Object, len(diagnoses))
	for i := range diagnoses {
		coding := codingDisplay("http://purl.obolibrary.org/obo/hp.owl", diagnoses[i][0], diagnoses[i][1])
		codings[i] = coding
	}
	return codings
}

func diagnosticCodeableConcept(r *rand.Rand) Object {
	codings :=diagnosisCodings(r)
	concept := codeableConcepts(codings)
	return concept
}
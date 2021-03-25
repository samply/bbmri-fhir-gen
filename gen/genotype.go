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

func Genotype(r *rand.Rand, patientIdx int, time time.Time) Object {
	Variant := make(map[string]interface{})
	codeableConcepts := variantCodeableConcept(r)
	Variant["resourceType"] = "Observation"
	Variant["id"] = fmt.Sprintf("bbmri-%d-genotype", patientIdx)
	Variant["meta"] = meta("https://fhir.bbmri.de/StructureDefinition/Observation")
	Variant["status"] = "final"
	Variant["subject"] = patientReference(patientIdx)
	Variant["effectiveDateTime"] = time.Format("2006-01-02")
	Variant["extension"] = extension("https://varnomen.hgvs.org/", "valueCodeableConcept", codeableConcepts)

	return Variant
}

// This is supposed to be HGVS
var variants = []string{
	"(14q32.33)(A_14_P117532-->A_14_p135856)x1",
	"(14q32.33)(A_14_P135580-->A_14_P138389-->A_14_P135856)x1",
	"14q32.33(SNP-A-2449096>SNP_A-1227677)x1n",
	"ANO5:c.1627dupA#etero",
	"ANO5:c.2498T>A#etero",
	"ANO5:fs*10aaSTOP",
	"c.2659dupA",
	"c.2777A>C",
	"c.369_370delAG",
	"c.691T>A",
	"CAPN3:c.1193-6T>A",
	"CAPN3:c.1746-20C>G",
	"DMD:del45",
	"DMD:delex3-12",
	"Dystrophin:del45",
	"EARS2:c.1278_1279insCTC",
	"EARS2:c.332C>T",
	"EARS2:c.502A>G",
	"MTND6:m.14484T>C",
	"TRIM32:c.C1786G:",
}

func randVariants(r *rand.Rand) []string {
	const maxVariants int = 3
	variantCount := r.Intn(maxVariants) + 1
	entries := make([]string, variantCount)
	duplicateCount := 0
	for i := range entries {
		idx := r.Intn(len(variants))
		variant := variants[idx]
		duplicate := false
		for j := 0; j < i; j++ {
			if entries[j] == variant {
				duplicate = true
				break
			}
		}
		if duplicate {
			duplicateCount++
			entries[i] = "ZZZZ"
			continue
		}
		entries[i] = variant
	}
	uniqueCount := variantCount - duplicateCount
	uniqueEntries := make([]string, uniqueCount)
	uniqueNum := 0
	for i := range entries {
		if entries[i] != "ZZZZ" {
			uniqueEntries[uniqueNum] = entries[i]
			uniqueNum++
		}
	}
	return uniqueEntries
}

func variantCodings(r *rand.Rand) []Object {
	variants := randVariants(r)
	codings := make([] Object, len(variants))
	for i := range variants {
		coding := codingDisplay("https://varnomen.hgvs.org/", variants[i], "Genotype")
		codings[i] = coding
	}
	return codings
}

func variantCodeableConcept(r *rand.Rand) Object {
	codings :=variantCodings(r)
	concept := codeableConcepts(codings)
	return concept
}
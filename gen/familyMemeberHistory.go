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
	"strings"
)

func FamilyMemberHistory(r *rand.Rand, patientIdx int, patient Object) Object {
	familyMemeberHistory := make(map[string]interface{})
	familyMemeberHistory["resourceType"] = "FamilyMemberHistory"
	familyMemeberHistory["id"] = fmt.Sprintf("bbmri-%d-relative", patientIdx)
	familyMemeberHistory["patient"] = patientReference(patientIdx)
	birthDate := randDate(r, 1950, 2010)
	familyMemeberHistory["bornDate"] = birthDate.Format("2006-01-02")
	familyMemeberHistory["sex"] = generateSex(r)
	familyMemeberHistory["relationship"] = generateRelationship(r)
	familyMemeberHistory["name"] = generateRelativeName(r, patient)
	familyMemeberHistory["condition"] = generateCondition(r)
	familyNumber := generateFamilyNumber(r)
	// Give patient and family mamber the same family code
	familyMemeberHistory["extension"] = familyNumber
	patient["extension"] = familyNumber

	return familyMemeberHistory
}

var sexes = []string{
	"male",
	"female",
}

func randSex(r *rand.Rand) string {
	return sexes[r.Intn(len(sexes))]
}

func generateSex(r *rand.Rand) Object {
	sex := randSex(r)
	coding := make(map[string]interface{})
	coding["coding"] = wrapInArray(codeableConcept(codingDisplay("http://hl7.org/fhir/ValueSet/administrative-gender", sex, camelizer(sex))))
	return coding
}

func camelizer(str string) string {
	var strCamel string
	for c, l := range str {
		if c == 0 {
			strCamel = strings.ToUpper(string(l))
		} else {
			strCamel += string(l)
		}
	}
	return strCamel
}

var relationships = []string{
	"FAMMEMB",
	"DAUC",
	"SON",
	"STPCHLD",
	"AUNT",
	"COUSN",
	"GGRPRN",
	"GRNDDAU",
	"GRNDSON",
	"NEPHEW",
	"NIECE",
	"UNCLE",
	"FTH",
	"MTH",
}

func randRelationship(r *rand.Rand) string {
	return relationships[r.Intn(len(relationships))]
}

func generateRelationship(r *rand.Rand) Object {
	relationship := randRelationship(r)
	obj := make(map[string]interface{})
	obj["coding"] = wrapInArray(codeableConcept(coding("http://terminology.hl7.org/ValueSet/v3-FamilyMember", relationship)))
	return obj
}

func generateRelativeName(r *rand.Rand, patient Object) string {
	forenames := randForenames(r)
	// I was hoping to extract a family name from the patient, but couldnt figure out how
	familyName := generateFamilyName(r)
	name := strings.Join(forenames, " ")
	name = strings.Join([]string{name, familyName}, " ")
	return name
}

var familyNumbers = []string{
	"F00695",
	"10990",
	"10762",
	"9915",
	"N0005/BMD",
	"N0001/DMD",
	"4152",
	"5321",
}

func generateFamilyNumber(r *rand.Rand) Array {
	number := familyNumbers[r.Intn(len(familyNumbers))]
	familyNumber := make(map[string]interface{})
	familyNumber["url"] = "https://fhir.bbmri.de/StructureDefinition/FamilyNumber"
	familyNumber["valueString"] = number
	return wrapInArray(familyNumber)
}

var conditions = [][]string{
	{"423461","Mucolipidosis type III alpha/beta","Joint stiffness; Thick skull base; Umbilical hernia"},
	{"2322","Kabuki syndrome"},
	{"43","X-linked adrenoleukodystrophy","Seizures; Abnormality of the cerebral white matter"},
}

func randCondition(r *rand.Rand) []string {
	return conditions[r.Intn(len(conditions))]
}

func generateCondition(r *rand.Rand) Array {
	condition := make(map[string]interface{})
	conditionParts := randCondition(r)
	orphaCode := conditionParts[0]
	orphaDesc := conditionParts[1]
	condition["code"] = codeableConcept(codingDisplay("http://www.orpha.net", orphaCode, orphaDesc))
	condition["note"] = generateNote(conditionParts)
	return wrapInArray(condition)
}

func generateNote(conditionParts []string) []map[string]string {
	var length = len(conditionParts) - 1
	var notes = make([]map[string]string, length)
	note1 := map[string]string{"text": conditionParts[1]}
	notes[0] = note1
	if len(conditionParts) > 2 {
		note2 := map[string]string{"text": conditionParts[2]}
		notes[1] = note2
	}
	return notes
}
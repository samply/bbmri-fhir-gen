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

import "fmt"

func Collection(idx int) Object {
	return Object{
		"resourceType": "Organization",
		"id":           fmt.Sprintf("collection-%d", idx),
		"meta":         meta("https://fhir.bbmri.de/StructureDefinition/Collection"),
		"extension": Array{
			bbmriExtensionString(
				"OrganizationDescription",
				"The Collection consists of blood from patients at the Klinikum Neustadt."),
			bbmriExtensionCodeableConcept(
				"CollectionType",
				codeableConcept(bbmriCoding("CollectionType", "HOSPITAL"))),
			bbmriExtensionCodeableConcept(
				"CollectionType",
				codeableConcept(bbmriCoding("CollectionType", "SAMPLE"))),
			bbmriExtensionCodeableConcept(
				"DataCategory",
				codeableConcept(bbmriCoding("DataCategory", "BIOLOGICAL_SAMPLES"))),
		},
		"identifier": Array{Object{
			"system": "http://www.bbmri-eric.eu/",
			"value":  "bbmri-eric:ID:de_12345:collection:vampir",
		}},
		"name":  "Venous and arterial blood from misc patients internal reserve",
		"alias": Array{"VAMPIR"},
		"partOf": Object{
			"reference": "Organization/biobank-0",
		},
		"contact": Array{Object{
			"purpose": Object{
				"coding": Array{Object{
					"system":  "http://terminology.hl7.org/CodeSystem/contactentity-type",
					"code":    "ADMIN",
					"display": "Administrative",
				}},
			},
			"name": Object{
				"family": "Dracula",
				"given":  Array{"Vladimir"},
				"prefix": Array{"Graf"},
			},
			"telecom": Array{Object{
				"system": "fax",
				"value":  "0452 4624-66",
			}},
		}},
	}
}

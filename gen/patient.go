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

func Patient(r *rand.Rand, idx int, ejprd bool) Object {
	patient := make(map[string]interface{})
	patient["resourceType"] = "Patient"
	patient["id"] = fmt.Sprintf("bbmri-%d", idx)
	patient["meta"] = meta("https://fhir.bbmri.de/StructureDefinition/Patient")
	patient["gender"] = randGender(r)

	birthDate := randDate(r, 1950, 2000)
	patient["birthDate"] = birthDate.Format("2006-01-02")

	deceasedDate := birthDate.Add(randAge(r))
	if deceasedDate.Before(time.Now()) {
		patient["deceasedDateTime"] = deceasedDate.Format("2006-01-02")
	}

	if ejprd {
		patient["identifier"] = generateIdentifier(r)
		patient["name"] = generateName(r)
		patient["address"] = generateAddress(r)
	}

	return patient
}

var genders = []string{"male", "female"}

func randGender(r *rand.Rand) string {
	return genders[r.Intn(len(genders))]
}

func randDate(r *rand.Rand, startYear int, endYear int) time.Time {
	start := time.Date(startYear, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()
	end := time.Date(endYear, time.January, 1, 0, 0, 0, 0, time.UTC).Unix()
	return time.Unix(start+r.Int63n(end-start), 0)
}

func randAge(r *rand.Rand) time.Duration {
	return time.Duration(r.Intn(80*365)+10*365) * 24 * time.Hour
}

func generateIdentifier(r *rand.Rand) Array {
	identifier := make(map[string]interface{})
	identifier["system"] = "dkfz-heidelberg.de"
	identifier["value"] = fmt.Sprintf("%d",r.Intn(10000))
	return wrapInArray(identifier)
}

var familyNames = []string{
	"Ackermann",
	"Adam",
	"Ahrens",
	"Albert",
	"Albrecht",
	"Arndt",
	"Arnold",
	"Bauer",
	"Baumann",
	"Beck",
	"Becker",
	"Berger",
	"Braun",
	"Claus",
	"Clemens",
	"Conrad",
	"Cordes",
	"Cramer",
	"Cremer",
	"Decker",
	"Diehl",
	"Dietrich",
	"Dietz",
	"Dittrich",
	"Ebert",
	"Eckert",
	"Eder",
	"Engel",
	"Engelhardt",
	"Ernst",
	"Fischer",
	"Frank",
	"Franke",
	"Franz",
	"Friedrich",
	"Fuchs",
	"Goebel",
	"Graf",
	"Grimm",
	"Haas",
	"Hahn",
	"Hartmann",
	"Herrmann",
	"Hoffmann",
	"Hofmann",
	"Huber",
	"Ibrahim",
	"Ihle",
	"Ihrig",
	"Imhof",
	"Irmer",
	"Irrgang",
	"Isele",
	"Jahn",
	"Jakob",
	"Jansen",
	"Janßen",
	"John",
	"Jung",
	"Kaiser",
	"Klein",
	"Koch",
	"Lang",
	"Lange",
	"Lehmann",
	"Lindner",
	"Lorenz",
	"Ludwig",
	"Lutz",
	"Maier",
	"Martin",
	"Mayer",
	"Meier",
	"Meyer",
	"Nagel",
	"Naumann",
	"Neubauer",
	"Neumann",
	"Nickel",
	"Nolte",
	"Nowak",
	"Obermeier",
	"Opitz",
	"Orth",
	"Oswald",
	"Ott",
	"Otte",
	"Otto",
	"Paul",
	"Peter",
	"Peters",
	"Petersen",
	"Pfeifer",
	"Pfeiffer",
	"Pohl",
	"Quaas",
	"Quade",
	"Quandt",
	"Quante",
	"Quast",
	"Queck",
	"Quint",
	"Reinhardt",
	"Reuter",
	"Richter",
	"Riedel",
	"Ritter",
	"Roth",
	"Rudolph",
	"Schmidt",
	"Schmitt",
	"Schneider",
	"Schroeder",
	"Schulz",
	"Schwarz",
	"Theis",
	"Thiel",
	"Thiele",
	"Thoma",
	"Thomas",
	"Thomsen",
	"Timm",
	"Uhlig",
	"Ulbrich",
	"Ullmann",
	"Ullrich",
	"Ulrich",
	"Unger",
	"Urban",
	"Vetter",
	"Vogel",
	"Vogt",
	"Voigt",
	"Volk",
	"Völker",
	"Vollmer",
	"Wagner",
	"Walter",
	"Weber",
	"Werner",
	"Winkler",
	"Wolf",
	"Xander",
	"Xanthopoulos",
	"Xavier",
	"Xhaferi",
	"Xhelili",
	"Xhemajli",
	"Xiao",
	"Yalcin",
	"Yasar",
	"Yavuz",
	"Yildirim",
	"Yildiz",
	"Yilmaz",
	"Yüksel",
	"Zander",
	"Zeller",
	"Ziegler",
	"Zimmermann",
	"Zink",
	"Zorn",
}

func generateFamilyName(r *rand.Rand) string {
	return familyNames[r.Intn(len(familyNames))]
}

var forenames = []string{
	"Aaron",
	"Alexander",
	"Alina",
	"Alma",
	"Amalia",
	"Amelia",
	"Amelie",
	"Anna",
	"Anton",
	"Arian",
	"Aurelia",
	"Ava",
	"Barbara",
	"Ben",
	"Benedikt",
	"Benjamin",
	"Benno",
	"Bente",
	"Bettina",
	"Bianca",
	"Birgit",
	"Bjarne",
	"Brigitte",
	"Bruno",
	"Carla",
	"Carlotta",
	"Cataleya",
	"Celine",
	"Charlotte",
	"Chiara",
	"Christian",
	"Christina",
	"Clara",
	"Claudia",
	"Cleo",
	"Cornelia",
	"Damian",
	"Daniel",
	"Daniela",
	"Dante",
	"Daria",
	"Dario",
	"David",
	"Denise",
	"Dennis",
	"Diana",
	"Dilara",
	"Dominik",
	"Elea",
	"Elena",
	"Elian",
	"Elias",
	"Elin",
	"Elina",
	"Elisa",
	"Ella",
	"Emil",
	"Emilia",
	"Emily",
	"Emma",
	"Fabian",
	"Fabienne",
	"Felix",
	"Fiete",
	"Finja",
	"Finn",
	"Fiona",
	"Florian",
	"Franziska",
	"Frederik",
	"Frieda",
	"Fynn",
	"Gabriel",
	"Gabriele",
	"Gaylord",
	"Georg",
	"Gerhard",
	"Gertrud",
	"Gesa",
	"Gina",
	"Grace",
	"Gregor",
	"Greta",
	"Gudrun",
	"Hailey",
	"Hanna",
	"Hannah",
	"Hans",
	"Heike",
	"Helena",
	"Helene",
	"Henri",
	"Henrik",
	"Henry",
	"Holly",
	"Hugo",
	"Ida",
	"Ilaria",
	"Ilja",
	"Ilvy",
	"Ilyas",
	"Ina",
	"Ines",
	"Iris",
	"Isabel",
	"Isabell",
	"Isabella",
	"Ivy",
	"Jakob",
	"Jan",
	"Jana",
	"Jaron",
	"Johanna",
	"Jona",
	"Jonas",
	"Jonathan",
	"Jonte",
	"Julia",
	"Julian",
	"Juna",
	"Karin",
	"Karl",
	"Kasimir",
	"Katharina",
	"Kaya",
	"Keno",
	"Kerstin",
	"Khai",
	"Kian",
	"Kilian",
	"Kira",
	"Klara",
	"Laura",
	"Lea",
	"Lena",
	"Leonie",
	"Levi",
	"Lia",
	"Liam",
	"Lima",
	"Lina",
	"Linus",
	"Lobo",
	"Lukas",
	"Malea",
	"Malia",
	"Malou",
	"Malu",
	"Maria",
	"Marie",
	"Matteo",
	"Mia",
	"Michael",
	"Mila",
	"Milo",
	"Mira",
	"Nadine",
	"Nael",
	"Nala",
	"Nalu",
	"Nele",
	"Nicole",
	"Nika",
	"Niklas",
	"Nina",
	"Noah",
	"Noemi",
	"Nora",
	"Odin",
	"Okka",
	"Ole",
	"Olga",
	"Oliver",
	"Olivia",
	"Omar",
	"Oona",
	"Ophelia",
	"Oskar",
	"Otis",
	"Otto",
	"Pascal",
	"Patricia",
	"Patrick",
	"Paul",
	"Paula",
	"Paulina",
	"Pauline",
	"Pepe",
	"Peter",
	"Petra",
	"Philipp",
	"Pia",
	"Qazim",
	"Quentin",
	"Quinn",
	"Quinten",
	"Quirin",
	"Ragnar",
	"Rahel",
	"Raphael",
	"Robin",
	"Romy",
	"Ronja",
	"Rosa",
	"Rosalie",
	"Ruben",
	"Ruby",
	"Runa",
	"Ruth",
	"Sabine",
	"Sabrina",
	"Samuel",
	"Sandra",
	"Sarah",
	"Sebastian",
	"Selina",
	"Silas",
	"Simon",
	"Sophia",
	"Sophie",
	"Stefan",
	"Tanguy",
	"Tanja",
	"Tara",
	"Tarek",
	"Thea",
	"Theo",
	"Theresa",
	"Thomas",
	"Tilda",
	"Tim",
	"Tobias",
	"Tom",
	"Ubbe",
	"Udo",
	"Ulf",
	"Uli",
	"Ulrich",
	"Ulrik",
	"Ulrike",
	"Urban",
	"Ursula",
	"Uta",
	"Ute",
	"Uwe",
	"Valentin",
	"Valentina",
	"Valeria",
	"Valerie",
	"Vanessa",
	"Vera",
	"Verena",
	"Veronika",
	"Victoria",
	"Viktoria",
	"Vincent",
	"Viola",
	"Waldemar",
	"Walter",
	"Wanda",
	"Werner",
	"Wilhelm",
	"Willi",
	"William",
	"Wilma",
	"Wim",
	"Wolfgang",
	"Wolke",
	"Wotan",
	"Xaver",
	"Xavier",
	"Xenia",
	"Xhenis",
	"Yael",
	"Yannick",
	"Yannik",
	"Yannis",
	"Yara",
	"Ylva",
	"Ylvi",
	"Yuki",
	"Yuma",
	"Yuna",
	"Yuri",
	"Yvonne",
	"Zacharias",
	"Zayn",
	"Zelda",
	"Zelko",
	"Zeno",
	"Zenzi",
	"Zita",
	"Zoe",
	"Zoey",
	"Zoro",
}

func randForenames(r *rand.Rand) []string {
	const maxForenames int = 2
	ForenameCount := r.Intn(maxForenames) + 1
	entries := make([]string, ForenameCount)
	duplicateCount := 0
	for i := range entries {
		idx := r.Intn(len(forenames))
		Forename := forenames[idx]
		duplicate := false
		for j := 0; j < i; j++ {
			if entries[j] == Forename {
				duplicate = true
				break
			}
		}
		if duplicate {
			duplicateCount++
			entries[i] = "ZZZZ"
			continue
		}
		entries[i] = Forename
	}
	uniqueCount := ForenameCount - duplicateCount
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

func generateName(r *rand.Rand) Array {
	name := make(map[string]interface{})
	name["family"] = generateFamilyName(r)
	name["given"] = randForenames(r)
	return wrapInArray(name)
}

var streets = []string{
	"Condrada dell'orso, 22",
	"Via G. Bovio 46",
	"Via Borgo Sant'Antonio, 23",
	"Via del Tulipano, 5",
	"Via aleardi 3",
	"Via Padova, 3",
	"Corso Italia, 1",
	"via dei giacinti, 2",
	"Via Torino 2",
	"Piazza Duomo, 8",
	"Via Ariosto 9",
	"Largo del Ginepro,6",
	"via Togliatti, 13",
	"Sancha de Lara, 8",
	"45, Triq Dun Karm",
	"12, Triq Ix-Xemx",
	"205 Hills Road",
	"Wall Street",
}

var towns = []string{
	"Monopoli BA",
	"Taranto TA",
	"Pordenone PN",
	"Forlinpopoli FC",
	"Roma",
	"Padova",
	"Mestrino",
	"Calambrone",
	"Rodano",
	"Milano",
	"Manfredonia FG",
	"New York",
	"Cambridge",
	"Bologna BO",
	"Malaga",
	"Birkirkara",
	"Marsaskala",
}

var countries = []string{
	"Italy",
	"UK",
	"USA",
	"Spain",
	"Malta",
}

func generateAddress(r *rand.Rand) Array {
	address := make(map[string]interface{})
	address["line"] = Array{streets[r.Intn(len(streets))], towns[r.Intn(len(towns))], countries[r.Intn(len(countries))],}
	return wrapInArray(address)
}

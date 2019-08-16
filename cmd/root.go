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

package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/samply/bbmri-fhir-gen/gen"
	"github.com/spf13/cobra"
	"math/rand"
	"os"
	"path/filepath"
)

var start, n, txSize int

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bbmri-fhir-gen [directory]",
	Short: "BBMRI FHIR Test Data Generator",
	Long: `Generates FHIR® Bundles with test data. Currently a fix set of 
FHIR® Patient, Observation and Specimen resources are generated.`,
	Version:   "0.1.0",
	ValidArgs: []string{"directory"},
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires a directory argument")
		}
		return checkDir(args[0])
	},
	Run: func(cmd *cobra.Command, args []string) {
		dir := args[0]

		r := rand.New(rand.NewSource(0))

		for i := 0; i < n/txSize; i++ {
			err := genTxFile(dir, r, start+i*txSize, txSize)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
		if 0 < n%txSize {
			err := genTxFile(dir, r, start+(n/txSize)*txSize, n%txSize)
			if err != nil {
				fmt.Println(err)
				os.Exit(1)
			}
		}
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().IntVarP(&start, "start", "s", 0, "patient index to start with")
	rootCmd.Flags().IntVarP(&n, "num", "n", 100, "number of patients to generate")
	rootCmd.Flags().IntVar(&txSize, "tx-size", 100, "number of patients per transaction")
}

func checkDir(dir string) error {
	if info, err := os.Stat(dir); os.IsNotExist(err) {
		return fmt.Errorf("directory `%s` doesn't exist", dir)
	} else if !info.IsDir() {
		return fmt.Errorf("`%s` isn't a directory", dir)
	} else {
		return nil
	}
}

func genTxFile(dir string, r *rand.Rand, start, n int) error {
	filename := fmt.Sprintf("transaction-%d.json", start)
	f, err := os.OpenFile(filepath.Join(dir, filename), os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}
	e := json.NewEncoder(f)
	e.SetIndent("", "  ")
	err = e.Encode(gen.Bundle(r, start, n))
	if err1 := f.Close(); err == nil {
		err = err1
	}
	return err
}

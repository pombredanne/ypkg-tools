//
// Copyright © 2016 Ikey Doherty
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Fprintf(os.Stderr, "Not yet implemented\n")

	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "Usage: %s [archive URIs]\n", os.Args[0])
		os.Exit(1)
	}

	archives := os.Args[1:]
	fmt.Fprintf(os.Stderr, "Fetching: %s\n", strings.Join(archives, ", "))

	os.Exit(1)
}
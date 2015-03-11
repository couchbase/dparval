//  Copyright (c) 2013 Couchbase, Inc.
//  Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
//  except in compliance with the License. You may obtain a copy of the License at
//    http://www.apache.org/licenses/LICENSE-2.0
//  Unless required by applicable law or agreed to in writing, software distributed under the
//  License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
//  either express or implied. See the License for the specific language governing permissions
//  and limitations under the License.

package main

import (
	"fmt"

	"github.com/couchbase/dparval"
)

func main() {

	// read some JSON
	bytes := []byte(`{"type":["test", "this"]}`)

	// create a Value object
	doc := dparval.NewValueFromBytes(bytes)

	// attempt to access a nested Value
	docType, err := doc.Path("type")
	if err != nil {
		panic("no property type exists")
	}

	// convert docType to a native go value
	docTypeValue := docType.Value()

	// display the value
	fmt.Printf("document type is %v\n", docTypeValue)

	docType, err = docType.Index(1)
	if err != nil {
		panic("no index exists")
	}
	fmt.Printf("index 1 of type is %v\n", docType.Value())
}

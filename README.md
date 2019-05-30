# Go Protobuf Reader

## Motivation
Getting access to the data stored in .proto files can be very handy, 
allowing you to get single source of truth definitions about your services.

Here at SecureNative, we use protobufs extensively to define our micro-services and with the use of [packman](github.com/securenative/packman)
we were able to generate all the  boilerplate parts allowing you to focus on solving the actual problem and not to mess around with infrastructural issues.

## Quick Example

```go
package main

import (
	"fmt"
	"github.com/securenative/GoProtobufReader/external"
	"io/ioutil"
)

func main() {
	// Create an instance of the protobuf reader:
	reader := protobuf.NewReader()
	
	// Read a protobuf file:
	bytes, _ := ioutil.ReadFile("path-to-file.proto")
	
	// Read the file, the result is the `struct` version of the protobuf file
	definition, _ := reader.Read(string(bytes))
	
	fmt.Printf("%v", definition)
}   
``` 

The resulting variable is a typed and easy-to-access go's struct that represents the file we've just read.


## Known Issues:
* Enums aren't supported yet.
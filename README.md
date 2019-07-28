# Go Protobuf Reader

## Motivation
Getting access to the data stored in .proto files can be very handy, 
allowing you to have a single source of your microservices meta-data.

Here at SecureNative, we use protobufs extensively to define our micro-services meta-data, and with the use of [packman](github.com/securenative/packman)
we are able to generate all the boilerplate needed, allowing the developer to focus on solving the actual business problem need to be solved without the need to mess-around and wiring all the parts together.

## Quick Example

```go
package main

import (
	"fmt"
	"github.com/securenative/GoProtobufReader/proto_reader"
	"io/ioutil"
)

func main() {
	// Create an instance of the protobuf reader:
	reader := proto_reader.NewReader()
	
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

// There are two types of embedding in Go
// 1. Embedding interface when defining an interface
// 2. Embedding fields in a struct

// We will look at embedding an field with an interface type inside of a struct

package main

type Speaker interface {
	Speak() string
}

type DemoSpeaker struct {
	// Speaker is embedded in the DemoSpeaker
	Speaker
}



// Embedding is occurring here because we didn't define a filed name for the Speaker field.
// Instead, the field name is defaulted to the same name as the type. That means that right now if we had a DemoSpeaker instance, the following two lines of code would effectively be the same.

demoSpeaker.Speak()
demoSpeker.Speaker.Speak()

// As long as we haven’t defined a Speak method on the DemoSpeaker, the first line will be the same as the second in our code. But more importantly, by embedding the Speaker type, our DemoSpeaker now implements the Speaker interface without us actually defining the Speak method. That means the following code would be valid.

var speaker Speaker = DemoSpeaker{}

// By embedding interfaces in a struct, you can ensure that the struct implements the entire interface even if it doesn't implement any of the methods on its own.


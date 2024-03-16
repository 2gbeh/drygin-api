package faker_util

import (
	"fmt"

	"github.com/bxcodec/faker/v3"
)

// USAGE
/*
	import "github.com/2gbeh/gothic-api/utils/faker_util"

	type FakeType faker_util.FakeUser
	var fakeData = faker_util.GetCollection[FakeType](FakeType{}, 10)

	return response_helper.OK[[]FakeType](context, fakeData)
*/

func GetStructWithTags() {
	a := FakeStructWithTags{}

	err := faker.FakeData(&a)
	
	if err != nil {
		fmt.Println(err)
	}
	
	fmt.Printf("%+v", a)
}

func GetResource[S any](newStruct S) S {
	// newStruct := S{}

	err := faker.FakeData(&newStruct)
	
	if err != nil {
		fmt.Println(err)
	}
	
	return newStruct
}

func GetCollection[S any](newStruct S, size int) []S {
	newSlice := []S{}

	for i := 0; i < size; i++ {
		newSlice = append(newSlice, GetResource(newStruct))
	}
	
	// Forget all generated unique values. 
	// Allows to start generating another unrelated dataset.
	faker.ResetUnique()

	return newSlice
}
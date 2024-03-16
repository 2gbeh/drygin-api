package utils

import (
	// "errors"

	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	zerolog "github.com/rs/zerolog/log"
	// "gorm.io/gorm"
)

func Uuid()string{
	return uuid.New().String()
}

func Snooze(){
	time.Sleep(time.Second * 3)
}

func LogFatal(err error, msg string){
	if err != nil {
	 	str := fmt.Errorf("GoError: %s", msg)
    log.Fatal(str)
  }
}

func LogInfo(str string){
	zerolog.Info().Msg(str)
}

func LoadJson[T any](dir string, model T) (obj T, err error) {
	file, err := os.Open(dir)	
	if err != nil {
		return &model{}, err
	}
	defer file.Close()
	bytes, _ := ioutil.ReadAll(file)

	json.Unmarshal(bytes, &model)
	return model["Data"], err
}

// func ErrorPanic(err error){
// 	if err != nil {
// 		panic(err)
// 	}
// }

// func ErrorsPanicIF(res *gorm.DB, str string){
// 	if res == nil {
// 		panic(errors.New(str))
// 	} 
// }

// func ErrorsPanic(str string){
// 	panic(errors.New(str))
// }



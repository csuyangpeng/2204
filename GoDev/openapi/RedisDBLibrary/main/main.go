package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	//testUri := "UriList"
	////testFilter := "Amf"
	//////mapA := make(map[string]interface{})
	//////mapA["testHref"] = "test2"
	//////RedisDBLibrary.RestfulAPIPutOne(testUri,testFilter,mapA)
	////result := RedisDBLibrary.RestfulAPIGetMany(testUri,testFilter)
	////fmt.Println("the result is:",result)
	//a := fmt.Sprintf("%s%s", testUri, "String")
	//fmt.Println(a)
	//collName := "NfProfile"
	//nfInstanceId := "4947a69a-f61b-4bc1-b9da-47c9c5d14b64"
	//var originNf models.NfProfile
	//nfProfilesRaw :=RedisDBLibrary.RestfulAPIGetMany(collName, nfInstanceId)
	//for _,v := range nfProfilesRaw {
	//	err := json.Unmarshal(v.([]byte), &originNf)
	//	if err != nil {
	//		fmt.Println("error:", err)
	//	}
	//}

	printStr := setsubscriptionId()
	fmt.Println("test data:", printStr)

}

func setsubscriptionId() string {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(100)
	return strconv.Itoa(x)
}

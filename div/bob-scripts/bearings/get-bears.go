package main

import(
    "fmt"
    "http"
    "log"
    "json"
)

func main(){

 	get()

}


func get() {
    fmt.Println("1. Performing Http Get...")
    resp, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
    if err != nil {
        log.Fatalln(err)
    }

    defer resp.Body.Close()
    bodyBytes, _ := ioutil.ReadAll(resp.Body)

    // Convert response body to string
    bodyString := string(bodyBytes)
    fmt.Println("API Response as String:\n" + bodyString)

    // Convert response body to Todo struct
    var todoStruct Todo
    json.Unmarshal(bodyBytes, &todoStruct)
    fmt.Printf("API Response as struct %+v\n", todoStruct)
	
}
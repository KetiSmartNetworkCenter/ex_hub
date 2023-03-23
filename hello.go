// package main

// import (
// 	"encoding/json"
// 	"log"
// 	"math/rand"
// )

// // type Rule struct {
// // 	Id      string                   `json:"id"`
// // 	Sql     string                   `json:"sql"`
// // 	Actions []map[string]interface{} `json:"actions"`
// // 	Options map[string]interface{}   `json:"options"`
// // }

// type simple_Rule struct {
// 	Id  string `json:"id"`
// 	Sql string `json:"sql"`
// }

// const ex_json = `{"id":"1", "sql":"select"}`

// const exampleJson = `{"kind":"Event","apiVersion":"events.k8s.io/v1","metadata":{"name":"swlee-event-test-53009236","namespace":"default"}}`

// const kuiper_Json = `{
//     "id":"rule1111q1",
//     "sql":"SELECT * FROM EdgeXStream WHERE meta(device)=\"Simple-Device01\"",
//     "actions":[
//         {
//             "mqtt":{
//                 "server":"tcp://172.21.0.1:1883",
//                 "topic":"mqtt_1"
//             }
//         },
//         {
//             "mqtt":{
//                 "server":"tcp://172.21.0.1:1883",
//                 "topic":"mqtt_2"
//             }
//         },
//         {
//             "log":{}
//         }
//     ]
// }`

// type Rule struct {
// 	Id      string                   `json:"id"`
// 	Sql     string                   `json:"sql"`
// 	Actions []map[string]interface{} `json:"actions"`
// 	Options map[string]interface{}   `json:"options"`
// }

// func main() {

// 	var result_struct Rule
// 	json.Unmarshal([]byte(kuiper_Json), &result_struct)
// 	log.Println(result_struct)

// 	result_struct_Action := result_struct.Actions
// 	for _, action := range result_struct_Action {
// 		log.Println(action)
// 		if _, ok := action["mqtt"]; ok {
// 			log.Println("mqtt is exist")
// 		}
// 	}

// 	// for i := 0; i < len(result_struct.Actions); i++ {
// 	// 	log.Println("%q", result_struct.Actions[i].(type))
// 	// 	// log.Println(result_struct.Actions[i].(map[string]interface{})["mqtt"])
// 	// }

// 	// var result map[string]interface{}
// 	// json.Unmarshal([]byte(ex_json), &result)

// 	var exampleJson_result map[string]interface{}
// 	json.Unmarshal([]byte(exampleJson), &exampleJson_result)

// 	// log.Println("kind: ", exampleJson_result["kind"])
// 	log.Println("map[metadata] : ", exampleJson_result["metadata"].(map[string]interface{})["name"])
// 	// // ok -> if key is exist return true
// 	// // 	  -> if key is not exist return false
// 	// if _, ok := exampleJson_result["metadata"].(map[string]interface{})["name"]; ok {
// 	// 	log.Println("name is existed")
// 	// }

// 	// log.Println(result["id"])

// 	// var param int8 = 5
// 	// var check int8
// 	// var exArray []int8 = ReturnArrayEx(param)
// 	// for check = 1; check < param; check++ {
// 	// 	exArray[check] = int8(rand.Intn(100))
// 	// 	fmt.Println(exArray[check])
// 	// }
// 	// return

// 	// var rule Rule
// 	// rule.Id = "rule1"
// 	// rule.Sql = " SELECT"
// 	// rule.Actions[0]["mqtt"] = "a"
// 	// rule.Options["option"] = "b"
// 	// log.Println(rule)

// 	// ex := make(map[int]interface{})
// 	// ex[1] = "a"
// 	// // log.Println(ex)
// 	// if A, ok := ex[2]; ok {
// 	// 	log.Println("A", A)
// 	// 	log.Println("ok", ok)
// 	// }
// }

// func ReturnArrayEx(paramint int8) []int8 {
// 	var n int8
// 	var exArray []int8 = make([]int8, paramint)
// 	for n = 0; n < paramint; n++ {
// 		exArray[n] = int8(rand.Intn(100))
// 	}
// 	return exArray
// }

// func ReturnArray(paramArray []int8) []int8 {
// 	var n int = 0
// 	for n < len(paramArray) {
// 		paramArray[n] = int8(rand.Intn(100))
// 	}
// 	return paramArray
// }
//////////////////////////////////////////////////////////////////////////////////////////////////////
package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

https://www.linkedin.com/pulse/closing-go-channel-written-several-goroutines-leo-lara
var UserList []User

type User struct {
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Signal bool
	// Signal string `json:"signal",omitempty`
}
type User_Signal struct {
	Name   string `json:"name"`
	Signal bool   `json:"signal"`
}

// wg := new(sync.WaitGroup)
func main() {
	log.Println("start")

	// var wg sync.WaitGroup
	wg := new(sync.WaitGroup)
	// a := User{Name: "a", Age: 2}
	// b := User{Name: "b", Age: 5}
	// Walk(wg, a.Age)
	// Walk(wg, b.Age)
	// a.Walk(wg)
	// b.Walk(wg)
	// wg.Add(1)
	// go a.Walk(wg)
	// wg.Add(1)
	// go b.Walk(wg)

	// http.Handle("/", r)
	r := mux.NewRouter()
	UserRouter(r, wg)

	header := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE"})
	origins := handlers.AllowedOrigins([]string{"*"})
	wg.Add(1)
	log.Println("web Server Start")
	http.ListenAndServe(":7700", handlers.CORS(header, methods, origins)(r))
	wg.Wait()
	log.Println("fi")

}

func UserRouter(router *mux.Router, wg *sync.WaitGroup) {
	router.HandleFunc("/api/v1/users", Post_User(*wg)).Methods("POST")
	router.HandleFunc("/api/v1/ping", Get_test()).Methods("GET")
	router.HandleFunc("/api/v1/stop", Stop_User(*wg)).Methods("POST")
	// router.HandleFunc("/api/v1/start", Start_User(*wg)).Methods("POST")
}

// func Post_User(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Ping health check")
// }
// func Get_test(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Ping Check Handler Webserver test")
// }
func Post_User(wg sync.WaitGroup) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var n_user User
		json.NewDecoder(r.Body).Decode(&n_user)
		n_user.Signal = true
		UserList = append(UserList, n_user)

		user_list_save_index := len(UserList) - 1
		log.Println("user_list_save_index", user_list_save_index)
		wg.Add(1)
		go UserList[user_list_save_index].Walk(&wg)
		// go n_user.Walk(&wg)

		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
	}
}
func Get_test() http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		log.Println("Ping Check Handler Webserver test")
		rw.Header().Set("Content-Type", "application/json")
		rw.WriteHeader(http.StatusOK)
		json.NewEncoder(rw).Encode(UserList)
	}
}

func Stop_User(wg sync.WaitGroup) http.HandlerFunc {
	return func(rw http.ResponseWriter, r *http.Request) {
		var user_signal User_Signal
		json.NewDecoder(r.Body).Decode(&user_signal)
		for i := 0; i < len(UserList); i++ {
			if user_signal.Name == UserList[i].Name {
				log.Println("UserList Index", i, "User Info", UserList[i])
				UserList[i].Signal = user_signal.Signal
				break
			}

		}
	}
}

// func Start_User(wg sync.WaitGroup) http.HandlerFunc {
// 	return func(rw http.ResponseWriter, r *http.Request) {
// 		var user_signal User_Signal
// 		json.NewDecoder(r.Body).Decode(&user_signal)
// 		for i := 0; i < len(UserList); i++ {
// 			if user_signal.Name == UserList[i].Name {
// 				log.Println("UserList Index", i, "User Info", UserList[i])
// 				UserList[i].Signal = user_signal.Signal
// 				wg.Add(1)
// 				UserList[i].Walk(&wg)
// 				break
// 			}
// 		}
// 	}
// }

func (u *User) Walk(wg *sync.WaitGroup) {
	// wg.Add(1)
	cnt := 0
	defer wg.Done()
	for {
		if !u.Signal {
			log.Println("User Name", u.Name, "stop")
			break
		}
		cnt = cnt + 1

		log.Println("User Name", u.Name, "cnt", cnt, "user_signal:", u.Signal)

		time.Sleep(time.Millisecond * 1000) // if cnt == 3 {
		// 	u.Event(cnt)
		// }
	}
	// go func() {
	// 	for {
	// 		cnt = cnt + 1
	// 		time.Sleep(time.Millisecond * 1000)
	// 		log.Println("User Name", u.Name, "cnt", cnt)
	// 		if cnt == u.Age {
	// 			break
	// 		}
	// 	}
	// }()

}

func (u *User) Event(cnt int) {
	log.Println("Evnet start", " User Name :", u.Name, "Event count", cnt)
}

// func Walk(wg_g *sync.WaitGroup, age int) {
// 	wg_g.Add(1)
// 	cnt := 0
// 	defer wg_g.Done()
// 	go func() {
// 		for {
// 			cnt = cnt + 1
// 			time.Sleep(time.Millisecond * 1000)
// 			log.Println("User Name", "cnt", cnt)

// 			if cnt == age {
// 				break
// 			}
// 		}
// 	}()
// }

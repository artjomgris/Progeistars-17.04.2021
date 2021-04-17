package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"strconv"
	"time"
)

func main() {
	/*hash, err := bcrypt.GenerateFromPassword([]byte("qwerty123"), bcrypt.MinCost)
	if err != nil {
		log.Println(err)
	}
	dbuser.Pass = string(hash)*/

	http.HandleFunc("/login", LoginHandler)
	http.HandleFunc("/data", DataHandler)
	err := http.ListenAndServe("127.0.0.1:5000", nil)
	if err != nil {
		panic(err.Error())
	}

}
type session struct {
	Id int
	Userid int
	Ip string
	Opened time.Time
	Expire time.Time
}
type user struct{
	Login string `json:"login"`
	Pass  string `json:"pass"`
}

var dbuser = user{"mylogin", "qwerty123"}

var sessions []session


func LoginHandler(w http.ResponseWriter, req *http.Request) {

	header := w.Header()
	header.Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Headers", "Content-Type")
	if req.Method == "POST" {
		data, err := io.ReadAll(req.Body)
		req.Body.Close()
		if err != nil {
			panic(err.Error())
		}

		var User user
		err = json.Unmarshal(data, &User)
		fmt.Println(User)
		if err != nil {
			panic(err.Error())
		}
		if User.Login == dbuser.Login {
			/*hash, err := bcrypt.GenerateFromPassword([]byte(User.Pass), bcrypt.MinCost)
			if err != nil {
				panic(err.Error())
			}*/
			if dbuser.Pass == User.Pass {
				ip, _, err := net.SplitHostPort(req.RemoteAddr)
				if err != nil {
					panic(err.Error())
				}
				var sess = session{
					Id:     len(sessions),
					Userid: 0,
					Ip:     ip,
					Opened: time.Now(),
					Expire: time.Now().Add(time.Hour),
				}
				sessions = append(sessions, sess)
				w.Write([]byte(fmt.Sprint(sess.Id)))
			} else {
				w.Write([]byte("Incorrect password!"))
			}
		} else {
			fmt.Println(string(data))
			w.Write([]byte("Incorrect login!"))
		}

	}  else if req.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(405)
	}

}
func DataHandler(w http.ResponseWriter, req *http.Request) {
	header := w.Header()
	header.Add("Access-Control-Allow-Methods", "POST, GET, OPTIONS")
	header.Add("Access-Control-Allow-Origin", "*")
	header.Add("Access-Control-Allow-Headers", "Content-Type, Session-Id")
	if req.Method == "GET" {
		id := req.Header.Get("Session-Id")
		for i, v := range sessions {
			intid, err := strconv.Atoi(id)
			if err != nil {
				panic(err.Error())
			}
			if v.Id == intid {
				ip, _, err := net.SplitHostPort(req.RemoteAddr)
				if err != nil {
					panic(err.Error())
				}
				if ip == v.Ip {
					if time.Now().Before(v.Expire) {
						fmt.Println("GET SUCCSESS!")
						w.Write([]byte("GET request success!"))
						break
					} else {
						sessions = append(sessions[:i], sessions[i+1:]...)
						w.Write([]byte("666;Session expired!"))
						break
					}
				} else {
					w.Write([]byte("666;Session expired!"))
					break
				}
			}
		}
	} else if req.Method == "OPTIONS" {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(405)
	}

}

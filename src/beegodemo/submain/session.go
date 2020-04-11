package main

import (
	"fmt"
	"github.com/astaxie/beego/session"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

var sessionManager *session.Manager

func init() {
	sessionManagerConfig := &session.ManagerConfig{
		CookieName:              "beegoSessionID",
		EnableSetCookie:         true,
		Gclifetime:              3600,
		Maxlifetime:             3600,
		DisableHTTPOnly:         true,
		Secure:                  false,
		CookieLifeTime:          3600,
		ProviderConfig:          "",
		Domain:                  "/custom",
		SessionIDLength:         64,
		EnableSidInHTTPHeader:   false,
		SessionNameInHTTPHeader: "xxxxx",
		EnableSidInURLQuery:     false,
		SessionIDPrefix:         "myPrefix",
	}
	sessionManager, _ = session.NewManager("memory", sessionManagerConfig)
	go sessionManager.GC()
}

func main() {
	router := httprouter.New()
	router.GET("/sessionHandle", sessionHandle)
	server := http.Server{
		Addr:    "localhost:8080",
		Handler: router,
	}
	err := server.ListenAndServe()
	fmt.Println("start server failed, err :", err)
}

func sessionHandle(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	ses, err := sessionManager.SessionStart(w, r)
	if err != nil {
		w.Write([]byte("start session err!"))
		return
	}
	username := ses.Get("username")
	if username == nil || username == "" {
		ses.Set("username", "caozhen")
	} else {
		fmt.Fprint(w, username)
	}
}

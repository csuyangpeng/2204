package sbicmn

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"lite5gc/cmn/rlogger"
	"lite5gc/cmn/types"
	"net/http"
)

type SbiInterface struct {
	serIpAddr string
	serPort   int
	router    *gin.Engine
}

func (p *SbiInterface) Initialize(sbiIpaddr string, sbiport int) {
	p.serIpAddr = sbiIpaddr
	p.serPort = sbiport
	p.router = gin.New()
	p.router.Use(gin.Recovery())
}

func (p *SbiInterface) SetSbiLogger() {
	p.router.Use(LogGin)
}
func (p *SbiInterface) GetRouter() *gin.Engine {
	return p.router
}

func (p *SbiInterface) Start() {
	rlogger.FuncEntry(types.ModCmn, nil)
	addr := fmt.Sprintf("%s:%d", p.serIpAddr, p.serPort)
	h1s := &http.Server{
		Addr:    addr,
		Handler: h2c.NewHandler(p.router, &http2.Server{}),
	}
	fmt.Println("Listening and serving HTTPS on ", addr)
	go h1s.ListenAndServe()
}
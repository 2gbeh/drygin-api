package gin_config

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func Start(routes *gin.Engine) {
	GO_API_PORT := os.Getenv("GO_API_PORT")
  // 	
	server := &http.Server{
		Addr:    fmt.Sprintf(":%s", GO_API_PORT),
		Handler: routes,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

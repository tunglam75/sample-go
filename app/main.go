package main

import (
	"net/http"
	"sample/app/infrastructure"
	"sample/app/router"

	mMiddleware "sample/app/shared/middleware"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	// start internal server
	go startInt()

	// start external server
	startExt()
}

func startExt() {
	// http.HandleFunc("/hello", HelloServer)
	// _ = http.ListenAndServe(":8080", nil)

	// sql new.
	// sqlHandler, _ := infrastructure.NewSQL()
	// s3 new.
	// s3Handler := infrastructure.NewS3()
	// cache new.
	// cacheHandler := infrastructure.NewCache()
	// logger new.
	loggerHandler := infrastructure.NewLogger()
	// translation new.
	// translationHandler := infrastructure.NewTranslation()
	// 3rd search api setup
	// searchAPIHandler := infrastructure.NewSearchAPI()

	// // monitoring setup
	// mLogger := infrastructure.NewLoggerWithType("monitoring")
	// monitoring.Setup(mLogger)

	mux := chi.NewRouter()
	r := &router.Router{
		Mux: mux,
		// SQLHandler: sqlHandler,
		// CacheHandler:       cacheHandler,
		LoggerHandler: loggerHandler,
		// TranslationHandler: translationHandler,
	}

	r.InitializeRouter()
	r.SetupHandler()

	// after process
	defer infrastructure.CloseLogger(r.LoggerHandler.Logfile)
	// defer infrastructure.CloseRedis(r.CacheHandler.Conn)
	// defer infrastructure.CloseLogger(mLogger.Logfile)

	_ = http.ListenAndServe(":8080", mux)
}

func startInt() {
	mux := chi.NewRouter()
	logger := infrastructure.NewLogger()
	mux.Use(mMiddleware.Logger(logger))

	defer infrastructure.CloseLogger(logger.Logfile)

	// profile
	mux.Mount("/debug", middleware.Profiler())
	_ = http.ListenAndServe(":18080", mux)
}

// //HelloServer hello world, the web server
// func HelloServer(w http.ResponseWriter, req *http.Request) {
// 	io.WriteString(w, "hello, world!\n")
// }

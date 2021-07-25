package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
	"github.com/p1ck0/TODOms/pkg/database/mem"
	"github.com/p1ck0/TODOms/pkg/endpoints"
	"github.com/p1ck0/TODOms/pkg/grpctransport"
	httptransport "github.com/p1ck0/TODOms/pkg/http"
	"github.com/p1ck0/TODOms/pkg/pb"
	"github.com/p1ck0/TODOms/pkg/repository"
	"github.com/p1ck0/TODOms/pkg/service"
	"google.golang.org/grpc"
)

func main() {
	httpAddr := flag.String("http.addr", ":8080", "HTTP listen address")
	flag.Parse()

	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewSyncLogger(logger)
		logger = level.NewFilter(logger, level.AllowDebug())
		logger = log.With(logger,
			"svc", "order",
			"ts", log.DefaultTimestampUTC,
			"caller", log.DefaultCaller,
		)
	}
	level.Info(logger).Log("msg", "service started")
	defer level.Info(logger).Log("msg", "service ended")

	db, err := mem.NewMemDB()
	if err != nil {
		os.Exit(-1)
	}

	rep := repository.NewRepo(db)
	serv := service.NewServ(*rep, logger)
	endpoints := endpoints.MakeEndpoints(*serv)
	grpcServer := grpctransport.NewGRPCServer(endpoints, logger)

	var h http.Handler
	h = httptransport.MakeHTTPHandler(*serv, log.With(logger, "component", "HTTP"))

	errs := make(chan error)
	go func() {
		c := make(chan os.Signal)
		signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
		errs <- fmt.Errorf("%s", <-c)
	}()

	grpcListener, err := net.Listen("tcp", ":8081")
	if err != nil {
		logger.Log("during", "Listen", "err", err)
		os.Exit(-1)
	}
	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		level.Info(logger).Log("transport", "HTTP", "addr", *httpAddr)
		server := &http.Server{
			Addr:    *httpAddr,
			Handler: h,
		}
		errs <- server.ListenAndServe()
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		server := grpc.NewServer()
		pb.RegisterTODOServiceServer(server, grpcServer)
		level.Info(logger).Log("msg", "Server started successfully")
		server.Serve(grpcListener)
	}()

	wg.Wait()
	level.Error(logger).Log("exit", <-errs)
}

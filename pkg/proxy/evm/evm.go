package evm

import (
	"context"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/INFURA/infrakit/pkg/proxy/server"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/go-logr/logr"
)

const (
	EndpointRefreshInterval = time.Second * 60
)

type Proxy struct {
	server        *server.Server
	log           logr.Logger
	ethclient     *ethclient.Client
	endpoints     []string
	endpointsLock sync.RWMutex
}

type Opts struct {
	Addr                     string
	EthBootstrapEndpoint     string
	Log                      logr.Logger
	NodeRegistryContract     string
	CustomerRegistryContract string
}

func NewProxy(opts Opts) (*Proxy, error) {
	serv := server.NewServer(server.Opts{
		Addr: opts.Addr,
	})

	ethclient, err := ethclient.Dial(opts.EthBootstrapEndpoint)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to ethereum: %w", err)
	}

	proxy := &Proxy{
		server:    serv,
		log:       opts.Log,
		ethclient: ethclient,
		endpoints: []string{},
	}

	serv.Router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		blockNumber, err := ethclient.BlockNumber(context.TODO())
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		fmt.Printf("Endpoints: %+v\n", proxy.endpoints)

		response := fmt.Sprintf("Hello infrakit. Current block number is %d", blockNumber)
		_, _ = w.Write([]byte(response))
	})

	serv.Router.Post("/rpc", func(w http.ResponseWriter, r *http.Request) {
		proxy.endpointsLock.RLock()

		// Pick a random endpoint
		endpoint := proxy.endpoints[rand.Intn(len(proxy.endpoints))]
		opts.Log.Info("selected endpoint", "endpoint", endpoint)

		proxy.endpointsLock.RUnlock()

		// Forward request to the endpoint
		resp, err := http.Post(endpoint, "application/json", r.Body)
		if err != nil {
			opts.Log.Error(err, "failed to forward request to endpoint")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		// Copy response from endpoint to the original request
		_, err = io.Copy(w, resp.Body)
		if err != nil {
			opts.Log.Error(err, "failed to copy response from endpoint")
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	})

	go func() {
		for {
			endpoints, err := getRegisteredNodes(ethclient, opts.NodeRegistryContract)
			if err != nil {
				opts.Log.Error(err, "failed to get registered nodes")
			}

			opts.Log.Info("got registered nodes", "endpoints", endpoints)

			proxy.endpointsLock.Lock()
			proxy.endpoints = endpoints
			proxy.endpointsLock.Unlock()

			time.Sleep(EndpointRefreshInterval)
		}
	}()

	return proxy, nil
}

func (p *Proxy) Start() error {
	err := p.server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("evm proxy Start returned error: %w", err)
	}

	p.log.Info("evm proxy server closed successfully")

	return nil
}

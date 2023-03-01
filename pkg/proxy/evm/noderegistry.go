package evm

import (
	"fmt"

	"github.com/INFURA/infrakit/pkg/contracts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func getRegisteredNodes(ethclient *ethclient.Client, nodeRegistryContract string) ([]string, error) {
	fmt.Printf("Contract: %+v\n", nodeRegistryContract)
	address := common.HexToAddress(nodeRegistryContract)
	nodeRegistry, _ := contracts.NewContract(address, ethclient)

	registeredIter, err := nodeRegistry.FilterRegistered(nil)
	if err != nil {
		return nil, fmt.Errorf("Error getting registered nodes: %v", err)
	}

	endpoints := []string{}

	for registeredIter.Next() {
		providerAddress := registeredIter.Event.Provider

		registeredDetails, err := nodeRegistry.Registered(nil, providerAddress)
		if err != nil {
			return nil, fmt.Errorf("Error getting registered node details: %v", err)
		}

		endpoints = append(endpoints, registeredDetails.Rpc)
	}

	return endpoints, nil
}

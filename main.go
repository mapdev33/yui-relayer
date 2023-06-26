package main

import (
	"log"

	corda "github.com/mapdev33/yui-relayer/chains/corda/module"
	ethereum "github.com/mapdev33/yui-relayer/chains/ethereum/module"
	fabric "github.com/mapdev33/yui-relayer/chains/fabric/module"
	tendermint "github.com/mapdev33/yui-relayer/chains/tendermint/module"
	"github.com/mapdev33/yui-relayer/cmd"
	mock "github.com/mapdev33/yui-relayer/provers/mock/module"
)

func main() {
	if err := cmd.Execute(
		tendermint.Module{},
		fabric.Module{},
		corda.Module{},
		ethereum.Module{},
		mock.Module{},
	); err != nil {
		log.Fatal(err)
	}
}

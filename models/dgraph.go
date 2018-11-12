package models

import (
	"context"

	"github.com/astaxie/beego"

	//	"errors"
	"flag"
	"fmt"

	"github.com/astaxie/beego/logs"

	"github.com/dgraph-io/dgo"
	"github.com/dgraph-io/dgo/protos/api"
	"google.golang.org/grpc"
)

var (
	Dg     *dgo.Dgraph
	dgraph = flag.String("d", beego.AppConfig.String("Dgraph::addr")+":"+beego.AppConfig.String("Dgraph::port"), "Dgraph server address")
)

func init() {
	var err error
	if Dg, err = Conn(); err != nil {
		logs.Error("dgraph connect filed")
	}
}

func Conn() (*dgo.Dgraph, error) {
	flag.Parse()
	conn, err := grpc.Dial(*dgraph, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}
	//defer conn.Close()

	Dg := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	return Dg, nil

}

func DgraphClient() {

	resp, err := Dg.NewTxn().Query(context.Background(), `{
  me(func: has(starring)) {
    name
  }
}`)

	if err != nil {
		logs.Error(err)
	}
	fmt.Printf("Response: %s\n", resp.Json)

}

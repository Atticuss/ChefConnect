package main

import (
	"log"

	"github.com/atticuss/chefconnect/repositories/dgraph"
	"github.com/dgraph-io/dgo/v2"
	"github.com/dgraph-io/dgo/v2/protos/api"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("ec2-34-238-150-16.compute-1.amazonaws.com:9080", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}

	client := dgo.NewDgraphClient(api.NewDgraphClient(conn))
	utilRepo := dgraph.NewDgraphRepositoryUtility(client)

	utilRepo.ClearDatastore()
	utilRepo.InitializeSchema()
	utilRepo.InitializeBaseData()
}

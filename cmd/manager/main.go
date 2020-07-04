package main

import (
	"github.com/atticuss/chefconnect/repositories/dgraph"
)

func main() {
	dgraphConfig := dgraph.Config{
		Host: "ec2-34-238-150-16.compute-1.amazonaws.com:9080",
	}
	utilRepo := dgraph.NewDgraphRepositoryUtility(&dgraphConfig)

	utilRepo.ClearDatastore()
	utilRepo.InitializeSchema()
	utilRepo.InitializeBaseData()
}

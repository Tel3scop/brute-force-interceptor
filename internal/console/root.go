package console

import (
	"github.com/Tel3scop/brute-force-interceptor/internal/connector"
	"github.com/Tel3scop/brute-force-interceptor/internal/connector/access"
	"github.com/Tel3scop/helpers/logger"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

var (
	conn          *grpc.ClientConn
	accessClient  *access.Client
	err           error
	serverAddress string
	rootCmd       = &cobra.Command{
		Use:   "anti-bruteforce-cli",
		Short: "CLI tool for managing AntiBruteforce service",
		PersistentPreRun: func(_ *cobra.Command, _ []string) {
			conn = connector.GetGRPCConnection(serverAddress)
			accessClient, err = access.New(conn)
			if err != nil {
				logger.Fatal(err.Error())
			}
		},
	}
)

func init() {
	rootCmd.PersistentFlags().StringVarP(
		&serverAddress,
		"connection",
		"c",
		"127.0.0.1:50051",
		"gRPC connection string",
	)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}

package cmd

import (
        "fmt"
	"log"

        "github.com/spf13/cobra"

	exporter "github.com/Mr-K-Validator/Validator-info/exporter"

)

var (

	chainList []string

	// command로 안 받을 경우 defult 값 지정
	chain string = ""
        restAddr string = "localhost:1317"
	listenPort string = "26661"

	operAddr string= ""
)

// versionCmd represents the version command
var runCmd = &cobra.Command{
        Use:   "run",
        Short: "Validator Operator Address",
        Long: `Be sure to enter either Validator Operator Address
ex#1_Local REST Server) Validator-info run --chain rizon --oper-addr rizonvaloper1tqv36eh27pnkta5pmlzrzmcp4zc09qcz7x086z
ex#2_Remote REST Server) Validator-info run --chain rizon --oper-addr rizonvaloper1tqv36eh27pnkta5pmlzrzmcp4zc09qcz7x086z --rest-server 192.168.0.10:1317 
`,
        Run: func(cmd *cobra.Command, args []string) {
		check_chain()
		set_config()
		run()
        },
}

func init() {
        rootCmd.AddCommand(runCmd)

	runCmd.Flags().StringVarP(&chain, "chain", "c", "", "Chain name of the monitoring node")
        runCmd.MarkFlagRequired("chain")

	runCmd.Flags().StringVarP(&restAddr, "rest-server", "", "localhost:1317", "<host>:<port> to Rest-Server(LCD-Server) interface for the selected chain")
	runCmd.Flags().StringVarP(&listenPort, "port", "p", "26661", "Port to listen for Prometheus collector connections")
	runCmd.Flags().StringVarP(&operAddr, "oper-addr", "", "", "Operator address for Validator")
}

func check_chain() {

	// chain check
	chainCheck := false
	for _, c := range chainList {
		if chain == c {
			chainCheck = true
		}
	}
	if !chainCheck {
		log.Fatal(fmt.Sprintf("[Error] %s is not supported", chain) +fmt.Sprint("\nList of supported chains: ", chainList))
	}


	if (operAddr == "") {
		log.Fatal("[Error] --oper-addr was not entered.")
	}
}

func run() {
	switch chain {
		default:
			exporter.Go(chain, listenPort)

        }
}

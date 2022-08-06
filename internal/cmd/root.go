package cmd

import (
	"fmt"
	"os"

	"github.com/burakkaraceylan/xapi-go/pkg/client"
	"github.com/spf13/cobra"
)

var (
	endpoint string
	username string
	password string
	auth     string
	version  string
	rootCmd  = &cobra.Command{
		Use:   "xapi-go",
		Short: "xapi-go is a xApi client written in Go",
		Long:  `A Fast xApi client implementation written in Go by Burak Karaceylan.`,
		Run: func(cmd *cobra.Command, args []string) {
			// Do Stuff Here
		},
	}
	getStatement = &cobra.Command{
		Use:  "getStatement [OPTIONS]",
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			var lrs *client.RemoteLRS
			var err error

			if len(username) > 0 {

				if len(password) == 0 {
					panic("You have to provide both username and password")
				}

				lrs, err = client.NewRemoteLRS(endpoint, version, username, password)

			} else if len(auth) > 0 {
				lrs, err = client.NewRemoteLRS(endpoint, version, auth)
			} else {
				panic("You have to provide either a username/password pair or an authorization header")
			}

			if err != nil {
				panic(err)
			}

			stmt, _, err := lrs.GetStatement(args[0])

			if err != nil {
				panic(err)
			}

			str, _ := stmt.ToJson(true)
			fmt.Println(str)

		},
	}
	about = &cobra.Command{
		Use: "about",
		Run: func(cmd *cobra.Command, args []string) {
			var lrs *client.RemoteLRS
			var err error

			if len(username) > 0 {

				if len(password) == 0 {
					panic("You have to provide both username and password")
				}

				lrs, err = client.NewRemoteLRS(endpoint, version, username, password)

			} else if len(auth) > 0 {
				lrs, err = client.NewRemoteLRS(endpoint, version, auth)
			} else {
				panic("You have to provide either a username/password pair or an authorization header")
			}

			if err != nil {
				panic(err)
			}

			stmt, err := lrs.About()

			if err != nil {
				panic(err)
			}

			str, _ := stmt.ToJson(true)
			fmt.Println(str)

		},
	}
)

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&endpoint, "endpoint", "", "URL of the API endpoint")
	rootCmd.MarkPersistentFlagRequired("endpoint")
	rootCmd.PersistentFlags().StringVar(&version, "version", "", "API version")
	rootCmd.MarkPersistentFlagRequired("version")
	rootCmd.PersistentFlags().StringVar(&username, "username", "", "API user's username")
	rootCmd.PersistentFlags().StringVar(&password, "password", "", "API user's password")
	rootCmd.MarkFlagsRequiredTogether("username", "password")
	rootCmd.PersistentFlags().StringVar(&auth, "auth", "", "Authentication header (Basic, Bearer etc...)")
	rootCmd.MarkFlagsMutuallyExclusive("username", "auth")

	rootCmd.AddCommand(getStatement)
	rootCmd.AddCommand(about)
}

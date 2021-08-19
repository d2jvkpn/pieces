package cmd

import (
	"log"

	"github.com/d2jvkpn/pieces/pkg/go/misc"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func NewServerCmd(name string) (command *cobra.Command) {
	var (
		addr string
		fSet *pflag.FlagSet
	)

	command = &cobra.Command{
		Use:   name,
		Short: `ntp server`,
		Long:  `network time protocol server`,

		Run: func(cmd *cobra.Command, args []string) {
			if addr == "" {
				log.Fatalln("not server address provided")
			}

			ser, err := misc.NewNetworkTimeServer(addr, 10)
			if err != nil {
				log.Fatalln(err)
			}

			log.Printf(">>> Network Time Server listening on: %q\n", addr)
			log.Fatalln(ser.Run())
		},
	}

	fSet = command.Flags()
	fSet.StringVar(&addr, "addr", "", "request addres")

	return command
}

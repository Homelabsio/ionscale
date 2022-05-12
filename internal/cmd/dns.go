package cmd

import (
	"context"
	"fmt"
	"github.com/jsiebens/ionscale/pkg/gen/api"
	"github.com/muesli/coral"
	"strings"
)

func getDNSConfig() *coral.Command {
	command := &coral.Command{
		Use:          "get-dns",
		Short:        "Get DNS configuration",
		SilenceUsage: true,
	}

	var tailnetID uint64
	var tailnetName string
	var target = Target{}

	target.prepareCommand(command)
	command.Flags().StringVar(&tailnetName, "tailnet", "", "")
	command.Flags().Uint64Var(&tailnetID, "tailnet-id", 0, "")

	command.RunE = func(command *coral.Command, args []string) error {
		client, c, err := target.createGRPCClient()
		if err != nil {
			return err
		}
		defer safeClose(c)

		tailnet, err := findTailnet(client, tailnetName, tailnetID)
		if err != nil {
			return err
		}

		req := api.GetDNSConfigRequest{TailnetId: tailnet.Id}
		resp, err := client.GetDNSConfig(context.Background(), &req)

		if err != nil {
			return err
		}
		config := resp.Config

		var allNameservers = config.Nameservers

		for i, j := range config.Routes {
			for _, n := range j.Routes {
				allNameservers = append(allNameservers, fmt.Sprintf("%s:%s", i, n))
			}
		}

		fmt.Printf("%-*v%v\n", 25, "Magic DNS Enabled:", config.MagicDns)
		fmt.Printf("%-*v%v\n", 25, "Override Local DNS:", config.OverrideLocalDns)
		fmt.Printf("%-*v%v\n", 25, "Nameservers:", strings.Join(allNameservers, ","))

		return nil
	}

	return command
}

func setDNSConfig() *coral.Command {
	command := &coral.Command{
		Use:          "set-dns",
		Short:        "Set DNS config",
		SilenceUsage: true,
	}

	var nameservers []string
	var magicDNS bool
	var overrideLocalDNS bool
	var tailnetID uint64
	var tailnetName string
	var target = Target{}

	target.prepareCommand(command)
	command.Flags().StringVar(&tailnetName, "tailnet", "", "")
	command.Flags().Uint64Var(&tailnetID, "tailnet-id", 0, "")
	command.Flags().StringSliceVarP(&nameservers, "nameserver", "", []string{}, "")
	command.Flags().BoolVarP(&magicDNS, "magic-dns", "", false, "")
	command.Flags().BoolVarP(&overrideLocalDNS, "override-local-dns", "", false, "")

	command.RunE = func(command *coral.Command, args []string) error {
		client, c, err := target.createGRPCClient()
		if err != nil {
			return err
		}
		defer safeClose(c)

		tailnet, err := findTailnet(client, tailnetName, tailnetID)
		if err != nil {
			return err
		}

		var globalNameservers []string
		var routes = make(map[string]*api.Routes)

		for _, n := range nameservers {
			split := strings.Split(n, ":")
			if len(split) == 2 {
				r, ok := routes[split[0]]
				if ok {
					r.Routes = append(r.Routes, split[1])
				} else {
					routes[split[0]] = &api.Routes{Routes: []string{split[1]}}
				}
			} else {
				globalNameservers = append(globalNameservers, n)
			}
		}

		req := api.SetDNSConfigRequest{
			TailnetId: tailnet.Id,
			Config: &api.DNSConfig{
				MagicDns:         magicDNS,
				OverrideLocalDns: overrideLocalDNS,
				Nameservers:      nameservers,
				Routes:           routes,
			},
		}
		resp, err := client.SetDNSConfig(context.Background(), &req)

		if err != nil {
			return err
		}

		config := resp.Config

		var allNameservers = config.Nameservers

		for i, j := range config.Routes {
			for _, n := range j.Routes {
				allNameservers = append(allNameservers, fmt.Sprintf("%s:%s", i, n))
			}
		}

		fmt.Printf("%-*v%v\n", 25, "Magic DNS Enabled:", config.MagicDns)
		fmt.Printf("%-*v%v\n", 25, "Override Local DNS:", config.OverrideLocalDns)
		fmt.Printf("%-*v%v\n", 25, "Nameservers:", strings.Join(allNameservers, ","))

		return nil
	}

	return command
}

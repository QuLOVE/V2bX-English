package cmd

import (
	"fmt"

	"github.com/QuLOVE/V2bX-English/common/systime"
	"github.com/beevik/ntp"
	"github.com/spf13/cobra"
)

var ntpServer string

var commandSyncTime = &cobra.Command{
	Use:   "synctime",
	Short: "Sync time from ntp server",
	Args:  cobra.NoArgs,
	Run:   synctimeHandle,
}

func init() {
	commandSyncTime.Flags().StringVar(&ntpServer, "server", "time.apple.com", "ntp server")
	command.AddCommand(commandSyncTime)
}

func synctimeHandle(_ *cobra.Command, _ []string) {
	t, err := ntp.Time(ntpServer)
	if err != nil {
		fmt.Println(Err("get time from server error: ", err))
		fmt.Println(Err("Failure to sync time"))
		return
	}
	err = systime.SetSystemTime(t)
	if err != nil {
		fmt.Println(Err("set system time error: ", err))
		fmt.Println(Err("Failure to sync time"))
		return
	}
	fmt.Println(Ok("Sync time successfully"))
}

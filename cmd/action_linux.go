package cmd

import (
	"fmt"
	"time"

	"github.com/QuLOVE/V2bX-English/common/exec"
	"github.com/spf13/cobra"
)

var (
	startCommand = cobra.Command{
		Use:   "start",
		Short: "Start V2bX service",
		Run:   startHandle,
	}
	stopCommand = cobra.Command{
		Use:   "stop",
		Short: "Stop V2bX service",
		Run:   stopHandle,
	}
	restartCommand = cobra.Command{
		Use:   "restart",
		Short: "Restart V2bX service",
		Run:   restartHandle,
	}
	logCommand = cobra.Command{
		Use:   "log",
		Short: "Output V2bX log",
		Run: func(_ *cobra.Command, _ []string) {
			exec.RunCommandStd("journalctl", "-u", "V2bX.service", "-e", "--no-pager", "-f")
		},
	}
)

func init() {
	command.AddCommand(&startCommand)
	command.AddCommand(&stopCommand)
	command.AddCommand(&restartCommand)
	command.AddCommand(&logCommand)
}

func startHandle(_ *cobra.Command, _ []string) {
	r, err := checkRunning()
	if err != nil {
		fmt.Println(Err("check status error: ", err))
		fmt.Println(Err("Failed to start V2bX"))
		return
	}
	if r {
		fmt.Println(Ok("V2bX is already running, no need to start it again. If you need to restart, please select restart."))
	}
	_, err = exec.RunCommandByShell("systemctl start V2bX.service")
	if err != nil {
		fmt.Println(Err("exec start cmd error: ", err))
		fmt.Println(Err("Failed to start V2bX"))
		return
	}
	time.Sleep(time.Second * 3)
	r, err = checkRunning()
	if err != nil {
		fmt.Println(Err("check status error: ", err))
		fmt.Println(Err("Failed to start V2bX"))
	}
	if !r {
		fmt.Println(Err("V2bX may fail to start, please use V2bX log to view the log information later"))
		return
	}
	fmt.Println(Ok("V2bX started successfully, please use V2bX log to view the running log"))
}

func stopHandle(_ *cobra.Command, _ []string) {
	_, err := exec.RunCommandByShell("systemctl stop V2bX.service")
	if err != nil {
		fmt.Println(Err("exec stop cmd error: ", err))
		fmt.Println(Err("Failed to stop V2bX"))
		return
	}
	time.Sleep(2 * time.Second)
	r, err := checkRunning()
	if err != nil {
		fmt.Println(Err("check status error:", err))
		fmt.Println(Err("Failed to stop V2bX"))
		return
	}
	if r {
		fmt.Println(Err("Failed to stop V2bX, it may be because the stop time exceeds two seconds, please check the log information later"))
		return
	}
	fmt.Println(Ok("V2bX stopped successfully"))
}

func restartHandle(_ *cobra.Command, _ []string) {
	_, err := exec.RunCommandByShell("systemctl restart V2bX.service")
	if err != nil {
		fmt.Println(Err("exec restart cmd error: ", err))
		fmt.Println(Err("Failed to restart V2bX"))
		return
	}
	r, err := checkRunning()
	if err != nil {
		fmt.Println(Err("check status error: ", err))
		fmt.Println(Err("Failed to restart V2bX"))
		return
	}
	if !r {
		fmt.Println(Err("V2bX may fail to start, please use V2bX log to view the log information later"))
		return
	}
	fmt.Println(Ok("V2bX restarted successfully"))
}

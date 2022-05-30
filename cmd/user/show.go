package user

import (
	"os"

	"github.com/gocarina/gocsv"
	"github.com/spf13/cobra"
	"github.com/tkeel-io/cli/fmtutil"
	"github.com/tkeel-io/cli/pkg/kubernetes"
	"github.com/tkeel-io/cli/pkg/print"
)

var UserInfoCmd = &cobra.Command{
	Use:   "show",
	Short: "Show user info.",
	Example: `
# Show user info by user id
tkeel user show <user-id> -t <tenant-id>

`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			print.WarningStatusEvent(os.Stdout, "Please specify the user id")
			print.WarningStatusEvent(os.Stdout, "For example, tkeel user show <user-id> -t <tenant-id>")
			os.Exit(1)
		}
		userID := args[0]
		data, err := kubernetes.TenantUserInfo(tenant, userID)
		if err != nil {
			print.FailureStatusEvent(os.Stdout, err.Error())
			os.Exit(1)
		}
		table, err := gocsv.MarshalString(data)
		if err != nil {
			print.FailureStatusEvent(os.Stdout, err.Error())
			os.Exit(1)
		}
		fmtutil.PrintTable(table)
	},
}

func init() {
	UserInfoCmd.Flags().BoolP("help", "h", false, "Print this help message")
	UserInfoCmd.Flags().StringVarP(&tenant, "tenant", "t", "", "Tenant ID")
	UserInfoCmd.MarkFlagRequired("tenant")
	UserCmd.AddCommand(UserInfoCmd)
}

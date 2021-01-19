package dns

import (
	"fmt"
	"github.com/MakeNowJust/heredoc"
	"github.com/antihax/optional"
	"github.com/masihyeganeh/ar-cli/pkg/api"
	"github.com/masihyeganeh/ar-cli/pkg/api/models"
	"github.com/masihyeganeh/ar-cli/pkg/utl"
	"github.com/spf13/cobra"
	"io"
)

// NewCmdDnsCreateA returns new cobra command to create DNS A record
func NewCmdDnsCreateA(in io.Reader, out, errout io.Writer) *cobra.Command {
	var ipFlag string
	var portFlag int
	var weightFlag int
	var countryFlag string
	var nameFlag string
	var ttlFlag int
	var cloudFlag bool
	var upstreamHttps string
	// Main command
	cmd := &cobra.Command{
		Use:   "a domain",
		Short: "create A record DNS for domain name (Example: example.com)",
		// TODO
		Long: heredoc.Doc(`
    Log in to Arvan API and save login for subsequent use
    First-time users of the client should run this command to connect to a Arvan API,
    establish an authenticated session, and save connection to the configuration file.`),
		ValidArgs: []string{"domain"},
		Args:      cobra.MinimumNArgs(1),
		Run: func(c *cobra.Command, args []string) {
			explainOut := utl.NewResponsiveWriter(out)
			c.SetOutput(explainOut)

			domain := args[0]

			record := models.DnsRecord{
				Type:  "a",
				Cloud: cloudFlag,
			}

			if len(nameFlag) > 0 {
				record.Name = nameFlag
			}
			if len(ipFlag) > 0 {
				aRecord := models.ARecord{
					Ip: ipFlag,
				}
				if portFlag > 0 && portFlag < 65536 {
					aRecord.Port = int32(portFlag)
				}
				if weightFlag > 0 {
					aRecord.Weight = int32(weightFlag)
				}
				if len(countryFlag) > 0 {
					aRecord.Country = countryFlag
				}

				record.Value = &models.OneOfDnsRecordValue{ARecord: aRecord}
			}

			if ttlFlag == 120 || ttlFlag == 180 || ttlFlag == 300 || ttlFlag == 600 || ttlFlag == 900 || ttlFlag == 1800 || ttlFlag == 3600 || ttlFlag == 7200 || ttlFlag == 18000 || ttlFlag == 43200 || ttlFlag == 86400 || ttlFlag == 172800 || ttlFlag == 432000 {
				record.Ttl = int32(ttlFlag)
			}
			if upstreamHttps == "default" || upstreamHttps == "auto" || upstreamHttps == "http" || upstreamHttps == "https" {
				record.UpstreamHttps = upstreamHttps
			}
			// TODO:
			//record.IpFilterMode = &models.DnsRecordIpFilterMode{
			//	Count:     "",
			//	Order:     "",
			//	GeoFilter: "",
			//}

			options := &api.DNSApiDnsRecordsCreateOpts{
				Body: optional.NewInterface(record),
			}
			res, _, err := api.GetAPIClient().DNSApi.DnsRecordsCreate(c.Context(), domain, options)
			utl.CheckApiErr(err)

			fmt.Fprintf(explainOut, "%s\n", res.Message)
		},
	}

	cmd.Flags().StringVarP(&ipFlag, "ip", "i", "", "")
	cmd.Flags().IntVarP(&portFlag, "port", "p", 0, "")
	cmd.Flags().IntVarP(&weightFlag, "weight", "w", 0, "")
	cmd.Flags().StringVarP(&countryFlag, "country", "c", "", "")
	cmd.Flags().StringVarP(&nameFlag, "name", "n", "", "<= 250 characters")
	cmd.Flags().IntVarP(&ttlFlag, "ttl", "t", 0, "120 or 180 or 300 or 600 or 900 or 1800 or 3600 or 7200 or 18000 or 43200 or 86400 or 172800 or 432000")
	cmd.Flags().BoolVarP(&cloudFlag, "cloud", "l", false, "")
	cmd.Flags().StringVarP(&upstreamHttps, "upstream_https", "u", "default", `"default" or "auto" or "http" or "https"`)

	return cmd
}

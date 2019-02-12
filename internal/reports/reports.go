package reports

import (
	"os"
	//"fmt"
	"net/http"
	"time"
	"net/url"
	"strconv"
	"io/ioutil"
	"encoding/json"
	"github.com/urfave/cli"
	"github.com/briandowns/spinner"
	"github.com/olekukonko/tablewriter"
)

type output struct{
	Data	[]result_data	`json:"data"`
	Status	string		`json:"status"`
	TotRec	int64		`json:"totalRecords"`
}

type result_data struct{
	TransId		int64	`json:"trid"`
	Rcptemail	string	`json:"rcptEmail"`
	Fromid		string	`json:"fromaddress"`
	ReqTime		string	`json:"requestedTime"`
	DelTime		string	`json:"deliveryTime"`
	ModTime		string	`json:"modifiedTime"`
	Status		string	`json:"status"`
	Size		int64	`json:"size"`
	Remarks		string	`json:"remarks"`
	Subject		string	`json:"subject"`
	Xheader		string	`json:"xapiheader"`
	Tags		string	`json:"tags"`
	Opens		[]result_open	`json:"opens"`
	Clicks		[]result_click	`json:"clicks"`
}

type result_open struct{
	IP	string	`json:"ip"`
	Time	string	`json:"time"`
}

type result_click struct{
	IP      string  `json:"ip"`
	Link	string	`json:"link"`
	Time    string  `json:"time"`
}

var (
	reportingflags []cli.Flag
)

func init(){
	reportingflags = []cli.Flag{
		cli.StringFlag{ //apikey
			Name:  "secretkey, K",
			Usage: "Pass api-Key here. NOTE :: we highly recommend to set your apikey as enviroment variable for security reason",
		},
		cli.StringFlag{ //fromid
			Name:  "events, e",
			Usage: "Filter based on different email events ie.processed,open,click,unsubscribe,bounce,softbounce,spam,invalid,dropped,hardbounce",
		},cli.StringFlag{ //fromid
			Name:  "sort, srt",
			Value: "asc",
			Usage: "Sort based on email sent time ie.asc,desc",
		},cli.StringFlag{ //fromid
			Name:  "startdate, stdate",
			Usage: "Starting date of statistics to retrive data. Format YYYY-MM-DD",
		},cli.StringFlag{ //fromid
			Name:  "enddate, edate",
			Usage: "End date of statistics to retrive data by default today date. Format YYYY-MM-DD",
		},cli.StringFlag{ //fromid
			Name:  "subject, s",
			Usage: "Filter logs based on subject",
		},cli.StringFlag{ //fromid
			Name:  "limit, l",
			Usage: "Number of results to return",
		},cli.StringFlag{ //fromid
			Name:  "xheader, x",
			Usage: "Filter logs based on x-apiheader",
		},cli.StringFlag{ //fromid
			Name:  "email, m",
			Usage: "Filter logs based on email",
		},cli.StringFlag{ //fromid
			Name:  "fromaddress, fa",
			Usage: "Filter logs based on fromaddress",
		},cli.StringFlag{ //fromid
			Name:  "offset, o",
			Usage: "Point in the list to begin retrieving data",
		},
	}
}

func Getflags() []cli.Flag{
	return reportingflags
}

func Fetchlogs(a *cli.Context) string{
	if a.Args().Present() {
		spin := spinner.New(spinner.CharSets[9], 100*time.Millisecond)
		spin.Start()
		ApiKey := os.Getenv("PEPIKEY")
		if len(a.String("K")) != 0 {
			ApiKey = a.String("K")
		}

		api_events := a.String("e")
		sort_ordr := a.String("srt")
		st_date := time.Now().AddDate(0, 0, -90).Format("2006-01-02")
		if len(a.String("stdate")) != 0{
			st_date = a.String("stdate")
		}
		ed_date := a.String("edate")
		sub := a.String("s")
		email := a.String("m")
		xheader := a.String("x")
		offset := a.String("o")
		limit := a.String("l")
		fromadd := a.String("fa")



		URL := "https://api.pepipost.com/v2/logs?sort=" + sort_ordr + "&startdate=" + st_date + "&enddate=" + ed_date + "&limit=" + limit + "&email=" + email + "&offset=" + offset + "&events=" + api_events + "&xapiheader=" + xheader + "&subject" + sub + "&fromaddress" + fromadd

		req, _ := http.NewRequest("GET", URL, nil)
		req.Header.Add("api_key",ApiKey )
		res, _ := http.DefaultClient.Do(req)
		defer res.Body.Close()
		body, _ := ioutil.ReadAll(res.Body)

		arr := output{}
		json.Unmarshal([]byte(body), &arr)
		tbl_data := [][]string{}

		for i := range arr.Data{
			mytid := strconv.FormatInt(arr.Data[i].TransId,10)
			tidsize := strconv.FormatInt(arr.Data[i].Size,10)


			myrow := []string{mytid,arr.Data[i].Rcptemail,arr.Data[i].Fromid,arr.Data[i].ReqTime,arr.Data[i].DelTime,arr.Data[i].ModTime,arr.Data[i].Status,tidsize,url.QueryEscape(arr.Data[i].Subject),arr.Data[i].Xheader,arr.Data[i].Tags,"","","","","",url.QueryEscape(arr.Data[i].Remarks)}

			//Opens status
			if arr.Data[i].Status == "open"{
				for l := range arr.Data[i].Opens{
					myrow:= []string{mytid,arr.Data[i].Rcptemail,arr.Data[i].Fromid,arr.Data[i].ReqTime,arr.Data[i].DelTime,arr.Data[i].ModTime,arr.Data[i].Status,tidsize,url.QueryEscape(arr.Data[i].Subject),arr.Data[i].Xheader,arr.Data[i].Tags,arr.Data[i].Opens[l].IP,arr.Data[i].Opens[l].Time,"","","",url.QueryEscape(arr.Data[i].Remarks)}
					tbl_data = append(tbl_data,myrow)
				}
			}

			//Click Status
			if arr.Data[i].Status == "click"{
				for l := range arr.Data[i].Clicks{
					myrow:= []string{mytid,arr.Data[i].Rcptemail,arr.Data[i].Fromid,arr.Data[i].ReqTime,arr.Data[i].DelTime,arr.Data[i].ModTime,arr.Data[i].Status,tidsize,url.QueryEscape(arr.Data[i].Subject),arr.Data[i].Xheader,arr.Data[i].Tags,"","",arr.Data[i].Clicks[l].IP,arr.Data[i].Clicks[l].Time,arr.Data[i].Clicks[l].Link,url.QueryEscape(arr.Data[i].Remarks)}
					tbl_data = append(tbl_data,myrow)
				}
			}

			tbl_data = append(tbl_data,myrow)
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetHeader([]string{"TransId", "rcptEmail", "fromaddress", "requestedTime","deliveryTime","modifiedTime","status","size","subject","xapiheader","tags","OpenIP","Opentime","ClickIP","ClickTime","ClickLink","Remarks"})
		table.SetAlignment(tablewriter.ALIGN_LEFT)   // Set Alignment
		for _, v := range tbl_data {
			table.Append(v)
		}
		table.SetFooter([]string{"", "","","","","","","","","","","Status",arr.Status,"LimitPass",limit,"TotalRecords",strconv.FormatInt(arr.TotRec,10)}) // Add Footer
		spin.Stop()
		table.Render()
	}
	return "NO Arguments Passed for fetchLogs\nTry pepipost fetchLogs -h [arguments]...\n\n"
}

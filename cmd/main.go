package main

import (
	"encoding/json"
	"flag"
	"github.com/StartOpsTools/diagHttpSystem/v1/pkg/httpRequest"
	"net"
	"regexp"
	"strings"
)

var url string
var reportUrl string

// 上报 trace 消息
type ReportBody struct {
	ErrorMessage string `json:"error_message"`
	Url string `json:"url"`
	Domain string `json:"domain"`
	DomainIpAddr string `json:"domain_ip_addr"`
	LocalDns string `json:"local_dns"`
	RequestBody string `json:"request_body"`
	TcpBody  string `json:"tcp_body"`
}

func init(){
	flag.StringVar(&url, "url", "", "")
	flag.StringVar(&reportUrl, "reportUrl", "http://hook.startops.com.cn/v1/hook/diag/http", "")
	flag.Parse()
}

func main(){
	var reportBody ReportBody
	var domainIpAddr string
	
	reportBody.Url = url
	
	urlCompile, err := regexp.Compile("(http[s]?)(://)(.*)")
	if err != nil {
		reportBody.ErrorMessage = err.Error()
		
		reportTrace(reportUrl, reportBody)
		return
	}
	ipCompile, err := regexp.Compile("\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}\\.\\d{1,3}")
	if err != nil {
		
		reportBody.ErrorMessage = err.Error()
		reportTrace(reportUrl, reportBody)
		return
	}
	
	urlStrings := urlCompile.FindStringSubmatch(url)
	if len(urlStrings) != 4 {
		reportBody.ErrorMessage = "正则表达式解析url失败."
		return
	}
	
	requestUrl := urlStrings[3]
	requestUrlSlice := strings.Split(requestUrl, "/")
	serverName := requestUrlSlice[0]
	domainSlice := strings.Split(serverName, ":")
	domain := domainSlice[0]
	
	// 访问域名解析
	domainIpAddr = ipCompile.FindString(domain)
	if domainIpAddr == "" {
		domainIpAddr, err := lookUpDns(domain)
		if err != nil {
			reportBody.DomainIpAddr = "解析域名失败, Domain: " + domain + ". err: " + err.Error()
		} else {
			reportBody.DomainIpAddr = domainIpAddr
		}
	}
	
	// local dns
	localDns, err := lookUpDns("whoami.akamai.net")
	if err != nil {
		reportBody.LocalDns = "localDns whoami.akamai.net failed, err: " + err.Error()
	} else {
		reportBody.LocalDns = localDns
	}
	
	reportBody.Domain = domain
	
	requestBody := httpRequest.Get(url)
	reportBody.RequestBody = requestBody
	
	// bpf
	/*
	domainIpAddrs := strings.Split(reportBody.DomainIpAddr, ",")
	domainIpAddrs = append(domainIpAddrs, domain)
	bpfString := strings.Join(domainIpAddrs, " or host ")
	bpfString = "( host " + bpfString + " )"
	bpfFilter := fmt.Sprintf("tcp and port %d and %s", 80, bpfString)
	fmt.Println("bpfString: ", bpfString)
	sniffer.OpenLive(bpfFilter)
	 */
	
	reportTrace(reportUrl, reportBody)
}


func reportTrace(reportUrl string, reportBody ReportBody) {
	
	reportData, err := json.Marshal(&reportBody)
	if err != nil {
		return
	}
	
	httpRequest.ReportTrace(reportUrl, reportData)
}


func lookUpDns(domain string) (string, error){
	
	lookupAddrs, err := net.LookupHost(domain)
	
	if err != nil {
		return "", err
	}
	
	lookUpString := strings.Join(lookupAddrs, ",")
	return lookUpString, nil
}


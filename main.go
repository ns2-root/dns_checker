package main

import (
	"fmt"
	"html/template"
	"net"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func DNS(service string) ([]string, error) {
	ipaddrs, err := net.LookupIP(service)
	if err != nil {
		return nil, err
	}

	var ipStrList []string
	for _, ip := range ipaddrs {
		ipStr := ip.String()
		ipStrList = append(ipStrList, ipStr)
		fmt.Println(ipStr)
	}

	return ipStrList, nil
}
func addrs() ([]string, error) {
	hostname, err := os.Hostname()
	if err != nil {
		return nil, err
	}

	addrs, err := net.LookupHost(hostname)
	if err != nil {
		return nil, err
	}
	var serverIpStrList []string
	for _, serverIp := range addrs {
		serverIpStr := serverIp
		serverIpStrList = append(serverIpStrList, serverIpStr)
		fmt.Println(serverIpStr)
	}
	return serverIpStrList, nil
	//fmt.Println(addrs)
}

func main() {
	router := gin.Default()
	router.ForwardedByClientIP = true
	router.SetTrustedProxies([]string{"89.252.140.72"})
	htmlTemplate := template.Must(template.New("index").Parse(`
	<!DOCTYPE html>
	<html>
	<head>
		<title>DNS Checker</title>
	</head>
	<body>
		<h3>DNS Checker</h3>
		<form action="/" method="get">
			<label for="service">Servis Adı:</label>
			<input type="text" name="service" required>
			<button type="submit">Sorgula</button>
		</form>
		{{ if .Values }}
			<h4>IP Adresleri:</h4>
			<ul>
				{{ range .Values }}
					<li>{{ . }}</li>
				{{ end }}
			</ul>
		{{ end }}
	</body>
	</html>
	`))

	router.GET("/", func(c *gin.Context) {
		service := c.Query("service")
		if service == "" {
			c.HTML(http.StatusBadRequest, "index", gin.H{"error": "Servis adı boş olamaz"})
			return
		}

		values, err := DNS(service)
		if err != nil {
			c.HTML(http.StatusInternalServerError, "index", gin.H{"error": fmt.Sprintf("IP adresi bulunamadı: %v", err)})
			return
		}

		c.HTML(http.StatusOK, "index", gin.H{
			"Values": values,
		})
	})

	router.SetHTMLTemplate(htmlTemplate)
	err := router.Run(":5457")
	if err != nil {
		return
	}
}

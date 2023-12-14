package main

import (
	"fmt"
	"html/template"
	"net"
	"net/http"

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

func main() {
	checker := gin.Default()
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

	checker.GET("/", func(c *gin.Context) {
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

	checker.SetHTMLTemplate(htmlTemplate)
	err := checker.Run(":8080")
	if err != nil {
		return
	}
}

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

var (
	certificates = []string{"ibiscybernetics.com"}
)

const (
	defaultTimeout = 3
	webroot        = "/usr/local/apache2/htdocs/"
	indexPage      = webroot + "index.html"
)

func main() {

	htmlContent := ""

	// check if the webroot directory exists
	pages, err := ioutil.ReadDir(webroot)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	if len(pages) == 0 {
		fmt.Println("Unable to locate files in: " + webroot)
		os.Exit(1)
	}

	htmlContent = start_template

	for _, cert := range certificates {

		certOutput := "Warning: Unable to obtain certificate."

		// assume that DNS is correctly setup for wildcard domains to have proper defaults
		certGenericDomain := strings.Replace(cert, "*", "default", -1)

		out, err := exec.Command("get_certs_from_domain", certGenericDomain).Output()
		if err == nil && len(out) > 0 {
			certOutput = string(out)
		}

		htmlContent += "    <p>"
		htmlContent += "        <h4>" + cert + "</h4>"
		htmlContent += "        <pre>"
		htmlContent += certOutput
		htmlContent += "        </pre>"
		htmlContent += "    </p>"
	}

	htmlContent += "<table>"

	// append the end_template component
	htmlContent += end_template

	// attempt to write out the contents to the webroot
	_, err = os.Create(indexPage)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	bytes := []byte(htmlContent)
	err = ioutil.WriteFile(indexPage, bytes, 0644)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	os.Exit(0)
}

/*
 * Templates
 */

const start_template = `
<!DOCTYPE html>
<html>
    <head>
        <meta charset="UTF-8">
        <title>Ibis Cybernetics - Service Health Dashboard</title>
        <style>
            body {
                font-family: Tahoma, Verdana, Arial, sans-serif;
            }
	    #content {
                margin: 0px auto 0px auto;
		padding: 0px;
		overflow: hidden;
                width: 70em;
	    }
            #control_panel {
                border-bottom: 1px solid rgb(40,40,40);
                margin-bottom: 40px;
                margin-top: 50px;
                padding-bottom: 20px;
            }
	    #left_hand_side_logo {
		float: left;
	    }
	    #table_header_name {
		text-align: left;
		width: 20em;
	    }
	    #table_header_url {
		text-align: left;
		width: 25em;
	    }
	    #table_header_server {
		text-align: left;
		width: 22em;
	    }
	    #table_header_type {
		text-align: left;
		width: 12em;
	    }
	    #table_header_is_active {
		text-align: left;
		width: 5em;
	    }
            .button_selected {
                border-bottom: 3px solid black;
                color: black;
                font-weight: 600;
                margin-right: 36px;
                padding-bottom: 20px;
                text-decoration: none;
            }
            .button_unselected {
                border-bottom: 0px solid black;
                color: black;
                font-weight: 300;
                margin-right: 36px;
                padding-bottom: 20px;
                text-decoration: none;
            }
	    .green {
		    color: green;
	    }
	    .red {
		    color: red;
	    }
        </style>
    </head>

    <body>
	<img id="left_hand_side_logo" alt="Welcome to the status page." src="logo.png" style="width: 150px;" />

        <div id="content">
	    <h1><a href="https://ibiscybernetics.com">Ibis Cybernetics</a> &gt;&gt; Service Health Dashboard</h1>

	    <div id="control_panel">
               <a href="index.html" class="button_selected">TLS Certificates</a>
	    </div>
`

const end_template = ` 
	    </table>
        </div>
    </body>
</html>
`

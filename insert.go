package main

import (
	"bytes"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"
	"text/template"
)

var snippets = `
{{ define "about" }}
## About Signal Sciences

[Signal Sciences](https://signalsciences.com/) empowers security and engineering teams by providing visible and effective web application security protecting against real-world attacks.  With our unique hybrid on-premise and cloud architecture, we process, protect and report on billions of requests per day for some of the most sophisticated companies in the world ranging from Adobe to Vimeo, Taser to Under Armour. Our goal is making a more secure Web, with tools that people love to use, written by people who love to make them.

{{ end }}

{{ define "apply" }}
## Apply

To apply, send the following to careers@signalsciences.com

1. Your resume, preferably in PDF, plaintext or markdown format.
2. Your github or other social-coding handle, or a URL to your personal site
   or blog.
3. A brief introduction to yourself, and why the job and Signal Sciences
   are right you.

Didn't see quite the right job?  Email us anyways.
{{ end }}
`

func main() {
	base := template.Must(template.New("snippets").Parse(snippets))

	// get every file in "tmpl/*"
	fds, err := ioutil.ReadDir("tmpl")
	if err != nil {
		log.Fatalf("Unable to read directory: %s", err)
	}
	for _, fd := range fds {
		name := fd.Name()
		log.Printf("Processing %q", name)
		if !strings.HasSuffix(name, ".md") {
			// if not a markdown file, skip
			continue
		}

		// read in file
		raw, err := ioutil.ReadFile(filepath.Join("tmpl", name))
		if err != nil {
			log.Fatalf("Unable to read file %q: %s", name, err)
		}

		buf := bytes.Buffer{}
		tmpl, err := base.Parse(string(raw))
		if err != nil {
			log.Fatalf("Unable to process page %q: %s", name, err)
		}

		err = tmpl.Execute(&buf, nil)
		if err != nil {
			log.Fatalf("Unable to process template for %q: %s", name, err)
		}

		// write it back out
		err = ioutil.WriteFile(name, buf.Bytes(), 0644)
		if err != nil {
			log.Fatalf("Unable to write file %q: %s", name, err)
		}
	}

}

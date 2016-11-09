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

blah blah

{{ end }}

{{ define "apply" }}
## Apply

To apply, send the following to careers@signalsciences.com

1. Your resume, preferably in plaintext, markdown or PDF formats.
2. Your github or other social-coding handle, or a URL to your personal site
   or blog.
3. A brief introduction to yourself, and why the job and Signal Sciences
   is the right place you.

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

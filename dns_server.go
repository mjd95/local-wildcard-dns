package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"github.com/miekg/dns"
)

const wildcardHostsFilename = "hosts"

func main() {
	dns.HandleFunc(".", handler)
	server := &dns.Server{Addr: ":53", Net: "udp"}
	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}

func handler(w dns.ResponseWriter, r *dns.Msg) {
	names := getNamesFromMessage(r)
	for _, name := range names {
		if result := checkForMatchInFile(name); result != "" {
			m := new(dns.Msg)
			m.SetReply(r)
			rr, err := dns.NewRR(fmt.Sprintf("%s A %s", name, result))
			if err == nil {
				m.Answer = append(m.Answer, rr)
			}
			w.WriteMsg(m)
		}
	}

	defaultMux := dns.NewServeMux()
	defaultMux.ServeDNS(w, r)
}

func getNamesFromMessage(r *dns.Msg) []string {
	questions := r.Question
	var names []string
	for _, question := range questions {
		names = append(names, question.Name)
	}
	return names
}

func checkForMatchInFile(name string) string {
	file, err := os.Open(wildcardHostsFilename)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		entry := strings.Fields(scanner.Text())
		if matches(name, entry[1]) {
			return entry[0]
		}
	}
	return ""
}

func matches(name, entry string) bool {
	// add trailing `.` to entry if necessary
	if entry[len(entry)-1] != byte('.') {
		entry = entry + "."
	}
	// remove leading `*` for matching
	if entry[0] == byte ('*') {
		entry = entry[1:]
	}
	return strings.HasSuffix(name, entry)
}

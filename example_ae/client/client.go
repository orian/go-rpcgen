// +build !appengine

package main

import (
	"github.com/kylelemons/go-rpcgen/example_ae/whoami"
	"github.com/kylelemons/go-rpcgen/webrpc"
	"log"
	"net/url"
	"os"
)

func main() {
	for _, arg := range os.Args[1:] {
		url, err := url.Parse(arg)
		if err != nil {
			log.Printf("invald url %q: %s", arg, err)
			continue
		}

		svc := whoami.NewWhoamiServiceWebClient(webrpc.JSON, url)

		in, out := whoami.Empty{}, whoami.YouAre{}
		if err := svc.Whoami(&in, &out); err != nil {
			log.Printf("whoami(%q): %s", url, err)
			continue
		}
		log.Printf("You are %s", *out.IpAddr)
	}
}

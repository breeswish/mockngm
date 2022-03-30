package main

import (
	"context"
	"flag"
	"net/url"
	"strings"
	"sync"

	"github.com/pingcap/log"
	"go.uber.org/zap"

	"github.com/breeswish/mockngm/topsql"
	"github.com/breeswish/mockngm/utils"
)

func main() {
	targets := flag.String("targets", "tidb://127.0.0.1:10080", "a list of targets separated by comma")
	flag.Parse()

	wg := sync.WaitGroup{}
	targetsArr := strings.Split(*targets, ",")
	for _, target := range targetsArr {
		parsed, err := url.Parse(target)
		if err != nil {
			log.Fatal("Parse target address failed", zap.String("target", target), zap.Error(err))
		}
		wg.Add(1)

		if parsed.Scheme != "tidb" && parsed.Scheme != "tikv" {
			log.Fatal("Unsupported component", zap.String("component", parsed.Scheme), zap.String("target", target))
		}

		c := utils.Component{
			Kind: utils.ComponentKind(parsed.Scheme),
			Addr: parsed.Host,
		}
		s := topsql.NewScraper(context.Background(), c, nil)
		wg.Add(1)
		go func() {
			s.Run()
		}()
	}

	wg.Wait()
}

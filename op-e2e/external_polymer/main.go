package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/ethereum-optimism/optimism/op-e2e/external"
	"github.com/ethereum/go-ethereum/common"
	"github.com/onsi/gomega/gbytes"
	"github.com/onsi/gomega/gexec"
)

func main() {
	var configPath string
	flag.StringVar(&configPath, "config", "", "Execute based on the config in this file")
	flag.Parse()
	if err := run(configPath); err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	os.Exit(0)
}

// relative to where the shim binary lives
const opPolymerBin = "polymer-peptide"

func run(configPath string) error {
	if configPath == "" {
		return fmt.Errorf("must supply a '--config <path>' flag")
	}

	configFile, err := os.Open(configPath)
	if err != nil {
		return fmt.Errorf("could not open config: %w", err)
	}

	var config external.Config
	if err := json.NewDecoder(configFile).Decode(&config); err != nil {
		return fmt.Errorf("could not decode config file: %w", err)
	}

	binPath, err := filepath.Abs(opPolymerBin)
	if err != nil {
		return fmt.Errorf("could not get absolute path of op-polymer")
	}
	if _, err := os.Stat(binPath); err != nil {
		return fmt.Errorf("could not locate op-polymer in working directory")
	}

	fmt.Printf("==================    op-polymer shim executing op-polymer     ==========================\n")
	sess, err := execute(binPath)
	if err != nil {
		return fmt.Errorf("could not execute polymer: %w", err)
	}
	defer sess.Close()

	fmt.Printf("==================    op-polymer shim encoding ready-file  ==========================\n")
	if err := external.AtomicEncode(config.EndpointsReadyPath, sess.endpoints); err != nil {
		return fmt.Errorf("could not encode endpoints")
	}

	fmt.Printf("==================    op-polymer shim awaiting termination  ==========================\n")
	select {
	case <-sess.session.Exited:
		return fmt.Errorf("geth exited")
	case <-time.After(30 * time.Minute):
		return fmt.Errorf("exiting after 30 minute timeout")
	}
}

type gethSession struct {
	session   *gexec.Session
	endpoints *external.Endpoints
}

func (es *gethSession) Close() {
	es.session.Terminate()
	select {
	case <-time.After(5 * time.Second):
		es.session.Kill()
	case <-es.session.Exited:
	}
}

func execute(binPath string) (*gethSession, error) {
	cmd := exec.Command(binPath, "start", "--app-rpc-address", "localhost:0")
	sess, err := gexec.Start(cmd, os.Stdout, os.Stderr)
	if err != nil {
		return nil, fmt.Errorf("could not start op-polymer session: %w", err)
	}
	matcher := gbytes.Say("Listening")
	var url string
	var httpPort int64
	var hash common.Hash
	httpPortRE := regexp.MustCompile(`\saddr=([^:]+):(\d+)\s`)
	genesisBlockRE := regexp.MustCompile(`\shash=(\w+)`)
	for (httpPort == 0 && hash == common.Hash{}) {
		match, err := matcher.Match(sess.Out)
		if err != nil {
			return nil, fmt.Errorf("could not execute matcher")
		}
		if !match {
			if sess.Out.Closed() {
				return nil, fmt.Errorf("op-polymer exited before announcing http ports")
			}
			// Wait for a bit more output, then try again
			time.Sleep(10 * time.Millisecond)
			continue
		}

		for _, line := range strings.Split(string(sess.Out.Contents()), "\n") {
			found := httpPortRE.FindStringSubmatch(line)
			if len(found) == 3 {
				url = found[1]
				httpPort, err = strconv.ParseInt(found[2], 10, 32)
				if err != nil {
					return nil, err
				}
				continue
			}

			found = genesisBlockRE.FindStringSubmatch(line)
			if len(found) == 2 {
				hash = common.HexToHash(found[1])
				continue
			}
		}
	}

	return &gethSession{
		session: sess,
		endpoints: &external.Endpoints{
			HTTPEndpoint:     fmt.Sprintf("http://%s:%d/", url, httpPort),
			WSEndpoint:       fmt.Sprintf("ws://%s:%d/websocket", url, httpPort),
			HTTPAuthEndpoint: fmt.Sprintf("http://%s:%d/", url, httpPort),
			WSAuthEndpoint:   fmt.Sprintf("ws://%s:%d/websocket", url, httpPort),
			GenesisBlockHash: hash.Hex(),
		},
	}, nil
}

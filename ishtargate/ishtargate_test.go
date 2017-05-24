package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os/exec"
	"sabey.co/unittest"
	"testing"
)

func TestIshtar(t *testing.T) {
	fmt.Println("TestIshtar")

	// execute ishtargate
	// you must "go build" first
	unittest.IsNil(t, exec.Command("./ishtargate").Run())

	// compare output, make sure its deterministic

	// Hostname

	// MyPC
	first, err := ioutil.ReadFile("unittest/ishtar/hostname/MyPC.hostname")
	unittest.Equals(t, len(first) > 0, true)
	unittest.IsNil(t, err)
	second, err := ioutil.ReadFile("unittest/results/hostname/MyPC.hostname")
	unittest.IsNil(t, err)
	unittest.Equals(t, bytes.Equal(first, second), true)
	// MediaServer
	first, err = ioutil.ReadFile("unittest/ishtar/hostname/MediaServer.hostname")
	unittest.Equals(t, len(first) > 0, true)
	unittest.IsNil(t, err)
	second, err = ioutil.ReadFile("unittest/results/hostname/MediaServer.hostname")
	unittest.IsNil(t, err)
	unittest.Equals(t, bytes.Equal(first, second), true)

	// Hosts

	// MyPC
	first, err = ioutil.ReadFile("unittest/ishtar/hosts/MyPC.hosts")
	unittest.Equals(t, len(first) > 0, true)
	unittest.IsNil(t, err)
	second, err = ioutil.ReadFile("unittest/results/hosts/MyPC.hosts")
	unittest.IsNil(t, err)
	unittest.Equals(t, bytes.Equal(first, second), true)
	// MediaServer
	first, err = ioutil.ReadFile("unittest/ishtar/hosts/MediaServer.hosts")
	unittest.Equals(t, len(first) > 0, true)
	unittest.IsNil(t, err)
	second, err = ioutil.ReadFile("unittest/results/hosts/MediaServer.hosts")
	unittest.IsNil(t, err)
	unittest.Equals(t, bytes.Equal(first, second), true)

	// SSH

	// MyPC-mysql
	first, err = ioutil.ReadFile("unittest/ishtar/ssh/MyPC-mysql.sh")
	unittest.Equals(t, len(first) > 0, true)
	unittest.IsNil(t, err)
	second, err = ioutil.ReadFile("unittest/results/ssh/MyPC-mysql.sh")
	unittest.IsNil(t, err)
	unittest.Equals(t, bytes.Equal(first, second), true)
	// MyPC-ssh
	first, err = ioutil.ReadFile("unittest/ishtar/ssh/MyPC-ssh.sh")
	unittest.Equals(t, len(first) > 0, true)
	unittest.IsNil(t, err)
	second, err = ioutil.ReadFile("unittest/results/ssh/MyPC-ssh.sh")
	unittest.IsNil(t, err)
	unittest.Equals(t, bytes.Equal(first, second), true)
	// MediaServer-http
	first, err = ioutil.ReadFile("unittest/ishtar/ssh/MediaServer-http.sh")
	unittest.Equals(t, len(first) > 0, true)
	unittest.IsNil(t, err)
	second, err = ioutil.ReadFile("unittest/results/ssh/MediaServer-http.sh")
	unittest.IsNil(t, err)
	unittest.Equals(t, bytes.Equal(first, second), true)

	// Firewall

	// MyPC
	first, err = ioutil.ReadFile("unittest/ishtar/firewall/MyPC.iptables")
	unittest.Equals(t, len(first) > 0, true)
	unittest.IsNil(t, err)
	second, err = ioutil.ReadFile("unittest/results/firewall/MyPC.iptables")
	unittest.IsNil(t, err)
	unittest.Equals(t, bytes.Equal(first, second), true)
	// MediaServer
	first, err = ioutil.ReadFile("unittest/ishtar/firewall/MediaServer.iptables")
	unittest.Equals(t, len(first) > 0, true)
	unittest.IsNil(t, err)
	second, err = ioutil.ReadFile("unittest/results/firewall/MediaServer.iptables")
	unittest.IsNil(t, err)
	unittest.Equals(t, bytes.Equal(first, second), true)
	// remote
	first, err = ioutil.ReadFile("unittest/ishtar/firewall/remote.iptables")
	unittest.Equals(t, len(first) > 0, true)
	unittest.IsNil(t, err)
	second, err = ioutil.ReadFile("unittest/results/firewall/remote.iptables")
	unittest.IsNil(t, err)
	unittest.Equals(t, bytes.Equal(first, second), true)
	// solo
	first, err = ioutil.ReadFile("unittest/ishtar/firewall/solo.iptables")
	unittest.Equals(t, len(first) > 0, true)
	unittest.IsNil(t, err)
	second, err = ioutil.ReadFile("unittest/results/firewall/solo.iptables")
	unittest.IsNil(t, err)
	unittest.Equals(t, bytes.Equal(first, second), true)

}

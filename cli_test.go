package cardanocli

import (
	"strings"
	"testing"
)

func TestCli(t *testing.T) {
	cli := New()
	cli.Path = "echo"

	var out string

	err := cli.Cmd().Address().Transaction().OptTxOut("hello").Run(CollectStdout(&out))

	if err != nil {
		t.Fatal(err.Error())
	}

	if out != "address transaction --tx-out hello" {
		t.Fatal("output not collected")
	}

}

func TestCliEnv(t *testing.T) {
	cli := New()
	cli.Path = "env"
	cli.SocketPath = "/node.socket"

	var out string

	err := cli.Cmd().Run(CollectStdout(&out))

	if err != nil {
		t.Fatal(err.Error())
	}

	if !strings.Contains(out, "CARDANO_NODE_SOCKET_PATH=/node.socket") {
		t.Fatal("environment not set")
	}

}

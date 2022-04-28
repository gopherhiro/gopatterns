package responsibility

import (
	"testing"
)

func TestChainOfResponsibility(t *testing.T) {
	chain := &HandlerChain{}
	chain.addHandler(&HandlerA{})
	chain.addHandler(&HandlerB{})
	chain.addHandler(&HandlerC{})
	chain.handle()
}

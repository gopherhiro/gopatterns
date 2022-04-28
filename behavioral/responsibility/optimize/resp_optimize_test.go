package optimize

import "testing"

func TestChainOfRespOptimize(t *testing.T) {
	chain := &HandlerChain{}
	chain.addHandler(&HandlerA{})
	chain.addHandler(&HandlerB{})
	chain.addHandler(&HandlerC{})
	chain.handle()
}

func TestAChainOfRespOptimize(t *testing.T) {
	chain := &AHandlerChain{}
	chain.addHandler(&HandlerA{})
	chain.addHandler(&HandlerB{})
	chain.addHandler(&HandlerC{})
	chain.handle()
}

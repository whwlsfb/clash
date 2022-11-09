package rules

import (
	"strconv"
	"strings"

	C "github.com/Dreamacro/clash/constant"
)

type InPort struct {
	port    string
	adapter string
}

func (g *InPort) RuleType() C.RuleType {
	return C.InPort
}

func (g *InPort) Match(metadata *C.Metadata) bool {
	return strings.EqualFold(metadata.InPort, g.port)
}

func (g *InPort) Adapter() string {
	return g.adapter
}

func (g *InPort) Payload() string {
	return g.port
}

func (g *InPort) ShouldResolveIP() bool {
	return false
}

func (g *InPort) ShouldFindProcess() bool {
	return false
}

func NewInPort(port string, adapter string) (*InPort, error) {
	_, err := strconv.ParseUint(port, 10, 16)
	if err != nil {
		return nil, errPayload
	}
	inPort := &InPort{
		port:    port,
		adapter: adapter,
	}

	return inPort, nil
}

package traffic

import (
	"fmt"
	"log"
	"net"
)

type GeneratorConfig struct {
	Interface *net.Interface
	Targets   []string
}

func (r GeneratorConfig) String() string {
	return fmt.Sprintf("Interface: %v, Targets: %v", r.Interface, r.Targets)
}

// GenerateTraffic generates traffic based on the provided configuration.
func GenerateTraffic(cfg GeneratorConfig) error {
	log.Default().Println("Generating traffic with config: ", cfg)
	return nil
}

package standard

import (
	"fmt"

	"github.com/cilium/ebpf/rlimit"
	"github.com/microsoft/retina/pkg/utils"
)

func (d *Daemon) RemoveMemlock() error {
	return rlimit.RemoveMemlock()
}

func (d *Daemon) TuneSysctls() error {
	if err := utils.TuneSysctls(); err != nil {
		return fmt.Errorf("failed to tune sysctls: %w", err)
	}
	return nil
}

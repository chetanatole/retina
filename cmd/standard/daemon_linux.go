package standard

import (
	"github.com/cilium/ebpf/rlimit"
	"github.com/microsoft/retina/pkg/utils"
)

func (d *Daemon) RemoveMemlock() error {
	return rlimit.RemoveMemlock()
}

func (d *Daemon) TuneSysctls() error {
	return utils.TuneSysctls()
}

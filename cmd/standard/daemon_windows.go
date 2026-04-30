package standard

func (d *Daemon) RemoveMemlock() error {
	// This function is a no-op on Windows.
	return nil
}

func (d *Daemon) TuneSysctls() error {
	// This function is a no-op on Windows.
	return nil
}

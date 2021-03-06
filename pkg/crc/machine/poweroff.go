package machine

import "github.com/pkg/errors"

func (client *client) PowerOff() error {
	libMachineAPIClient, cleanup, err := createLibMachineClient(false)
	defer cleanup()
	if err != nil {
		return errors.Wrap(err, "Cannot initialize libmachine")
	}

	host, err := libMachineAPIClient.Load(client.name)
	if err != nil {
		return errors.Wrap(err, "Cannot load machine")
	}

	if err := host.Kill(); err != nil {
		return errors.Wrap(err, "Cannot kill machine")
	}
	return nil
}

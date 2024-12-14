package notifysend

import (
	"fmt"
	"os/exec"
)

type NotifySendLinux struct {
	executable string
	appName string
	icon 	string
	wait 	uint

	_ struct {}
}

func Init(appName, ico string, waitTime uint) *NotifySendLinux {
	exe, err := exec.LookPath("sh")

	if err != nil {
		exe = "/sbin/sh"
	}

	return &NotifySendLinux{
		executable: exe,
		appName: appName,
		icon: ico,
		wait: waitTime,
	}
}

func (notify *NotifySendLinux) Push (title, message string) error {
	exe, err := exec.LookPath("notify-send")

	if err != nil {
		exe = "notify-send"
	}

	execute := fmt.Sprintf(
		"%s --app-name '%s' '%s' '%s' ",
		exe,
		notify.appName,
		title, message,
	)

	fmt.Println(execute)

	comm := exec.Command(notify.executable, "-c", execute)

	_, err = comm.Output()
	return err
}
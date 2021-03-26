package config

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"github.com/pkg/errors"
	"os/exec"
	"runtime"
	"strings"
)

func (c *config) inputSlackToken() (string, error) {
	validate := func(input string) error {
		if strings.HasPrefix(input, "xoxp-") {
			return nil
		}
		return errors.New("slack tokenを入力してください")
	}
	prompt := promptui.Prompt{
		Label:    "slack tokenを入力してください",
		Validate: validate,
	}
	if err := browserOpen("https://api.slack.com/start/overview#creating"); err != nil {
		return "", err
	}
	result, err := prompt.Run()
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return "", err
	}
	return result, nil
}

func browserOpen(url string) error {
	switch runtime.GOOS {
	case "windows":
		return exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "linux":
		return exec.Command("xdg-open", url).Start()
	case "darwin":
		return exec.Command("open", url).Start()
	default:
		return fmt.Errorf("unsupported platform")
	}
}

package main

import (
	"errors"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"os"
	"os/exec"
	"strings"
)

type ImageConvert struct {
	push             bool
	pull             bool
	tagDockerHubUser string
	sourceImage      string
	rootCmd          *cobra.Command
}

func New() ImageConvert {
	ic := ImageConvert{}
	var rootCmd = &cobra.Command{
		Use:   "image-convert",
		Short: "镜像转换",
		Run: func(cmd *cobra.Command, args []string) {
			err := Convert(ic.sourceImage, ic.tagDockerHubUser, ic.push, ic.pull)
			if err != nil {
				logrus.Error(err)
				return
			}
		},
	}
	ic.rootCmd = rootCmd
	rootCmd.PersistentFlags().StringVar(&ic.sourceImage, "s-image", "", "原始镜像名")
	rootCmd.PersistentFlags().StringVar(&ic.tagDockerHubUser, "docker-hub-user", "1083298593", "原始镜像名")
	rootCmd.PersistentFlags().BoolVar(&ic.push, "push", false, "推送镜像，默认为 true")
	rootCmd.PersistentFlags().BoolVar(&ic.pull, "pull", false, "拉取镜像，默认为 false")
	return ImageConvert{
		rootCmd: rootCmd,
	}
}

func (ic *ImageConvert) Execute() {
	if err := ic.rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
func main() {
	ic := New()
	ic.Execute()
}

func Convert(sourceImage, tagDockerHubUser string, push, pull bool) error {
	if len(sourceImage) == 0 {
		return errors.New("需要指定 --s-image ")
	}
	if !push && !pull {
		return errors.New("需要指定 --pull 或 --push ")
	}
	list1 := strings.Split(sourceImage, ":")
	version := list1[1]
	list2 := strings.Split(list1[0], "/")
	imageName := list2[len(list2)-1]
	trgImageName := fmt.Sprintf("%s/%s:%s", tagDockerHubUser, imageName, version)
	if push {
		logrus.Infof("拉取镜像(%s)中...", sourceImage)
		pullCmd := exec.Command("docker", "pull", sourceImage)
		stdout, err := pullCmd.Output()
		if err != nil {
			return err
		}
		logrus.Infof("镜像(%s -> %s)转换中...", sourceImage, trgImageName)
		tagCmd := exec.Command("docker", "tag", sourceImage, trgImageName)
		stdout, err = tagCmd.Output()
		if err != nil {
			return err
		}
		logrus.Info(string(stdout))

		logrus.Infof("镜像(%s)推送中...", trgImageName)
		pushCmd := exec.Command("docker", "push", trgImageName)
		stdout, err = pushCmd.Output()
		if err != nil {
			return err
		}
		logrus.Info(string(stdout))
		return nil
	}

	if pull {
		logrus.Infof("拉取镜像(%s)中...", trgImageName)
		pushCmd := exec.Command("docker", "pull", trgImageName)
		stdout, err := pushCmd.Output()
		if err != nil {
			return err
		}
		logrus.Info(string(stdout))

		logrus.Infof("镜像(%s -> %s)转换中...", trgImageName, sourceImage)
		tagCmd := exec.Command("docker", "tag", trgImageName, sourceImage)
		stdout, err = tagCmd.Output()
		if err != nil {
			return err
		}
		logrus.Info(string(stdout))
		return err
	}
	return nil
}

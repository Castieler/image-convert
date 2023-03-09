package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/spf13/cobra"
)

var (
	push             bool
	pull             bool
	tagDockerHubUser string
	sourceImage      string
)

var rootCmd = &cobra.Command{
	Use:   "hugo",
	Short: "Hugo is a very fast static site generator",
	Long: `A Fast and Flexible Static Site Generator built with
                love by spf13 and friends in Go.
                Complete documentation is available at https://gohugo.io/documentation/`,
	Run: func(cmd *cobra.Command, args []string) {
		list1 := strings.Split(sourceImage, ":")
		version := list1[1]
		list2 := strings.Split(list1[0], "/")
		imageName := list2[len(list2)-1]
		trgImageName := fmt.Sprintf("%s/%s:%s", tagDockerHubUser, imageName, version)
		if push {
			tagCmd := exec.Command("docker", "tag", sourceImage, trgImageName)
			stdout, err := tagCmd.Output()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(string(stdout))
			pushCmd := exec.Command("docker", "push", trgImageName)
			stdout, err = pushCmd.Output()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(string(stdout))
			return
		}

		if pull {
			pushCmd := exec.Command("docker", "pull", trgImageName)
			stdout, err := pushCmd.Output()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(string(stdout))

			tagCmd := exec.Command("docker", "tag", trgImageName, sourceImage)
			stdout, err = tagCmd.Output()
			if err != nil {
				fmt.Println(err.Error())
				return
			}
			fmt.Println(string(stdout))
			return
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
func main() {
	rootCmd.PersistentFlags().StringVar(&sourceImage, "s-image", "", "原始镜像名")
	rootCmd.PersistentFlags().StringVar(&tagDockerHubUser, "docker-hub-user", "1083298593", "原始镜像名")
	rootCmd.PersistentFlags().BoolVar(&push, "push", false, "推送镜像，默认为 true")
	rootCmd.PersistentFlags().BoolVar(&pull, "pull", false, "拉取镜像，默认为 false")
	rootCmd.Execute()
}

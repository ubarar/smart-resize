package main

import (
	"flag"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strconv"

	"github.com/davidbyttow/govips/v2/vips"
	"github.com/schollz/progressbar/v3"
	"github.com/ubarar/smart-resize/images"
)

func main() {
	threads := flag.Int("threads", 4, "vips concurrency level")
	tasks := createTasks()

	if len(tasks) == 0 {
		return
	}

	vips.LoggingSettings(func(_ string, _ vips.LogLevel, _ string) {}, vips.LogLevelError)
	vips.Startup(&vips.Config{ConcurrencyLevel: *threads, })
	defer vips.Shutdown()


	bar := progressbar.Default(int64(len(tasks)))

	for _, task := range tasks {
		im := images.ResizeImage(task.Name, task.Target)
		images.SaveImage(im, filepath.Join("..", strconv.Itoa(task.Target), task.Name))
		bar.Add(1)
	}
}

var extensionsRegex = regexp.MustCompile(`\.((jpg)|(jpeg)|(png))$`)
var targetsRegex = regexp.MustCompile(`^[1-9][0-9]+$`)

// for given directory dir, list all files that are images
func getFiles(dir string, regex *regexp.Regexp) []string{
	files, err := os.ReadDir(dir)

	if err != nil {
		log.Fatal("failed to get files", err)
	}

	results := []string{}

	for _,file := range files {
		if regex.MatchString(file.Name()) {
			results = append(results, file.Name())
		}
	}

	return results
}

func getOriginalFiles() []string {
	return getFiles(".", extensionsRegex)
}

type ResizeTask struct {
	Name string
	Target int
}

func setDifference(a []string, b []string) []string{
	diff := []string{}

	mb := make(map[string]struct{}, len(b))

	for _, y := range b {
		mb[y] = struct{}{}
	}

	for _, x := range a {
		if _, ok := mb[x]; !ok {
			diff = append(diff, x)
		}
	}

	return diff
}

func createTasks() []ResizeTask {
	originals := getOriginalFiles()
	targets := getFiles("..", targetsRegex)

	tasks := []ResizeTask{}

	for _, target := range targets {
		targetFiles := getFiles(".." + "/" + target, extensionsRegex)

		newFiles := setDifference(originals, targetFiles)

		targetInt, _ := strconv.Atoi(target)

		for _, file := range newFiles {
			tasks = append(tasks, ResizeTask{file, int(targetInt)})
		}
	}
	return tasks
}






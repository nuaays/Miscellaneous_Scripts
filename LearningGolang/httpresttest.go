package main

import (
	"bytes"
	"errors"
	"fmt"
	"math/rand"
	"os/exec"
	_ "regexp"
	"strconv"
	"strings"
	"time"
)

var p = fmt.Println

func GitCheckout(projectId int64, addr string) (string, error) {
	if strings.TrimSpace(addr) == "" {
		return "", errors.New("git repo address cannot be null")
	}
	tmp_addr := strings.TrimLeft(strings.TrimLeft(addr, "git@"), "http://")

	gitRepoAddr := fmt.Sprintf("http://%s:%s@%s", "httptest", "***", tmp_addr)

	dir := strconv.FormatInt(projectId+time.Now().Unix()+rand.Int63n(10000), 10)
	path := fmt.Sprintf("%s/%s", "/tmp", dir)

	//beego.Info(fmt.Sprintf("To pull files from %s into %s for project %d", addr, path, projectId))

	cmd := exec.Command("git",
		"clone", gitRepoAddr,
		path,
	)

	_, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return "", errors.New(string(ee.Stderr))
		}
		return "", err
	}

	return path, nil
}

func SvnCheckout(projectId int64, addr string) (string, error) {
	if strings.TrimSpace(addr) == "" {
		return "", errors.New("svn address cannot be null")
	}

	dir := strconv.FormatInt(projectId+time.Now().Unix()+rand.Int63n(10000), 10)
	path := fmt.Sprintf("%s/%s", "/tmp", dir)

	cmd := exec.Command("svn",
		"co", addr,
		"--username", "httptest",
		"--password", "*****",
		"--non-interactive", path,
	)

	// var out bytes.Buffer
	// cmd.Stdout = &out
	// err := cmd.Run()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("in all caps: %q\n", out.String())　　//in all caps: "SOME INPUT"
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run() //cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			return "", errors.New(string(ee.Stderr))
		}
		return "", err
	} else {
		p(out.String())
	}

	return path, nil
}

func main() {
	dir := strconv.FormatInt(7255+time.Now().Unix()+rand.Int63n(10000), 10)
	fmt.Println(dir)

	repoPath := "git@gitlab.***.com/**/*.git"
	//repoPath := "http://1**:8888/repos/**"

	tmp_addr := strings.TrimLeft(strings.TrimLeft(strings.TrimLeft(repoPath, "git"), "@"), "http://")

	p(tmp_addr)
	p(strings.TrimLeft(strings.TrimLeft(repoPath, "git@"), "http://"))

	// var f = SvnCheckout

	// if strings.HasSuffix(repoPath, ".git") {
	// 	p("Git Repo!")
	// 	f = GitCheckout
	// } else {
	// 	f = SvnCheckout
	// }

	// path, err := f(1234, repoPath)
	// if err != nil {
	// 	p("ERROR!", repoPath)

	// } else {
	// 	p(path)

	// }

}

func test() {
	mystr := "http://**.**.**.**:8888/repos/***********"
	tmp_arr := strings.Split(mystr, "//")
	username := "httptest"
	password := "***"
	tmp_arr[0] += fmt.Sprintf("//%s:%s@", username, password)
	//p(tmp_arr[0])
	//p(strings.Join(tmp_arr, ""))
	repoPath := strings.Join(tmp_arr, "")
	p(repoPath)

	targetPath := "/tmp/repo"

	//reg = regexp.MustCompile(`(?i:^hello).*Go`)
	//fmt.Printf("%q\n", reg.FindAllString(text, -1))
	// ["Hello 世界！123 Go"]

	//fmt.Println(regexp.MatchString("^http.*.git ", repoPath))
	//fmt.Println(regexp.MustCompile(`(?i:^hello).*git`, []byte(repoPath)))

	cmd := exec.Command("git",
		"clone", repoPath,
		//"--username", beego.AppConfig.String("SVN_Username"),
		//"--password", beego.AppConfig.String("SVN_Password"),
		//"--non-interactive", path,
		targetPath,
	)

	_, err := cmd.Output()
	if err != nil {
		if ee, ok := err.(*exec.ExitError); ok {
			p(errors.New(string(ee.Stderr)))
		}
		//return err
	}
}

package lib

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
)

func init() {
	fmt.Println("Loaded env modle")
}

type Conf struct {
	Last_Opt time.Time
	File     string
	Stroes   map[string]interface{}
}

func NewConf(fs string) (*Conf, error) {
	f, err := os.OpenFile(fs, os.O_RDONLY, 0644)
	defer f.Close()
	if err != nil {
		return nil, err
	}

	stores, err := _init(f)
	if err != nil {
		return nil, err
	}

	return &Conf{Last_Opt: time.Now(), File: fs, Stroes: stores}, nil
}

func _init(f *os.File) (map[string]interface{}, error) {
	val := make(map[string]interface{})

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		l := scanner.Text()
		if ok, err := regexp.MatchString("(^\\s*#)", l); err != nil || ok {
			fmt.Println("read conf skip lien : ", l)
			continue
		}
		arr := strings.Split(l, "=")
		if len(arr) == 2 {
			k := strings.TrimSpace(arr[0])
			v := strings.TrimSpace(arr[1])
			val[k] = v
		}
	}
	return val, nil
}

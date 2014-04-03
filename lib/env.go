package lib

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func init() {
	//fmt.Println("Initial env model")
}

var (
	CONF_MGR = make(map[string]*Conf)

	err_key_is_nil = errors.New("Paramter key is empty")
	key_not_exists = "Can not find the corresponding key[%s]."
)

type Conf struct {
	Last_Opt time.Time
	File     string
	Stroes   map[string]string
}

func NewConf(fs string) (*Conf, error) {
	if v, ok := CONF_MGR[fs]; ok {
		return v, nil
	}
	this := &Conf{Last_Opt: time.Now(), File: fs}
	return this.load_conf()
}

func (self *Conf) Get(k string) (string, error) {
	if k == "" {
		return "", err_key_is_nil
	}
	if v := self.Stroes[k]; v == "" {
		return "", errors.New(fmt.Sprintf(key_not_exists, k))
	} else {
		return v, nil
	}
}

func (self *Conf) AssertInt(k string) int {
	v, err := self.Get(k)
	if err != nil {
		panic(err)
	}
	if r, err := strconv.Atoi(v); err != nil {
		panic(err)
	} else {
		return r
	}
}

func (self *Conf) load_conf() (*Conf, error) {
	val := make(map[string]string)

	f, err := os.OpenFile(self.File, os.O_RDONLY, 0644)
	defer f.Close()
	if err != nil {
		return nil, err
	}

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
	self.Stroes = val
	return regiter_conf(self)
}

func regiter_conf(conf *Conf) (*Conf, error) {
	CONF_MGR[conf.File] = conf
	return conf, nil
}

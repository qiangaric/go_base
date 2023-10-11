// base_demo/split.go

package split

import (
	"strings"

	// "github.com/pingcap/log"
	// "github.com/pingcap/log"
	log "github.com/sirupsen/logrus"
)

// Split 把字符串s按照给定的分隔符sep进行分割返回字符串切片
func Split(s, sep string) (result []string) {
	i := strings.Index(s, sep)
	log.Info("i = ",i)
	for i > -1 {
		result = append(result, s[:i])
		log.Info(result)
		s = s[i+1:]
		i = strings.Index(s, sep)
	}
	result = append(result, s)
	return
}

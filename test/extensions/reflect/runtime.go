package reflect

import (
	"reflect"
	"regexp"
	"runtime"
	"strings"
)

var gccgoRE = regexp.MustCompile(`\.pN\d+_`)

// GetSimpleFuncName Цельноутащено отсюда: Called в mock testify
func GetSimpleFuncName(funcInfo interface{}) string {
	ptr := reflect.ValueOf(funcInfo).Pointer()

	functionPath := runtime.FuncForPC(ptr).Name()
	// Next four lines are required to use GCCGO function naming conventions.
	// For Ex:  github_com_docker_libkv_store_mock.WatchTree.pN39_github_com_docker_libkv_store_mock.Mock
	// uses interface information unlike golang github.com/docker/libkv/store/mock.(*Mock).WatchTree
	// With GCCGO we need to remove interface information starting from pN<dd>.
	if gccgoRE.MatchString(functionPath) {
		functionPath = gccgoRE.Split(functionPath, -1)[0]
	}
	parts := strings.Split(functionPath, ".")
	functionName := parts[len(parts)-1]

	return strings.Split(functionName, "-")[0]
}

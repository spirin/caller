package caller

import (
	"regexp"
	"runtime"
	"strings"
)

var cbRegexp = regexp.MustCompile(`func\d+`)

// Get raw caller name from runtime
func Get(skip int) string {
	pc, _, _, ok := runtime.Caller(skip)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		return details.Name()
	}

	return "unknown"
}

// Get full caller name, e.g., "testNameFullObj.Fn @ github.com/spirin/caller"
func NameFull() string {
	return getNameFull()
}

// Get caller name, e.g., "caller.testNameObj.Fn"
func Name() string {
	return getName()
}

// Get caller package name, e.g., "caller"
func Package() string {
	return getPackage()
}

func getNameFull() string {
	pkg, callerName := parseName(Get(3))

	if len(callerName) == 0 {
		return pkg
	}

	return callerName + " @ " + pkg
}

func getName() string {
	pkg, callerName := parseName(Get(3))

	pkgParts := strings.Split(pkg, "/")
	pkgName := pkgParts[len(pkgParts)-1]

	if len(callerName) == 0 {
		return pkgName
	}

	return pkgName + "." + callerName
}

func getPackage() string {
	callerName := Get(3)

	pos := strings.LastIndexByte(callerName, '/')
	if pos > -1 {
		callerName = callerName[pos+1:]
	}

	return strings.Split(callerName, ".")[0]
}

func parseName(rawName string) (string, string) {
	var callerName, pkgPath string

	pos := strings.LastIndexByte(rawName, '/')
	if pos > -1 {
		pkg, rawCaller := rawName[:pos], rawName[pos+1:]
		parts := strings.SplitN(rawCaller, ".", 2)
		callerName = parts[1]
		pkgPath = pkg + "/" + parts[0]
	} else {
		parts := strings.SplitN(rawName, ".", 2)
		callerName = parts[1]
		pkgPath = parts[0]
	}

	filtered := make([]string, 0)

	for _, v := range strings.Split(callerName, ".") {
		if v == "" {
			continue
		}
		if v == "init" {
			continue
		}
		if v == "glob" {
			continue
		}
		if cbRegexp.MatchString(v) {
			continue
		}
		if strings.IndexByte(v, '(') > -1 {
			v = v[2 : len(v)-1]
		}
		filtered = append(filtered, v)
	}

	return pkgPath, strings.Join(filtered, ".")
}

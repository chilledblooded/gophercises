package middleware

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"runtime/debug"
	"strings"
)

//RecoveryMid function will recover from the panic situation.
//If any fatal error or panic occurs it will recover error.
func RecoveryMid(app http.Handler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Println(err)
				stack := debug.Stack()
				log.Println(string(stack))
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "<h1>panic: %v</h1><pre>%s</pre>", err, errLinks(string(stack)))
			}
		}()
		app.ServeHTTP(w, r)
	}
}

//errLinks function is helper function to generate the links of file where error has occur.
func errLinks(stack string) string {
	lines := strings.Split(stack, "\n")
	for li, line := range lines {
		if len(line) == 0 || line[0] != '\t' {
			continue
		}
		filePath := ""
		for i, ch := range line {
			if ch == ':' {
				filePath = line[1:i]
				break
			}
		}
		var lineStr strings.Builder
		for i := len(filePath) + 2; i < len(line); i++ {
			if line[i] < '0' || line[i] > '9' {
				break
			}
			lineStr.WriteByte(line[i])
		}
		v := url.Values{}
		v.Set("path", filePath)
		v.Set("line", lineStr.String())
		lines[li] = "\t<a href=\"/debug/?" + v.Encode() + "\">" + filePath + ":" + lineStr.String() + "</a>" + line[len(filePath)+2+len(lineStr.String()):]
	}
	return strings.Join(lines, "\n")
}

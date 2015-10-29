package tricorder

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

const (
	browseMetricsUrl = "/tricorder"
	htmlTemplateStr  = `
	<html>
	<head>
	  <link rel="stylesheet" type="text/css" href="/tricorderstatic/theme.css">
	</head>
	<body>
	{{with $top := .}}
	  {{range .List}}
	    {{if .Directory}}
	      <a href="{{$top.Link .Directory}}">{{.Directory.AbsPath}}</a><br>
            {{else}}
	      {{if $top.IsDistribution .Metric.Value.Type}}
	        {{.Metric.AbsPath}} (distribution: {{.Metric.Description}})<br>
	        {{with .Metric.Value.AsDistribution.Snapshot}}
		  <table>
	          {{range .Breakdown}}
		    <tr>
  	            {{if .First}}
	              <td align="right">&lt;{{.End}}:</td><td align="right">{{.Count}}</td>
	            {{else if .Last}}
	              <td align="right">&gt;={{.Start}}:</td><td align="right"> {{.Count}}</td>
	            {{else}}
		    <td align="right">{{.Start}}-{{.End}}:</td> <td align="right">{{.Count}}</td>
	            {{end}}
		    </tr>
		  {{end}}
		  </table>
	          {{if .Count}}
		  <span class="summary"> min: {{.Min}} max: {{.Max}} avg: {{.Average}} &#126;median: {{$top.ToFloat32 .Median}} count: {{.Count}}</span><br><br>
	          {{end}}
		{{end}}
	      {{else}}
	        {{.Metric.AbsPath}} {{.Metric.Value.AsHtmlString}} ({{.Metric.Value.Type}}: {{.Metric.Description}})<br>
	      {{end}}
	    {{end}}
	  {{end}}
	{{end}}
	</body>
	</html>
	  `

	themeCss = `
	.summary {color:#999999; font-style: italic;}
	  `
)

var (
	htmlTemplate = template.Must(template.New("browser").Parse(htmlTemplateStr))
	errLog       *log.Logger
	appStartTime time.Time
)

type view struct {
	*directory
}

func (v *view) Link(d *directory) string {
	return browseMetricsUrl + d.AbsPath()
}

func (v *view) IsDistribution(t valueType) bool {
	return t == Dist
}

func (v *view) ToFloat32(f float64) float32 {
	return float32(f)
}

func emitDirectoryAsHtml(d *directory, w io.Writer) error {
	v := &view{d}
	if err := htmlTemplate.Execute(w, v); err != nil {
		return err
	}
	return nil
}

func browseFunc(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	d := root.GetDirectory(path)
	if d == nil {
		fmt.Fprintf(w, "Path does not exist.")
		return
	}
	if err := emitDirectoryAsHtml(d, w); err != nil {
		fmt.Fprintln(w, "Error in template.")
		errLog.Printf("Error in template: %v\n", err)
	}
}

func newStatic() http.Handler {
	result := http.NewServeMux()
	addStatic(result, "/theme.css", themeCss)
	return result
}

func addStatic(mux *http.ServeMux, path, content string) {
	addStaticBinary(mux, path, []byte(content))
}

func addStaticBinary(mux *http.ServeMux, path string, content []byte) {
	mux.Handle(
		path,
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			http.ServeContent(
				w,
				r,
				path,
				appStartTime,
				bytes.NewReader(content))
		}))
}

func getProgramArgs() string {
	return strings.Join(os.Args[1:], " ")
}

func registerDefaultMetrics() {
	RegisterMetric("/name", &os.Args[0], None, "Program name")
	RegisterMetric("/args", getProgramArgs, None, "Program args")
}

func initHttpFramework() {
	appStartTime = time.Now()
	errLog = log.New(os.Stderr, "", log.LstdFlags|log.Lmicroseconds)
}

func registerBrowserHandlers() {
	http.Handle(browseMetricsUrl+"/", http.StripPrefix(browseMetricsUrl, http.HandlerFunc(browseFunc)))
	http.Handle("/tricorderstatic/", http.StripPrefix("/tricorderstatic", newStatic()))
}

func init() {
	registerDefaultMetrics()
	initHttpFramework()
	registerBrowserHandlers()

}

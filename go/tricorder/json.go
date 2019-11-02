package tricorder

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Cloud-Foundations/tricorder/go/tricorder/messages"
)

var (
	jsonUrl = "/metricsapi"
)

func jsonAsMetric(m *metric, s *session) *messages.Metric {
	var result messages.Metric
	m.InitJsonMetric(s, &result)
	return &result
}

type jsonMetricsCollector messages.MetricList

func (c *jsonMetricsCollector) Collect(m *metric, s *session) (err error) {
	*c = append(*c, jsonAsMetric(m, s))
	return nil
}

func jsonSetUpHeaders(h http.Header) {
	h.Set("Content-Type", "application/json")
	h.Set("X-Tricorder-Media-Type", "tricorder.v1")
}

func jsonHandlerFunc(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	jsonSetUpHeaders(w.Header())
	path := r.URL.Path
	var content []byte
	var err error
	if r.Form.Get("singleton") != "" {
		m := root.GetMetric(path)
		if m == nil {
			httpError(w, http.StatusNotFound)
			return
		}
		content, err = json.Marshal(jsonAsMetric(m, nil))
	} else {
		collector := make(jsonMetricsCollector, 0)
		root.GetAllMetricsByPath(path, &collector, nil)
		content, err = json.Marshal(collector)
	}
	if err != nil {
		handleError(w, err)
		return
	}
	var buffer bytes.Buffer
	json.Indent(&buffer, content, "", "\t")
	buffer.WriteTo(w)
}

func initJsonHandlers() {
	http.Handle(jsonUrl+"/", http.StripPrefix(jsonUrl, gzipHandler{http.HandlerFunc(jsonHandlerFunc)}))
}

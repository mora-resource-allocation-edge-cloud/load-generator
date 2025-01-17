package conf

import (
	"os"
	"strconv"
)

var (
	ServiceUrl          = getEnvString("SERVICE_URL", "http://edge-vp-1.master.particles.dieei.unict.it")
	ZipfS               = getEnvFloat64("ZIPF_S", 1.01)
	ZipfV               = getEnvFloat64("ZIPF_V", 1)
	ExpLambda           = getEnvFloat64("EXP_AVG", 0.1) // Average requests per second
	PostMetricsEndPoint = getEnvString("POST_METRICS_ENDPOINT",
		"http://video-metrics-collector.zion.alessandrodistefano.eu:8080/v1/video-reproduction")
	ClientUrl 			= getEnvString("CLIENT_URL",
		"http://video-metrics-collector.zion.alessandrodistefano.eu:8880/samples/dash-if-reference-player-api-metrics-push/index.html")
	MaxExecutionTime    = getEnvInt64("MAX_TIME_PER_REQUEST", 900)
	MaxExposedPorts 	= getEnvInt64("MAX_EXPOSED_PORTS", 128)
)

func getEnvString(key string, defaultValue string) string {
	if str, ok := os.LookupEnv(key); ok {
		return str
	}
	return defaultValue
}

func getEnvFloat64(key string, defaultValue float64) float64 {
	if str, ok := os.LookupEnv(key); ok {
		if i, e := strconv.ParseFloat(str, 64); e != nil {
			return i
		}
	}
	return defaultValue
}

func getEnvInt64(key string, defaultValue int64) int64 {
	if str, ok := os.LookupEnv(key); ok {
		if i, e := strconv.ParseInt(str, 10, 64); e != nil {
			return i
		}
	}
	return defaultValue
}

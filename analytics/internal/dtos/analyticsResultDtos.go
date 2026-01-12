package dtos

type AnalyticsResult struct {
	UseKw     float64
	GenKw     float64
	NetKw     float64
	Timestamp int64
	Model     string
}

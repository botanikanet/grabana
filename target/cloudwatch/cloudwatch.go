package cloudwatch

// Option represents an option that can be used to configure a cloudwatch query.
type Option func(target *Loki)

// Cloudwatch represents a loki query.
type Cloudwatch struct {
	Ref          string
	Namespace    string
	MetricName   string
	Region       string
	Statistics   []string
	Dimensions   map[string]string
	Period       string
	LegendFormat string
}

// New creates a new cloudwatch query.
func New(metric string, namespace string, options ...Option) *Loki {
	cloudwatch := &Cloudwatch{
		MetricName: metric,
		Namespace:  namespace,
	}

	for _, opt := range options {
		opt(cloudwatch)
	}

	return cloudwatch
}

// Legend sets the legend format.
func Legend(legend string) Option {
	return func(cloudwatch *Cloudwatch) {
		cloudwatch.LegendFormat = legend
	}
}

// Region sets the region.
func Region(region string) Option {
	return func(cloudwatch *Cloudwatch) {
		cloudwatch.Region = region
	}
}

// Statistic sets the statistic.
func Statistic(statistic []string) Option {
	return func(cloudwatch *Cloudwatch) {
		cloudwatch.Statistic = statistic
	}
}

// Dimensions sets the dimensions.
func Dimensions(dimensions map[string]string) Option {
	return func(cloudwatch *Cloudwatch) {
		cloudwatch.Dimensions = dimensions
	}
}

// Ref sets the reference ID for this query.
func Ref(ref string) Option {
	return func(cloudwatch *Cloudwatch) {
		cloudwatch.Ref = ref
	}
}

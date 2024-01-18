package cloudwatch

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestNewCloudwatchTargetCanBeCreated(t *testing.T) {
	req := require.New(t)
	query := "{app=\"cloudwatch\"}"

	target := New(query)

	req.Equal(query, target.Expr)
}

func TestLegendCanBeConfigured(t *testing.T) {
	req := require.New(t)
	legend := "lala"

	target := New("", Legend(legend))

	req.Equal(legend, target.LegendFormat)
}

func TestRefCanBeConfigured(t *testing.T) {
	req := require.New(t)

	target := New("", Ref("A"))

	req.Equal("A", target.Ref)
	req.False(target.Hidden)
}

func TestTargetCanBeHidden(t *testing.T) {
	req := require.New(t)

	target := New("", Hide())

	req.True(target.Hidden)
}

func TestMatchExactCanBeSet(t *testing.T) {
	req := require.New(t)

	target := New("", "", MatchExact("B"))

	req.Equal("B", target.MatchExact)
}

func TestRegionCanBeSet(t *testing.T) {
	req := require.New(t)

	target := New("", "", Region("C"))

	req.Equal("C", target.Region)
}

func TestSqlExpressionCanBeSet(t *testing.T) {
	req := require.New(t)

	target := New("", "", SqlExpression("D"))

	req.Equal("D", target.SqlExpression)
}

func TestStatisticCanBeSet(t *testing.T) {
	req := require.New(t)

	target := New("", "", Statistic("E"))

	req.Equal("E", target.SqlExpression)
}

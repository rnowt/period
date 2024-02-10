package period

import (
	"errors"
	"time"
)

type Period struct {
	Start time.Time
	End   time.Time
}

var ErrIllegalPeriod = errors.New("period: start before end")

func NewPeriod(start, end time.Time) (*Period, error) {
	if end.Before(start) {
		return nil, ErrIllegalPeriod
	}
	return &Period{
		Start: start,
		End:   end,
	}, nil
}

func (p *Period) Contains(dt time.Time) bool {
	return !dt.Before(p.Start) && !dt.After(p.End)
}

func (p *Period) Overlaps(op *Period) bool {
	return !p.Start.After(op.End) && !p.End.Before(op.Start)
}

// TODO
func (p *Period) Overlap(op *Period) time.Duration {
	return 0
}

func (p *Period) Split(n uint8) []Period {
	duration := p.End.Sub(p.Start)
	step := duration / time.Duration(n)
	periods := make([]Period, 0, n)
	start := p.Start
	for range n {
		periods = append(periods, Period{start, start.Add(step)})
		start = start.Add(step)
	}
	return periods
}

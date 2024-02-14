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

func New(start, end time.Time) (*Period, error) {
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

func (p *Period) Overlap(op *Period) *Period {
	maxStart := p.Start
	if op.Start.After(p.Start) {
		maxStart = op.Start
	}
	minEnd := p.End
	if op.End.Before(p.End) {
		minEnd = op.End
	}
	return &Period{maxStart, minEnd}
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

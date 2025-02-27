package series

import (
	"errors"
	"github.com/PDA-Khacer/go-dataframe/common"
)

func Apply[S common.Frame, D common.Frame](s *Series[S], fn func(*Series[S]) *Series[D]) (*Series[D], error) {
	if s == nil {
		return nil, errors.New("source series is nil")
	}

	re := fn(s)

	re.UpdateDType()
	return re, nil
}

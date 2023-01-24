package config

import (
	"fmt"
	"math/rand"
)

type (
	ComicYear struct {
		Year              int
		SequenceStart     int
		SequenceEnd       int
		SequencePrefix    string
		SequencePrefixEnd int
	}
)

func (c ComicYear) RandomImageFilename() string {
	r := rand.Intn(c.SequenceEnd-c.SequenceStart) + c.SequenceStart
	if c.SequencePrefixEnd > 0 && r > c.SequencePrefixEnd {
		return fmt.Sprintf("%d_%d.jpg", c.Year, r)
	}

	return fmt.Sprintf("%d_%s%d.jpg", c.Year, c.SequencePrefix, r)
}

var (
	AllComics = []*ComicYear{
		&ComicYear{
			Year:          1980,
			SequenceStart: 8,
			SequenceEnd:   89,
		},
		&ComicYear{
			Year:          1981,
			SequenceStart: 5,
			SequenceEnd:   88,
		},
		&ComicYear{
			Year:          1982,
			SequenceStart: 5,
			SequenceEnd:   88,
		},
		&ComicYear{
			Year:           1983,
			SequenceStart:  5,
			SequenceEnd:    88,
			SequencePrefix: "00",
		},
		&ComicYear{
			Year:          1984,
			SequenceStart: 5,
			SequenceEnd:   90,
		},
		&ComicYear{
			Year:          1985,
			SequenceStart: 7,
			SequenceEnd:   94,
		},
		&ComicYear{
			Year:              1986,
			SequenceStart:     5,
			SequenceEnd:       107,
			SequencePrefix:    "0", // everything pre-100 has 0 prefix, 100+ no prefix
			SequencePrefixEnd: 99,
		},
		&ComicYear{
			Year:              1987,
			SequenceStart:     10,
			SequenceEnd:       123,
			SequencePrefix:    "0", // everything pre-100 has 0 prefix, 100+ no prefix
			SequencePrefixEnd: 99,
		},
		&ComicYear{
			Year:          1988,
			SequenceStart: 5,
			SequenceEnd:   100,
		},
		// was on haitus in '89
		//&ComicYear{
		//	Year: 1989,
		//},
		&ComicYear{
			Year:          1990,
			SequenceStart: 5,
			SequenceEnd:   76,
		},
		&ComicYear{
			Year:          1991,
			SequenceStart: 5,
			SequenceEnd:   76,
		},
		&ComicYear{
			Year:          1992,
			SequenceStart: 5,
			SequenceEnd:   76,
		},
		&ComicYear{
			Year:          1993,
			SequenceStart: 6,
			SequenceEnd:   74,
		},
		&ComicYear{
			Year:          1994,
			SequenceStart: 5,
			SequenceEnd:   80,
		},
	}
)

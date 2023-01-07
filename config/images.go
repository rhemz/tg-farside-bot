package config

type (
	ComicYear struct {
		Year           int
		SequenceStart  int
		SequenceEnd    int
		SequencePrefix string
	}
)

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
			Year:           1986,
			SequenceStart:  5,
			SequenceEnd:    107,
			SequencePrefix: "0",
		},
		&ComicYear{
			Year:          1987,
			SequenceStart: 10,
			SequenceEnd:   123,
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

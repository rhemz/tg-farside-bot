package logger

import (
	"github.com/labstack/echo/v4"
	"github.com/rs/zerolog"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

func TestInitLogger1(t *testing.T) {
	type (
		args struct {
			level   zerolog.Level
			logJSON bool
		}
		want struct {
			level      zerolog.Level
			hookLen    int
			hookType   zerolog.Hook
			timeFormat string
		}
	)
	tests := []struct {
		name string
		args args
		want want
	}{
		{
			name: "jsondebug",
			args: args{level: zerolog.DebugLevel, logJSON: true},
			want: want{
				level:      zerolog.DebugLevel,
				hookLen:    1,
				timeFormat: zerolog.TimeFormatUnix,
			},
		},
		{
			name: "nojsoninfo",
			args: args{level: zerolog.InfoLevel, logJSON: false},
			want: want{
				level:      zerolog.InfoLevel,
				hookLen:    1,
				timeFormat: time.RFC3339,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			InitLogger(tt.args.level, tt.args.logJSON)
			assert.IsType(t, zerolog.Logger{}, Logger)

			refl := reflect.ValueOf(Logger)
			assert.Equal(t, tt.want.level, zerolog.Level(refl.FieldByName("level").Int()))
			assert.Equal(t, tt.want.hookLen, refl.FieldByName("hooks").Len())
			assert.Equal(t, tt.want.timeFormat, zerolog.TimeFieldFormat)
		})
	}
}

func TestL(t *testing.T) {
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/", nil)

	type args struct {
		c echo.Context
	}
	tests := []struct {
		name string
		args args
		want *zerolog.Logger
	}{
		{name: "background", args: args{c: e.NewContext(req, httptest.NewRecorder())}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := L(tt.args.c)
			assert.IsType(t, zerolog.Logger{}, *l)

			refl := reflect.ValueOf(*l)
			assert.Equal(t, int64(7), refl.FieldByName("level").Int())
		})
	}
}

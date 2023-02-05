package web

import (
	"fmt"
	"html/template"
	"path/filepath"
	"strconv"
	"time"

	"github.com/gin-gonic/gin/render"
	"github.com/gofrs/uuid"
)

type DynamicRender map[string][]string

var funcMap = template.FuncMap{

	"add": func(first_value, second_value int) int {
		return first_value + second_value
	},

	"processDuration": func(t1, t2 *time.Time) string {
		if t1 != nil && t2 != nil {
			if t1.Equal(*t2) {
				return "0s"
			} else {
				diff := t2.Sub(*t1)
				return fmt.Sprintf("%v", diff)
			}
		} else if t1 != nil && t2 == nil {
			diff := time.Since(*t1)
			return fmt.Sprintf("%v", diff)
		}
		return "-"
	},

	"date": func(time time.Time) string {
		date := time.Format("2006-01-02")
		return date
	},

	"dateUnix": func(unixDate int64) string {
		date := time.Unix(unixDate, 0).Format("2006-01-02")
		return date
	},

	"dateUnixDay": func(unixDate int64) int {
		day := time.Unix(unixDate, 0).Day()
		return day
	},

	"dateUnixDayWithOrder": func(unixDate int64) string {
		var ordinalDictionary = map[int]string{
			0: "th",
			1: "st",
			2: "nd",
			3: "rd",
			4: "th",
			5: "th",
			6: "th",
			7: "th",
			8: "th",
			9: "th",
		}
		day := time.Unix(unixDate, 0).Day()
		return strconv.Itoa(day) + ordinalDictionary[day%10]
	},

	"dateUnixDayName": func(unixDate int64) time.Weekday {
		day := time.Unix(unixDate, 0).Weekday()
		return day
	},

	"dateUnixMonthName": func(unixDate int64) time.Month {
		month := time.Unix(unixDate, 0).Month()
		return month
	},

	"dateUnixWeekInMonth": func(unixDate int64) int {
		_, currentWeekNumber := time.Unix(unixDate, 0).ISOWeek()
		return currentWeekNumber
	},

	"dateAndTime": func(time time.Time) string {
		date := time.Format("2006-01-02 15:04:05")
		return date
	},

	"dateAndTimeUnix": func(unixDate int64) string {
		date := time.Unix(unixDate, 0).Format("2006-01-02 15:04:05")
		return date
	},

	"type": func(i interface{}) string {
		value := fmt.Sprintf("%T", i)
		return value
	},

	"idUint": func(i interface{}) uint {
		value := fmt.Sprintf("%T", i)
		s, _ := strconv.Atoi(value)
		return uint(s)
	},

	"idUintPointer": func(i *uint) uint {
		if i == nil {
			return 0
		}
		return *i
	},

	"intPointer": func(i *int) int {
		if i == nil {
			return 0
		}
		return *i
	},

	"toString": func(id uuid.UUID) string {
		return id.String()
	},

	"bytesToString": func(b []byte) string {
		return string(b)
	},

	"isHistorical": func(first, last int64) string {
		from := time.Unix(first, 0)
		to := time.Unix(last, 0)

		if int(to.Sub(from).Hours()/24) > 30 {
			return "Yes"
		} else {
			return "No"
		}
	},
}

func (r *repository) NewMultiTemplate(templateMap map[string]map[string][]string) render.HTMLRender {
	renderedTemplates := make(map[string][]string)
	for path, templates := range templateMap {
		for base, views := range templates {
			for i, view := range views {
				views[i] = fmt.Sprintf("%s/%s.html", path, view)
			}

			renderedTemplates[base] = views
		}

	}

	return DynamicRender(renderedTemplates)
}

func (r DynamicRender) Instance(name string, data interface{}) render.Render {
	views, ok := r[name]
	if !ok {
		panic(fmt.Sprintf("Dynamic template with name %s not found", name))
	}

	baseName := filepath.Base(views[0])
	tmpl := template.Must(template.New(baseName).Funcs(funcMap).ParseFiles(views...))

	return render.HTML{
		Template: tmpl,
		Data:     data,
	}
}

var Views = map[string][]string{
	// "login": {"base", "login"},
	"login":    {"login"},
	"overview": {"base", "overview"},
}

func (r *repository) GetViews() map[string][]string {
	return Views
}

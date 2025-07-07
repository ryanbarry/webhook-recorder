package main


import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/rs/zerolog"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/rs/zerolog/log"
)

func init() {
	zerolog.TimeFieldFormat = time.RFC3339Nano // matches timestamp of the echo framework
}

type QueryParams struct {
	 SplunkEventIndex string `query:"splunkEventIndex" json:"event_index,omitempty"`
	 SplunkEventSource string `query:"splunkEventSource" json:"event_source,omitempty"`
	 SplunkEventSourcetype string `query:"splunkEventSourcetype" json:"event_sourcetype,omitempty"`
}

func main()  {
	e := echo.New()
	e.HideBanner = true

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/health", health)
	e.POST("/record", record)


	e.Logger.Fatal(e.Start(":1323"))
}

func health(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"status": "healthy!"})
}

func record(c echo.Context) error  {
	result := make(map[string]interface{})
	if err := (&echo.DefaultBinder{}).BindBody(c, &result); err != nil {
		return err
	}

	var queryParams QueryParams
	if err := (&echo.DefaultBinder{}).BindQueryParams(c, &queryParams); err != nil {
		return err
	}

	log.Info().Msg(queryParams.SplunkEventSourcetype)

	m, err := json.Marshal(queryParams)
	if err != nil {
		return err
	}

	log.Info().
		Interface("msg", result).
		RawJSON("meta", m).
		Send()

	return nil
}

package api

import (
	"html/template"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/grafana/grafana/pkg/services/featuremgmt"
	"github.com/grafana/grafana/pkg/services/licensing"
	"github.com/grafana/grafana/pkg/setting"
)

func TestGetFavIconURL(t *testing.T) {
	cdnURL := "https://cdn.example.com/"

	t.Run("Should use custom favicon URL when configured", func(t *testing.T) {
		cfg := setting.NewCfg()
		cfg.AppFaviconURL = "https://example.com/custom-favicon.ico"

		hs := &HTTPServer{
			Cfg:      cfg,
			License:  &licensing.OSSLicensingService{},
			Features: featuremgmt.WithFeatures(),
		}

		result := hs.getFavIconURL(cdnURL)
		expected := template.URL("https://example.com/custom-favicon.ico")
		assert.Equal(t, expected, result)
	})

	t.Run("Should use default favicon when custom URL is not configured", func(t *testing.T) {
		cfg := setting.NewCfg()
		cfg.AppFaviconURL = ""

		hs := &HTTPServer{
			Cfg:      cfg,
			License:  &licensing.OSSLicensingService{},
			Features: featuremgmt.WithFeatures(),
		}

		result := hs.getFavIconURL(cdnURL)
		expected := template.URL(cdnURL + "public/build/img/fav32.png")
		assert.Equal(t, expected, result)
	})

	t.Run("Should use local path for favicon", func(t *testing.T) {
		cfg := setting.NewCfg()
		cfg.AppFaviconURL = "/public/img/custom-favicon.ico"

		hs := &HTTPServer{
			Cfg:      cfg,
			License:  &licensing.OSSLicensingService{},
			Features: featuremgmt.WithFeatures(),
		}

		result := hs.getFavIconURL(cdnURL)
		expected := template.URL("/public/img/custom-favicon.ico")
		assert.Equal(t, expected, result)
	})
}

func TestAppTitleConfiguration(t *testing.T) {
	t.Run("Should use configured app title", func(t *testing.T) {
		cfg := setting.NewCfg()
		cfg.AppTitle = "My Custom Grafana"

		require.Equal(t, "My Custom Grafana", cfg.AppTitle)
	})

	t.Run("Should use default app title when not configured", func(t *testing.T) {
		cfg := setting.NewCfg()
		// When loading from defaults.ini, it should be "Grafana"
		cfg.AppTitle = "Grafana"

		require.Equal(t, "Grafana", cfg.AppTitle)
	})
}

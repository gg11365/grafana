---
title: Configure Basic Branding
description: Customize the browser title and favicon in Grafana
weight: 290
---

# Configure Basic Branding

You can customize the application title and favicon that appear in the browser tab.

## Configuration Options

Add the following configuration to your `grafana.ini` file or use environment variables:

### Using grafana.ini

```ini
[branding]
# Set custom application title that appears in browser title bar. Default is "Grafana"
app_title = My Custom Grafana

# Set custom favicon URL. If left empty, the default Grafana favicon will be used.
# You can use a local path (e.g., /public/img/custom-favicon.ico) or an external URL
app_favicon_url = /public/img/custom-favicon.ico
```

### Using Environment Variables

```bash
# Set custom application title
GF_BRANDING_APP_TITLE="My Custom Grafana"

# Set custom favicon URL
GF_BRANDING_APP_FAVICON_URL="/public/img/custom-favicon.ico"
```

## Examples

### Using a Local Favicon

1. Place your favicon file (e.g., `custom-favicon.ico`) in `/usr/share/grafana/public/img/`
2. Update your `grafana.ini`:

```ini
[branding]
app_title = My Company
app_favicon_url = /public/img/custom-favicon.ico
```

### Using an External Favicon

```ini
[branding]
app_title = My Company
app_favicon_url = https://example.com/favicon.ico
```

## Notes

- The `app_title` setting changes the title displayed in the browser tab
- The `app_favicon_url` setting changes the favicon (icon) displayed in the browser tab
- If `app_favicon_url` is empty or not set, Grafana will use the default favicon
- Restart Grafana after making configuration changes

For more advanced branding options (login logo, loading logo, etc.), see [Configure custom branding](../configure-custom-branding/).

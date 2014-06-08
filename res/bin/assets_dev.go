// +build debug

package res

import (
	"fmt"
	"io/ioutil"
	"strings"
)

// bindata_read reads the given file from disk. It returns an error on failure.
func bindata_read(path, name string) ([]byte, error) {
	buf, err := ioutil.ReadFile(path)
	if err != nil {
		err = fmt.Errorf("Error reading asset %s at %s: %v", name, path, err)
	}
	return buf, err
}

// css_main_css reads file data from disk. It returns an error on failure.
func css_main_css() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/css/main.css",
		"css/main.css",
	)
}

// index_html reads file data from disk. It returns an error on failure.
func index_html() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/index.html",
		"index.html",
	)
}

// vendor_ds_store reads file data from disk. It returns an error on failure.
func vendor_ds_store() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/.DS_Store",
		"vendor/.DS_Store",
	)
}

// vendor_angular_1_3_0_beta_11_ds_store reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_ds_store() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/.DS_Store",
		"vendor/angular-1.3.0-beta.11/.DS_Store",
	)
}

// vendor_angular_1_3_0_beta_11_angular_animate_js reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_animate_js() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-animate.js",
		"vendor/angular-1.3.0-beta.11/angular-animate.js",
	)
}

// vendor_angular_1_3_0_beta_11_angular_animate_min_js reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_animate_min_js() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-animate.min.js",
		"vendor/angular-1.3.0-beta.11/angular-animate.min.js",
	)
}

// vendor_angular_1_3_0_beta_11_angular_animate_min_js_map reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_animate_min_js_map() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-animate.min.js.map",
		"vendor/angular-1.3.0-beta.11/angular-animate.min.js.map",
	)
}

// vendor_angular_1_3_0_beta_11_angular_cookies_js reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_cookies_js() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-cookies.js",
		"vendor/angular-1.3.0-beta.11/angular-cookies.js",
	)
}

// vendor_angular_1_3_0_beta_11_angular_cookies_min_js reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_cookies_min_js() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-cookies.min.js",
		"vendor/angular-1.3.0-beta.11/angular-cookies.min.js",
	)
}

// vendor_angular_1_3_0_beta_11_angular_cookies_min_js_map reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_cookies_min_js_map() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-cookies.min.js.map",
		"vendor/angular-1.3.0-beta.11/angular-cookies.min.js.map",
	)
}

// vendor_angular_1_3_0_beta_11_angular_csp_css reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_csp_css() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-csp.css",
		"vendor/angular-1.3.0-beta.11/angular-csp.css",
	)
}

// vendor_angular_1_3_0_beta_11_angular_loader_js reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_loader_js() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-loader.js",
		"vendor/angular-1.3.0-beta.11/angular-loader.js",
	)
}

// vendor_angular_1_3_0_beta_11_angular_loader_min_js reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_loader_min_js() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-loader.min.js",
		"vendor/angular-1.3.0-beta.11/angular-loader.min.js",
	)
}

// vendor_angular_1_3_0_beta_11_angular_loader_min_js_map reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_loader_min_js_map() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-loader.min.js.map",
		"vendor/angular-1.3.0-beta.11/angular-loader.min.js.map",
	)
}

// vendor_angular_1_3_0_beta_11_angular_messages_js reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_messages_js() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-messages.js",
		"vendor/angular-1.3.0-beta.11/angular-messages.js",
	)
}

// vendor_angular_1_3_0_beta_11_angular_messages_min_js reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_messages_min_js() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-messages.min.js",
		"vendor/angular-1.3.0-beta.11/angular-messages.min.js",
	)
}

// vendor_angular_1_3_0_beta_11_angular_messages_min_js_map reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_messages_min_js_map() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-messages.min.js.map",
		"vendor/angular-1.3.0-beta.11/angular-messages.min.js.map",
	)
}

// vendor_angular_1_3_0_beta_11_angular_mocks_js reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_mocks_js() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-mocks.js",
		"vendor/angular-1.3.0-beta.11/angular-mocks.js",
	)
}

// vendor_angular_1_3_0_beta_11_angular_resource_js reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_resource_js() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-resource.js",
		"vendor/angular-1.3.0-beta.11/angular-resource.js",
	)
}

// vendor_angular_1_3_0_beta_11_angular_resource_min_js reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_resource_min_js() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-resource.min.js",
		"vendor/angular-1.3.0-beta.11/angular-resource.min.js",
	)
}

// vendor_angular_1_3_0_beta_11_angular_resource_min_js_map reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_resource_min_js_map() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-resource.min.js.map",
		"vendor/angular-1.3.0-beta.11/angular-resource.min.js.map",
	)
}

// vendor_angular_1_3_0_beta_11_angular_route_js reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_route_js() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-route.js",
		"vendor/angular-1.3.0-beta.11/angular-route.js",
	)
}

// vendor_angular_1_3_0_beta_11_angular_route_min_js reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_route_min_js() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-route.min.js",
		"vendor/angular-1.3.0-beta.11/angular-route.min.js",
	)
}

// vendor_angular_1_3_0_beta_11_angular_route_min_js_map reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_route_min_js_map() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-route.min.js.map",
		"vendor/angular-1.3.0-beta.11/angular-route.min.js.map",
	)
}

// vendor_angular_1_3_0_beta_11_angular_sanitize_js reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_sanitize_js() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-sanitize.js",
		"vendor/angular-1.3.0-beta.11/angular-sanitize.js",
	)
}

// vendor_angular_1_3_0_beta_11_angular_sanitize_min_js reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_sanitize_min_js() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-sanitize.min.js",
		"vendor/angular-1.3.0-beta.11/angular-sanitize.min.js",
	)
}

// vendor_angular_1_3_0_beta_11_angular_sanitize_min_js_map reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_sanitize_min_js_map() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-sanitize.min.js.map",
		"vendor/angular-1.3.0-beta.11/angular-sanitize.min.js.map",
	)
}

// vendor_angular_1_3_0_beta_11_angular_scenario_js reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_scenario_js() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-scenario.js",
		"vendor/angular-1.3.0-beta.11/angular-scenario.js",
	)
}

// vendor_angular_1_3_0_beta_11_angular_touch_js reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_touch_js() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-touch.js",
		"vendor/angular-1.3.0-beta.11/angular-touch.js",
	)
}

// vendor_angular_1_3_0_beta_11_angular_touch_min_js reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_touch_min_js() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-touch.min.js",
		"vendor/angular-1.3.0-beta.11/angular-touch.min.js",
	)
}

// vendor_angular_1_3_0_beta_11_angular_touch_min_js_map reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_touch_min_js_map() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular-touch.min.js.map",
		"vendor/angular-1.3.0-beta.11/angular-touch.min.js.map",
	)
}

// vendor_angular_1_3_0_beta_11_angular_js reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_js() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular.js",
		"vendor/angular-1.3.0-beta.11/angular.js",
	)
}

// vendor_angular_1_3_0_beta_11_angular_min_js reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_min_js() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular.min.js",
		"vendor/angular-1.3.0-beta.11/angular.min.js",
	)
}

// vendor_angular_1_3_0_beta_11_angular_min_js_map reads file data from disk. It returns an error on failure.
func vendor_angular_1_3_0_beta_11_angular_min_js_map() ([]byte, error) {
	return bindata_read(
		"/Users/advanderveer/Documents/Projects/go/src/github.com/adminibar/boatyard/res/src/vendor/angular-1.3.0-beta.11/angular.min.js.map",
		"vendor/angular-1.3.0-beta.11/angular.min.js.map",
	)
}

// Asset loads and returns the asset for the given name.
// It returns an error if the asset could not be found or
// could not be loaded.
func Asset(name string) ([]byte, error) {
	cannonicalName := strings.Replace(name, "\\", "/", -1)
	if f, ok := _bindata[cannonicalName]; ok {
		return f()
	}
	return nil, fmt.Errorf("Asset %s not found", name)
}

// AssetNames returns the names of the assets.
func AssetNames() []string {
	names := make([]string, 0, len(_bindata))
	for name := range _bindata {
		names = append(names, name)
	}
	return names
}

// _bindata is a table, holding each asset generator, mapped to its name.
var _bindata = map[string]func() ([]byte, error){
	"css/main.css": css_main_css,
	"index.html": index_html,
	"vendor/.DS_Store": vendor_ds_store,
	"vendor/angular-1.3.0-beta.11/.DS_Store": vendor_angular_1_3_0_beta_11_ds_store,
	"vendor/angular-1.3.0-beta.11/angular-animate.js": vendor_angular_1_3_0_beta_11_angular_animate_js,
	"vendor/angular-1.3.0-beta.11/angular-animate.min.js": vendor_angular_1_3_0_beta_11_angular_animate_min_js,
	"vendor/angular-1.3.0-beta.11/angular-animate.min.js.map": vendor_angular_1_3_0_beta_11_angular_animate_min_js_map,
	"vendor/angular-1.3.0-beta.11/angular-cookies.js": vendor_angular_1_3_0_beta_11_angular_cookies_js,
	"vendor/angular-1.3.0-beta.11/angular-cookies.min.js": vendor_angular_1_3_0_beta_11_angular_cookies_min_js,
	"vendor/angular-1.3.0-beta.11/angular-cookies.min.js.map": vendor_angular_1_3_0_beta_11_angular_cookies_min_js_map,
	"vendor/angular-1.3.0-beta.11/angular-csp.css": vendor_angular_1_3_0_beta_11_angular_csp_css,
	"vendor/angular-1.3.0-beta.11/angular-loader.js": vendor_angular_1_3_0_beta_11_angular_loader_js,
	"vendor/angular-1.3.0-beta.11/angular-loader.min.js": vendor_angular_1_3_0_beta_11_angular_loader_min_js,
	"vendor/angular-1.3.0-beta.11/angular-loader.min.js.map": vendor_angular_1_3_0_beta_11_angular_loader_min_js_map,
	"vendor/angular-1.3.0-beta.11/angular-messages.js": vendor_angular_1_3_0_beta_11_angular_messages_js,
	"vendor/angular-1.3.0-beta.11/angular-messages.min.js": vendor_angular_1_3_0_beta_11_angular_messages_min_js,
	"vendor/angular-1.3.0-beta.11/angular-messages.min.js.map": vendor_angular_1_3_0_beta_11_angular_messages_min_js_map,
	"vendor/angular-1.3.0-beta.11/angular-mocks.js": vendor_angular_1_3_0_beta_11_angular_mocks_js,
	"vendor/angular-1.3.0-beta.11/angular-resource.js": vendor_angular_1_3_0_beta_11_angular_resource_js,
	"vendor/angular-1.3.0-beta.11/angular-resource.min.js": vendor_angular_1_3_0_beta_11_angular_resource_min_js,
	"vendor/angular-1.3.0-beta.11/angular-resource.min.js.map": vendor_angular_1_3_0_beta_11_angular_resource_min_js_map,
	"vendor/angular-1.3.0-beta.11/angular-route.js": vendor_angular_1_3_0_beta_11_angular_route_js,
	"vendor/angular-1.3.0-beta.11/angular-route.min.js": vendor_angular_1_3_0_beta_11_angular_route_min_js,
	"vendor/angular-1.3.0-beta.11/angular-route.min.js.map": vendor_angular_1_3_0_beta_11_angular_route_min_js_map,
	"vendor/angular-1.3.0-beta.11/angular-sanitize.js": vendor_angular_1_3_0_beta_11_angular_sanitize_js,
	"vendor/angular-1.3.0-beta.11/angular-sanitize.min.js": vendor_angular_1_3_0_beta_11_angular_sanitize_min_js,
	"vendor/angular-1.3.0-beta.11/angular-sanitize.min.js.map": vendor_angular_1_3_0_beta_11_angular_sanitize_min_js_map,
	"vendor/angular-1.3.0-beta.11/angular-scenario.js": vendor_angular_1_3_0_beta_11_angular_scenario_js,
	"vendor/angular-1.3.0-beta.11/angular-touch.js": vendor_angular_1_3_0_beta_11_angular_touch_js,
	"vendor/angular-1.3.0-beta.11/angular-touch.min.js": vendor_angular_1_3_0_beta_11_angular_touch_min_js,
	"vendor/angular-1.3.0-beta.11/angular-touch.min.js.map": vendor_angular_1_3_0_beta_11_angular_touch_min_js_map,
	"vendor/angular-1.3.0-beta.11/angular.js": vendor_angular_1_3_0_beta_11_angular_js,
	"vendor/angular-1.3.0-beta.11/angular.min.js": vendor_angular_1_3_0_beta_11_angular_min_js,
	"vendor/angular-1.3.0-beta.11/angular.min.js.map": vendor_angular_1_3_0_beta_11_angular_min_js_map,
}

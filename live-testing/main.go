package main

import (
	"bytes"
	"compress/gzip"
	"crypto/tls"
	"io"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"strconv"
	"strings"
)

const (
	httpPort  = ":80"
	httpsPort = ":443"
	certFile  = "server.crt"
	keyFile   = "server.key"
)

var (
	TargetDomain    = "example.com"
	LocalDomain     = "localhost"
	TargetEndpoints = []TargetEndpoint{}
	transport       = &http.Transport{TLSClientConfig: &tls.Config{InsecureSkipVerify: true}}
	tempUrl         = make(map[string]*url.URL)
)

func main() {
	config, err := LoadConfig()
	if err != nil {
		log.Fatal("Failed to load configuration:", err)
	}

	TargetDomain = config.TargetDomain
	LocalDomain = config.LocalDomain
	TargetEndpoints = config.TargetEndpoints

	handle := http.NewServeMux()
	handle.HandleFunc("/", proxyHandler)

	go func() {
		log.Println("Listening on http://" + LocalDomain + httpPort)
		err := http.ListenAndServe(LocalDomain+httpPort, http.HandlerFunc(proxyHandler))
		if err != nil {
			log.Fatal(err)
		}
	}()

	log.Println("Listening on https://" + LocalDomain + httpsPort)
	err = http.ListenAndServeTLS(LocalDomain+httpsPort, certFile, keyFile, http.HandlerFunc(proxyHandler))
	if err != nil {
		log.Fatal(err)
	}
}

func proxyHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request matches any local endpoint
	if ok, localURL := checkLocalEndpoint(r.URL.Path); ok {
		log.Printf("[REQUEST] %s %s%s", r.Method, localURL.String(), r.URL.Path)
		r.Host = localURL.Host
		proxy := httputil.NewSingleHostReverseProxy(localURL)
		proxy.Transport = transport
		proxy.ServeHTTP(w, r)
		return
	}

	// If no local endpoint matches, proxy to the target domain
	hostname, _, err := net.SplitHostPort(r.Host)
	if err != nil {
		hostname = r.Host // fallback if no port present
	}
	subdomain := strings.Replace(hostname, LocalDomain, TargetDomain, 1)

	protocol := "http://"
	if r.TLS != nil {
		protocol = "https://"
	}

	remoteURL, err := urlParse(protocol + subdomain)
	if err != nil {
		http.Error(w, "Invalid remote URL", http.StatusInternalServerError)
		log.Println("[ERROR] Failed to parse remote URL:", err)
		return
	}

	log.Printf("[REQUEST] %s %s%s", r.Method, remoteURL.String(), r.URL.Path)
	r.Host = remoteURL.Host
	proxy := httputil.NewSingleHostReverseProxy(remoteURL)
	proxy.Transport = transport
	proxy.ModifyResponse = modifyResponse
	proxy.ServeHTTP(w, r)
}

func modifyResponse(resp *http.Response) error {
	contentType := resp.Header.Get("Content-Type")

	if resp.StatusCode != http.StatusOK ||
		(!strings.Contains(contentType, "application/javascript") &&
			!strings.Contains(contentType, "text/html") &&
			!strings.HasSuffix(resp.Request.URL.Path, ".js")) {
		return nil
	}

	var reader io.ReadCloser
	var err error

	switch resp.Header.Get("Content-Encoding") {
	case "gzip":
		reader, err = gzip.NewReader(resp.Body)
		if err != nil {
			return err
		}
	default:
		reader = resp.Body
	}
	defer reader.Close()

	body, err := io.ReadAll(reader)
	if err != nil {
		return err
	}

	body = bytes.ReplaceAll(body, []byte(LocalDomain), []byte("nfieunfei"))
	body = bytes.ReplaceAll(body, []byte(TargetDomain), []byte(LocalDomain))

	resp.Body = io.NopCloser(bytes.NewReader(body))
	resp.ContentLength = int64(len(body))
	resp.Header.Set("Content-Length", strconv.Itoa(len(body)))
	resp.Header.Del("Content-Encoding")

	return nil
}

func checkLocalEndpoint(path string) (bool, *url.URL) {
	for _, endpoint := range TargetEndpoints {
		if !strings.HasPrefix(path, endpoint.Prefix) {
			continue
		}

		urlParsed, err := urlParse(endpoint.Local)
		if err != nil {
			log.Printf("[ERROR] Failed to parse local URL for endpoint %s: %v", endpoint.Prefix, err)
			return false, nil
		}
		return true, urlParsed
	}
	return false, nil
}

func urlParse(rawURL string) (*url.URL, error) {
	if parsed, ok := tempUrl[rawURL]; ok {
		return parsed, nil
	}

	if !strings.HasPrefix(rawURL, "http://") && !strings.HasPrefix(rawURL, "https://") {
		rawURL = "http://" + rawURL
	}

	parsed, err := url.Parse(rawURL)
	if err != nil {
		log.Println("[ERROR] Failed to parse URL:", err)
		return nil, err
	}
	tempUrl[rawURL] = parsed
	return parsed, nil
}

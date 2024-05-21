package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"github.com/fatih/color"
	"golang.org/x/net/html"
)

func main() {
	fmt.Println(`
            __  ____        __  
  ___ ____ / /_/ __/__  ___/ /__
 / _  / -_) __/ _// _ \/ _  (_-<
 \_, /\__/\__/___/_//_/\_,_/___/
/___/      - Links Extractor                     
	`)
	var (
		singleURL  string
		listFile   string
		outputFile string
	)

	flag.StringVar(&singleURL, "u", "", "Single URL to fetch")
	flag.StringVar(&listFile, "l", "", "Text file containing list of URLs")
	flag.StringVar(&outputFile, "o", "extracted.txt", "Output file to write extracted URLs")
	flag.Parse()

	if singleURL == "" && listFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var urls []string

	if singleURL != "" {
		urls = append(urls, singleURL)
	}

	if listFile != "" {
		urlsFromFile, err := readURLsFromFile(listFile)
		if err != nil {
			fmt.Println("Error reading URLs from file:", err)
			os.Exit(1)
		}
		urls = append(urls, urlsFromFile...)
	}

	extractedURLs := make([]string, 0)
	for _, u := range urls {
		userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.99 Safari/537.36"
		acceptHeader := "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"

		client := &http.Client{}

		req, err := http.NewRequest("GET", u, nil)
		if err != nil {
			fmt.Println("Error creating request for", u, ":", err)
			continue
		}

		req.Header.Set("User-Agent", userAgent)
		req.Header.Set("Accept", acceptHeader)

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error fetching", u, ":", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Error response for", u, ":", resp.Status)
			continue
		}

		links := extractLinks(resp.Body, u)
		fmt.Println("--------------------------------------------------")
		fmt.Println(color.CyanString("[INFO] Extracted URLs from"), color.YellowString(u))
		fmt.Println("--------------------------------------------------")
		for _, link := range links {
			fmt.Println(color.GreenString("[EXTRACTED] "+link))
			if strings.HasPrefix(link, "/") {
				baseURL, _ := url.Parse(u)
				extractedURL := strings.TrimSuffix(baseURL.String(), "/") + link
				extractedURLs = append(extractedURLs, extractedURL)
			} else if strings.Contains(link, "http://") || strings.Contains(link, "https://") {
				parsedURL, err := url.Parse(link)
				if err != nil {
					fmt.Println("Error parsing URL:", err)
					continue
				}
				if strings.Contains(parsedURL.Host, getHostname(u)) {
					extractedURLs = append(extractedURLs, link)
				}
			}
		}
	}

	err := writeURLsToFile(outputFile, extractedURLs)
	if err != nil {
		fmt.Println("Error writing extracted URLs to file:", err)
	} else {
		fmt.Println(color.MagentaString("[OUTPUT] Extracted URLs written to"), color.YellowString(outputFile))
	}
}

func extractLinks(body io.Reader, baseURL string) []string {
	links := make([]string, 0)

	z := html.NewTokenizer(body)
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			return links
		case tt == html.StartTagToken:
			token := z.Token()
			if token.Data == "a" || token.Data == "script" {
				for _, attr := range token.Attr {
					if attr.Key == "href" || attr.Key == "src" {
						links = append(links, attr.Val)
					}
				}
			}
		}
	}
}

func readURLsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, strings.TrimSpace(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return urls, nil
}

func getHostname(u string) string {
	parsedURL, err := url.Parse(u)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return ""
	}
	return parsedURL.Hostname()
}

func writeURLsToFile(filename string, urls []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, url := range urls {
		_, err := writer.WriteString(url + "\n")
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}
root@imranhudaa:~/urls# 
root@imranhudaa:~/urls# cat getEnds.go 
package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"strings"
	"github.com/fatih/color"
	"golang.org/x/net/html"
)

func main() {
	fmt.Println(`
            __  ____        __  
  ___ ____ / /_/ __/__  ___/ /__
 / _  / -_) __/ _// _ \/ _  (_-<
 \_, /\__/\__/___/_//_/\_,_/___/
/___/      - Links Extractor                     
	`)
	var (
		singleURL  string
		listFile   string
		outputFile string
	)

	flag.StringVar(&singleURL, "u", "", "Single URL to fetch")
	flag.StringVar(&listFile, "l", "", "Text file containing list of URLs")
	flag.StringVar(&outputFile, "o", "extracted.txt", "Output file to write extracted URLs")
	flag.Parse()

	if singleURL == "" && listFile == "" {
		flag.PrintDefaults()
		os.Exit(1)
	}

	var urls []string

	if singleURL != "" {
		urls = append(urls, singleURL)
	}

	if listFile != "" {
		urlsFromFile, err := readURLsFromFile(listFile)
		if err != nil {
			fmt.Println("Error reading URLs from file:", err)
			os.Exit(1)
		}
		urls = append(urls, urlsFromFile...)
	}

	extractedURLs := make([]string, 0)
	for _, u := range urls {
		userAgent := "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.99 Safari/537.36"
		acceptHeader := "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8"

		client := &http.Client{}

		req, err := http.NewRequest("GET", u, nil)
		if err != nil {
			fmt.Println("Error creating request for", u, ":", err)
			continue
		}

		req.Header.Set("User-Agent", userAgent)
		req.Header.Set("Accept", acceptHeader)

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error fetching", u, ":", err)
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != http.StatusOK {
			fmt.Println("Error response for", u, ":", resp.Status)
			continue
		}

		links := extractLinks(resp.Body, u)
		fmt.Println("--------------------------------------------------")
		fmt.Println(color.CyanString("[INFO] Extracted URLs from"), color.YellowString(u))
		fmt.Println("--------------------------------------------------")
		for _, link := range links {
			fmt.Println(color.GreenString("[EXTRACTED] "+link))
			if strings.HasPrefix(link, "/") {
				baseURL, _ := url.Parse(u)
				extractedURL := strings.TrimSuffix(baseURL.String(), "/") + link
				extractedURLs = append(extractedURLs, extractedURL)
			} else if strings.Contains(link, "http://") || strings.Contains(link, "https://") {
				parsedURL, err := url.Parse(link)
				if err != nil {
					fmt.Println("Error parsing URL:", err)
					continue
				}
				if strings.Contains(parsedURL.Host, getHostname(u)) {
					extractedURLs = append(extractedURLs, link)
				}
			}
		}
	}

	err := writeURLsToFile(outputFile, extractedURLs)
	if err != nil {
		fmt.Println("Error writing extracted URLs to file:", err)
	} else {
		fmt.Println(color.MagentaString("[OUTPUT] Extracted URLs written to"), color.YellowString(outputFile))
	}
}

func extractLinks(body io.Reader, baseURL string) []string {
	links := make([]string, 0)

	z := html.NewTokenizer(body)
	for {
		tt := z.Next()
		switch {
		case tt == html.ErrorToken:
			return links
		case tt == html.StartTagToken:
			token := z.Token()
			if token.Data == "a" || token.Data == "script" {
				for _, attr := range token.Attr {
					if attr.Key == "href" || attr.Key == "src" {
						links = append(links, attr.Val)
					}
				}
			}
		}
	}
}

func readURLsFromFile(filename string) ([]string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var urls []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		urls = append(urls, strings.TrimSpace(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return urls, nil
}

func getHostname(u string) string {
	parsedURL, err := url.Parse(u)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return ""
	}
	return parsedURL.Hostname()
}

func writeURLsToFile(filename string, urls []string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for _, url := range urls {
		_, err := writer.WriteString(url + "\n")
		if err != nil {
			return err
		}
	}
	writer.Flush()
	return nil
}

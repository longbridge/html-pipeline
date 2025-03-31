package pipeline

import (
	"net/url"

	"github.com/PuerkitoBio/goquery"
)

// ExternalLinkFilter a filter to match external links to add rel="nofollow" target="_blank"
type ExternalLinkFilter struct {
	// IgnoreHosts hosts will ignore
	IgnoreHosts []string
}

// Call render
func (f ExternalLinkFilter) Call(doc *goquery.Document) (err error) {
	doc.Find("a").Each(func(i int, node *goquery.Selection) {
		src := node.AttrOr("href", "")
		// Fix src that not URL Schema
		srcURL, err := url.Parse(src)
		if err != nil {
			return
		}
		if srcURL.Scheme == "" {
			srcURL.Scheme = "https"
			src = srcURL.String()
		}

		if isHost(f.IgnoreHosts, src) {
			return
		}

		node.SetAttr("rel", "nofollow")
		node.SetAttr("target", "_blank")
	})

	return
}

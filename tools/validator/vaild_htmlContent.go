package validator

import (
	md "github.com/JohannesKaufmann/html-to-markdown"
	"github.com/PuerkitoBio/goquery"
	"github.com/russross/blackfriday"
	"gvb/internal/global"
	"strings"
)

func FilterScriptTag(content string) (string, error) {
	unsafe := blackfriday.MarkdownCommon([]byte(content))
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(string(unsafe)))
	if err != nil {
		global.Log.Error(err)
		return "", err
	}
	nodes := doc.Find("script").Nodes
	if len(nodes) > 0 {
		// 有script标签
		doc.Find("script").Remove()
		converter := md.NewConverter("", true, nil)
		html, _ := doc.Html()
		markdown, _ := converter.ConvertString(html)
		return markdown, nil
	}
	return content, nil
}

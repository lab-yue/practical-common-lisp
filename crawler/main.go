package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"sync"
)

var (
	linkReg    = regexp.MustCompile(`href='(.+?)'[\s\S]*?>([\s\S]+?)<`)
	articleReg = regexp.MustCompile(`<P>[\s\S]+</P>`)
)

//SideBar type
type SideBar struct {
	Docs Docs `json:"docs"`
}

//Docs type
type Docs struct {
	Chapters []string
}

// Chapter type
type Chapter struct {
	id      string
	url     string
	title   string
	content []byte
}

func (c *Chapter) writeToFile() {
	filename := `docs/` + c.title + `.md`
	ioutil.WriteFile(filename, c.content, 0644)
}

func main() {
	nav := `<ol>
	<li><a href='introduction-why-lisp.html'>Introduction: Why Lisp?</a></li>
	<li><a href='lather-rinse-repeat-a-tour-of-the-repl.html'>Lather, Rinse, Repeat: A Tour of the REPL</a></li>
	<li><a href='practical-a-simple-database.html'>Practical: A Simple Database</a></li>
	<li><a href='syntax-and-semantics.html'>Syntax and Semantics</a></li>
	<li><a href='functions.html'>Functions</a></li>
	<li><a href='variables.html'>Variables</a></li>
	<li><a href='macros-standard-control-constructs.html'>Macros: Standard Control Constructs</a></li>
	<li><a href='macros-defining-your-own.html'>Macros: Defining Your Own</a></li>
	<li><a href='practical-building-a-unit-test-framework.html'>Practical: Building a Unit Test Framework</a></li>
	<li><a href='numbers-characters-and-strings.html'>Numbers, Characters, and Strings</a></li>
	<li><a href='collections.html'>Collections</a></li>
	<li><a href='they-called-it-lisp-for-a-reason-list-processing.html'>They Called It LISP for a Reason: List Processing</a></li>
	<li><a href='beyond-lists-other-uses-for-cons-cells.html'>Beyond Lists: Other Uses for Cons Cells</a></li>
	<li><a href='files-and-file-io.html'>Files and File I/O</a></li>
	<li><a href='practical-a-portable-pathname-library.html'>Practical: A Portable Pathname Library</a></li>
	<li><a href='object-reorientation-generic-functions.html'>Object Reorientation: Generic Functions</a></li>
	<li><a href='object-reorientation-classes.html'>Object Reorientation: Classes</a></li>
	<li><a href='a-few-format-recipes.html'>A Few FORMAT Recipes</a></li>
	<li><a href='beyond-exception-handling-conditions-and-restarts.html'>Beyond Exception Handling: Conditions and Restarts</a></li>
	<li><a href='the-special-operators.html'>The Special Operators</a></li>
	<li><a href='programming-in-the-large-packages-and-symbols.html'>Programming in the Large: Packages and Symbols</a></li>
	<li><a href='loop-for-black-belts.html'>LOOP for Black Belts</a></li>
	<li><a href='practical-a-spam-filter.html'>Practical: A Spam Filter</a></li>
	<li><a href='practical-parsing-binary-files.html'>Practical: Parsing Binary Files</a></li>
	<li><a href='practical-an-id3-parser.html'>Practical: An ID3 Parser</a></li>
	<li><a href='practical-web-programming-with-allegroserve.html'>Practical: Web Programming with AllegroServe</a></li>
	<li><a href='practical-an-mp3-database.html'>Practical: An MP3 Database</a></li>
	<li><a href='practical-a-shoutcast-server.html'>Practical: A Shoutcast Server</a></li>
	<li><a href='practical-an-mp3-browser.html'>Practical: An MP3 Browser</a></li>
	<li><a href='practical-an-html-generation-library-the-interpreter.html'>Practical: An HTML Generation Library, the Interpreter</a></li>
	<li><a href='practical-an-html-generation-library-the-compiler.html'>Practical: An HTML Generation Library, the Compiler</a></li>
	<li><a href='conclusion-whats-next.html'>Conclusion: What's Next?</a></li>
  </ol>`
	var wg sync.WaitGroup
	client := &http.Client{}
	sidebars := &SideBar{}
	docList := &Docs{}
	links := linkReg.FindAllStringSubmatch(nav, -1)

	for _, link := range links {
		url, title := link[1], link[2]
		id := buildID(title)

		docList.Chapters = append(docList.Chapters, id)

		chapter := Chapter{
			id,
			`http://www.gigamonkeys.com/book/` + url,
			title,
			[]byte{},
		}
		wg.Add(1)
		go func() {
			defer wg.Done()
			fetchChapter(client, &chapter)
		}()
	}

	wg.Wait()
	sidebars.Docs = *docList
	sidebarsJSON, _ := json.Marshal(sidebars)

	err := ioutil.WriteFile("website/sidebars.json", sidebarsJSON, 0644)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Done")
}

func fetchChapter(client *http.Client, chapter *Chapter) error {
	r, err := client.Get(chapter.url)

	if err != nil {
		fmt.Println(err)
		return err
	}
	defer r.Body.Close()

	b, _ := ioutil.ReadAll(r.Body)
	s := articleReg.FindStringSubmatch(string(b))[0]
	chapter.content = []byte(buildMD(chapter, s))
	chapter.writeToFile()
	fmt.Println(`Done ` + chapter.title)
	return nil
}

var (
	aReg    = regexp.MustCompile(`<A[\s\S]+?><H2>([\s\S]+?)</H2></A>`)
	bReg    = regexp.MustCompile(`<B>([\s\S]+?)</B>`)
	iReg    = regexp.MustCompile(`<I>([\s\S]+?)</I>`)
	pReg    = regexp.MustCompile(`<P>([\s\S]+?)</P>`)
	divReg  = regexp.MustCompile(`<DIV[\s\S]+?>`)
	codeReg = regexp.MustCompile(`<CODE>([\s\S]+?)</CODE>`)
	preReg  = regexp.MustCompile(`<PRE>([\s\S]+?)</PRE>`)
)

func replaceHTML(s string) string {

	s = pReg.ReplaceAllString(s, "\n\n$1\n\n")
	s = bReg.ReplaceAllString(s, "$1")
	s = iReg.ReplaceAllString(s, "$1")
	s = codeReg.ReplaceAllString(s, "`$1`")
	s = preReg.ReplaceAllString(s, "\n```lisp\n$1\n```\n")
	s = aReg.ReplaceAllString(s, "\n## $1\n")
	s = divReg.ReplaceAllString(s, "")
	s = strings.Replace(s, "&quot;", `"`, -1)
	s = strings.Replace(s, "&amp;", `&`, -1)
	s = strings.Replace(s, "&lt;", `<`, -1)
	s = strings.Replace(s, "&gt;", `>`, -1)
	return s
}

func buildMD(c *Chapter, article string) string {

	title := fmt.Sprintf(`---
id: %s
title: "%s"
---`,
		buildID(c.title),
		c.title,
	)

	return title + replaceHTML(article)

}

var (
	special = regexp.MustCompile(`[\?\(\)\/\., :]`)
	hyphens = regexp.MustCompile(`--+`)
	tailing = regexp.MustCompile(`-$`)
)

func buildID(s string) string {
	s = strings.ToLower(s)
	s = special.ReplaceAllString(s, "-")
	s = hyphens.ReplaceAllString(s, "-")
	s = tailing.ReplaceAllString(s, "")
	return s
}

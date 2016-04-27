// Package elems contains definition for different html element types and some custom node types

//go:generate go run generate.go

// Documentation source: "HTML element reference" by Mozilla Contributors, https://developer.mozilla.org/en-US/docs/Web/HTML/Element, licensed under CC-BY-SA 2.5.

// Documentation for custom types lies within the  "github.com/influx6/faux/domtrees" package

package elems

import (
	"github.com/influx6/gu/gutrees"
)

// Text provides the concrete implementation for using the domtrees.Text struct
func Text(txt string) *gutrees.Element {
	return gutrees.NewText(txt)
}

// Anchor provides the following for html elements ->
// The HTML Anchor Element (<a>) defines a hyperlink to a location on the same page or any other page on the Web. It can also be used (in an obsolete way) to create an anchor point—a destination for hyperlinks within the content of a page, so that links aren't limited to connecting simply to the top of a page.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/a
func Anchor(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("a", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Abbreviation provides the following for html elements ->
// The HTML <abbr> element (or HTML Abbreviation Element) represents an abbreviation and optionally provides a full description for it. If present, the title attribute must contain this full description and nothing else.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/abbr
func Abbreviation(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("abbr", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Address provides the following for html elements ->
// The HTML <address> element supplies contact information for its nearest <article> or <body> ancestor; in the latter case, it applies to the whole document.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/address
func Address(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("address", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Area provides the following for html elements ->
// The HTML <area> element defines a hot-spot region on an image, and optionally associates it with a hypertext link. This element is used only within a <map> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/area
func Area(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("area", true)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Article provides the following for html elements ->
// The HTML <article> element represents a self-contained composition in a document, page, application, or site, which is intended to be independently distributable or reusable (e.g., in syndication). This could be a forum post, a magazine or newspaper article, a blog entry, an object, or any other independent item of content. Each <article> should be identified, typically by including a heading (<h1>-<h6> element) as a child of the <article> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/article
func Article(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("article", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Aside provides the following for html elements ->
// The HTML <aside> element represents a section of the page with content connected tangentially to the rest, which could be considered separate from that content. These sections are often represented as sidebars or inserts. They often contain the definitions on the sidebars, such as definitions from the glossary; there may also be other types of information, such as related advertisements; the biography of the author; web applications; profile information or related links on the blog.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/aside
func Aside(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("aside", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Audio provides the following for html elements ->
// The HTML <audio> element is used to embed sound content in documents. It may contain one or more audio sources, represented using the src attribute or the <source> element; the browser will choose the most suitable one.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/audio
func Audio(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("audio", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Bold provides the following for html elements ->
// The HTML <b> Element represents a span of text stylistically different from normal text, without conveying any special importance or relevance. It is typically used for keywords in a summary, product names in a review, or other spans of text whose typical presentation would be boldfaced. Another example of its use is to mark the lead sentence of each paragraph of an article.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/b
func Bold(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("b", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Base provides the following for html elements ->
// The HTML <base> element specifies the base URL to use for all relative URLs contained within a document. There can be only one <base> element in a document.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/base
func Base(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("base", true)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// BidirectionalIsolation provides the following for html elements ->
// The HTML <bdi> Element (or Bi-Directional Isolation Element) isolates a span of text that might be formatted in a different direction from other text outside it.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/bdi
func BidirectionalIsolation(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("bdi", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// BidirectionalOverride provides the following for html elements ->
// The HTML <bdo> Element (or HTML bidirectional override element) is used to override the current directionality of text. It causes the directionality of the characters to be ignored in favor of the specified directionality.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/bdo
func BidirectionalOverride(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("bdo", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// BlockQuote provides the following for html elements ->
// The HTML <blockquote> Element (or HTML Block Quotation Element) indicates that the enclosed text is an extended quotation. Usually, this is rendered visually by indentation (see Notes for how to change it). A URL for the source of the quotation may be given using the cite attribute, while a text representation of the source can be given using the <cite> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/blockquote
func BlockQuote(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("blockquote", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Break provides the following for html elements ->
// The HTML element line break <br> produces a line break in text (carriage-return). It is useful for writing a poem or an address, where the division of lines is significant.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/br
func Break(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("br", true)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Button provides the following for html elements ->
// The HTML <button> Element represents a clickable button.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/button
func Button(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("button", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Canvas provides the following for html elements ->
// The HTML <canvas> Element can be used to draw graphics via scripting (usually JavaScript). For example, it can be used to draw graphs, make photo compositions or even perform animations. You may (and should) provide alternate content inside the <canvas> block. That content will be rendered both on older browsers that don't support canvas and in browsers with JavaScript disabled.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/canvas
func Canvas(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("canvas", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Caption provides the following for html elements ->
// The HTML <caption> Element (or HTML Table Caption Element) represents the title of a table. Though it is always the first descendant of a <table>, its styling, using CSS, may place it elsewhere, relative to the table.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/caption
func Caption(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("caption", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Citation provides the following for html elements ->
// The HTML Citation Element (<cite>) represents a reference to a creative work. It must include the title of a work or a URL reference, which may be in an abbreviated form according to the conventions used for the addition of citation metadata.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/cite
func Citation(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("cite", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Code provides the following for html elements ->
// The HTML Code Element (<code>) represents a fragment of computer code. By default, it is displayed in the browser's default monospace font.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/code
func Code(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("code", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Column provides the following for html elements ->
// The HTML Table Column Element (<col>) defines a column within a table and is used for defining common semantics on all common cells. It is generally found within a <colgroup> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/col
func Column(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("col", true)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// ColumnGroup provides the following for html elements ->
// The HTML Table Column Group Element (<colgroup>) defines a group of columns within a table.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/colgroup
func ColumnGroup(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("colgroup", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Data provides the following for html elements ->
// The HTML <data> Element links a given content with a machine-readable translation. If the content is time- or date-related, the <time> must be used.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/data
func Data(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("data", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// DataList provides the following for html elements ->
// The HTML Datalist Element (<datalist>) contains a set of <option> elements that represent the values available for other controls.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/datalist
func DataList(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("datalist", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Description provides the following for html elements ->
// The HTML <dd> element (HTML Description Element) indicates the description of a term in a description list (<dl>) element. This element can occur only as a child element of a description list and it must follow a <dt> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dd
func Description(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("dd", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// DeletedText provides the following for html elements ->
// The HTML Deleted Text Element (<del>) represents a range of text that has been deleted from a document. This element is often (but need not be) rendered with strike-through text.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/del
func DeletedText(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("del", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Details provides the following for html elements ->
// The HTML Details Element (<details>) is used as a disclosure widget from which the user can retrieve additional information.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/details
func Details(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("details", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Definition provides the following for html elements ->
// The HTML Definition Element (<dfn>) represents the defining instance of a term.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dfn
func Definition(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("dfn", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Dialog provides the following for html elements ->
// The HTML <dialog> element represents a dialog box or other interactive component, such as an inspector or window. <form> elements can be integrated within a dialog by specifying them with the attribute method="dialog". When such a form is submitted, the dialog is closed with a returnValue attribute set to the value of the submit button used.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dialog
func Dialog(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("dialog", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Div provides the following for html elements ->
// The HTML <div> element (or HTML Document Division Element) is the generic container for flow content, which does not inherently represent anything. It can be used to group elements for styling purposes (using the class or id attributes), or because they share attribute values, such as lang. It should be used only when no other semantic element (such as <article> or <nav>) is appropriate.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/div
func Div(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("div", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// DescriptionList provides the following for html elements ->
// The HTML <dl> element (or HTML Description List Element) encloses a list of pairs of terms and descriptions. Common uses for this element are to implement a glossary or to display metadata (a list of key-value pairs).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dl
func DescriptionList(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("dl", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// DefinitionTerm provides the following for html elements ->
// The HTML <dt> element (or HTML Definition Term Element) identifies a term in a definition list. This element can occur only as a child element of a <dl>. It is usually followed by a <dd> element; however, multiple <dt> elements in a row indicate several terms that are all defined by the immediate next <dd> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/dt
func DefinitionTerm(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("dt", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Element provides the following for html elements ->
// The HTML <element> element is used to define new custom DOM elements.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/element
func Element(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("element", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Emphasis provides the following for html elements ->
// The HTML element emphasis  <em> marks text that has stress emphasis. The <em> element can be nested, with each level of nesting indicating a greater degree of emphasis.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/em
func Emphasis(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("em", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Embed provides the following for html elements ->
// The HTML <embed> Element represents an integration point for an external application or interactive content (in other words, a plug-in).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/embed
func Embed(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("embed", true)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// FieldSet provides the following for html elements ->
// The HTML <fieldset> element is used to group several controls as well as labels (<label>) within a web form.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/fieldset
func FieldSet(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("fieldset", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// FigureCaption provides the following for html elements ->
// The HTML <figcaption> element represents a caption or a legend associated with a figure or an illustration described by the rest of the data of the <figure> element which is its immediate ancestor which means <figcaption> can be the first or last element inside a <figure> block. Also, the HTML Figcaption Element is optional; if not provided, then the parent figure element will have no caption.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figcaption
func FigureCaption(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("figcaption", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Figure provides the following for html elements ->
// The HTML <figure> element represents self-contained content, frequently with a caption (<figcaption>), and is typically referenced as a single unit. While it is related to the main flow, its position is independent of the main flow. Usually this is an image, an illustration, a diagram, a code snippet, or a schema that is referenced in the main text, but that can be moved to another page or to an appendix without affecting the main flow.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/figure
func Figure(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("figure", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Footer provides the following for html elements ->
// The HTML <footer> element represents a footer for its nearest sectioning content or sectioning root element. A footer typically contains information about the author of the section, copyright data or links to related documents.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/footer
func Footer(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("footer", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Form provides the following for html elements ->
// The HTML <form> element represents a document section that contains interactive controls to submit information to a web server.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/form
func Form(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("form", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Header provides the following for html elements ->
// The HTML <header> element represents a group of introductory or navigational aids. It may contain some heading elements but also other elements like a logo, wrapped section's header, a search form, and so on.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/header
func Header(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("header", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// HeadingsGroup provides the following for html elements ->
// The HTML <hgroup> Element (HTML Headings Group Element) represents the heading of a section. It defines a single title that participates in the outline of the document as the heading of the implicit or explicit section that it belongs to.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/hgroup
func HeadingsGroup(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("hgroup", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// HorizontalRule provides the following for html elements ->
// The HTML <hr> element represents a thematic break between paragraph-level elements (for example, a change of scene in a story, or a shift of topic with a section). In previous versions of HTML, it represented a horizontal rule. It may still be displayed as a horizontal rule in visual browsers, but is now defined in semantic terms, rather than presentational terms.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/hr
func HorizontalRule(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("hr", true)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Italic provides the following for html elements ->
// The HTML <i> Element represents a range of text that is set off from the normal text for some reason, for example, technical terms, foreign language phrases, or fictional character thoughts. It is typically displayed in italic type.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/i
func Italic(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("i", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// InlineFrame provides the following for html elements ->
// The HTML Inline Frame Element (<iframe>) represents a nested browsing context, effectively embedding another HTML page into the current page. In HTML 4.01, a document may contain a head and a body or a head and a frame-set, but not both a body and a frame-set. However, an <iframe> can be used within a normal document body. Each browsing context has its own session history and active document. The browsing context that contains the embedded content is called the parent browsing context. The top-level browsing context (which has no parent) is typically the browser window.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/iframe
func InlineFrame(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("iframe", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Image provides the following for html elements ->
// The HTML <img> element represents an image in the document.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/img
func Image(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("img", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Input provides the following for html elements ->
// The HTML element <input> is used to create interactive controls for web-based forms in order to accept data from the user. How an <input> works varies considerably depending on the value of its type attribute.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/input
func Input(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("input", true)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// InsertedText provides the following for html elements ->
// The HTML <ins> Element (or HTML Inserted Text) HTML represents a range of text that has been added to a document.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ins
func InsertedText(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("ins", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// KeyboardInput provides the following for html elements ->
// The HTML Keyboard Input Element (<kbd>) represents user input and produces an inline element displayed in the browser's default monospace font.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/kbd
func KeyboardInput(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("kbd", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Label provides the following for html elements ->
// The HTML Label Element (<label>) represents a caption for an item in a user interface. It can be associated with a control either by placing the control element inside the <label> element, or by using the for attribute. Such a control is called the labeled control of the label element. One input can be associated with multiple labels.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/label
func Label(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("label", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Legend provides the following for html elements ->
// The HTML <legend> Element (or HTML Legend Field Element) represents a caption for the content of its parent <fieldset>.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/legend
func Legend(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("legend", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// ListItem provides the following for html elements ->
// The HTML <li> element (or HTML List Item Element) is used to represent an item in a list. It must be contained in a parent element: an ordered list (<ol>), an unordered list (<ul>), or a menu (<menu>). In menus and unordered lists, list items are usually displayed using bullet points. In ordered lists, they are usually displayed with an ascending counter on the left, such as a number or letter.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/li
func ListItem(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("li", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Link provides the following for html elements ->
// The HTML <link> element specifies relationships between the current document and an external resource. Possible uses for this element include defining a relational framework for navigation. This Element is most used to link to style sheets.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/link
func Link(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("link", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Main provides the following for html elements ->
// The HTML <main> element represents the main content of  the <body> of a document or application. The main content area consists of content that is directly related to, or expands upon the central topic of a document or the central functionality of an application. This content should be unique to the document, excluding any content that is repeated across a set of documents such as sidebars, navigation links, copyright information, site logos, and search forms (unless the document's main function is as a search form).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/main
func Main(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("main", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Map provides the following for html elements ->
// The HTML <map> element is used with <area> elements to define an image map (a clickable link area).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/map
func Map(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("map", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Mark provides the following for html elements ->
// The HTML Mark Element (<mark>) represents highlighted text, i.e., a run of text marked for reference purpose, due to its relevance in a particular context. For example it can be used in a page showing search results to highlight every instance of the searched-for word.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/mark
func Mark(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("mark", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Menu provides the following for html elements ->
// The HTML <menu> element represents a group of commands that a user can perform or activate. This includes both list menus, which might appear across the top of a screen, as well as context menus, such as those that might appear underneath a button after it has been clicked.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/menu
func Menu(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("menu", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// MenuItem provides the following for html elements ->
// The HTML <menuitem> element represents a command that a user is able to invoke through a popup menu. This includes context menus, as well as menus that might be attached to a menu button.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/menuitem
func MenuItem(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("menuitem", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Meta provides the following for html elements ->
// The HTML <meta> element represents any metadata information that cannot be represented by one of the other HTML meta-related elements (<base>, <link>, <script>, <style> or <title>).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meta
func Meta(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("meta", true)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Meter provides the following for html elements ->
// The HTML <meter> Element represents either a scalar value within a known range or a fractional value.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/meter
func Meter(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("meter", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Navigation provides the following for html elements ->
// The HTML <nav> element (HTML Navigation Element) represents a section of a page that links to other pages or to parts within the page: a section with navigation links.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/nav
func Navigation(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("nav", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// NoFrames provides the following for html elements ->
// <noframes> is an HTML element which is used to supporting browsers which are not able to support <frame> elements or configured to do so.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/noframes
func NoFrames(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("noframes", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// NoScript provides the following for html elements ->
// The HTML <noscript> Element defines a section of html to be inserted if a script type on the page is unsupported or if scripting is currently turned off in the browser.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/noscript
func NoScript(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("noscript", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Object provides the following for html elements ->
// The HTML Embedded Object Element (<object>) represents an external resource, which can be treated as an image, a nested browsing context, or a resource to be handled by a plugin.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/object
func Object(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("object", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// OrderedList provides the following for html elements ->
// The HTML <ol> Element (or HTML Ordered List Element) represents an ordered list of items. Typically, ordered-list items are displayed with a preceding numbering, which can be of any form, like numerals, letters or Romans numerals or even simple bullets. This numbered style is not defined in the HTML description of the page, but in its associated CSS, using the list-style-type property.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ol
func OrderedList(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("ol", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// OptionsGroup provides the following for html elements ->
// In a Web form, the HTML <optgroup> element  creates a grouping of options within a <select> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/optgroup
func OptionsGroup(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("optgroup", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Option provides the following for html elements ->
// In a Web form, the HTML <option> element is used to create a control representing an item within a <select>, an <optgroup> or a <datalist> HTML5 element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/option
func Option(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("option", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Output provides the following for html elements ->
// The HTML <output> element represents the result of a calculation or user action.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/output
func Output(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("output", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Paragraph provides the following for html elements ->
// The HTML <p> element (or HTML Paragraph Element) represents a paragraph of text.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/p
func Paragraph(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("p", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Parameter provides the following for html elements ->
// The HTML <param> Element (or HTML Parameter Element) defines parameters for <object>.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/param
func Parameter(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("param", true)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Picture provides the following for html elements ->
// The HTML <picture> element is a container used to specify multiple <source> elements for a specific <img> contained in it. The browser will choose the most suitable source according to the current layout of the page (the constraints of the box the image will appear in) and the device it will be displayed on (e.g. a normal or hiDPI device.)
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/picture
func Picture(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("picture", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Preformatted provides the following for html elements ->
// The HTML <pre> element (or HTML Preformatted Text) represents preformatted text. Text within this element is typically displayed in a non-proportional ("monospace") font exactly as it is laid out in the file. Whitespace inside this element is displayed as typed.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/pre
func Preformatted(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("pre", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Progress provides the following for html elements ->
// The HTML <progress> Element is used to view the completion progress of a task. While the specifics of how it's displayed is left up to the browser developer, it's typically displayed as a progress bar. Javascript can be used to manipulate the value of progress bar.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/progress
func Progress(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("progress", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Quote provides the following for html elements ->
// The HTML Quote Element (<q>) indicates that the enclosed text is a short inline quotation. This element is intended for short quotations that don't require paragraph breaks; for long quotations use <blockquote> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/q
func Quote(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("q", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// RubyParenthesis provides the following for html elements ->
// The HTML <rp> element is used to provide fall-back parenthesis for browsers non-supporting ruby annotations. Ruby annotations are for showing pronounciation of East Asian characters, like using Japanese furigana or Taiwainese bopomofo characters. The <rp> element is used in the case of lack of <ruby> element support its content has what should be displayed in order to indicate the presence of a ruby annotation, usually parentheses.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rp
func RubyParenthesis(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("rp", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// RubyText provides the following for html elements ->
// The HTML <rt> Element embraces pronunciation of characters presented in a ruby annotations, which are used to describe the pronunciation of East Asian characters. This element is always used inside a <ruby> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rt
func RubyText(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("rt", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Rtc provides the following for html elements ->
// The HTML <rtc> Element embraces semantic annotations of characters presented in a ruby of <rb> elements used inside of <ruby> element. <rb> elements can have both pronunciation (<rt>) and semantic (<rtc>) annotations.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/rtc
func Rtc(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("rtc", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Ruby provides the following for html elements ->
// The HTML <ruby> Element represents a ruby annotation. Ruby annotations are for showing pronunciation of East Asian characters.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ruby
func Ruby(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("ruby", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Strikethrough provides the following for html elements ->
// The HTML Strikethrough Element (<s>) renders text with a strikethrough, or a line through it. Use the <s> element to represent things that are no longer relevant or no longer accurate. However, <s> is not appropriate when indicating document edits; for that, use the <del> and <ins> elements, as appropriate.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/s
func Strikethrough(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("s", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Sample provides the following for html elements ->
// The HTML <samp> element is an element intended to identify sample output from a computer program. It is usually displayed in the browser's default monotype font (such as Lucida Console).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/samp
func Sample(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("samp", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Script provides the following for html elements ->
// The HTML Script Element (<script>) is used to embed or reference an executable script within an HTML or XHTML document.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/script
func Script(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("script", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Section provides the following for html elements ->
// The HTML <section> element represents a generic section of a document, i.e., a thematic grouping of content, typically with a heading. Each <section> should be identified, typically by including a heading (<h1>-<h6> element) as a child of the <section> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/section
func Section(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("section", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Select provides the following for html elements ->
// The HTML select (<select>) element represents a control that presents a menu of options. The options within the menu are represented by <option> elements, which can be grouped by <optgroup> elements. Options can be pre-selected for the user.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/select
func Select(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("select", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Shadow provides the following for html elements ->
// The HTML <shadow> element is used as a shadow DOM insertion point. You might use it if you have created multiple shadow roots under a shadow host. It is not useful in ordinary HTML. It is used with Web Components.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Shadow
func Shadow(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("shadow", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Small provides the following for html elements ->
// The HTML Small Element (<small>) makes the text font size one size smaller (for example, from large to medium, or from small to x-small) down to the browser's minimum font size.  In HTML5, this element is repurposed to represent side-comments and small print, including copyright and legal text, independent of its styled presentation.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/small
func Small(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("small", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Source provides the following for html elements ->
// The HTML <source> element specifies multiple media resources for either the <picture>, the <audio> or the <video> element. It is an empty element. It is commonly used to serve the same media content in multiple formats supported by different browsers.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/source
func Source(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("source", true)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Span provides the following for html elements ->
// The HTML <span> element is a generic inline container for phrasing content, which does not inherently represent anything. It can be used to group elements for styling purposes (using the class or id attributes), or because they share attribute values, such as lang.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/span
func Span(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("span", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Strong provides the following for html elements ->
// The HTML Strong Element (<strong>) gives text strong importance, and is typically displayed in bold.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/strong
func Strong(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("strong", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Style provides the following for html elements ->
// The HTML <style> element contains style information for a document, or part of a document. By default, the style instructions written inside that element are expected to be CSS.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/style
func Style(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("style", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Subscript provides the following for html elements ->
// The HTML Subscript Element (<sub>) defines a span of text that should be displayed, for typographic reasons, lower, and often smaller, than the main span of text.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/sub
func Subscript(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("sub", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Summary provides the following for html elements ->
// The HTML summary element (<summary>) is used as a summary, caption, or legend for the content of a <details> element.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/summary
func Summary(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("summary", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Superscript provides the following for html elements ->
// The HTML Superscript Element (<sup>) defines a span of text that should be displayed, for typographic reasons, higher, and often smaller, than the main span of text.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/sup
func Superscript(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("sup", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Table provides the following for html elements ->
// The HTML Table Element (<table>) represents data in two dimensions or more.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/table
func Table(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("table", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// TableBody provides the following for html elements ->
// The HTML Table Body Element (<tbody>) defines one or more <tr> element data-rows to be the body of its parent <table> element (as long as no <tr> elements are immediate children of that table element.)  In conjunction with a preceding <thead> and/or <tfoot> element, <tbody> provides additional semantic information for devices such as printers and displays. Of the parent table's child elements, <tbody> represents the content which, when longer than a page, will most likely differ for each page printed; while the content of <thead> and <tfoot> will be the same or similar for each page printed. For displays, <tbody> will enable separate scrolling of the <thead>, <tfoot>, and <caption> elements of the same parent <table> element.  Note that unlike the <thead>, <tfoot>, and <caption> elements however, multiple <tbody> elements are permitted (if consecutive), allowing the data-rows in long tables to be divided into different sections, each separately formatted as needed.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tbody
func TableBody(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("tbody", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// TableData provides the following for html elements ->
// The Table cell HTML element (<td>) defines a cell of a table that contains data. It participates in the table model.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/td
func TableData(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("td", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Template provides the following for html elements ->
// The HTML template element <template> is a mechanism for holding client-side content that is not to be rendered when a page is loaded but may subsequently be instantiated during runtime using JavaScript.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/template
func Template(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("template", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// TextArea provides the following for html elements ->
// The HTML <textarea> element represents a multi-line plain-text editing control.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/textarea
func TextArea(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("textarea", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// TableFoot provides the following for html elements ->
// The HTML Table Foot Element (<tfoot>) defines a set of rows summarizing the columns of the table.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tfoot
func TableFoot(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("tfoot", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// TableHeader provides the following for html elements ->
// The HTML element table header cell <th> defines a cell as a header for a group of cells of a table. The group of cells that the header refers to is defined by the scope and headers attribute.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/th
func TableHeader(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("th", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// TableHead provides the following for html elements ->
// The HTML Table Head Element (<thead>) defines a set of rows defining the head of the columns of the table.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/thead
func TableHead(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("thead", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Time provides the following for html elements ->
// The HTML <time> element represents either a time on a 24-hour clock or a precise date in the Gregorian calendar (with optional time and timezone information).
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/time
func Time(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("time", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Title provides the following for html elements ->
// The HTML <title> element defines the title of the document, shown in a browser's title bar or on the page's tab. It can only contain text, and any contained tags are ignored.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/title
func Title(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("title", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// TableRow provides the following for html elements ->
// The HTML element table row <tr> defines a row of cells in a table. Those can be a mix of <td> and <th> elements.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/tr
func TableRow(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("tr", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Track provides the following for html elements ->
// The HTML <track> element is used as a child of the media elements—<audio> and <video>. It lets you specify timed text tracks (or time-based data), for example to automatically handle subtitles. The tracks are formatted in WebVTT format (.vtt files) — Web Video Text Tracks.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/track
func Track(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("track", true)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Underline provides the following for html elements ->
// The HTML Underline Element (<u>) renders text with an underline, a line under the baseline of its content.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/u
func Underline(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("u", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// UnorderedList provides the following for html elements ->
// The HTML <ul> element (or HTML Unordered List Element) represents an unordered list of items, namely a collection of items that do not have a numerical ordering, and their order in the list is meaningless. Typically, unordered-list items are displayed with a bullet, which can be of several forms, like a dot, a circle or a squared. The bullet style is not defined in the HTML description of the page, but in its associated CSS, using the list-style-type property.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/ul
func UnorderedList(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("ul", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Variable provides the following for html elements ->
// The HTML Variable Element (<var>) represents a variable in a mathematical expression or a programming context.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/var
func Variable(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("var", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Video provides the following for html elements ->
// Use the  HTML <video> element to embed video content in a document. The video element contains one or more video sources. To specify a video source, use either the src attribute or the <source> element; the browser will choose the most suitable one.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/video
func Video(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("video", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// WordBreakOpportunity provides the following for html elements ->
// The HTML element word break opportunity <wbr> represents a position within text where the browser may optionally break a line, though its line-breaking rules would not otherwise create a break at that location.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/wbr
func WordBreakOpportunity(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("wbr", true)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Header1 provides the following for html elements ->
// Heading elements implement six levels of document headings, <h1> is the most important and <h6> is the least. A heading element briefly describes the topic of the section it introduces. Heading information may be used by user agents, for example, to construct a table of contents for a document automatically.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header1(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("h1", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Header2 provides the following for html elements ->
// Heading elements implement six levels of document headings, <h1> is the most important and <h6> is the least. A heading element briefly describes the topic of the section it introduces. Heading information may be used by user agents, for example, to construct a table of contents for a document automatically.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header2(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("h2", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Header3 provides the following for html elements ->
// Heading elements implement six levels of document headings, <h1> is the most important and <h6> is the least. A heading element briefly describes the topic of the section it introduces. Heading information may be used by user agents, for example, to construct a table of contents for a document automatically.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header3(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("h3", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Header4 provides the following for html elements ->
// Heading elements implement six levels of document headings, <h1> is the most important and <h6> is the least. A heading element briefly describes the topic of the section it introduces. Heading information may be used by user agents, for example, to construct a table of contents for a document automatically.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header4(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("h4", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Header5 provides the following for html elements ->
// Heading elements implement six levels of document headings, <h1> is the most important and <h6> is the least. A heading element briefly describes the topic of the section it introduces. Heading information may be used by user agents, for example, to construct a table of contents for a document automatically.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header5(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("h5", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

// Header6 provides the following for html elements ->
// Heading elements implement six levels of document headings, <h1> is the most important and <h6> is the least. A heading element briefly describes the topic of the section it introduces. Heading information may be used by user agents, for example, to construct a table of contents for a document automatically.
// https://developer.mozilla.org/en-US/docs/Web/HTML/Element/Heading_Elements
func Header6(markup ...gutrees.Appliable) *gutrees.Element {
	e := gutrees.NewElement("h6", false)
	for _, m := range markup {
		m.Apply(e)
	}
	return e
}

package highlight

import (
	"bytes"
	"io"

	"github.com/fatih/color"
)

// Term highlights a piece of code for rendering in the terminal.
func Term(lang, code string) (string, error) {
	h, err := makeAndHighlight(lang, code)
	if err != nil {
		return "", err
	}
	return h.renderTerm()
}

func (h *highlighter) renderTerm() (string, error) {
	var buf bytes.Buffer
	h.render(&buf, func(w io.Writer, class string, body []byte) {
		color, ok := termColors[class]
		if !ok {
			color = resetColor
		}
		w.Write([]byte(color("%s", body)))
	})
	return buf.String(), nil
}

// Theme borrowed from Felix Frederick Becker
// See THEME-LICENSE.txt
// https://github.com/felixfbecker/cli-highlight/blob/master/src/theme.ts

var resetColor = color.New(color.Reset).SprintfFunc()
var greyColor = color.New(color.FgBlack, color.Faint).SprintfFunc()
var termColors = map[string]func(format string, a ...interface{}) string{

	/**
	 * keyword in a regular Algol-style language
	 */
	"keyword": color.BlueString,

	/**
	 * built-in or library object (constant, class, function)
	 */
	"built_in": color.CyanString,

	/**
	 * user-defined type in a language with first-class syntactically significant types, like
	 * Haskell
	 */
	"type": color.New(color.FgCyan, color.Faint).SprintfFunc(),

	/**
	 * special identifier for a built-in value ("true", "false", "null")
	 */
	"literal": color.BlueString,

	/**
	 * number, including units and modifiers, if any.
	 */
	"number": color.GreenString,

	/**
	 * literal regular expression
	 */
	"regexp": color.RedString,

	/**
	 * literal string, character
	 */
	"string": color.RedString,

	/**
	 * parsed section inside a literal string
	 */
	"subst": resetColor,

	/**
	 * symbolic constant, interned string, goto label
	 */
	"symbol": resetColor,

	/**
	 * class or class-level declaration (interfaces, traits, modules, etc)
	 */
	"class": color.BlueString,

	/**
	 * function or method declaration
	 */
	"function": color.YellowString,

	/**
	 * name of a class or a function at the place of declaration
	 */
	"title": resetColor,

	/**
	 * block of function arguments (parameters) at the place of declaration
	 */
	"params": resetColor,

	/**
	 * comment
	 */
	"comment": color.GreenString,

	/**
	 * documentation markup within comments
	 */
	"doctag": color.GreenString,

	/**
	 * flags, modifiers, annotations, processing instructions, preprocessor directive, etc
	 */
	"meta": greyColor,

	/**
	 * keyword or built-in within meta construct
	 */
	"meta-keyword": resetColor,

	/**
	 * string within meta construct
	 */
	"meta-string": resetColor,

	/**
	 * heading of a section in a config file, heading in text markup
	 */
	"section": resetColor,

	/**
	 * XML/HTML tag
	 */
	"tag ": greyColor,

	/**
	 * name of an XML tag, the first word in an s-expression
	 */
	"name": color.BlueString,

	/**
	 * s-expression name from the language standard library
	 */
	"builtin-name": resetColor,

	/**
	 * name of an attribute with no language defined semantics (keys in JSON, setting names in
	 * .ini), also sub-attribute within another highlighted object, like XML tag
	 */
	"attr": color.CyanString,

	/**
	 * name of an attribute followed by a structured value part, like CSS properties
	 */
	"attribute": resetColor,

	/**
	 * variable in a config or a template file, environment var expansion in a script
	 */
	"variable": resetColor,

	/**
	 * list item bullet in text markup
	 */
	"bullet": resetColor,

	/**
	 * code block in text markup
	 */
	"code": resetColor,

	/**
	 * emphasis in text markup
	 */
	"emphasis": color.New(color.Italic).SprintfFunc(),

	/**
	 * strong emphasis in text markup
	 */
	"strong": color.New(color.Bold).SprintfFunc(),

	/**
	 * mathematical formula in text markup
	 */
	"formula": resetColor,

	/**
	 * hyperlink in text markup
	 */
	"link": color.New(color.Underline).SprintfFunc(),

	/**
	 * quotation in text markup
	 */
	"quote": resetColor,

	/**
	 * tag selector in CSS
	 */
	"selector-tag": resetColor,

	/**
	 * #id selector in CSS
	 */
	"selector-id": resetColor,

	/**
	 * .class selector in CSS
	 */
	"selector-class": resetColor,

	/**
	 * [attr] selector in CSS
	 */
	"selector-attr": resetColor,

	/**
	 * :pseudo selector in CSS
	 */
	"selector-pseudo": resetColor,

	/**
	 * tag of a template language
	 */
	"template-tag": resetColor,

	/**
	 * variable in a template language
	 */
	"template-variable": resetColor,

	/**
	 * added or changed line in a diff
	 */
	"addition": color.GreenString,

	/**
	 * deleted line in a diff
	 */
	"deletion": color.RedString,
}

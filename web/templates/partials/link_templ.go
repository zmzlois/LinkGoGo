// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package partials

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

import (
	"fmt"
	"strings"
)

func ImageTranform(image string, path templ.SafeURL, name string) string {

	pathString := string(path)
	if image == " " {
		if strings.Contains(pathString, "twitter") || strings.Contains(pathString, "x.com") {
			return (fmt.Sprintf(`<img src="%s" class="w-14 h-14 rounded" alt="avatar"/>`, "/dist/twitter.svg"))
		} else if strings.Contains(pathString, "github") {
			return "/dist/github.svg"
		} else if strings.Contains(pathString, "instagram") {
			return "/dist/instagram.svg"
		}
		return string(name[0])
	} else {
		return image
	}
}

func Link(image string, href templ.SafeURL, text string, isLogIn bool) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, templ_7745c5c3_W io.Writer) (templ_7745c5c3_Err error) {
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templ_7745c5c3_W.(*bytes.Buffer)
		if !templ_7745c5c3_IsBuffer {
			templ_7745c5c3_Buffer = templ.GetBuffer()
			defer templ.ReleaseBuffer(templ_7745c5c3_Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<style type=\"text/css\">\n        .w-full {\n            width: 100%;\n        } \n        .p-1 {\n            padding: 0.25rem;\n        }\n        .px-3 {\n            padding-left: 0.75rem;\n            padding-right: 0.75rem;\n        }\n        .p-2 {\n            padding: 0.5rem;\n        }\n        .rounded {\n            border-radius: 0.25rem;\n        }\n        .bg-zinc-100 {\n            background-color: #f5f7f9;\n        }\n        .drop-shadow-sm {\n            box-shadow: 0 1px 2px 0 rgba(0, 0, 0, 0.05);\n        }\n        .content-center {\n            align-content: center;\n        }\n        .items-center {\n            align-items: center;\n        }\n        .justify-center {\n            justify-content: center;\n        }\n        .flex {\n            display: flex;\n        }\n        .justify-between {\n            justify-content: space-between;\n        }\n        .animate-pop{\n            transition-timing-function: ease-in;\n           transition: all 300ms cubic-bezier(.47,1.64,.41,.8);\n\t        &:hover {\n\t\t        transform: scale(1.04);\n\t        } \n        }\n\n        .hidden {\n            display: none;\n        }\n        .group {\n            position: relative;\n            &:hover {\n                .group-hover\\:block {\n                    z-index: 10;\n                    display: block;\n                }\n                .group-hover\\:bg-zinc-300 {\n                    background-color: #e5e7eb;\n                }\n            }\n        }\n        .h-14 {\n           height: 4rem;\n        }\n        .w-14 {\n            width: 4rem;\n        }\n        .h-10 {\n            height: 2.5rem;\n        }\n        .w-10 {\n            width: 2.5rem;\n        }\n        .h-6 {\n            height: 1rem;\n        }\n        .w-6 {\n            width: 1rem;\n        }\n        .rounded {\n            border-radius: 0.25rem;\n        }\n        .rounded-full {\n            border-radius: 9999px;\n        }\n        .bg-inherit {\n            background-color: inherit;\n        }\n        .transition-colors {\n            transition-property: background-color, border-color, color, fill, stroke;\n            transition-timing-function: cubic-bezier(0.4, 0, 0.2, 1);\n            transition-duration: 150ms;\n        }\n\n        .bg-zinc-900 {\n            background-color: #111827;\n        }\n        .text-zinc-100 {\n            color: #f9fafb;\n        }\n        .border-white {\n            border-color:1px solid #fff;\n        }\n        \n\n    </style><a href=\"")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var2 templ.SafeURL = href
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(string(templ_7745c5c3_Var2)))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("\" class=\"w-full p-2 rounded bg-zinc-100 drop-shadow-sm content-center items-center flex group justify-between animate-pop\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if strings.Contains(string(href), "twitter") || strings.Contains(string(href), "x.com") {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<img src=\"/dist/twitter.svg\" class=\"w-14 h-14 rounded\" alt=\"avatar\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else if strings.Contains(string(href), "github") {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<img src=\"/dist/github.svg\" class=\"w-14 h-14 rounded\" alt=\"avatar\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else if image != " " && strings.Contains(string(href), "instagram") {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<img src=\"/dist/instagram.svg\" class=\"w-14 h-14 rounded\" alt=\"avatar\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-14 h-14 flex p-1 bg-zinc-900 rounded text-zinc-100  items-center justify-center\"><p class=\"text-zinc-100 font-semibold text-3xl\">")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			var templ_7745c5c3_Var3 string
			templ_7745c5c3_Var3, templ_7745c5c3_Err = templ.JoinStringErrs(ImageTranform(image, href, text))
			if templ_7745c5c3_Err != nil {
				return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/templates/partials/link.templ`, Line: 140, Col: 86}
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var3))
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		var templ_7745c5c3_Var4 string
		templ_7745c5c3_Var4, templ_7745c5c3_Err = templ.JoinStringErrs(text)
		if templ_7745c5c3_Err != nil {
			return templ.Error{Err: templ_7745c5c3_Err, FileName: `web/templates/partials/link.templ`, Line: 143, Col: 11}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString(templ.EscapeString(templ_7745c5c3_Var4))
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</p>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if isLogIn {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-10 h-10 p-2 flex rounded-full bg-inherit group-hover:bg-zinc-300 items-center transition-colors content-center\"><img src=\"/dist/dot-horizontal.svg\" class=\"group-hover:block hidden w-6 h-6\"></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		} else {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"w-10 h-10 p-2 flex  items-center  content-center\"></div>")
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</a>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}

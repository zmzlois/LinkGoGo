// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.543
package partials

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

type LinkType struct {
	Image string
	Path  templ.SafeURL
	Name  string
}

var links = []LinkType{
	{Image: " ", Path: "https://loiszhao.com", Name: "Website"},
	{Image: " ", Path: "https://twitter.com/zmzlois", Name: "Twitter"},
	{Image: " ", Path: "https://github.com/zmzlois", Name: "Github"},
	{Image: " ", Path: "https://github.com/zmzlois/LinkGoGo", Name: "The repo made this shit"},
}

func LinkSection(isLogIn bool) templ.Component {
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
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<div class=\"flex flex-col  items-center  py-8\"><div class=\"lg:w-3/5 w-11/12 justify-between items-center flex flex-col gap-8\">")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		for _, link := range links {
			templ_7745c5c3_Err = Link(link.Image, link.Path, link.Name, isLogIn).Render(ctx, templ_7745c5c3_Buffer)
			if templ_7745c5c3_Err != nil {
				return templ_7745c5c3_Err
			}
		}
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("</div></div>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		if !templ_7745c5c3_IsBuffer {
			_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteTo(templ_7745c5c3_W)
		}
		return templ_7745c5c3_Err
	})
}
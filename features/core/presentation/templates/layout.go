package templates

import data_repos "github.com/SabianF/gopher-web/features/core/data/repos"

type LayoutData struct {
	Title           string
	MetaDescription string
	PageHeader      string
	PageBody        string
	PageFooter      string
}

func Layout(data LayoutData) string {
	return data_repos.ParseFiles(data, data_repos.GetRootPath() + "/features/core/presentation/templates/layout.html")
}

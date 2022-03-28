package appmodule

// import (
// 	"fmt"
// 	"strings"

// 	"github.com/nurcahyaari/kite/lib/impl"
// 	"github.com/nurcahyaari/kite/templates"
// 	"github.com/nurcahyaari/kite/utils/fs"
// )

// type RepositoryOption struct {
// 	impl.KiteOptions
// 	ModuleName           string
// 	ModulePath           string
// 	ModuleRepositoryPath string
// }

// func NewRepository(opt RepositoryOption) (impl.AppGenerator, error) {
// 	repositoryPath := fs.ConcatDirPath(opt.ModulePath, "repository")
// 	return RepositoryOption{
// 		KiteOptions:          opt.KiteOptions,
// 		ModuleName:           opt.ModuleName,
// 		ModulePath:           opt.ModulePath,
// 		ModuleRepositoryPath: repositoryPath,
// 	}, nil
// }

// func (o RepositoryOption) Run() error {
// 	err := o.createRepositoryPath()
// 	if err != nil {
// 		return err
// 	}

// 	err = o.createBaseFile()
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (o RepositoryOption) createRepositoryPath() error {
// 	err := fs.CreateFolderIsNotExist(o.ModuleRepositoryPath)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

// func (o RepositoryOption) createBaseFile() error {
// 	tmpl := templates.NewTemplate(templates.Template{
// 		PackageName: "repository",
// 		Import: []templates.ImportedPackage{
// 			{
// 				FilePath: fs.ConcatDirPath(o.GoModName, "infrastructure"),
// 			},
// 		},
// 		IsDependency: true,
// 		Dependency: templates.Dependency{
// 			HaveInterface:  true,
// 			DependencyName: fmt.Sprintf("%sRepository", strings.Title(o.ModuleName)),
// 			FuncParams: []templates.DependencyFuncParam{
// 				{
// 					ParamName:     "db",
// 					ParamDataType: "*infrastructure.MysqlImpl",
// 				},
// 			},
// 			DependencyMethod: []templates.DependencyMethod{},
// 		},
// 	})

// 	templateString, err := tmpl.Render()
// 	if err != nil {
// 		return err
// 	}

// 	err = fs.CreateFileIfNotExist(o.ModuleRepositoryPath, "repository.go", templateString)
// 	if err != nil {
// 		return err
// 	}

// 	return nil
// }

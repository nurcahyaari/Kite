package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/nurcahyaari/kite/library/generator"
	"github.com/nurcahyaari/kite/library/generator/appmodule"
	"github.com/nurcahyaari/kite/library/generator/misc"

	"github.com/nurcahyaari/kite/utils"
	"github.com/nurcahyaari/kite/utils/logger"
	cli "github.com/urfave/cli/v2"
)

func main() {
	var err error
	app := cli.NewApp()
	app.Name = "kite"
	app.Description = "Projects Generator for Golang inspired by Clean Code Arch"

	// path, err := os.Getwd()
	// if err != nil {
	// 	logger.Errorln(err)
	// 	return
	// }

	app.Commands = []*cli.Command{
		{
			Name:        "new",
			Description: "make new Apps",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Value: "kite",
					Usage: "creating new Apps with name",
				},
			},
			Action: func(c *cli.Context) error {
				var option generator.GeneratorOptions
				var newAppOption generator.NewAppOption

				option.GoModName = c.String("name")
				path := "/Users/nurcahyaari/Documents/projects/tools/testkite/"
				option.Path = path
				option.ProjectPath = path + option.GoModName
				option.DefaultDBDialeg = "mysql"

				splitPath := strings.Split(option.GoModName, "/")
				option.AppName = utils.CapitalizeFirstLetter(splitPath[len(splitPath)-1])

				newAppOption.GeneratorOptions = option

				newApp := generator.NewApp(newAppOption)
				err := newApp.Run()
				if err != nil {
					logger.Errorln(err)
				}

				return err
			},
		},
		{
			Name:        "modules",
			Description: "Make a new module",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Value: "",
					Usage: "module name",
				},
			},
			Action: func(c *cli.Context) error {
				var err error
				var option generator.GeneratorOptions
				var moduleOption appmodule.ModulesOption

				path := "/Users/nurcahyaari/Documents/projects/tools/testkite/test1/"
				option.Path = path
				option.ProjectPath = path

				moduleOption.Options = option
				moduleOption.IsModule = true
				moduleOption.ModuleName = c.String("name")
				if moduleOption.ModuleName == "" {
					err = fmt.Errorf("module name cannot be empty")
					logger.Errorln(err.Error())
					return err
				}

				module := appmodule.NewModules(moduleOption)
				module.Run()

				return nil
			},
		},
		{
			Name:        "middleware",
			Description: "Make a new middleware",
			Flags: []cli.Flag{
				&cli.StringFlag{
					Name:  "name",
					Value: "",
					Usage: "middleware name",
				},
			},
			Action: func(c *cli.Context) error {
				var option generator.GeneratorOptions
				var middlewareOption misc.MiddlewareOption
				path := "/Users/nurcahyaari/Documents/projects/tools/testkite/test1/"

				option.Path = path
				option.ProjectPath = path

				middlewareOption.Options = option
				middlewareOption.MiddlewareName = c.String("name")
				middlewareOption.InternalPath = path + "internal/"

				middleware := misc.NewMiddleware(middlewareOption)
				middleware.Run()

				return nil
			},
		},
	}

	err = app.Run(os.Args)
	if err != nil {
		fmt.Errorf("Error ", err)
	}
}

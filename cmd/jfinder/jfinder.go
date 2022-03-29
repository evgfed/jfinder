package main

import (
	"bytes"
	"fmt"
	"github.com/gookit/config/v2"
	"github.com/gookit/config/v2/yaml"
	"os"
)

func main() {

	/*
		// open a new index
		mapping := bleve.NewIndexMapping()
		//index, err := bleve.New("1index.bleve", mapping)
		index, err := bleve.Open("1index.bleve", mapping)
		if err != nil {
			fmt.Println(err)
			return
		}

		data := struct {
			Name string
		}{
			Name: "text",
		}

		// index some data
		index.Index("id", data)

		// search for some text
		query := bleve.NewMatchQuery("text")
		search := bleve.NewSearchRequest(query)
		searchResults, err := index.Search(search)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Println(searchResults)
	*/

	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	fmt.Println("Current dir is ", pwd)

	config.WithOptions(config.ParseEnv)

	// add driver for support yaml content
	config.AddDriver(yaml.Driver)
	confFile1 := pwd + "/configs/jfinder_config.yaml"
	confFile2 := pwd + "/configs/adding_config.yaml"
	fmt.Println("Load config at:", confFile1)
	err = config.LoadFiles(confFile1, confFile2)
	if err != nil {
		panic(err)
	}

	fmt.Printf("config data: \n %#v\n", config.Data())
	app := config.String("name")
	debug := config.Bool("debug")
	fmt.Printf("app:%v name: %v\n", app, debug)

	// load env config
	config.LoadOSEnv([]string{"APP_NAME", "APP_DEBUG"}, true)

	// set value
	config.Set("name", "new name")
	name := config.String("name")
	fmt.Printf("- set string\n  val: %v\n", name)

	// read
	debugEnv := config.Bool("APP_DEBUG") // true
	appEnv := config.String("app_name")  // "config"
	fmt.Printf("appEnv:%v nameEnv: %v\n", appEnv, debugEnv)
	os.Setenv("APP_NAME", "Test")
	for _, s := range os.Environ() {
		fmt.Println("-->", s)
	}

	// if you want export config data
	buf := new(bytes.Buffer)
	_, err = config.DumpTo(buf, config.Yaml)
	if err != nil {
		panic(err)
	}
	fmt.Printf("export config:\n%s", buf.String())

	//	ctx := context.Background()
	/*
		cfg := config.NewConfig(ctx, "testConfig.yaml")
		cfg.SetValue(ctx, "val1", "data1", config.EmTypeNone)
		fmt.Println(cfg.GetValue(ctx, "val1"))
	*/
}

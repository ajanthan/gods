package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"sort"

	"github.com/ajanthan/gods/pkg/conf"
	"github.com/ajanthan/gods/pkg/db"
	"github.com/ajanthan/gods/pkg/model"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
)

type byIndex []model.Params

func (b byIndex) Len() int           { return len(b) }
func (b byIndex) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }
func (b byIndex) Less(i, j int) bool { return b[i].Ordinal < b[j].Ordinal }
func requestHandler(repo *db.Repository, query model.Query) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		log.Println("Handling ", req.URL.Path, "Method ", req.Method)
		sql := query.SQL
		var args []interface{}
		//repo.DB.Query()
		params := query.Params
		sort.Sort(byIndex(params))
		for _, param := range params {
			arg := mux.Vars(req)[param.Name]
			switch param.SQLType {
			case "string":
				args = append(args, fmt.Sprint(arg))
			case "int":
				//args = append(args, strconv.ParseInt(arg, 10, 64))
			}

		}
		log.Println("SQL: ", sql, "Args: ", args)
		log.Println("DB: ", repo.DB)
		rows, dbErr := repo.DB.Query(sql, args...)
		defer rows.Close()
		if dbErr != nil {
			log.Fatal(dbErr)
		}
		results := make([]interface{}, len(query.Result.Fields))
		resultPtrs := make([]interface{}, len(query.Result.Fields))

		for i := range query.Result.Fields {
			results[i] = ""
			resultPtrs[i] = &results[i]
		}
		res.Header().Set("Content-Type", "application/json")
		switch query.Result.Type {
		case "object":
			results := make([]interface{}, len(query.Result.Fields))
			resultPtrs := make([]interface{}, len(query.Result.Fields))

			for i := range query.Result.Fields {
				results[i] = ""
				resultPtrs[i] = &results[i]
			}
			result := make(map[string]interface{}, len(query.Result.Fields))
			log.Println("Inside object Result")
			for rows.Next() {

				scnErr := rows.Scan(resultPtrs...)
				log.Println("Scanning row for object")
				if scnErr != nil {
					log.Fatal(scnErr)
				}
			}
			for i, field := range query.Result.Fields {
				switch results[i].(type) {
				case string:
					result[field.Name] = results[i].(string)
				case int64:
					result[field.Name] = results[i].(int64)
				case []byte:
					result[field.Name] = string(results[i].([]byte))
				default:
					log.Println("Assiging default for ", field.Column)
					log.Printf("Output type %T", results[i])
					result[field.Name] = results[i]
				}

			}
			encoder := json.NewEncoder(res)
			encoErr := encoder.Encode(result)

			if encoErr != nil {
				log.Fatal(encoErr)
			}
		case "array":
			results := make([]interface{}, len(query.Result.Fields))
			resultPtrs := make([]interface{}, len(query.Result.Fields))

			for i := range query.Result.Fields {
				results[i] = ""
				resultPtrs[i] = &results[i]
			}
			var arrayOfMap []map[string]interface{}
			for rows.Next() {
				result := make(map[string]interface{}, len(query.Result.Fields))
				scnErr := rows.Scan(resultPtrs...)
				log.Println("Scanning row for objects in array")
				if scnErr != nil {
					log.Fatal(scnErr)
				}

				for i, field := range query.Result.Fields {
					switch results[i].(type) {
					case string:
						result[field.Name] = results[i].(string)
					case int64:
						result[field.Name] = results[i].(int64)
					case []byte:
						result[field.Name] = string(results[i].([]byte))
					default:
						log.Println("Assiging default for ", field.Column)
						log.Printf("Output type %T", results[i])
						result[field.Name] = results[i]
					}

				}
				arrayOfMap = append(arrayOfMap, result)
			}
			encoder := json.NewEncoder(res)
			encoErr := encoder.Encode(arrayOfMap)

			if encoErr != nil {
				log.Fatal(encoErr)
			}
		}

	}
}
func main() {
	godsConf, err := conf.LoadGoDs("../spec/sample/sample-1.yaml")
	if err != nil {
		log.Fatal(err)
	}
	repository := &db.Repository{}
	dbErr := repository.Init(godsConf.Datasource.Type, godsConf.Datasource.Url)
	if dbErr != nil {
		log.Fatal(dbErr)
	}
	log.Println("DB in main: ", repository.DB)
	router := mux.NewRouter()
	//requestHandler := func(res http.ResponseWriter, req *http.Request) {}
	for _, resource := range godsConf.Resources {
		path := resource.Path
		method := resource.Method
		router.HandleFunc(path, requestHandler(repository, resource.Query)).Methods(method)
	}

	log.Fatal(http.ListenAndServe(":9090", router))

}

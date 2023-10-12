package main

import (
    "encoding/json"
    "fmt"
    "github.com/jedib0t/go-pretty/v6/table"
    "io/ioutil"
    "net/http"
    "os"
    "strings"
)

func main() {
    if len(os.Args) < 2 {
        printUsage()
        return
    }

    if len(os.Args) == 2 && (os.Args[1] == "--help" || os.Args[1] == "-h") {
        printUsage()
        return
    }

    if len(os.Args) == 2 && (os.Args[1] == "--version" || os.Args[1] == "-v") {
        printVersion()
        return
    }

    if len(os.Args) == 3 && os.Args[1] == "--read" {
        jsonFileName := os.Args[2]
        if strings.HasSuffix(jsonFileName, ".json") {
            readTable(jsonFileName)
        } else {
            fmt.Println("Error: The specified file is not a JSON file.")
        }
        return
    }

    if len(os.Args) == 3 && os.Args[1] == "--URL" {
        jsonURL := os.Args[2]
        processURL(jsonURL)
        return
    }

    jsonFileName := os.Args[1]
    displayTable(jsonFileName)
}

func displayTable(jsonFileName string) {
    jsonData, err := os.ReadFile(jsonFileName)
    if err != nil {
        fmt.Println("Error reading JSON file:", err)
        return
    }

    var jsonDataSlice []map[string]interface{}
    err = json.Unmarshal(jsonData, &jsonDataSlice)
    if err == nil {
        displayJSONArray(jsonDataSlice)
    } else {
        jsonDataMap := make(map[string]interface{})
        err = json.Unmarshal(jsonData, &jsonDataMap)
        if err != nil {
            fmt.Println("Error unmarshaling JSON data:", err)
            return
        }
        displayJSONObject(jsonDataMap)
    }
}

func displayJSONArray(data []map[string]interface{}) {
    t := table.NewWriter()
    t.SetOutputMirror(os.Stdout)

    var header []interface{}
    if len(data) > 0 {
        for key := range data[0] {
            header = append(header, interface{}(key))
        }
        t.AppendHeader(header)
    }

    for _, row := range data {
        var rowData []interface{}
        for _, colName := range header {
            colNameStr, _ := colName.(string)
            cellValue := row[colNameStr]

            if _, isNumber := cellValue.(float64); isNumber {
                cellValue = fmt.Sprintf("\x1b[94m%.2f\x1b[0m", cellValue)
            }

            rowData = append(rowData, cellValue)
        }
        t.AppendRow(rowData)
    }

    t.SetStyle(table.StyleLight)
    t.Render()
}

func displayJSONObject(data map[string]interface{}) {
    t := table.NewWriter()
    t.SetOutputMirror(os.Stdout)

    var header []interface{}
    for key := range data {
        header = append(header, interface{}(key))
    }
    t.AppendHeader(header)

    var rowData []interface{}
    for _, colName := range header {
        colNameStr, _ := colName.(string)
        rowData = append(rowData, data[colNameStr])
    }
    t.AppendRow(rowData)

    t.SetStyle(table.StyleLight)
    t.Render()
}

func readTable(jsonFileName string) {
    jsonData, err := os.ReadFile(jsonFileName)
    if err != nil {
        fmt.Println("Error reading JSON file:", err)
        return
    }

    fmt.Println(string(jsonData))
}

func processURL(jsonURL string) {
    resp, err := http.Get(jsonURL)
    if err != nil {
        fmt.Println("Error fetching data from URL:", err)
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        fmt.Printf("HTTP request failed with status code: %d\n", resp.StatusCode)
        return
    }

    jsonData, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading data from URL:", err)
        return
    }

    displayTableFromJSON(jsonData)
}

func displayTableFromJSON(jsonData []byte) {
    var jsonDataSlice []map[string]interface{}
    err := json.Unmarshal(jsonData, &jsonDataSlice)
    if err == nil {
        displayJSONArray(jsonDataSlice)
    } else {
        jsonDataMap := make(map[string]interface{})
        err = json.Unmarshal(jsonData, &jsonDataMap)
        if err != nil {
            fmt.Println("Error unmarshaling JSON data:", err)
            return
        }
        displayJSONObject(jsonDataMap)
    }
}

func printUsage() {
    fmt.Println("Usage:")
    fmt.Println("To display a table: jtab <JSON file>")
    fmt.Println("To read the JSON content: jtab --read <JSON file>")
    fmt.Println("To fetch and display JSON data from a URL: jtab --URL <URL>")
}

func printVersion() {
    fmt.Println("Current version of Jtab is: V1.8.0")
}

# jtab - JSON Table Formatter From Wlezz

`jtab` is a command-line tool written in Go for formatting JSON data as tables. It reads JSON data from a file and presents it in a tabular format. This tool is especially useful for visualizing structured JSON data in the terminal.

## Usage

### Prerequisites

Before using `jtab`, make sure you have Go installed on your system.

### Installation

Clone this repository and build the executable:

```bash
git clone https://github.com/Ghalib-craftLeProgrammeur/Jtab
cd jtab
go build jtab.go
```
Usage
To use jtab, provide a JSON file as a command-line argument:

```bash
jtab.exe <JSON file>
```
Replace <JSON file> with the path to your JSON file.

### Features
Reads JSON data from a file.
Formats the data as a table.
Supports displaying structured data in a more readable way.
Example
Assuming you have a JSON file named data.json with the following content:

```json

[
  {
    "name": "John",
    "age": 30,
    "city": "New York"
  },
  {
    "name": "Alice",
    "age": 25,
    "city": "San Francisco"
  }
]
```
You can format it as a table like this:

```bash
jtab.exe data.json
```
This will display the JSON data in a tabular format:

```sql

┌────────┬─────────────┬──────┬───────────────┬─────────┬─────────┐
│ COLOR  │ PRICE       │ ID   │ MAKE          │ MODEL   │ YEAR    │
├────────┼─────────────┼──────┼───────────────┼─────────┼─────────┤
│ Blue   │ 25000000.00 │ 1.00 │ Toyota        │ Camry   │ 2022.00 │
│ Red    │ 20000000.00 │ 2.00 │ Honda         │ Civic   │ 2021.00 │
│ Silver │ 40000000.00 │ 3.00 │ Ford          │ Escape  │ 2023.00 │
│ Black  │ 60000000.00 │ 4.00 │ Chevrolet     │ Malibu  │ 2022.00 │
│ White  │ 15000000.59 │ 5.00 │ Volkswagen    │ Jetta   │ 2021.00 │
│ Gray   │ 22000000.00 │ 6.00 │ BMW           │ X5      │ 2023.00 │
│ Green  │ 40000000.00 │ 7.00 │ Mercedes-Benz │ C-Class │ 2022.00 │
│ Silver │ 25000000.00 │ 8.00 │ Audi          │ A4      │ 2021.00 │
└────────┴─────────────┴──────┴───────────────┴─────────┴─────────┘
```
### Customization
You can customize the jtab tool to fit your needs. For example, you can modify the table style using the SetStyle method or add more formatting options.

### Dependencies:

[go-pretty](https://pkg.go.dev/github.com/jedib0t/go-pretty/v6/table) - A library for rendering and formatting tables in the terminal.
License
This project is licensed under the MIT License - see the LICENSE file for details.

Acknowledgments
Thanks to the go-pretty library for providing the table rendering functionality.
Feel free to contribute to this project or use it as a template for your own Go command-line tools!

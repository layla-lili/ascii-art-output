
go run . --output=test00.txt
Error: no input string specified
Usage: go run . [OPTION] [STRING] [BANNER]
EX: go run . --output=<fileName.txt> something standard

go run . "--output test00.txt banana standard"
Error: Must have equal '=' after the flag
Usage: go run . [OPTION] [STRING] [BANNER]
EX: go run . --output=<fileName.txt> something standard

go run . "--output=test00.txt banana standard"
Error: no input string specified
Usage: go run . [OPTION] [STRING] [BANNER]
EX: go run . --output=<fileName.txt> something standard
test() {
    echo "Test (0): go run . "--output test00.txt banana standard ""
    go run . "--output test00.txt banana standard "
    echo 
    echo "Test (0): go run . --output=test00.txt "First\\nTest" shadow"
    go run . --output=test00.txt "First\nTest" shadow
    echo
    echo "Test (1): go run . --output=test01.txt "hello" standard"
    go run . --output=test01.txt "hello" standard
    echo
    echo "Test (2): go run . --output=test02.txt "123 \-\> \#\$\%" standard"
    go run . --output=test02.txt "123 -> #$%" standard
    echo
    echo "Test (3): go run . --output=test03.txt "432 \-\> \#\$\%\&\@" shadow " 
    go run . --output=test03.txt "432 -> #$%&@" shadow
    echo
    echo "Test (4): go run . --output=test04.txt "There" shadow"
    go run . --output=test04.txt "There" shadow
    echo
    echo "Test (5): go run . --output=test05.txt "123 \-\> \"\#\$\%\@" thinkertoy"
    go run . --output=test05.txt "123 -> \"#$%@" thinkertoy
    echo
    echo "Test (6): go run . --output=test06.txt "2 you" thinkertoy"
    go run . --output=test06.txt "2 you" thinkertoy
    echo
    echo "Test (7): go run . --output=test07.txt 'Testing long output!' standard"
    go run . --output=test07.txt 'Testing long output!' standard
}
test
# Reguläre Ausdrücke in Go

In diesem Kursteil geht es um Behandlung regulärer Ausdrücke (fortan: «Regexp»)
in der Programmiersprache Go.

## Offizielle Dokumentation

Go verfügt über sein eigenes Dokumentationssystem. Folgende Einträge (als Befehl
angegeben, verlinkt auf die HTML-Dokumentation) sind besonders empfehlenswert.

- [`go doc regexp`](https://pkg.go.dev/regexp)
- [`go doc regexp.Regexp`](https://pkg.go.dev/regexp#Regexp)
- [`go doc regexp/syntax`](https://pkg.go.dev/regexp/syntax)

Für eine Übung wird die Ausgabe von `go doc` benötigt. Von daher ist es besser,
sich gleich mit dem Befehl vertraut zu machen, statt sich nur auf HTML-Version
zu verlassen.

## Dateien zum Download

Da Go über keine REPL verfügt, sind die vorgestellten Programme etwas
umfassender als etwa Python-Beispiele. Der Einfachheit halber können die ganzen
Beispiele unter folgendem Link als Zip-Datei heruntergeladen werden:

    TODO: Zip-Datei verlinken

## Syntax

Go verwendet grundsätzlich die gleiche regexp-Syntax wie Perl oder Python. Genau
genommen handelt es sich um die Syntax von
[RE2](https://pkg.go.dev/regexp/syntax), mit kleineren Ausnahmen.

## Implementierung

Die regexp-Implementierung von Go basiert auf der Arbeit von Ken Thompson in den
1960er-Jahren. Diese Implementierung wird auch in `grep` und `awk` verwendet.
Sie basiert auf endlichen Automaten und soll in ca. 400 Zeilen C-Code umsetzbar
sein. Die Laufzeitkomplexität dieser Implementierung (_Thompson NFA_: Thompson
Non-Deterministic Finite Automaton) wächst linear zur Eingabe. Andere
Implementierungen haben eine wesentlich höhere Laufzeitkomplexität. Siehe dazu
den Beitrag von Russ Cox: [Regular Expression Matching Can Be Simple And
Fast](https://swtch.com/~rsc/regexp/regexp1.html) (Der Artikel ist von 2007: das
Jahr, in dem die Entwicklung von Go lanciert worden ist ‒ u.a. von Ken Thompson.
Russ Cox gehört mittlerweile zum Kernteam von Go.)

## Passt es? Einfaches Matching

Ein einfaches Matching kann mit der Funktion `regexp.Match` umgesetzt werden:

    func Match(pattern string, b []byte) (matched bool, err error)
        Match reports whether the byte slice b contains any match of the regular
        expression pattern. More complicated queries need to use Compile and the
        full Regexp interface.

Liegen die zu prüfenden Daten als String und nicht als Byte-Array vor, kann man
`regexp.MatchString` verwenden:

    func MatchString(pattern string, s string) (matched bool, err error)
        MatchString reports whether the string s contains any match of the regular
        expression pattern. More complicated queries need to use Compile and the
        full Regexp interface.

Das folgende Beispiel `reglexp` (eine Kombination aus «REPL» und «Regexp»)
veranschaulicht die Verwendung der `Match`-Funktion:

```go
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [regexp]\n", os.Args[0])
		os.Exit(1)
	}
	pattern := os.Args[1]

	input := bufio.NewReader(os.Stdin)
	var line []byte
	var err error
	for ; err != io.EOF; line, err = input.ReadBytes('\n') {
		if len(line) == 0 {
			continue
		}
		if ok, _ := regexp.Match(pattern, line); ok {
			fmt.Print(string(line))
		}
	}
}
```

Das Beispiel lässt sich folgendermassen starten:

    $ go run replexp/main.go '#[A-Fa-f0-9]{6}'

Als Kommandozeilenargument wird die Regexp `#[A-Fa-f0-9]{6}` verwendet, womit
hexadezimale Farbangaben mit einleitendem Rautezeichen gematched werden können.
Die Interaktion mit dem Programm sieht dann etwa folgendermassen aus:

    #FFFFFF
    #FFFFFF
    #ffffff
    #ffffff
    #232323
    #232323
    #abc
    #foobar
    #deadbeef
    #deadbeef
    #999    

Eingabezeilen werden nach Betätigung von `[Return]` nur dann erneut ausgegeben,
sofern sie der Regexp genügen (`#FFFFFF`, `#ffffff`, `#232323` usw.),
andernfalls (`#abc`, `#foobar`, `#999`) jedoch nicht.

### Aufgabe 1

Probiere verschiedene Regexp als Kommandozeilenparameter aus. Gib anschliessend
Zeilen ein, und überlege dir vor der Betätigung von `[Return]`, ob die Zeile der
Regexp genügt oder nicht.

### Aufgabe 2

In der obigen Wiedergabe der REPL-Interaktion wird die Zeichenkette `#deadbeef`
ausgegeben, obwohl diese länger als die geforderten sechs Zeichen ist. Warum ist
das so, und wie lässt sich das korrigieren?

### Aufgabe 3

Das Programm verwendet die `regexp.Match`-Funktion, welche das Suchmuster als
String und den Text, in dem zu suchen ist, als Byte-Array erwartet. Die
`regexp.MatchString`-Funktion arbeitet mit zwei String-Argumenten. Was muss
alles umgestellt werden, damit `MatchString` verwendet werden kann? (Siehe
`input.ReadBytes` im Code.) Wird der Code durch die Umstellung insgesamt kürzer
und/oder besser lesbar, oder verschlechtert sich die Situation eher?

## Kompilierung und `Regexp`-Typ

Regexp können kompiliert und wiederverwendet werden. Das `regexp`-Paket bietet
vier Funktionen an, womit eine Regexp kompiliert und anschliessend als
`Regexp`-Typ wiederverwendet werden können:

- `func Compile(expr string) (*Regexp, error)`
- `func CompilePOSIX(expr string) (*Regexp, error)`
- `func MustCompile(str string) *Regexp`
- `func MustCompilePOSIX(str string) *Regexp`

Die Funktionen mit dem `Must`-Präfix werfen eine Runtime Panic, wenn der
angegebene Ausdruck nicht kompiliert werden kann. Das ist besonders bei hart
codierten regulären Ausdrücken sinnvoll, sodass fehlerhafter Code möglichst früh
und offensichtlich scheitert. Die beiden Funktionen _ohne_ `Must`-Prefix geben
stattdessen einen `error` zurück, wenn die Regexp nicht kompiliert werden kann.

Standardmässig wird die RE2-Syntax (PCRE mit kleinen Unterschieden) verwendet.
Die Funktionen mit dem `POSIX`-Suffix schränken die Syntax auf EREs ein.

Durch das optionale `Must`-Präfix und `POSIX`-Suffix ergeben sich die obigen
vier Funktionen.

Hat man durch gelungene Kompilierung eine `Regexp`-Struktur erhalten, bietet
diese eine Vielzahl von Methoden an. Hier sollen nur die `Find`-Methode und die
`Replace`-Methode mit ihren verschiedenen Ausprägungen von Interesse sein.
(`Match` und `MatchString` sind auf der `Regexp`-Struktur analog zu gebrauchen
wie vom `regexp`-Modul.)

Die Methode `Find` gibt es in verschiedenen Ausprägungen. Die folgenden Suffixe
können in der angegebenen Reihenfolge angefügt werden, um die jeweilige Semantik
zu erhalten. (Hier wird auf eine Eindeutschung der Dokumentation verzichtet):

- `All`: matches successive non-overlapping matches of the entire expression
- `String`: the argument is a string; otherwise it is a slice of bytes
- `Submatch`: the return value is a slice identifying the successive submatches of the expression
- `Index`: matches and submatches are identified by byte index pairs within the input string

Die Methodennamen folgen der Regexp:

    Find(All)?(String)?(Submatch)?(Index)?

Somit sind folgende Methoden vorhanden:

    func (re *Regexp) Find(b []byte) []byte
    func (re *Regexp) FindAll(b []byte, n int) [][]byte
    func (re *Regexp) FindAllIndex(b []byte, n int) [][]int
    func (re *Regexp) FindAllString(s string, n int) []string
    func (re *Regexp) FindAllStringIndex(s string, n int) [][]int
    func (re *Regexp) FindAllStringSubmatch(s string, n int) [][]string
    func (re *Regexp) FindAllStringSubmatchIndex(s string, n int) [][]int
    func (re *Regexp) FindAllSubmatch(b []byte, n int) [][][]byte
    func (re *Regexp) FindAllSubmatchIndex(b []byte, n int) [][]int
    func (re *Regexp) FindIndex(b []byte) (loc []int)
    func (re *Regexp) FindReaderIndex(r io.RuneReader) (loc []int)
    func (re *Regexp) FindReaderSubmatchIndex(r io.RuneReader) []int
    func (re *Regexp) FindString(s string) string
    func (re *Regexp) FindStringIndex(s string) (loc []int)
    func (re *Regexp) FindStringSubmatch(s string) []string
    func (re *Regexp) FindStringSubmatchIndex(s string) []int
    func (re *Regexp) FindSubmatch(b []byte) [][]byte
    func (re *Regexp) FindSubmatchIndex(b []byte) []int

Bei den `Replace`-Methoden gibt es wiederum Varianten zur Ersetzung des
gefundenen Textes mit Literalen (`Literal`) bzw. einer Funktion (`Func`), sowie
Ausprägungen für Byte-Arrays und Strings (`String`):

    func (re *Regexp) ReplaceAll(src, repl []byte) []byte
    func (re *Regexp) ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte
    func (re *Regexp) ReplaceAllLiteral(src, repl []byte) []byte
    func (re *Regexp) ReplaceAllLiteralString(src, repl string) string
    func (re *Regexp) ReplaceAllString(src, repl string) string
    func (re *Regexp) ReplaceAllStringFunc(src string, repl func(string) string) string

Für die nächste Übung kommt das Programm `godocfuncs/main.go` zum Einsatz:

```go
package main

import (
	"fmt"
	"os"
	"os/exec"

	dr "github.com/patrickbucher/dfdegoregexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [term]", os.Args[0])
		os.Exit(1)
	}

	cmd := exec.Command("go", "doc", os.Args[1])
	cmdOut, err := cmd.StdoutPipe()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	err = cmd.Start()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	dr.FilterLines(cmdOut, os.Stdout)

	err = cmd.Wait()
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
```

Dieses lässt sich folgendermassen ausführen:

    $ go run godocfuncs/main.go [package]

Es erwartet als Argument den Namen eines Go-Pakets, wie z.B. `string`, `regexp`
oder `regexp.Regexp`. Die Ausgabe der jeweiligen Dokumentationsseite (`go doc
string`, `go doc regexp` usw.) soll so gefiltert werden, dass nur die zum Paket
gehörigen Funktionsdeklarationen ausgegeben werden sollen. Hierzu wird die
Regexp `functionDeclaration` und die Funktion `dr.FilterLines` verwendet,
welche in der Datei `godocfuncs.go` definiert sind:

```go
package dfdegoregexp

import (
	"bufio"
	"io"
	"regexp"
)

const functionDeclaration = `` // TODO

func FilterLines(r io.Reader, w io.Writer) {
	p := regexp.MustCompilePOSIX(functionDeclaration)
	s := bufio.NewReader(r)
	var l []byte
	var err error
	for ; err != io.EOF; l, err = s.ReadBytes('\n') {
		if p.Match(l) {
			w.Write(l)
		}
	}
}
```

### Aufgabe 4

Die Regexp `functionDeclaration` (als String) muss ergänzt werden, damit das
Programm korrekt arbeitet. Hier ein Beispiel für eine korrekte Ausgabe:

    $ go run godocfuncs/main.go regexp
    func Match(pattern string, b []byte) (matched bool, err error)
    func MatchReader(pattern string, r io.RuneReader) (matched bool, err error)
    func MatchString(pattern string, s string) (matched bool, err error)
    func QuoteMeta(s string) string

Die Datei `godocfuncs_test.go` enthält einen Testfall, der folgendermassen
ausgeführt werden kann:

    $ go test -run TestFilterFuncLines

## Weitere Aufgaben

TODO: `manperf`

TODO: `emailextract`

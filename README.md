# Reguläre Ausdrücke in Go

In diesem Kursteil geht es um Behandlung regulärer Ausdrücke (fortan: «Regexp»)
in der Programmiersprache Go.

## Offizielle Dokumentation

Go verfügt über sein eigenes Dokumentationssystem (`go doc`). Folgende Einträge
(als Befehl angegeben, verlinkt auf die HTML-Dokumentation) sind für das
vorliegende Thema empfehlenswert:

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

- v0.0.1: [dfdegoregexp.zip](https://github.com/patrickbucher/dfdegoregexp/releases/download/v0.0.1/dfdegoregexp.zip)
- v0.1.0 (Release folgt): [dfdegoregexp.zip](https://github.com/patrickbucher/dfdegoregexp/releases/download/v0.1.0/dfdegoregexp.zip)

## Syntax

Go verwendet grundsätzlich die gleiche regexp-Syntax wie Perl oder Python. Genau
genommen handelt es sich um die Syntax von
[RE2](https://github.com/google/re2/wiki/Syntax), mit kleineren Ausnahmen. Es
ist aber auch möglich, die einfacheren ERE in Go zu verwenden. Hierzu stellt die
`regexp`-Bibliothek Funktionen mit dem `POSIX`-Suffix zur Verfügung (mehr dazu
später).

## Implementierung

Die regexp-Implementierung von Go basiert auf der Arbeit von Ken Thompson in den
1960er-Jahren. Sie basiert auf endlichen Automaten und soll in ca. 400 Zeilen
C-Code umsetzbar sein.

Die Laufzeitkomplexität dieser Implementierung (_Thompson NFA_: Thompson
Non-Deterministic Finite Automaton) wächst linear zur Eingabe. Andere
Implementierungen haben eine wesentlich höhere Laufzeitkomplexität. Siehe dazu
den Beitrag von Russ Cox: [Regular Expression Matching Can Be Simple And
Fast](https://swtch.com/~rsc/regexp/regexp1.html) (Der Artikel ist von 2007: das
Jahr, in dem die Entwicklung von Go lanciert worden ist; u.a. von Ken Thompson.
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
Darauf ist im Code entsprechend zu reagieren.

Standardmässig wird die RE2-Syntax (PCRE mit kleinen Unterschieden) verwendet.
Die Funktionen mit dem `POSIX`-Suffix schränken die Syntax auf EREs ein.

Durch die optionalen Präfixe `Must` und `POSIX` ergeben sich die obigen vier
Funktionen.

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

Die Methodennamen folgen (gemäss Dokumentation) der Regexp:

    Find(All)?(String)?(Submatch)?(Index)?

Somit sind folgende Methoden vorhanden (zwei Varianten, die auf einem
`io.RuneReader` basieren, sind der Vollständigkeit halber unten noch
aufgeführt):

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
    func (re *Regexp) FindString(s string) string
    func (re *Regexp) FindStringIndex(s string) (loc []int)
    func (re *Regexp) FindStringSubmatch(s string) []string
    func (re *Regexp) FindStringSubmatchIndex(s string) []int
    func (re *Regexp) FindSubmatch(b []byte) [][]byte
    func (re *Regexp) FindSubmatchIndex(b []byte) []int

    func (re *Regexp) FindReaderIndex(r io.RuneReader) (loc []int)
    func (re *Regexp) FindReaderSubmatchIndex(r io.RuneReader) []int

Den kleineren Rest der Funktionen, die _nicht_ dem genannten Muster entsprechen,
kann man bequem mittels `(e)grep` finden:

    $ go doc regexp.Regexp \
        | grep -Ev 'Find(All)?(String)?(Submatch)?(Index)?' \
        | grep '^func'

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

Es erwartet als Argument den Namen eines Go-Pakets oder -Symbols, wie z.B.
`string`, `regexp` oder `regexp.Regexp`. Die Ausgabe der jeweiligen
Dokumentationsseite (`go doc string`, `go doc regexp` usw.) soll so gefiltert
werden, dass nur die zum Paket gehörigen Funktionsdeklarationen ausgegeben
werden sollen. Hierzu wird die Regexp `functionDeclaration` und die Funktion
`dr.FilterLines` verwendet, welche in der Datei `godocfuncs.go` definiert sind:

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

## Sektionen von Manpages

Das Programm `manperf/main.go` erwartet als Argumente einen beliebigen Befehl
mit Kommandozeilenargumenten (z.B. `man 3 printf` oder `man go`). Der angegebene
Befehl wird ausgeführt, und die Ausgabe davon vom Programm abgefangen. Aus der
Ausgabe, die hierzu in der Form einer Manpage vorliegen muss, werden nun die
einzelnen Sektionstitel extrahiert:

```go
package main

import (
	"fmt"
	"os"
	"strings"

	dr "github.com/patrickbucher/dfdegoregexp"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "usage: %s [command] [arguments]", os.Args[0])
		os.Exit(1)
	}

	lines := dr.CommandOutput(os.Args[1], os.Args[1:]...)
	fmt.Print(strings.Join(dr.ExtractSectionsBad(lines), "\n"))

	// TODO: implement
	// fmt.Print(strings.Join(dr.ExtractSectionsBetter(lines), "\n"))
}
```

Hierzu werden drei Funktionen aus dem Paket `dfdegoregexp` verwendet (Importname
mit `dr` verkürzt), welche allesamt in `manperf.go` definiert sind:

- `CommandOutput`: Die Funktion führt das angegebene Programm mit Argumenten aus
  und liefert dessen Ausgabe als ein String-Slice zurück.
- `ExtractSectionsBad`: Die Funktion extrahiert die
  Sektionstitel aus der gegebenen Ausgabe und liefert sie als String-Slice
  zurück. **Die Implementierung ist fehlerhaft und inperformant!**
- `ExtractSectionsBetter`: Die Funktion ist noch zu implementieren und aus
  `manperf/main.go` aufzurufen.

### Aufgabe 5

Öffne die Datei `manperf_test.go` und betrachte den String-Slice
`expectedSections`. Rufe nun `man man` auf, und überprüfe, ob die in
`expectedSections` aufgelisteten Sektionen mit der tatsächlichen Ausgabe von
`man man` übereistimmen; auch in ihrer Reihenfolge. Nimm Korrekturen daran vor,
falls nötig.

### Aufgabe 6

Implementiere die Funktion `ExtractSectionsBetter`, sodass der entsprechende
Testfall `TestExtractSectionsBetter` durchläuft. Dieser kann folgendermassen
ausgeführt werden:

    $ go test -run TestExtractSectionsBetter

Läuft der Test durch, kannst du das Programm `manperf/main.go` so umschreiben,
dass die neue Funktion `ExtractSectionsBetter` anstelle von `ExtractSectionsBad`
verwendet wird. Rufe nun das Programm mit einer beliebigen Manpage aus, und
kontrolliere dessen Ausgabe, z.B.:

    $ go run manperf/main.go man 3 printf

### Aufgabe 7

Übertrage die funktioniertende Regexp aus der Funktion `ExtractSectionsBetter`
nach `ExtractSectionsBad`, sodass nun beide Testfälle durchlaufen:

    $ go test -run TestExtractSections.*

Verändere aber nicht die Struktur der Funktion `ExtractSectionsBad` und
verzichte auf eine kompilierte Regexp.

Starte nun die Benchmarks für die beiden Funktionen (ohne Tests):

    $ go test -bench . -run ^$

(Dem aufmerksamen Leser dürfte mittlerweile aufgefallen sein, dass das Go-Tool
auch reguläre Ausdrücke zur Selektion von Benchmarks und Testfällen akzeptiert.)

Versuche die Implementierung von `ExtractSectionsBetter` schnell zu machen.
(Tipp: verwende eine kompilierte `regexp.Regexp`, falls du noch nicht auf diese
Idee gekommen bist.)

Ist die Variante der `Compile`-Funktion mit dem `POSIX`-Suffix schneller? Warum (nicht)?

## Parsen von E-Mails; Gruppen

Im letzten Beispiel sollen Informationen aus E-Mail-Adressen ausgelesen werden.
Die E-Mail-Adressen liegen in verschiedenen Formaten vor. Diese Formate sollen
hier nicht definiert, sondern nur anhand vorgegebener Testfälle beschrieben
werden. (Die zu implementierenden Regeln sind also induktiv zu erschliessen.)

Die Datei `emailextract_test.go` definiert die Testfälle:

```go
package dfdegoregexp

import "testing"

type testCase struct {
	email, desc string
}

var tests = []testCase{
	{"joey@foobar.com", "Joey, FOOBAR"},
	{"harry.callahan@sfpd.gov", "Harry Callahan, SFPD"},
	{"homer.simpson69@aol.com", "Homer Simpson, *1969, AOL"},
	{"stan.marsh2012@southpark.com", "Stan Marsh, *2012, SOUTHPARK"},
	{"julius.caesar@rom.it", "Julius Caesar, ROM"},
}

func TestEmailExtract(t *testing.T) {
	for _, test := range tests {
		expected := test.desc
		actual := Extract(test.email)
		if actual != expected {
			t.Errorf(`Extract("%s"): expected "%s", got "%s"`, test.email, expected, actual)
		}
	}
}
```

Eine E-Mail-Adresse kann folgende Elemente enthalten:

1. einen Vornamen (immer vorhanden)
2. einen Nachnamen (optional, nach dem Punkt)
3. ein Geburtsdatum (optional, vor dem `@`-Zeichen; eine zweistellige Zahl
  impliziert das 20. Jahrhundert)
4. einen Firmennamen (immer vorhanden, Domain ohne TLD; in Grossbuchstaben)

Die Implementierung liegt in der Datei `emailextract.go` vor:

```go
package dfdegoregexp

import (
	"fmt"
	"regexp"
)

var (
	// TODO: This regexp must be written. Figure out group names according to switch/case below.
	r = ``
	p = regexp.MustCompile(r)
)

type emailInfo struct {
	first, last, company string
	year                 int
}

func (e emailInfo) String() string {
	if e.first == "" || e.company == "" {
		return ""
	}
	var y int
	if e.year != 0 {
		if e.year >= 100 {
			y = e.year
		} else {
			y = 1900 + e.year
		}
	}
	if e.last != "" && y != 0 {
		return fmt.Sprintf("%s %s, *%d, %s", e.first, e.last, y, e.company)
	}
	if e.last != "" && y == 0 {
		return fmt.Sprintf("%s %s, %s", e.first, e.last, e.company)
	}
	if e.last == "" && y == 0 {
		return fmt.Sprintf("%s, %s", e.first, e.company)
	}
	return ""
}

func Extract(email string) string {
	matches := p.FindStringSubmatch(email)
	if len(matches) == 0 {
		return ""
	}
	var ei emailInfo
	for i, name := range p.SubexpNames() {
		switch name {
		case "first":
			// TODO
		case "last":
			// TODO
		case "year":
			// TODO
		case "comp":
			// TODO
		}
	}
	return ei.String()
}
```

Die Funktion `Extract` nimmt einen String (die E-Mail-Adresse) entgegen, und
gibt einen neuen String zurück. Das ganze passiert in mehreren Schritten:

1. Das Pattern `p` wird auf die E-Mail-Adresse angewendet, um Matches zu
   erzeugen (siehe `go doc regexp.FindStringSubmatch`).
2. Die Matches von `p` werden abgearbeitet. Mit `p.SubexpNames` (siehe `go doc
   regexp.SubexpNames`) erhält man eine Map, deren Key den Index und Value den
   Namen des Matches bezeichnet (`i` kann als Index für `matches` verwendet
   werden).
3. Es wird eine `emailInfo`-Struktur erstellt, auf der die extrahierten
   Informationen gesammelt werden können. Anhand des Gruppennamens können die
   ausgelesenen Informationen mithilfe von `switch`/`case` ins richtige Feld der
   `emailInfo`-Struktur abgelegt werden. (Siehe `go doc strconv.Atoi` für
   Konvertierungen von String zu Integer).
4. Zuletzt wird die `String`-Methode von `emailInfo` verwendet, um die
   extrahierten Informationen in der Form auszugeben, wie sie vom Testfall
   erwartet werden. (Die `String`-Methode ist bereits korrekt implementiert.)

### Aufgabe 8

Rufe die Dokumentation zu Regexp-Syntax (`go doc regexp/syntax`) auf und finde
heraus, wie man «named capturing groups» definiert.

Definiere nun die Regexp (Variable `r` oben an der Datei `emailextract.go`) und
implementiere die Funktion `Extract` zu Ende, indem zu die `TODO`-Zeilen
innerhalb der `switch`/`case`-Kontrollstruktur abarbeitest.

Führe den Testfall aus, um die Implementierung zu überprüfen:

    $ go test -run TestEmailExtract

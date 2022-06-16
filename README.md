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
Russ Cox gehört zum Kernteam von Go.)

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
TODO: exercise version of replexp/main.go
```

Das Beispiel lässt sich folgendermassen starten:

    $ go run replexp/main.go '#[A-Fa-f0-9]{6}'

Als Kommandozeilenargument wird die Regexp `#[A-Fa-f0-9]{6}` verwendet, womit
hexadezimale Farbangaben mit einleitendem Rautezeichen gematched werden können.
Die Interaktion mit dem Programm sieht dann etwa folgendermassen aus:

    #fff
    #ffff
    #FFFFFF
    #FFFFFF
    #ffffff
    #ffffff
    #222
    #232323
    #232323
    #deaded
    #deaded
    #DeadEd
    #DeadEd
    #Deardr    

Eingabezeilen werden nach Betätigung von `[Return]` nur dann erneut ausgegeben,
sofern sie der Regexp genügen (`#FFFFFF`, `#ffffff`, `#232323` usw.).

### Aufgabe 1

Probiere verschiedene Regexp als Kommandozeilenparameter aus. Gib anschliessend
Zeilen ein, und überlege dir vor der Betätigung von `[Return]`, ob die Zeile der
Regexp genügt oder nicht.

### Aufgabe 2

Das Programm verwendet die ...

## Kompilierung und `Regexp`-Typ

- Compile
- CompilePOSIX
- MustCompile
- MustCompilePOSIX

Standardmässig wird die RE2-Syntax (PCRE mit kleinen Unterschieden) verwendet.
Die Funktionen mit dem `POSIX`-Suffix schränken die Syntax auf EREs ein.

Die Funktionen mit dem `Must`-Präfix werfen eine Runtime Panic, wenn der
angegebene Ausdruck nicht kompiliert werden kann. Das ist besonders bei hart
codierten regulären Ausdrücken sinnvoll, sodass fehlerhafter Code möglichst früh
und offensichtlich scheitert.

- Regexp

Methode `Find` kombinierbar mit:

- `All`: matches successive non-overlapping matches of the entire expression
- `Index`: matches and submatches are identified by byte index pairs within the input string
- `String`: the argument is a string; otherwise it is a slice of bytes
- `Submatch`: the return value is a slice identifying the successive submatches of the expression

Die Methodennamen folgen dem Ausdruck:

    Find(All)?(String)?(Submatch)?(Index)?

Das ergibt Methoden wie:

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

Weiter gibt es die `Replace`-Methoden:

    func (re *Regexp) ReplaceAll(src, repl []byte) []byte
    func (re *Regexp) ReplaceAllFunc(src []byte, repl func([]byte) []byte) []byte
    func (re *Regexp) ReplaceAllLiteral(src, repl []byte) []byte
    func (re *Regexp) ReplaceAllLiteralString(src, repl string) string
    func (re *Regexp) ReplaceAllString(src, repl string) string
    func (re *Regexp) ReplaceAllStringFunc(src string, repl func(string) string) string

Und `Split`:

    func (re *Regexp) Split(s string, n int) []string

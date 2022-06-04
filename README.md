# Reguläre Ausdrücke ("regexp") in Go

## Offizielle Dokumentation

- [`go doc regexp`](https://pkg.go.dev/regexp)
    - [`go doc regexp.Regexp`](https://pkg.go.dev/regexp#Regexp)
- [`go doc regexp/syntax`](https://pkg.go.dev/regexp/syntax)

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
Jahr, in dem die Entwicklung von Go lanciert worden ist ‒ u.a. von Ken Thompson.)

## Passt es? Einfaches Matching

TODO: regexp.Match, MatchReader, MatchString...

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

# tools

https://kahoot.it/

https://miro.com/

https://app.sli.do/

https://gist.github.com/asukakenji/ac8a05644a2e98f1d5ea8c299541fce9


# Pøipravit "conditional create"

Program který:

- akceptuje název adresáøe, nebo použije `cwd` pokud žádný není zadaný
- rekurzivnì projde adresáø, a sestaví seznam `*.sql` souborù, které v nìm jsou
- každý soubor otevøe pro ètení:
    - podívá se, zda v záhlaví je comment `/* conditional create */` - pokud ano, soubor se pøeskoèí
    - podívá se, zda jde o create skript (regulární výraz)
        - pokud ne, soubor se pøeskoèí
        - pokud ano, odloží si název vytváøené tabulky (pøedpoklad: tabulka je jen jedna)
- pokud soubor obsahuje komentáø 
- následnì se soubor otevøe pro zápis
- do souboru se doplní záhlaví, v záhlaví se dynamicky upraví název tabulky

```
-----------------------------------------------------------------------
/*--conditional_create--*/
select count(*) from dbc.tables 
where tablename='___daná__tabulka___' and databasename=( select database )
having count(*) > 0;

.IF ERRORCODE <> 0 THEN .QUIT 10;
.IF ACTIVITYCOUNT = 0 THEN .GOTO TABLE_DOES_NOT_EXIST;
.REMARK 'tabulka EXISTUJE'
.GOTO KONEC

.LABEL TABLE_DOES_NOT_EXIST
-----------------------------------------------------------------------
```


Pøedpoklady, omezení:

- každý create skript obsahuje pouze jeden CREATE TABLE statement (bonus points: ovìøte)
- každý skript je nakódovaný v `Windows-1250` - pozor, pokud jeho obsah naètete jako `string`,
  nelze nad tím stringem provádìt `range`, protože dojde k sanitizaci èeštiny! (použijte `[]byte`)

## Architektura

### main

```
func main() {    
    var path string             // cestu musíme nìjak pøipravit
    
    
    sniffer := sniffer.New(path).Work()                 // sends SQL files down the Chan()
    analyzer := analyzer.New(sniffer.Chan()).Work()     // sends create scripts down Chan()
    patcher := patcher.New(analyzer.Chan()).Work()      // patched create scripts
    
    // block and wait for all work to be done
    <- patcher.Done()
    
    // check for any errors
    if err := joinErr(sniffer.Err(), analyzer.Err(), patcher.Err()); err != nil {
        fmt.Log(err)
        os.Exit(1)
    }    
}

```

### data 

```
type Script struct {
    Path        string      // full path to the script
    Table       string      // name of the table this script creates
}
```

### sniffer

```
package sniffer

// Sniffer 
type Sniffer struct { 
    dir string
    out chan <- string
    err []error
    m   sync.Mutex
}

// New returns a new sniffer.
func (s *Scanner) New(path string)

// Work sniffs .sql files, and sends each file path to a channel.
// When no more files are available, the channel is closed.
// Work runs in the background (goroutine).
// Work returns self.
func (s *Scanner) Work() *Scanner

// Chan returns an output channel.
func (s *Scanner) Chan() chan <- string

// Err returns an error, if any was encountered. The function is mutex protected.
func (s *Scanner) Err() error
```

### analyzer

```
package analyzer

type Analyzer struct { 
    paths <- chan string            // incoming paths
    work  chan <- data.Script       // outgoing analyzed files
}

func New(paths <- chan string) *Analyzer

func (a *Analyzer) Work() *Analyzer

// Chan returns an output channel.
func (s *Scanner) Chan() chan <- data.Script

func (a *Analyzer) Err() error
```

### patcher

```
package patcher

type Patcher struct {
    work <- chan data.Script    // incoming scripts
    done chan <- struct {}      // done signal
}

// New returns a new patcher object.
func New(work <- chan data.Script) *Patcher

// Work patches files so that they contain the conditional create header.
// Work runs in the background (goroutine).
// When no more work is available, it closes the done channel (see Done()).
func (p *Patcher) Work() *Patcher

func (p *Patcher) Err() error

// Done returns the channel which is used by Work as completion signal.
func (p *Patcher) Done() chan <- struct{}
```


# Pøipravit generátor logù 

## Vstupy

- maska souboru formou akceptovatelná funkcí `os.Glob`
- typ logu který chceme generovat

## Výstupy

- log soubory v uvedeném formátu
- podpora dìlených souborù

## Bonus points

- paralelní zpracování - každý soubor v samostatné `goroutine`
- zapracovat do `bihelp` jako command

# Refactoring bihelp.exe

- místo knihovny `Cobra` použít knihovnu `docopt`

# git bisect

Napsat program, který na vstupu vezme commit (nebo sadu commitù), a do cíle uloží
všechny soubory, které daný commit upravil - ale v té podobì, v jaké je aktuální checkout.


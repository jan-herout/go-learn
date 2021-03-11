# tools

https://kahoot.it/

https://miro.com/

https://app.sli.do/

https://gist.github.com/asukakenji/ac8a05644a2e98f1d5ea8c299541fce9


# P�ipravit "conditional create"

Program kter�:

- akceptuje n�zev adres��e, nebo pou�ije `cwd` pokud ��dn� nen� zadan�
- rekurzivn� projde adres��, a sestav� seznam `*.sql` soubor�, kter� v n�m jsou
- ka�d� soubor otev�e pro �ten�:
    - pod�v� se, zda v z�hlav� je comment `/* conditional create */` - pokud ano, soubor se p�esko��
    - pod�v� se, zda jde o create skript (regul�rn� v�raz)
        - pokud ne, soubor se p�esko��
        - pokud ano, odlo�� si n�zev vytv��en� tabulky (p�edpoklad: tabulka je jen jedna)
- pokud soubor obsahuje koment�� 
- n�sledn� se soubor otev�e pro z�pis
- do souboru se dopln� z�hlav�, v z�hlav� se dynamicky uprav� n�zev tabulky

```
-----------------------------------------------------------------------
/*--conditional_create--*/
select count(*) from dbc.tables 
where tablename='___dan�__tabulka___' and databasename=( select database )
having count(*) > 0;

.IF ERRORCODE <> 0 THEN .QUIT 10;
.IF ACTIVITYCOUNT = 0 THEN .GOTO TABLE_DOES_NOT_EXIST;
.REMARK 'tabulka EXISTUJE'
.GOTO KONEC

.LABEL TABLE_DOES_NOT_EXIST
-----------------------------------------------------------------------
```


P�edpoklady, omezen�:

- ka�d� create skript obsahuje pouze jeden CREATE TABLE statement (bonus points: ov��te)
- ka�d� skript je nak�dovan� v `Windows-1250` - pozor, pokud jeho obsah na�tete jako `string`,
  nelze nad t�m stringem prov�d�t `range`, proto�e dojde k sanitizaci �e�tiny! (pou�ijte `[]byte`)

## Architektura

### main

```
func main() {    
    var path string             // cestu mus�me n�jak p�ipravit
    
    
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


# P�ipravit gener�tor log� 

## Vstupy

- maska souboru formou akceptovateln� funkc� `os.Glob`
- typ logu kter� chceme generovat

## V�stupy

- log soubory v uveden�m form�tu
- podpora d�len�ch soubor�

## Bonus points

- paraleln� zpracov�n� - ka�d� soubor v samostatn� `goroutine`
- zapracovat do `bihelp` jako command

# Refactoring bihelp.exe

- m�sto knihovny `Cobra` pou��t knihovnu `docopt`

# git bisect

Napsat program, kter� na vstupu vezme commit (nebo sadu commit�), a do c�le ulo��
v�echny soubory, kter� dan� commit upravil - ale v t� podob�, v jak� je aktu�ln� checkout.


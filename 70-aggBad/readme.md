# aggBad

## Popis problému

Validační komponenta nám zahazuje do tzv. "BAD" souborů ty vstupní záznamy, které při zpracování neodpovídají
specifikaci. Každý řádek v BAD souboru obsahuje informaci o tom, na jakém "portu" se problém vyskytnul,
chybový kód, a chybnou hodnotu.

Běžnou součástí vývoje je "agregace" těchto BAD souborů - po testovacím loadu je potřeba se podívat, co všechno
v BAD souborech máme, a posoudit, zda jde o chybu v naší implementaci, nebo zda jde o porušení extrakt specifikace.

## Očekávaný výstup

Napiště program, který na vstupu akceptuje cestu, a volitelně souborovou masku.
Program na základě dané cesty a jmenné konvence najde "správný" adresář (nebo adresáře), kde se mohou vyskytovat
BAD soubory. Program dále sestaví seznam BAD souborů, a jeden po druhém je načte. V případě, že soubor svojí velikostí
překročí stanovený limit (volitelný parametr z příkazové řádky), načte se pouze prvních N řádků (volitelný parametr
z příkazové řádky.

Program by měl běžet paralelně (goroutines).

Program následně sečte jednotlivé chyby. Pro každý soubor, port, a chybový kód se na výstupu vypíše počet jejich výskytů, a ukázkové hodnota.

## Procvičíte si: 

- parsing [argumentů](https://golang.org/pkg/flag/) z příkazové řádky
- práci se [soubory](https://golang.org/pkg/os/), práci s [bufferem a scanning](https://golang.org/pkg/bufio/)
- "portable" práci s [cestou](https://golang.org/pkg/path/filepath/) na souborovém systému, a rekurzivní scan
- orchestraci jednotlivých gorutin - channel, mutex

## Lekce

- [příprava prostředí](01-priprava-prostredi.md)

# Možné pokračování

Děkuji za to, že jste mi věnovali svůj čas. Doufám že tento bleskurychlý průlet základy jazyka
Vám dal hrubé základy, na kterých se dá stavět. Současně jsem si vědom, že opravdu se naučit 
nový programovací jazyk je možné pouze tím, že čteme a píšeme kód.

Na to konto bych si dovolil navrhnout vážným zájemcům **pokračování**. 

Moje představa je následující:

- podržel bych formát dvouhodinových schůzek
- představoval bych si, že čas budeme věnovat řešení konkrétních úloh, které:
    - vám dají možnost procvičit si základní témata, jako je například
      práce se soubory, práce s textem, práce s regulárními výrazy
    - nám na BI dají možnost vyřešit si konkrétní problémy, které často řešíme
    
Jinými slovy, mým cílem by bylo kolektivně pracovat na konkrétním problému, a to formou
párového programování.

Přepravil bych vždy konkrétní zadání (specifikaci) - které by nás provedlo step-by-step
procesem přípravy koknrétního programu. Součástí zadání by byl:

- hrubý popis datových struktur, které mají vzniknout
- výčet a navržená signatura jejich metod
- konkrétní test case, který "na konci dne" musí úspěšně projít

Nejlepší implementace by potom mohla být použita jako nástroj k řešení kokrétního problému
na BI.

V dalším textu je nástin konkrétních problémů, které by bylo možné řešit. Úlohy jsou
seřazené od těch "subjektivně jednodušších" k složitějším.

Všechny úlohy ve finále být řešeny tak, aby program běžel na MS Windows.

Všechny úlohy by měly být řešeny pouze s pomocí standardní knihovny.

## getLine

### Popis problému

Zpracováváme někdy i docela velké extrakty, a validační komponenta nám občas zahlásí problém na
konkrétním řádku někde "uprostřed" souboru. Pokud je soubor velký, je obtížné daný řádek najít, a
podívat se, "co je na něm špatně".

### Očekávaný výstup

Napište program, který na vstupu akceptuje název souboru, a číslo řádku, nebo "range" řádků.
Program načte soubor řádek za řádkem, a na standardní výstup zašle pouze ty řádky, které byly
byžadovány. V okamžiku, kdy program našel poslední "chtěný" řádek, zpracování se ukončí.

Program tedy **nebude** načítat celý soubor.

Součástí řešení je i vhodný návrh "API" z příkazové řádky.

### Procvičíte si:

- parsing [argumentů](https://golang.org/pkg/flag/) z příkazové řádky
- práci se [soubory](https://golang.org/pkg/os/), práci s [bufferem a scanning](https://golang.org/pkg/bufio/)


## aggBad

### Popis problému

Validační komponenta nám zahazuje do tzv. "BAD" souborů ty vstupní záznamy, které při zpracování neodpovídají
specifikaci. Každý řádek v BAD souboru obsahuje informaci o tom, na jakém "portu" se problém vyskytnul,
chybový kód, a chybnou hodnotu.

Běžnou součástí vývoje je "agregace" těchto BAD souborů - po testovacím loadu je potřeba se podívat, co všechno
v BAD souborech máme, a posoudit, zda jde o chybu v naší implementaci, nebo zda jde o porušení extrakt specifikace.

### Očekávaný výstup

Napiště program, který na vstupu akceptuje cestu, a volitelně souborovou masku.
Program na základě dané cesty a jmenné konvence najde "správný" adresář (nebo adresáře), kde se mohou vyskytovat
BAD soubory. Program dále sestaví seznam BAD souborů, a jeden po druhém je načte. V případě, že soubor svojí velikostí
překročí stanovený limit (volitelný parametr z příkazové řádky), načte se pouze prvních N řádků (volitelný parametr
z příkazové řádky.

Program by měl běžet paralelně (goroutines).

Program následně sečte jednotlivé chyby. Pro každý soubor, port, a chybový kód se na výstupu vypíše počet jejich výskytů, a ukázkové hodnota.

### Procvičíte si: 

- parsing [argumentů](https://golang.org/pkg/flag/) z příkazové řádky
- práci se [soubory](https://golang.org/pkg/os/), práci s [bufferem a scanning](https://golang.org/pkg/bufio/)
- "portable" práci s [cestou](https://golang.org/pkg/path/filepath/) na souborovém systému, a rekurzivní scan
- orchestraci jednotlivých gorutin - channel, mutex


## cc

### Popis problému

Běžnou součástí vývojového cyklu je likvidace prostředí v databázi, a jeho inicializace (create table).
Ruční příprava čistící dávky je časově náročná, a náchylná na chyby.

### Očekávaný výstup

Napište program, který na vstupu akceptuje cestu na souborovém systému, a "typ" operace, kterou chceme provést.

Program následně:

- rekurzivně projde adresář, a sestaví seznam `*.sql` souborů, které v něm jsou
- každý soubor otevře pro čtení:
    - podívá se, zda v záhlaví je comment `/*--conditional_create--*/` - pokud ano, soubor se přeskočí
    - podívá se, zda jde o create skript (regulární výraz)
        - pokud ne, soubor se přeskočí
        - pokud ano, odloží si název vytvářené tabulky (předpoklad: tabulka je jen jedna)    
- následně se soubor otevře pro zápis
- do souboru se doplní záhlaví, v záhlaví se dynamicky upraví název tabulky; záhlaví je svázané s typem operace,
  v principu vypadá zhruba takhle (s drobnými obměnami):
  
```  
-----------------------------------------------------------------------
/*--conditional_create--*/
select count(*) from dbc.tables 
where tablename='CEM_O2TV_EPG_KALT' and databasename=( select database )
having count(*) > 0;

.IF ERRORCODE <> 0 THEN .QUIT 10;
.IF ACTIVITYCOUNT = 0 THEN .GOTO TABLE_DOES_NOT_EXIST;
.REMARK 'tabulka EXISTUJE'
drop table CEM_O2TV_EPG_KALT;

.LABEL TABLE_DOES_NOT_EXIST
-----------------------------------------------------------------------
```

Program by měl běžet paralelně (goroutines).

### Procvičíte si

- parsing [argumentů](https://golang.org/pkg/flag/) z příkazové řádky
- práci se [soubory](https://golang.org/pkg/os/), práci s [bufferem a scanning](https://golang.org/pkg/bufio/)
- "portable" práci s [cestou](https://golang.org/pkg/path/filepath/) na souborovém systému, a rekurzivní scan
- práci s [regulárními výrazy](https://golang.org/pkg/regexp/)
- orchestraci jednotlivých gorutin - channel, mutex


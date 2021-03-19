# Libraries

Tento dokument obsahuje seznam vybraných open source knihoven, seřazený podle témat.
Existují další podobné seznamy, akceptované komunitou - cílem tohoto dokumentu není
tyto seznamy suplovat, ale držet si seznam těch knihoven, které považuji za nejvíc
relevantní pro řešení problémů, se kterými se běžně setkávám při své práci já.

Viz také:

- https://github.com/avelino/awesome-go

## Utilities

| url                                          | comment                            |
| -------------------------------------------- | ---------------------------------- |
| https://pkg.go.dev/github.com/ungerik/go-dry | Knihovna pomocných funkcí          |
| https://github.com/bitfield/script           | scripting, manipulace text souborů |

## Error handling

| url                                        | comment                                                                                       |
| ------------------------------------------ | --------------------------------------------------------------------------------------------- |
| https://github.com/hashicorp/go-multierror | Podporuje koncept "skupiny chyb".                                                             |
| https://github.com/emperror/emperror       | skupiny chyb, _handler_ rutiny, panic a recover, filtry, integrace na logovací knihovny, apod |
| https://github.com/palantir/stacktrace     | volitelný stack trace                                                                         |
| https://github.com/moshenahmias/failure    | chyba je custom `struct`, serializovaná jako JSON                                             |
| https://github.com/nsnikhil/erx            | obohacení o _kind_, _severity_, _operation_, stack trace, apod                                |

## Testing

| url                                       | comment                                                                                                |
| ----------------------------------------- | ------------------------------------------------------------------------------------------------------ |
| https://github.com/stretchr/testify       | assertions, mocking, ...                                                                               |
| https://github.com/franela/goblin         | assertions                                                                                             |
| https://github.com/maxatome/go-testdeep   | deep struct comparison, nadstavba nad `testing`, ala [Test::Deep](https://metacpan.org/pod/Test::Deep) |
| https://github.com/hexops/autogold        | testování s pomocí automaticky generovaných _golden_ _files_                                           |
| https://github.com/smartystreets/goconvey | velmi bohatá knihovna, se spoustou závislostí                                                          |
| https://github.com/yields/phony           | generátor umělých dat                                                                                  |

## Logging

| url                                     | comment                                                                                                    |
| --------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| https://github.com/golang/glog          | This is an efficient pure Go implementation of leveled logs                                                |
| https://github.com/sirupsen/logrus      | Logrus is a structured logger for Go (golang), completely API compatible with the standard library logger. |
| https://github.com/natefinch/lumberjack | Package lumberjack provides a rolling logger.                                                              |

## Konfigurace, env

| url                                | comment                                                 |
| ---------------------------------- | ------------------------------------------------------- |
| https://github.com/joho/godotenv   | Load environment variables from .env                    |
| https://github.com/subosito/gotenv | Load environment variables from .env or io.Reader in Go |

## CLI, prompts

| url                                 | comment                                                                                     |
| ----------------------------------- | ------------------------------------------------------------------------------------------- |
| https://github.com/docopt/docopt.go | http://http://docopt.org/                                                                   |
| https://github.com/spf13/cobra      | subjektivně poměrně složitý postup, jak definovat CLI rozhraní; ale je to de-facto standard |
| https://github.com/c-bata/go-prompt | autocomplete promts                                                                         |

## Dokumentace

| url                            | comment                                                                       |
| ------------------------------ | ----------------------------------------------------------------------------- |
| https://github.com/go101/golds | Golds is a Go local docs server, a Go docs generator, and a Go code reader. I |

## Deep copy

| url                                        | comment                                                     |
| ------------------------------------------ | ----------------------------------------------------------- |
| https://github.com/mitchellh/copystructure | copystructure is a Go library for deep copying values in Go |

## Templating

| url                              | comment                                                            |
| -------------------------------- | ------------------------------------------------------------------ |
| https://github.com/CloudyKit/jet | flexibilní, user friendly, a dobře dokumentovaný templating engine |

## Notifikace

| url                                  | comment                  |
| ------------------------------------ | ------------------------ |
| https://pkg.go.dev/gopkg.in/toast.v1 | Win10 notifikace (toast) |

## Text a validace vstupů

| url                                       | comment                      |
| ----------------------------------------- | ---------------------------- |
| https://github.com/hexops/gotextdiff/     | compare and diff             |
| https://github.com/asaskevich/govalidator | validátor stringů a struktur |

## YAML, JSON

| url                            | comment                                 |
| ------------------------------ | --------------------------------------- |
| https://github.com/ghodss/yaml | (de)serializace do YAML podle JSON tagů |

## Dev Tools & other

| url                                                | comment                                       |
| -------------------------------------------------- | --------------------------------------------- |
| https://taskfile.dev/                              | `task` místo `make` - konfigurace v YAML      |
| https://github.com/cortesi/modd                    | filesystem watcher                            |
| https://github.com/wtetsu/gaze                     | jako `modd`, ale jednodušší - bez konfigurace |
| https://tpaschalis.github.io/delve-debugging/      | jak použít `delve` debugger?                  |
| https://pkg.go.dev/golang.org/x/tools/cmd/stringer | generátor `fmt.Stringer` interface            |
| https://github.com/mattn/docx2md                   | `docx` to `md`                                |

## Databases

| url                              | comment                                                                                            |
| -------------------------------- | -------------------------------------------------------------------------------------------------- |
| https://github.com/jmoiron/sqlx  | wrapper nad `database/sql`, není to [DBI](https://metacpan.org/pod/DBI), ale zjednodušuje to život |
| https://github.com/godror/godror | Oracle database driver pro Go                                                                      |

## Git commit hooks

| url                                            | comment                   |
| ---------------------------------------------- | ------------------------- |
| https://github.com/TekWizely/pre-commit-golang | commit hooks pro git a Go |

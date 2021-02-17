#! C:/Strawberry/perl/bin/perl.EXE
#start,OMIT
use Path::Tiny;         # BLACK MAGIC!!!
foreach (path('c:\BI_Domain\learn-go\10-about\lines.pl')->lines) {
    print $x++," ", $_ if /^\s*(u|f)/
}
#end,OMIT
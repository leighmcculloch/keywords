# Keywords

A list and count of keywords in programming languages.

┌──────────────────────────────────────────────────────────────────────────────────────────────────┐
│                                                                                           ...... │
│                                                                                           ...... │
│                                                                                           ...... │
│                                                                                           ...... │
│                                                                                           ...... │
│                                                                                           ...... │
│                                                                                           ...... │
│                                                                                           ...... │
│                                                                                           ...... │
│                                                                                           ...... │
│                                                                                           ...... │
│                                                                                           ...... │
│                                                                                           ...... │
│                                                                             ...... ...... ...... │
│                                                                      ...... ...... ...... ...... │
│                                                 ...... ...... ...... ...... ...... ...... ...... │
│  22   ..25.. ..26.. ..32.. ..35.. ..36.. ..49.. ..51.. ..52.. ..54.. ..89.. .100.. .109.. .417.. │
│Lua    Go     Erlang C      Python Ruby   JS     Java   Rust   Dart   Swift  C#     C++    COBOL  │
└──────────────────────────────────────────────────────────────────────────────────────────────────┘

## Why does it matter?

The number of keywords in a programming language _can_ be an indication to it's simplicity/complexity, and that can impact the simplicity/complexity of the solutions that developers produce using it. Complex solutions can be more expensive to maintain and difficult to hire for. However, this is dependent on many factors and keyword count is only one; language idioms also play a massive part.

## Contribute

Don't see a language here? Please open a pull request adding it!

## Similar Work

* https://github.com/e3b0c442/keywords - An expanded version of this list, including a large collection of languages and formal references.

## Keyword List

### Lua (22 keywords)

and          break        do           else         elseif
end          false        for          function     goto
if           in           local        nil          not
or           repeat       return       then         true
until        while

### Golang (25 keywords)

break        default      func         interface    select
case         defer        go           map          struct
chan         else         goto         package      switch
const        fallthrough  if           range        type
continue     for          import       return       var

### Erlang (26 keywords)

after        and          andalso      begin        bnot
bor          bsl          bsr          bxor         case
catch        cond         div          end          fun
if           let          not          of           or
orelse       receive      rem          try          when
xor

### C (32 keywords)

auto         double       int          struct       break
else         long         switch       case         enum
register     typedef      char         extern       return
union        const        float        short        unsigned
continue     for          signed       void         default
goto         sizeof       volatile     do           if
static       while

### Python (35 keywords)

False        await        else         import       pass
None         break        except       in           raise
True         class        finally      is           return
and          continue     for          lambda       try
as           def          from         nonlocal     while
assert       del          global       not          with
async        elif         if           or           yield

### Ruby (36 keywords)

alias        and          begin        break        case
class        def          defined?     do           else
elsif        end          ensure       false        for
if           in           module       next         nil
not          or           redo         rescue       retry
return       self         super        then         true
undef        unless       until        when         while
yield

### JavaScript (49 keywords)

arguments    await        break        case         catch
class        const        continue     debugger     default
delete       do           else         enum         eval
export       extends      false        finally      for
from         function     if           implements   import
in           instanceof   interface    let          new
null         package      private      protected    public
return       static       super        switch       this
throw        true         try          typeof       var
void         while        with         yield

### Java (51 keywords)

abstract     continue     for          new          switch
assert       default      goto         package      synchronized
boolean      do           if           private      this
break        double       implements   protected    throw
byte         else         import       public       throws
case         enum         instanceof   return       transient
catch        extends      int          short        try
char         final        interface    static       void
class        finally      long         strictfp     volatile
const        float        native       super        while
_ (underscore)

### Rust (52 keywords)

abstract     alignof      as           become       box
break        const        continue     crate        do
else         enum         extern       false        final
fn           for          if           impl         in
let          loop         macro        match        mod
move         mut          offsetof     override     priv
proc         pub          pure         ref          return
Self         self         sizeof       static       struct
super        trait        true         type         typeof
unsafe       unsized      use          virtual      where
while        yield

### Dart (54 keywords)

abstract     as           assert       async        await
break        case         catch        class        const
continue     covariant    default      deferred     do
dynamic      else         enum         export       extends
external     factory      false        final        finally
for          get          if           implements   import
in           is           library      new          null
operator     part         rethrow      return       set
static       super        switch       sync         this
throw        true         try          typedef      var
void         while        with         yield

### Swift (89 keywords)

associatedtype class          deinit         enum
extension      func           import         init
inout          internal       let            operator
private        protocol       public         static
struct         subscript      typealias      var
break          case           continue       default
defer          do             else           fallthrough
for            guard          if             in
repeat         return         switch         where
while          as             catch          dynamicType
false          is             nil            rethrows
super          self           Self           throw
throws         true           try            #column
#file          #function      #line          #available
#column        #else#elseif   #endif         #file
#function      #if            #line          #selector
associativity  convenience    dynamic        didSet
final          get            infix          indirect
lazy           left           mutating       none
nonmutating    optional       override       postfix
precedence     prefix         Protocol       required
right          set            Type           unowned
weak           willSet

### C# (100 keywords)

abstract     add          alias        as           ascending
async        await        base         bool         break
byte         case         catch        char         checked
class        const        continue     decimal      default
delegate     descending   do           double       dynamic
else         enum         event        explicit     extern
false        finally      fixed        float        for
foreach      from         get          global       goto
group        if           implicit     in           int
interface    internal     into         is           join
let          lock         long         namespace    new
null         object       operator     orderby      out
override     params       partial      private      protected
public       readonly     ref          remove       return
sbyte        sealed       select       set          short
sizeof       stackalloc   static       string       struct
switch       this         throw        true         try
typeof       uint         ulong        unchecked    unsafe
ushort       using        value        var          virtual
void         volatile     where        while        yield

### C++ (109 keywords)

#define                  #defined                 #elif
#else                    #endif                   #error
#if                      #ifdef                   #ifndef
#include                 #line                    #pragma
#undef                   alignas                  alignof
and                      and_eq                   asm
atomic_cancel            atomic_commit            atomic_noexcept
auto                     bitand                   bitor
bool                     break                    case
catch                    char                     char16_t
char32_t                 class                    compl
concept                  const                    constexpr
const_cast               continue                 decltype
default                  delete                   do
double                   dynamic_cast             else
enum                     explicit                 export
extern                   false                    final
float                    for                      friend
goto                     if                       inline
int                      import                   long
module                   mutable                  namespace
new                      noexcept                 not
not_eq                   nullptr                  operator
or                       or_eq                    override
private                  protected                public
register                 reinterpret_cast         requires
return                   short                    signed
sizeof                   static                   static_assert
static_cast              struct                   switch
synchronized             template                 this
thread_local             throw                    transaction_safe
transaction_safe_dynamic true                     try
typedef                  typeid                   typename
union                    unsigned                 using
virtual                  void                     volatile
wchar_t                  while                    xor
xor_eq

### COBOL (z/OS platform) (417 keywords)

ACCEPT              ACCESS              ADD                 ADDRESS
ADVANCING           AFTER               ALL                 ALPHABET
ALPHABETIC          ALPHABETIC-LOWER    ALPHABETIC-UPPER    ALPHANUMERIC
ALPHANUMERIC-EDITED ALSO                ALTER               ALTERNATE
AND                 ANY                 APPLY               ARE
AREA                AREAS               ASCENDING           ASSIGN
AT                  AUTHOR              BASIS               BEFORE
BEGINNING           BINARY              BLANK               BLOCK
BOTTOM              BY                  CALL                CANCEL
CBL                 CD                  CF                  CH
CHARACTER           CHARACTERS          CLASS               CLASS-ID
CLOCK-UNITS         CLOSE               COBOL               CODE
CODE-SET            COLLATING           COLUMN              COM-REG
COMMA               COMMON              COMMUNICATION       COMP
COMP-1              COMP-2              COMP-3              COMP-4
COMP-5              COMPUTATIONAL       COMPUTATIONAL-1     COMPUTATIONAL-2
COMPUTATIONAL-3     COMPUTATIONAL-4     COMPUTATIONAL-5     COMPUTE
CONFIGURATION       CONTAINS            CONTENT             CONTINUE
CONTROL             CONTROLS            CONVERTING          COPY
CORR                CORRESPONDING       COUNT               CURRENCY
DATA                DATE-COMPILED       DATE-WRITTEN        DAY
DAY-OF-WEEK         DBCS                DE                  DEBUG-CONTENTS
DEBUG-ITEM          DEBUG-LINE          DEBUG-NAME          DEBUG-SUB-1
DEBUG-SUB-2         DEBUG-SUB-3         DEBUGGING           DECIMAL-POINT
DECLARATIVES        DELETE              DELIMITED           DELIMITER
DEPENDING           DESCENDING          DESTINATION         DETAIL
DISPLAY             DISPLAY-1           DIVIDE              DIVISION
DOWN                DUPLICATES          DYNAMIC             EGCS
EGI                 EJECT               ELSE                EMI
ENABLE              END                 END-ADD             END-CALL
END-COMPUTE         END-DELETE          END-DIVIDE          END-EVALUATE
END-IF              END-INVOKE          END-MULTIPLY        END-OF-PAGE
END-PERFORM         END-READ            END-RECEIVE         END-RETURN
END-REWRITE         END-SEARCH          END-START           END-STRING
END-SUBTRACT        END-UNSTRING        END-WRITE           ENDING
ENTER               ENTRY               ENVIRONMENT         EOP
EQUAL               ERROR               ESI                 EVALUATE
EVERY               EXCEPTION           EXIT                EXTEND
EXTERNAL            FALSE               FD                  FILE
FILE-CONTROL        FILLER              FINAL               FIRST
FOOTING             FOR                 FROM                FUNCTION
GENERATE            GIVING              GLOBAL              GO
GOBACK              GREATER             GROUP               HEADING
HIGH-VALUE          HIGH-VALUES         I-O                 I-O-CONTROL
ID                  IDENTIFICATION      IF                  IN
INDEX               INDEXED             INDICATE            INHERITS
INITIAL             INITIALIZE          INITIATE            INPUT
INPUT-OUTPUT        INSERT              INSPECT             INSTALLATION
INTO                INVALID             INVOKE              IS
JUST                JUSTIFIED           KANJI               KEY
LABEL               LAST                LEADING             LEFT
LENGTH              LESS                LIMIT               LIMITS
LINAGE              LINAGE-COUNTER      LINE                LINE-COUNTER
LINES               LINKAGE             LOCAL-STORAGE       LOCK
LOW-VALUE           LOW-VALUES          MEMORY              MERGE
MESSAGE             METACLASS           METHOD              METHOD-ID
MODE                MODULES             MORE-LABELS         MOVE
MULTIPLE            MULTIPLY            NATIVE              NATIVE_BINARY
NEGATIVE            NEXT                NO                  NOT
NULL                NULLS               NUMBER              NUMERIC
NUMERIC-EDITED      OBJECT              OBJECT-COMPUTER     OCCURS
OF                  OFF                 OMITTED             ON
OPEN                OPTIONAL            OR                  ORDER
ORGANIZATION        OTHER               OUTPUT              OVERFLOW
OVERRIDE            PACKED-DECIMAL      PADDING             PAGE
PAGE-COUNTER        PASSWORD            PERFORM             PF
PH                  PIC                 PICTURE             PLUS
POINTER             POSITION            POSITIVE            PRINTING
PROCEDURE           PROCEDURE-POINTER   PROCEDURES          PROCEED
PROCESSING          PROGRAM             PROGRAM-ID          PURGE
QUEUE               QUOTE               QUOTES              RANDOM
RD                  READ                READY               RECEIVE
RECORD              RECORDING           RECORDS             RECURSIVE
REDEFINES           REEL                REFERENCE           REFERENCES
RELATIVE            RELEASE             RELOAD              REMAINDER
REMOVAL             RENAMES             REPLACE             REPLACING
REPORT              REPORTING           REPORTS             REPOSITORY
RERUN               RESERVE             RESET               RETURN
RETURN-CODE         RETURNING           REVERSED            REWIND
REWRITE             RF                  RH                  RIGHT
ROUNDED             RUN                 SAME                SD
SEARCH              SECTION             SECURITY            SEGMENT
SEGMENT-LIMIT       SELECT              SELF                SEND
SENTENCE            SEPARATE            SEQUENCE            SEQUENTIAL
SERVICE             SET                 SHIFT-IN            SHIFT-OUT
SIGN                SIZE                SKIP1               SKIP2
SKIP3               SORT                SORT-CONTROL        SORT-CORE-SIZE
SORT-FILE-SIZE      SORT-MERGE          SORT-MESSAGE        SORT-MODE-SIZE
SORT-RETURN         SOURCE              SOURCE-COMPUTER     SPACE
SPACES              SPECIAL-NAMES       STANDARD            STANDARD-1
STANDARD-2          START               STATUS              STOP
STRING              SUB-QUEUE-1         SUB-QUEUE-2         SUB-QUEUE-3
SUBTRACT            SUM                 SUPER               SUPPRESS
SYMBOLIC            SYNC                SYNCHRONIZED        TABLE
TALLY               TALLYING            TAPE                TERMINAL
TERMINATE           TEST                TEXT                THAN
THEN                THROUGH             THRU                TIME
TIMES               TITLE               TO                  TOP
TRACE               TRAILING            TRUE                TYPE
UNIT                UNSTRING            UNTIL               UP
UPON                USAGE               USE                 USING
VALUE               VALUES              VARYING             WHEN
WHEN-COMPILED       WITH                WORDS               WORKING-STORAGE
WRITE               WRITE-ONLY          ZERO                ZEROES
ZEROS
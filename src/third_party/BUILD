go_get(
    name = 'go-sdl2',
    get = 'github.com/veandco/go-sdl2/sdl',
    visibility = ['PUBLIC'],
)

go_get(
    name = 'testify',
    get = 'github.com/stretchr/testify/assert',
    deps = [':spew', ':difflib'],
    visibility = ['PUBLIC'],
)

go_get(
    name = 'spew',
    get = 'github.com/davecgh/go-spew/spew',
    visibility = ['PUBLIC'],
)

go_get(
    name = 'difflib',
    get = 'github.com/pmezard/go-difflib/difflib',
    visibility = ['PUBLIC'],
)

go_get(
    name = 'gocui',
    get = 'github.com/jroimartin/gocui',
    deps = [':termbox-go'],
    visibility = ['PUBLIC'],
)

go_get(
    name = 'termbox-go',
    deps = [':go-runewidth'],
    get = 'github.com/nsf/termbox-go',
)

go_get(
    name = 'go-runewidth',
    get = 'github.com/mattn/go-runewidth',
)
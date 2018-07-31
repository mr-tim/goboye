go_library(
    name = 'cpu',
    srcs = [
        'opcodes.go',
        'registers.go',
    ],
    deps = [
    ],
)

go_test(
    name = 'cpu_test',
    srcs = [
        'registers_test.go',
    ],
    deps = [
        ':cpu',
        '//third_party:testify',
    ]
)
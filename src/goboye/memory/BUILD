go_library(
    name = 'memory',
    srcs = [
        'memory.go',
    ],
    visibility = [
        '//goboye/cpu:cpu',
    ]
)

go_test(
    name = 'memory_test',
    srcs = [
        'memory_test.go',
    ],
    deps = [
        ':memory',
        '//third_party:testify',
    ]
)
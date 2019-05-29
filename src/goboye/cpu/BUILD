go_library(
    name = 'cpu',
    srcs = [
        'opcodes.go',
        'opcode_handlers.go',
        'extended_opcode_handlers.go',
        'extended_opcodes.go',
        'interrupts.go',
        'processor.go',
        'registers.go',
        'disassembler.go',
    ],
    deps = [
        '//goboye/memory:memory',
        '//goboye/utils:utils',
    ],
    visibility = [
        '//goboye/debugger/...',
    ]
)

go_test(
    name = 'cpu_test',
    srcs = [
        'opcode_handlers_test.go',
        'extended_opcode_handlers_test.go',
        'registers_test.go',
    ],
    deps = [
        ':cpu',
        '//third_party:testify',
    ]
)
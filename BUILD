load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/arcenik/go-curiositybox-calendar-puzzle-solver
gazelle(name = "gazelle")

go_library(
    name = "go-curiositybox-calendar-puzzle-solver_lib",
    srcs = [
        "main.go",
        "print.go",
        "solver.go",
    ],
    importpath = "github.com/arcenik/go-curiositybox-calendar-puzzle-solver",
    visibility = ["//visibility:private"],
    deps = [
        "//board",
        "//dbsqlite3",
        "@com_github_alecthomas_kong//:go_default_library",
    ],
)

go_binary(
    name = "go-curiositybox-calendar-puzzle-solver",
    embed = [":go-curiositybox-calendar-puzzle-solver_lib"],
    visibility = ["//visibility:public"],
)

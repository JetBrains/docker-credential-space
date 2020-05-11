load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

licenses(["notice"])  # Apache 2.0

exports_files(["LICENSE"])

gazelle(
    name = "gazelle",
    command = "fix",
    external = "vendored",
    extra_args = [
        "-build_file_name",
        "BUILD,BUILD.bazel",  # Prioritize `BUILD` for newly added files.
    ],
    prefix = "github.com/JetBrains/docker-credential-space",
)

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "github.com/JetBrains/docker-credential-space",
    visibility = ["//visibility:private"],
    deps = [
        "//cli:go_default_library",
        "//vendor/github.com/google/subcommands:go_default_library",
    ],
)

go_binary(
    name = "docker-credential-space",
    embed = [":go_default_library"],
    pure = "on",
    visibility = ["//visibility:public"],
)

load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix bitbucket.com/42nerds/42nerdsctl
gazelle(name = "gazelle")

go_library(
    name = "go_default_library",
    srcs = ["main.go"],
    importpath = "bitbucket.com/42nerds/42nerdsctl",
    visibility = ["//visibility:private"],
    deps = ["//cmd:go_default_library"],
)

go_binary(
    name = "42nerdsctl",
    embed = [":go_default_library"],
    visibility = ["//visibility:public"],
)

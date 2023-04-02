load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/roronoadiogo/build-bazel-golang-hello
gazelle(name = "gazelle")

gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)

go_library(
    name = "project_lib",
    srcs = ["main.go"],
    importpath = "github.com/roronoadiogo/build-bazel-golang-hello",
    visibility = ["//visibility:private"],
    deps = [
        "@com_github_go_chi_chi_v5//:chi",
        "@com_github_go_chi_chi_v5//middleware",
    ],
)

go_binary(
    name = "bazel-project-example",
    embed = [":project_lib"],
    visibility = ["//visibility:public"],
)

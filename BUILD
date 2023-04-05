load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")
load("@bazel_gazelle//:def.bzl", "gazelle")
load("@io_bazel_rules_docker//go:image.bzl", "go_image")
load("@io_bazel_rules_docker//container:container.bzl", "container_push")

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

go_image(
    name = "image_docker_example",
    binary = ":bazel-project-example",
    importpath = "github.com/roronoadiogo/build-bazel-golang-hello",
    ports = ["3333"],
)

container_push(
    name = "push_image_docker_bazel",
    format = "Docker",
    image = ":image_docker_example",
    registry = "docker.io",
    repository = "roronoadiogo/bazel-project-example",
    tag = "dev",
)
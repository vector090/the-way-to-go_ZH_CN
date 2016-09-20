package(
  default_visibility = ["//visibility:public"],
)

load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")

go_library(
  name = "go_default_library",
  srcs = glob(["*.go"], exclude = ["*_test.go"]),
  deps = [
  ],
)

go_test(
  name = "uc_test",
  srcs = glob(["*_test.go"]),
  library = ":go_default_library",
  deps = [
  ],
  tags = [ "requires-network" ],
)

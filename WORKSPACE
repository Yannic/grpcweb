## Copyright 2018 Yannic Bonenberger. All rights reserved.
##
## Licensed under the Apache License, Version 2.0 (the "License");
## you may not use this file except in compliance with the License.
## You may obtain a copy of the License at
##
##    http://www.apache.org/licenses/LICENSE-2.0
##
## Unless required by applicable law or agreed to in writing, software
## distributed under the License is distributed on an "AS IS" BASIS,
## WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
## See the License for the specific language governing permissions and
## limitations under the License.

workspace(name = "com_github_yannic_grpcweb")

load("@bazel_tools//tools/build_defs/repo:git.bzl", "git_repository")

git_repository(
    name = "io_bazel_rules_go",
    commit = "0f0d007c89dc67a5a34490acafc5195b191f5045",
    remote = "https://github.com/bazelbuild/rules_go.git",
)

git_repository(
    name = "bazel_gazelle",
    commit = "6a1b93cc9b1c7e55e7d05a6d324bcf9d87ea3ab1",
    remote = "https://github.com/bazelbuild/bazel-gazelle.git",
)

load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains")
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies", "go_repository")

go_repository(
    name = "com_github_rs_cors",
    commit = "7af7a1e09ba336d2ea14b1ce73bf693c6837dbf6",
    importpath = "github.com/rs/cors",
)

go_repository(
    name = "org_golang_google_grpc",
    commit = "be077907e29fdb945d351e4284eb5361e7f8924e",
    importpath = "google.golang.org/grpc",
)

go_repository(
    name = "org_golang_x_text",
    commit = "be25de41fadfae372d6470bda81ca6beb55ef551",
    importpath = "golang.org/x/text",
)

go_repository(
    name = "com_github_mwitkow_go_conntrack",
    commit = "cc309e4a22231782e8893f3c35ced0967807a33e",
    importpath = "github.com/mwitkow/go-conntrack",
)

go_repository(
    name = "com_github_stretchr_testify",
    commit = "69483b4bd14f5845b5a1e55bca19e954e827f1d0",
    importpath = "github.com/stretchr/testify",
)

go_rules_dependencies()

go_register_toolchains()

gazelle_dependencies()

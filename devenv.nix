{ pkgs, lib, ... }:

{
  # Packages
  packages = [
    pkgs.golangci-lint
    pkgs.gofumpt
  ];

  # Environment

  # Golang
  languages.c.enable = true;
  languages.go.enable = true;

  # MySQLs

}

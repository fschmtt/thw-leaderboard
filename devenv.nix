{ pkgs, lib, ... }:

{
  # Packages
  packages = [
    pkgs.golangci-lint
    pkgs.gofumpt
    pkgs.yarn
  ];

  # Environment

  # Golang
  languages.c.enable = true;
  languages.go.enable = true;

  # JavaScript
  languages.javascript.enable = true;

  # TypeScript
  languages.typescript.enable = true;

  # MySQL
  services.mysql = {
      enable = true;
      package = pkgs.mysql80;
      initialDatabases = lib.mkDefault [{ name = "leaderboard"; }];
      ensureUsers = lib.mkDefault [
        {
          name = "leaderboard";
          password = "leaderboard";
          ensurePermissions = {
            "leaderboard.*" = "ALL PRIVILEGES";
          };
        }
      ];
      settings = {
        mysqld = {
          log_bin_trust_function_creators = 1;
          port = 35721;
        };
      };
    };

    # Environment
    env.DATABASE_HOST = "127.0.0.1";
    env.DATABASE_PORT = "35721";
    env.DATABASE_USER = "leaderboard";
    env.DATABASE_PASSWORD = "leaderboard";
    env.DATABASE_NAME = "leaderboard";

    env.SERVER_ADDRESS = ":8008";
}

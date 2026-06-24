{
  inputs = {
    nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";
    flake-parts.url = "github:hercules-ci/flake-parts";
    flake-utils.url = "github:numtide/flake-utils";
  };

  outputs = inputs @ {
    nixpkgs,
    flake-parts,
    flake-utils,
    ...
  }:
    flake-parts.lib.mkFlake {inherit inputs;} {
      systems = flake-utils.lib.defaultSystems;

      perSystem = {pkgs, ...}: let
        version = "1.0";

        runtimeDeps = pkgs.lib.optionals pkgs.stdenv.hostPlatform.isLinux [
          pkgs.libnotify
        ];

        pomodoro = pkgs.buildGoModule {
          pname = "pomodoro";
          inherit version;
          src = pkgs.lib.cleanSource ./.;

          vendorHash = "sha256-5pGJXObrNkUtIxBV3xJi0k04M+++olVRQJTqxbPZb4g=";

          nativeBuildInputs = [pkgs.makeWrapper];

          postInstall = ''
            mv $out/bin/pomodoro-go $out/bin/pomodoro
            wrapProgram $out/bin/pomodoro \
              --prefix PATH : ${pkgs.lib.makeBinPath runtimeDeps}
          '';
        };
      in {
        packages.default = pomodoro;
      };
    };
}

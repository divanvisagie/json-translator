{ pkgs ? import <nixpkgs> {} }:

pkgs.mkShell {
  buildInputs = [
    pkgs.go
    pkgs.pkg-config
    pkgs.gtk3
  ];

  shellHook = ''
    export CGO_ENABLED=1
    export PKG_CONFIG_PATH="${pkgs.gtk3.dev}/lib/pkgconfig"
  '';
}


{ pkgs ? import <nixpkgs> { } }:
pkgs.mkShell { buildInputs = [ pkgs.neo4j pkgs.neo4j-desktop ]; }

{
    description = "AnonymousSurvey";
    nixConfig = {
        bash-prompt = "anonsurvey > ";
    };
    inputs = {
        nixpkgs.url = "github:NixOS/nixpkgs/573c650e8a14b2faa0041645ab18aed7e60f0c9a";
        systems.url = "github:nix-systems/default/da67096a3b9bf56a91d16901293e51ba5b49a27e";
    };
    outputs = { self, nixpkgs, systems }: let
        eachSystem = nixpkgs.lib.genAttrs (import systems);
    in {
        devShells = eachSystem (system: let
          pkgs = nixpkgs.legacyPackages.${system};
          sqlDir = "/Users/noor/pet-projects/anonsurvey/local-db";
        in {
            default = with pkgs; mkShell {
                hardeningDisable = [ "fortify" ];
                # nativeBuildInputs = [ git sqlite ];
                # buildInputs = [ go zig ];
                nativeBuildInputs = [ git mariadb ];
                buildInputs = [ go ];
                shellHook = ''
                    export GOROOT=`go env GOROOT`
                    export GOVERSION=`go env GOVERSION`

                    # export CGO_ENABLED=1
                    # export CC="zig cc"
                    # export CXX="zig c++"
                '';
            };
        });
    };
}

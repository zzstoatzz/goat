{
  description = "Go AT Protocol CLI";

  inputs.nixpkgs.url = "github:NixOS/nixpkgs/nixos-unstable";

  outputs = { nixpkgs, ... }:
    let
      systems = [ "x86_64-linux" "aarch64-linux" "x86_64-darwin" "aarch64-darwin" ];
      forAllSystems = nixpkgs.lib.genAttrs systems;
    in
    {
      packages = forAllSystems (system:
        let
          pkgs = import nixpkgs { inherit system; };
          goat = pkgs.buildGoModule {
            pname = "goat";
            version = "0.2.3-space-spike";
            src = ./.;
            vendorHash = "sha256-ygXKxjgnNZeDmZNbYpT4SQB+UGgFgxYxiXRW7whBRYU=";

            ldflags = [ "-s" "-w" ];

            meta = {
              description = "General-purpose AT Protocol command-line tool";
              homepage = "https://github.com/bluesky-social/goat";
              license = with pkgs.lib.licenses; [ asl20 mit ];
              mainProgram = "goat";
            };
          };
        in
        {
          inherit goat;
          default = goat;
        });
    };
}

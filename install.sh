#!/bin/sh

set -e

os=$(echo "$(uname -s)" | tr '[:upper:]' '[:lower:]')
arch=$(uname -m)
version=${1:-latest}



if [ "$arch" = "x86_64" ]; then
    arch="amd64"
fi


regocli_url="https://github.com/drorIvry/rego-cli/releases/$version/download/rego-$os-$arch.tar.gz"

rego_install="${REGO_INSTALL:-$HOME/.rego}"
bin_dir="$rego_install/bin"
exe="$bin_dir/rego"

if [ ! -d "$bin_dir" ]; then
 	mkdir -p "$bin_dir"
fi

curl -q --fail --location --progress-bar --output "$exe.tar.gz" "$regocli_url"
cd "$bin_dir"
tar xzf "$exe.tar.gz"
chmod +x "$exe"
rm "$exe.tar.gz"

echo 'baseUrl: http://localhost:4004' > "$rego_install/.rego-cli"

echo "rego-cli was installed successfully to $exe"

if command -v rego >/dev/null; then
	echo "Run 'rego --help' to get started"
else
	case $SHELL in
	/bin/zsh) shell_profile=".zshrc" ;;
	*) shell_profile=".bash_profile" ;;
	esac
	echo "Manually add the directory to your \$HOME/$shell_profile (or similar)"
	echo "  export REGO_INSTALL=\"$rego_install\""
	echo "  export PATH=\"\$REGO_INSTALL/bin:\$PATH\""
	echo "Run '$exe --help' to get started"
fi
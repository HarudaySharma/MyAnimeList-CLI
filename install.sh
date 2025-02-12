# placing the binary files in right location
mkdir -p ~/.local/share/mal-cli
cp ./bin/main.bin ~/.local/share/mal-cli/mal-cli-daemon.bin

mkdir -p ~/.local/bin
cp ./bin/script.bin ~/.local/bin/mal-cli
export PATH="$PATH:~/.local/bin"

## dependencies check
install_fzf() {
    echo "installing fzf..."

    git clone --depth 1 https://github.com/junegunn/fzf.git ~/.fzf
    ~/.fzf/install
}

update_fzf() {
    echo "updating fzf..."

    # delete the already present fzf
    fzf_path=$(command -v fzf)
    if [ -n "$fzf_path" ]; then
        echo "Removing existing fzf at $fzf_path"
        rm -rf "$fzf_path"
    fi

    # Also remove ~/.fzf if it exists
    rm -rf ~/.fzf

    install_fzf
}

## check for fzf
if command -v fzf >/dev/null 2>&1; then

    fzf_min_version="0.54.0"
    installed_version=$(fzf --version | awk '{print $1}')

    if [ "$(printf '%s\n' "$fzf_min_version" "$installed_version" | sort -V | head -n1)" = "$fzf_min_version" ]; then
        echo "fzf is installed and meets the version requirement ($fzf_min_version+)."
    else
        echo "fzf is installed but doesn't meet the version requirement ($fzf_min_version+)."
        update_fzf
    fi
else
    echo "fzf not found"
    install_fzf
fi


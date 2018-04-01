export AUTOJUMP_SOURCED=1

if [[ -d ~/.jmp/functions  ]]; then
	fpath=(~/.jmp/functions ${fpath})
fi

# default autojump command
j() {
    if [[ ${1} == -* ]] && [[ ${1} != "--" ]]; then
        jmp ${@}
        return
    fi

    setopt localoptions noautonamedirs
    local output="$(jmp ${@})"
    if [[ -d "${output}" ]]; then
        if [ -t 1 ]; then  # if stdout is a terminal, use colors
                echo -e "\\033[32m${output}\\033[0m"
        else
                echo -e "${output}"
        fi
        cd "${output}"
    else
        echo "autojump: directory '${@}' not found"
        echo "\n${output}\n"
        echo "Try \`jmp -h\` for more information."
        false
    fi
}

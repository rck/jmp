export AUTOJUMP_SOURCED=1

# enable tab completion
_jmp() {
        local cur
        cur=${COMP_WORDS[*]:1}
        comps=$(jmp --c $cur)
        while read i; do
            COMPREPLY=("${COMPREPLY[@]}" "${i}")
        done <<EOF
        $comps
EOF
}
complete -F _jmp j


# default autojump command
j() {
    if [[ ${1} == -* ]] && [[ ${1} != "--" ]]; then
        jmp  ${@}
        return
    fi

    output="$(jmp ${@})"
    if [[ -d "${output}" ]]; then
        if [ -t 1 ]; then  # if stdout is a terminal, use colors
                echo -e "\\033[32m${output}\\033[0m"
        else
                echo -e "${output}"
        fi
        cd "${output}"
    else
        echo "jmp: directory '${@}' not found"
        echo "\n${output}\n"
        echo "Try \`jmp -h\` for more information."
        false
    fi
}

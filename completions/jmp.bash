export AUTOJUMP_SOURCED=1

# enable tab completion
_jmp() {
	local cur
	cur=${COMP_WORDS[*]:1}
	comps=$(jmp --c "$cur")
	while read -r i; do
		COMPREPLY=("${COMPREPLY[@]}" "${i}")
	done <<EOF
   $comps
EOF
}
complete -F _jmp j

# default jmp command
j() {
	if [[ ${1} == -* ]] && [[ ${1} != "--" ]]; then
		jmp "${@}"
		return
	fi

	output="$(jmp "${@}")"
	if [ -t 1 ]; then  # if stdout is a terminal, use colors
		echo -e "\\033[32m${output}\\033[0m"
	else
		echo -e "${output}"
	fi
	cd "${output}" || exit
}

# jump to child directory (subdirectory of current path)
jc() {
	if [[ ${1} == -* ]] && [[ ${1} != "--" ]]; then
		jmp "${@}"
		return
	else
		j "$(pwd)" "${@}"
	fi
}

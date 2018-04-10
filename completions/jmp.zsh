export AUTOJUMP_SOURCED=1

if [[ -d ~/.jmp/functions  ]]; then
	fpath=(~/.jmp/functions ${fpath})
fi

# default jmp  command
j() {
    if [[ ${1} == -* ]] && [[ ${1} != "--" ]]; then
        jmp ${@}
        return
    fi

    setopt localoptions noautonamedirs
	 local output="$(jmp ${@})"
	 if [ -t 1 ]; then  # if stdout is a terminal, use colors
		 echo -e "\\033[32m${output}\\033[0m"
	 else
		 echo -e "${output}"
	 fi
	 cd "${output}"
}

# jump to child directory (subdirectory of current path)
jc() {
	if [[ ${1} == -* ]] && [[ ${1} != "--" ]]; then
		jmp ${@}
		return
	else
		j $(pwd) ${@}
	fi
}

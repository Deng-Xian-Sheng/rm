package CLI

import "flag"

var (
	FlagRCommit = "Attempt to remove the file hierarchy rooted in each file argument.  The -R option implies the -d option.  If the -i option is specified, the user is prompted for confirmation before each directory's contents are processed (as well as before the attempt is made to remove the directory).  If the user does not respond affirmatively, the file hierarchy rooted in that directory is skipped."
	FlagFCommit = "Attempt to remove the files without prompting for confirmation, regardless of the file's permissions.  If the file does not exist, do not display a diagnostic message or modify the exit status to reflect an error.  The -f option overrides any previous -i options."
	FlagR bool // 删除文件夹
	FlagF bool // 不打印日志
)

func init() {
	flag.BoolVar(&FlagR, "r", false, FlagRCommit)
	flag.BoolVar(&FlagF, "f", false, FlagFCommit)
	flag.Parse()
}
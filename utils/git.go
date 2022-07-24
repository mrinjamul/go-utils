package utils

// ShortGitCommit returns the short form of the git commit hash
func ShortGitCommit(fullGitCommit string) (shortCommit string) {
	if len(fullGitCommit) < 8 {
		return fullGitCommit
	}
	if len(fullGitCommit) >= 7 {
		shortCommit = fullGitCommit[0:7]
	}

	return shortCommit
}

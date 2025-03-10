package tests

import (
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCompletion(t *testing.T) {
	if runtime.GOOS == "windows" {
		t.Skip("skipping test on windows.")
	}

	ts := newTester(t)
	defer ts.teardown()

	t.Run("completion help", func(t *testing.T) {
		out, err := ts.run("completion")
		assert.NoError(t, err)
		assert.Contains(t, out, "Source for auto completion in bash")
		assert.Contains(t, out, "Source for auto completion in zsh")
	})

	t.Run("bash completion", func(t *testing.T) {
		bash := `_gopass_bash_autocomplete() {
     local cur opts base
     COMPREPLY=()
     cur="${COMP_WORDS[COMP_CWORD]}"
     opts=$( ${COMP_WORDS[@]:0:$COMP_CWORD} --generate-bash-completion )
     local IFS=$'\n'
     COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
     return 0
 }

complete -F _gopass_bash_autocomplete gopass`

		out, err := ts.run("completion bash")
		assert.NoError(t, err)
		assert.Equal(t, bash, out)
	})

	t.Run("zsh completion", func(t *testing.T) {
		out, err := ts.run("completion zsh")
		assert.NoError(t, err)
		assert.Contains(t, out, "compdef gopass")
	})

	t.Run("fish completion", func(t *testing.T) {
		out, err := ts.run("completion fish")
		assert.NoError(t, err)
		assert.Contains(t, out, "complete")
	})
}

func TestCompletionNoPath(t *testing.T) {
	ts := newTester(t)
	defer ts.teardown()

	t.Setenv("PATH", "/tmp/foobar")

	t.Run("generate bash", func(t *testing.T) {
		_, err := ts.run("--generate-bash-completion")
		assert.NoError(t, err)
	})
}

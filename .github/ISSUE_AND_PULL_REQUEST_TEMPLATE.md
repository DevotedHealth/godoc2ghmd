{{if .IsIssue}}
<!-- comment so this doesn't show up in the UI 
Thanks for taking the time to report an issue.

Since this is a fork of a fork, we also recommend that you take a look
<a href="https://github.com/GandalfUK/godoc2ghmd">https://github.com/GandalfUK/godoc2ghmd</a>,
<a href="https://github.com/davecheney/godoc2md">https://github.com/davecheney/godoc2md</a>, or
<a href="https://github.com/wdamron/godoc2gh">https://github.com/wdamron/godoc2gh</a>.

It may be entirely possible that using one of those forks will solve your
immeadiate problem. That shouldn't stop you from reporting but might be faster
as we don't often check here.
-->

### Expected Behavior


### Actual Behavior


### Fix
<!-- If you have a theory on how you might fix this, that would be much
appreciated. Pull requests are welcome as well! -->


{{end}}
{{if .IsPullRequest}}
<!-- comment so this doesn't show up in the UI.

Thanks for taking the time to create a pull request.

Since this is a fork of a fork, it may make more sense to contribute to the upstream repository
in addition or instead of contributing here. The upstream forks are: 
<a href="https://github.com/GandalfUK/godoc2ghmd">https://github.com/GandalfUK/godoc2ghmd</a>,
<a href="https://github.com/davecheney/godoc2md">https://github.com/davecheney/godoc2md</a>, or
<a href="https://github.com/wdamron/godoc2gh">https://github.com/wdamron/godoc2gh</a>.

It may also be entirely possible that using one of those forks will solve your
immeadiate problem. That shouldn't stop you from creating the PR but might be faster.
-->

### Intended change


### Checklist

- [ ] Ran `make lint` and no linting issues
- [ ] Ran `make test` and all the tests passes
- [ ] Added tests to new added code
- [ ] Ran `make all` so all the examples are updated
{{end}}

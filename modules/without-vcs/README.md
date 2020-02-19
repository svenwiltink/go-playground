Normally the name of a module is the same as the version control server used: 
`github.com/svenwiltink/youtube-dl` for example. Sometimes it is desired to just
be able to use local directories for modules, not keeping a VCS in mind. This
module does just that. It is called `banana` and its packages will always start
with `banana/`. The only subpackage in this module is in the `more` directory: 
`banana/more`. The `banana.go` file, which is in the `main` package can simply
import the package and use the exported `BigBanana` variable.


Building a module that uses this system is the same as other ones:
```sh
code/go-playground/modules/without-vcs>go build .
code/go-playground/modules/without-vcs>./banana
BIG YELLOW UNIT
```
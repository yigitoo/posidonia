# posidonia
A project for saving Posidonia Oceanica and getting them in quarantina. (With under control and restrict areas)

# SUBMODULE OPERATIONS
<pre>
With version 2.13 of Git and later, `--recurse-submodules` can be used instead of `--recursive`:

    git clone --recurse-submodules -j8 git://github.com/foo/bar.git
    cd bar

<sup>Editor’s note: `-j8` is an optional performance optimization that became available in version 2.8, and fetches up to 8 submodules at a time in parallel — see `man git-clone`.</sup>

With version 1.9 of Git up until version 2.12 (`-j` flag only available in version 2.8+):

    git clone --recursive -j8 git://github.com/foo/bar.git
    cd bar

With version 1.6.5 of Git and later, you can use:

    git clone --recursive git://github.com/foo/bar.git
    cd bar

For already cloned repos, or older Git versions, use:

    git clone git://github.com/foo/bar.git
    cd bar
    git submodule update --init --recursive
</pre>

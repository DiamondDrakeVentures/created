# packdiff

packdiff parses two [Crash Assistant]'s `modlist.json` and shows the difference in a similar manner
to `diff`.

In the future, we may implement a functionality to parse [packwiz] packs directly.

## Usage

``` shell
Usage: packdiff [options...] <a> <b>

  -r        Show removal(s).
  -a        Show addition(s).
  -u        Show update(s).
  -n        Show unchanged.
  -q        Hide statistics.
```

## Reading output

Each line contains a mod ID and, optionally, one or more version strings.
If the mod ID could not be determined, the `mod.jar` name is used instead.

Changes are denoted by either `--`, `++`, or `^^` prefix.

This indicates no change to `aa4_atlas`.

``` text
   aa4_atlas: 1.1.2
```

This indicates a new addition with the mod ID `accessories`.

``` text
++ accessories: 1.1.0-beta.48+1.21.1
```

This indicates a removal of `elytratrims`.

``` text
-- elytratrims: 3.5.9
```

This indicates an update of `neoforge`.
Note that the current version and the new version are listed, seperated by `=>`.
A downgrade is also indicated as an update.

``` text
^^ neoforge: neoforge-21.1.194 => neoforge-21.1.196
```

## License

Unlike the rest of this repository, packdiff is under the [MIT License].

[Crash Assistant]: https://modrinth.com/mod/crash-assistant
[packwiz]: https://packwiz.infra.link/
[MIT License]: LICENSE

Blinky
==========

## Description

[Blinky][] could be used to generate to display system informations
and/or ASCII distributions logos.

![Screenshot](blinky-0.3.0.png)

```bash
$ blinky system infos
OS: Arch Linux x86_64
Hostname: nlamirault2
Kernel: 4.2.2-1-ARCH
Memory: 5931634688/8330813440 %!d(float64=9.205463422308974)
Processor: Intel(R) Core(TM) i5-2520M CPU @ 2.50GHz
Uptime: 1445397458
```

```bash
$ blinky logo list
 List logos
 - arch
 - debian
 - mint
 - osx
 - ubuntu
 - windows
```

```bash
$ blinky logo --name="arch" display
                   -
                  o+
                 ooo
                +oooo:
               +oooooo:
              --+oooooo+:
             /:-:++oooo+:
            /++++/+++++++:
           /++++++++++++++:
          /+++ooooooooooooo/
         ./ooosssso++osssssso+
        .oossssso-\\\\/ossssss+
       -osssssso.      :ssssssso.
      :osssssss/        osssso+++.
     /ossssssss/        +ssssooo/-
    /ossssso+/:-        -:/+osssso+-
   +sso+:-\                 \.-/+oso:
  ++:.                           \-/+/
 .\                                 \/
```


```bash
$ blinky display

                   -                     OS: Arch Linux x86_64
                  o+                     Kernel: 4.2.2-1-ARCH
                 ooo                     Hostname: nlamirault2
                +oooo:                   Uptime: 1445397458
               +oooooo:                  Processor: Intel(R) Core(TM) i5-2520M CPU @ 2.50GHz
              --+oooooo+:                Mem: 5930680320/8330813440 %!d(float64=9.210920944593857)
             /:-:++oooo+:
            /++++/+++++++:
           /++++++++++++++:
          /+++ooooooooooooo/
         ./ooosssso++osssssso+
        .oossssso-\\\\/ossssss+
       -osssssso.      :ssssssso.
      :osssssss/        osssso+++.
     /ossssssss/        +ssssooo/-
    /ossssso+/:-        -:/+osssso+-
   +sso+:-\                 \.-/+oso:
  ++:.                           \-/+/
 .\                                 \/

```

## Support / Contribute

See [here](CONTRIBUTING.md)


## Changelog

A changelog is available [here](ChangeLog.md).


## License

See [LICENSE](LICENSE).


## Contact

[Blinky]: https://github.com/nlamirault/blinky
[COPYING]: https://github.com/nlamirault/blinky/blob/master/COPYING
[Issue tracker]: https://github.com/nlamirault/blinky/issues

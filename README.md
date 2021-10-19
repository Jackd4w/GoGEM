# GoGEM

![GoGEM Logo](/goGEM-Logo.png)

## About

GoGEM is a tool designed to make the deployment process of your iGEM-Wiki as easy as possible.
It is able to automatically fetch a page from a WordPress Instance hosted by your team. If you use this tool please give credit to me and link to this repo (<https://github.com/Jackd4w/GoGEM>).

## Installation

To install this programm you can use the _go install github.com/Jackd4w/GoGEM_ command.
alternatively the cloning of this repo and _go run_ or _go build_ can be used.

A pre-compiled version will be available with each release.

## Usage

If installed the tool can be used by executing the _GoGEM_ command in your CLI.

## Examples

Upload: _GoGEM upload -u "[Your Username]" -y 2021 -t "TU_Darmstadt" -w "[Your WP Wiki]" -o "test"_

Save your WP Page: _GoGEM fetchWP [URL]_

Purge: _GoGEM purge -u "[Username]" -y [Wiki Year] -t "[Teamname]" -o "[Offset]"_

## Issues

Please report Issues to this repo (<https://github.com/Jackd4w/GoGEM>), this is where the development will continue.

## Contribution

I tried to comment the code reasonably. Please try writing verbose comments when contributing, as this is intended to be a project a beginner programmer can understand. As stated above development will only be continued on this repo (<https://github.com/Jackd4w/GoGEM>).
